package model

type CreditCardToken struct {
	CreditCardID string `json:"credit_card_id,omitempty"`
	PayerID      string `json:"payer_id,omitempty"`
	Last4        string `json:"last4,omitempty"`
	Type         string `json:"type,omitempty"`
	ExpireMonth  int    `json:"expire_month,omitempty"`
	ExpireYear   int    `json:"expire_year,omitempty"`
}
