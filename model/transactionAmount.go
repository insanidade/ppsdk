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

//NewAmount returns an instance of an amount json object
func NewDefaultTransactionAmount() *TransactionAmount {
	return &TransactionAmount{}
}

func (ta *TransactionAmount) GetValue() string {
	return ta.Value
}
func (ta *TransactionAmount) SetValue(value string) {
	ta.Value = value
}
func (ta *TransactionAmount) GetCurrency() string {
	return ta.Currency
}
func (ta *TransactionAmount) SetCurrency(value string) {
	ta.Currency = value
}
