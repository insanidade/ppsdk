package model

type PaymentSource struct {
	Token *Token `json:"token,omitempty"`
}

func (ps *PaymentSource) SetToken(t *Token) {
	ps.Token = t
}
