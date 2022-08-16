package store

import (
	"fmt"
	"strings"

	dynamicstruct "github.com/Ompluscator/dynamic-struct"
	"github.com/dereference-xyz/trickle/model"
	"gorm.io/gorm"
)

// We use this to represent an instance of a dynamically created struct of an account.
// This is used by Gorm which uses reflection.
type accountInstance struct {
	// Type name of the account.
	name string
	// Instance of the dynamically created struct.
	instance interface{}
}

// Driver that contains any database-specific functionality.
type Driver interface {
	// Return gorm dialector.
	Dialector() gorm.Dialector
	// Serialize the given value based on the data type.
	Serialize(dataType model.DataType, value interface{}) interface{}
	// Deserialize the given value based on the data type.
	Deserialize(dataType model.DataType, value interface{}) interface{}
}

// Store of account data.
type AccountStore struct {
	// Gorm db instance.
	db *gorm.DB
	// Driver to use for database-specific functionality.
	driver Driver
}

// Create a new account store with the given driver.
func NewAccountStore(driver Driver) (*AccountStore, error) {
	db, err := gorm.Open(driver.Dialector(), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &AccountStore{
		db:     db,
		driver: driver,
	}, nil
}

// Create instances of structs with fields that are dynamically determined by the accounts propert types.
func createDynamicStructs(accountTypes []*model.AccountType) []*accountInstance {
	instances := []*accountInstance{}
	for _, acc := range accountTypes {
		builder := dynamicstruct.NewStruct()
		for _, prop := range acc.PropertyTypes {
			zeroValue := zeroValue(prop.DataType)
			builder.AddField(
				strings.Title(prop.Name),
				zeroValue,
				fmt.Sprintf("gorm:\"column:%s\"", prop.Name))
		}
		inst := builder.Build().New()
		instances = append(instances, &accountInstance{
			name:     acc.Name,
			instance: inst,
		})
	}
	return instances
}

// Use gorm to create/update schema in underlying database based on the given program type.
func (st *AccountStore) AutoMigrate(programType *model.ProgramType) error {
	dynamicStructs := createDynamicStructs(programType.AccountTypes)

	for _, inst := range dynamicStructs {
		err := st.db.Table(inst.name).AutoMigrate(inst.instance)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create rows for given account data in database.
func (st *AccountStore) Create(accounts []*model.Account) error {
	rowsByAccountType := make(map[string][]map[string]interface{})
	for _, acc := range accounts {
		converted, err := st.serialize(acc.AccountType, acc.Data)
		if err != nil {
			return err
		}
		rowsByAccountType[acc.Type] = append(rowsByAccountType[acc.Type], converted)
	}
	for accountType, rows := range rowsByAccountType {
		result := st.db.Table(accountType).Create(rows)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// Return the accounts matching the given predicates.
// predicates: Map of property name to value. Accounts will be filtered on the conjunction of equality predicates.
func (st *AccountStore) Read(accountType *model.AccountType, predicates map[string]interface{}) ([]*model.Account, error) {
	rows := []map[string]interface{}{}
	result := st.db.Table(accountType.Name).Where(predicates).Find(&rows)
	if result.Error != nil {
		return nil, result.Error
	}

	accounts := []*model.Account{}
	for _, row := range rows {
		converted, err := st.deserialize(accountType, row)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &model.Account{
			AccountType: accountType,
			Type:        accountType.Name,
			Data:        converted,
		})
	}

	return accounts, nil
}

// Helper function to loop through the row values and use the driver to serialize them.
func (st *AccountStore) serialize(accountType *model.AccountType, row map[string]interface{}) (map[string]interface{}, error) {
	converted := make(map[string]interface{})
	for _, propertyType := range accountType.PropertyTypes {
		if value, ok := row[propertyType.Name]; ok {
			converted[propertyType.Name] = st.driver.Serialize(propertyType.DataType, value)
		}
	}
	return converted, nil
}

// Helper function to loop through the row values and use the driver to deserialize them.
func (st *AccountStore) deserialize(accountType *model.AccountType, row map[string]interface{}) (map[string]interface{}, error) {
	converted := make(map[string]interface{})
	for _, propertyType := range accountType.PropertyTypes {
		if value, ok := row[propertyType.Name]; ok {
			converted[propertyType.Name] = st.driver.Deserialize(propertyType.DataType, value)
		}
	}
	return converted, nil
}
