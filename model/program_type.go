package model

type ProgramType struct {
	AccountTypes       []*AccountType
	accountTypesByName map[string]*AccountType
}

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

func (prog *ProgramType) AccountType(name string) (*AccountType, bool) {
	accountType, ok := prog.accountTypesByName[name]
	return accountType, ok
}
