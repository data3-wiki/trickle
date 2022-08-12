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

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(dialector gorm.Dialector) (*AccountStore, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &AccountStore{
		db: db,
	}, nil
}

func createDynamicStructs(accountTypes []*model.AccountType) ([]*accountInstance, error) {
	instances := []*accountInstance{}
	for _, acc := range accountTypes {
		builder := dynamicstruct.NewStruct()
		for _, prop := range acc.PropertyTypes {
			zeroValue, err := prop.DataType.ZeroValue()
			if err != nil {
				return nil, err
			}
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
	return instances, nil
}

func (st *AccountStore) AutoMigrate(programType *model.ProgramType) error {
	dynamicStructs, err := createDynamicStructs(programType.AccountTypes)
	if err != nil {
		return err
	}

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

func (st *AccountStore) Read(accountType string, predicates map[string]interface{}) ([]*model.Account, error) {
	rows := []map[string]interface{}{}
	result := st.db.Table(accountType).Where(predicates).Find(&rows)
	if result.Error != nil {
		return nil, result.Error
	}

	accounts := []*model.Account{}
	for _, row := range rows {
		accounts = append(accounts, &model.Account{
			Type: accountType,
			Data: row,
		})
	}

	return accounts, nil
}
