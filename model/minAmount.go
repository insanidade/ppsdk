package model

type MinAmount struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Value         string `json:"value,omitempty"`
}

func NewMinAmount(currency string, value string) *MinAmount {
	return &MinAmount{
		CurreencyCode: currency,
		Value:         value,
	}
}
