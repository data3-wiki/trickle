package model

type AccountType struct {
	Name                string
	PropertyTypes       []*PropertyType
	propertyTypesByName map[string]*PropertyType
}

func NewAccountType(name string, propertyTypes []*PropertyType) *AccountType {
	propertyTypesByName := make(map[string]*PropertyType)
	for _, pt := range propertyTypes {
		propertyTypesByName[pt.Name] = pt
	}
	return &AccountType{
		Name:                name,
		PropertyTypes:       propertyTypes,
		propertyTypesByName: propertyTypesByName,
	}
}

func (at *AccountType) PropertyType(name string) (*PropertyType, bool) {
	pt, ok := at.propertyTypesByName[name]
	return pt, ok
}

type PropertyType struct {
	Name     string
	DataType DataType
}
