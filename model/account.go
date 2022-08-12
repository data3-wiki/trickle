package model

// TODO: Add slot.
type Account struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}
