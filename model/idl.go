package model

import (
	"encoding/json"
)

// Solana Anchor framework IDL definition.
type IDL struct {
	// Account definitions.
	Accounts []IDLAccount `json:"accounts"`
}

// IDL account definition.
type IDLAccount struct {
	// Type name of the account.
	Name string `json:"name"`
	// Type definition of the account.
	Type IDLAccountType `json:"type"`
}

// Type definition of an account.
type IDLAccountType struct {
	// Field definitions of the account.
	Fields []IDLAccountField `json:"fields"`
}

// Field definition of an account.
type IDLAccountField struct {
	// Name of the field.
	Name string `json:"name"`
	// Type of the field (could be string or json object).
	Type interface{} `json:"type"`
}

// Parse IDL json bytes into internally used ProgramType.
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

// Convert IDL field type to internally used DataType.
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
