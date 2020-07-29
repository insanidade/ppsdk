package model

type MonthlyPayment struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Curreency     string `json:"currency,omitempty"`
	Value         string `json:"value,omitempty"`
}

func (mp *MonthlyPayment) SetValue(value string) {
	mp.Value = value
}
func (mp *MonthlyPayment) GetValue() string {
	return mp.Value
}

func (mp *MonthlyPayment) SetCurrency(value string) {
	mp.Curreency = value
}
func (mp *MonthlyPayment) GetCurrency() string {
	return mp.Curreency
}

func (mp *MonthlyPayment) SetCurrencyCode(value string) {
	mp.CurreencyCode = value
}
func (mp *MonthlyPayment) GetCurrencyCode() string {
	return mp.CurreencyCode
}

func NewMonthlyPayment() *MonthlyPayment {
	return &MonthlyPayment{}
}

// func NewMonthlyPayment(currency string, value string) *MonthlyPayment {
// 	return &MonthlyPayment{
// 		CurreencyCode: currency,
// 		Value:         value,
// 	}
// }
