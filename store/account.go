package store

import (
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
	return &AccountStore{db: db}, nil
}

func (st *AccountStore) AutoMigrate(accountTypes []*model.AccountType) error {
	instances := []*accountInstance{}
	for _, acc := range accountTypes {
		builder := dynamicstruct.NewStruct()
		for _, prop := range acc.Properties {
			zeroValue, err := prop.DataType.ZeroValue()
			if err != nil {
				return err
			}
			// TODO: Add appropriate struct tag.
			builder.AddField(strings.Title(prop.Name), zeroValue, "")
		}
		inst := builder.Build().New()
		instances = append(instances, &accountInstance{
			name:     acc.Name,
			instance: inst,
		})
	}

	for _, inst := range instances {
		err := st.db.Table(inst.name).AutoMigrate(inst.instance)
		if err != nil {
			return err
		}
	}

	return nil
}
