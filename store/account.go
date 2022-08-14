package store

import (
	"fmt"
	"strings"

	dynamicstruct "github.com/Ompluscator/dynamic-struct"
	"github.com/dereference-xyz/trickle/model"
	"gorm.io/gorm"
)

type accountInstance struct {
	name     string
	instance interface{}
}

type Driver interface {
	Dialector() gorm.Dialector
	Serialize(dataType model.DataType, value interface{}) interface{}
	Deserialize(dataType model.DataType, value interface{}) interface{}
}

type AccountStore struct {
	db *gorm.DB
	dr Driver
}

func NewAccountStore(driver Driver) (*AccountStore, error) {
	db, err := gorm.Open(driver.Dialector(), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &AccountStore{
		db: db,
	}, nil
}

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

func (st *AccountStore) Create(accounts []*model.Account) error {
	rowsByAccountType := make(map[string][]map[string]interface{})
	for _, acc := range accounts {
		rowsByAccountType[acc.Type] = append(rowsByAccountType[acc.Type], acc.Data)
	}
	for accountType, rows := range rowsByAccountType {
		result := st.db.Table(accountType).Create(rows)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (st *AccountStore) Read(accountType *model.AccountType, predicates map[string]interface{}) ([]*model.Account, error) {
	rows := []map[string]interface{}{}
	result := st.db.Table(accountType.Name).Where(predicates).Find(&rows)
	if result.Error != nil {
		return nil, result.Error
	}

	accounts := []*model.Account{}
	for _, row := range rows {
		accounts = append(accounts, &model.Account{
			AccountType: accountType,
			Type:        accountType.Name,
			Data:        row,
		})
	}

	return accounts, nil
}
