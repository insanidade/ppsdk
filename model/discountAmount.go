package model

type DiscountAmount struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Curreency     string `json:"currency,omitempty"`
	Value         string `json:"value,omitempty"`
}

func NewDiscountAmount(currency string, value string) *DiscountAmount {
	return &DiscountAmount{
		CurreencyCode: currency,
		Value:         value,
	}
}
