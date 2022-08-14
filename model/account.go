package model

// TODO: Add slot.
type Account struct {
	AccountType *AccountType           `json:"-"`
	Type        string                 `json:"type"`
	Data        map[string]interface{} `json:"data"`
}
