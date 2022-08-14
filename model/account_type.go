package model

type AccountType struct {
	Name          string
	PropertyTypes []*PropertyType
}

type PropertyType struct {
	Name     string
	DataType DataType
}
