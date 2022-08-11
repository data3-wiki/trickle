package model

import (
	"encoding/json"
)

type IDL struct {
	Accounts []IDLAccount `json:"accounts"`
}

type IDLAccount struct {
	Name string         `json:"name"`
	Type IDLAccountType `json:"type"`
}

type IDLAccountType struct {
	Kind   string            `json:"struct"`
	Fields []IDLAccountField `json:"fields"`
}

type IDLAccountField struct {
	Name string      `json:"name"`
	Type interface{} `json:"type"`
}

func FromIDL(idlJson []byte) ([]*AccountType, error) {
	var idl IDL
	err := json.Unmarshal(idlJson, &idl)
	if err != nil {
		return nil, err
	}

	accountTypes := []*AccountType{}
	for _, idlAccount := range idl.Accounts {
		properties := []*PropertyType{}
		for _, field := range idlAccount.Type.Fields {
			var dataType string
			switch t := field.Type.(type) {
			case string:
				dataType = t
			default:
				bytes, err := json.Marshal(t)
				if err != nil {
					return nil, err
				}
				dataType = string(bytes)
			}
			properties = append(properties, &PropertyType{
				Name:     field.Name,
				DataType: DataType(dataType),
			})
		}
		accountTypes = append(accountTypes, &AccountType{
			Name:       idlAccount.Name,
			Properties: properties,
		})
	}

	return accountTypes, nil
}
