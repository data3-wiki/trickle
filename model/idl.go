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

func FromIDL(idlJson []byte) (*ProgramType, error) {
	var idl IDL
	err := json.Unmarshal(idlJson, &idl)
	if err != nil {
		return nil, err
	}

	accountTypes := []*AccountType{}
	for _, idlAccount := range idl.Accounts {
		properties := []*PropertyType{}
		for _, field := range idlAccount.Type.Fields {
			dataType, err := toDataType(field.Type)
			if err != nil {
				return nil, err
			}
			properties = append(properties, &PropertyType{
				Name:     field.Name,
				DataType: dataType,
			})
		}
		accountTypes = append(accountTypes, &AccountType{
			Name:          idlAccount.Name,
			PropertyTypes: properties,
		})
	}

	return NewProgramType(accountTypes), nil
}

func toDataType(rawFieldType interface{}) (DataType, error) {
	var fieldType string
	switch t := rawFieldType.(type) {
	case string:
		fieldType = t
	default:
		return Text, nil
	}

	switch fieldType {
	// TODO: Support more types.
	default:
		return Text, nil
	}
}
