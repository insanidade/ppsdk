package model

type PaymentSource struct {
	Token *Token `json:"token,omitempty"`
}

func NewDefaultPaymentSource() *PaymentSource {
	return &PaymentSource{}
}

func (ps *PaymentSource) SetToken(t *Token) {
	ps.Token = t
}
