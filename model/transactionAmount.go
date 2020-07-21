package model

//Amount defines an amount json object
type TransactionAmount struct {
	Currency string `json:"currency_code,omitempty"`
	Value    string `json:"value,omitempty"`
}

//NewAmount returns an instance of an amount json object
func NewTransactionAmount(currency string, value string) *TransactionAmount {
	return &TransactionAmount{
		Currency: currency,
		Value:    value,
	}
}
