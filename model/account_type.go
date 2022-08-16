package model

// Schema of an account.
type AccountType struct {
	// Type name of the account.
	Name string
	// Definitions of the properties in the account.
	PropertyTypes []*PropertyType
	// Property types indexed by name.
	propertyTypesByName map[string]*PropertyType
}

// Schema of an account's property.
type PropertyType struct {
	// Type name of the property.
	Name string
	// Data type of the property.
	DataType DataType
}

// Create a new account type with the given name and property types.
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

// Return the property type with the given name, if it exists.
// Defensively return a bool indicating if a property type was found (reminds the caller to check).
func (at *AccountType) PropertyType(name string) (*PropertyType, bool) {
	pt, ok := at.propertyTypesByName[name]
	return pt, ok
}
