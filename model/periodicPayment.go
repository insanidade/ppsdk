package model

type PeriodicPayment struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Value         string `json:"value,omitempty"`
}

func NewPeriodicPayment(currency string, value string) *PeriodicPayment {
	return &PeriodicPayment{
		CurreencyCode: currency,
		Value:         value,
	}
}
