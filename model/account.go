package model

// TODO: Add slot.
// Intermediate representation of account data.
type Account struct {
	// Type of the account.
	AccountType *AccountType `json:"-"`
	// Type name of the account.
	Type string `json:"type"`
	// Decoded data in the account.
	Data map[string]interface{} `json:"data"`
}
