package model

type TotalInterest struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Value         string `json:"value,omitempty"`
}

func NewTotalInterest(currency string, value string) *TotalInterest {
	return &TotalInterest{
		CurreencyCode: currency,
		Value:         value,
	}
}
