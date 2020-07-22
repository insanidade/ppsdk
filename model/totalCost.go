package model

type TotalCost struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Value         string `json:"value,omitempty"`
}

func NewTotalCost(currency string, value string) *TotalCost {
	return &TotalCost{
		CurreencyCode: currency,
		Value:         value,
	}
}
