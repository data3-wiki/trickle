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
			properties = append(properties, &PropertyType{
				Name:     field.Name,
				DataType: toDataType(field.Type),
			})
		}
		accountTypes = append(accountTypes, NewAccountType(idlAccount.Name, properties))
	}

	return NewProgramType(accountTypes), nil
}

func toDataType(rawFieldType interface{}) DataType {
	var fieldType string
	switch t := rawFieldType.(type) {
	case string:
		fieldType = t
	default:
		return TextDataType{}
	}

	switch fieldType {
	case "u8", "u16", "u32":
		return IntegerDataType{}
	case "f32", "f64":
		return RealDataType{}
	case "bool":
		return BooleanDataType{}
	default:
		return TextDataType{}
	}
}
