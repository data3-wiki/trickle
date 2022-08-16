package model

// Schema of a Solana program.
type ProgramType struct {
	// Account type definitions.
	AccountTypes []*AccountType
	// Account types indexed by name.
	accountTypesByName map[string]*AccountType
}

// Create a new program type with the given account types.
func NewProgramType(accountTypes []*AccountType) *ProgramType {
	accountTypesByName := make(map[string]*AccountType)
	for _, acc := range accountTypes {
		accountTypesByName[acc.Name] = acc
	}
	return &ProgramType{
		AccountTypes:       accountTypes,
		accountTypesByName: accountTypesByName,
	}
}

// Return the account type with the given name, if it exists.
// Defensively return a bool indicating if an account type was found (reminds the caller to check).
func (prog *ProgramType) AccountType(name string) (*AccountType, bool) {
	accountType, ok := prog.accountTypesByName[name]
	return accountType, ok
}
