package model

type MonthlyPayment struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Curreency     string `json:"currency,omitempty"`
	Value         string `json:"value,omitempty"`
}

func NewMonthlyPayment(currency string, value string) *MonthlyPayment {
	return &MonthlyPayment{
		CurreencyCode: currency,
		Value:         value,
	}
}
