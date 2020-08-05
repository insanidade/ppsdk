package model

type CreditCardToken struct {
	CreditCardID string `json:"credit_card_id,omitempty"`
	PayerID      string `json:"payer_id,omitempty"`
	Last4        string `json:"last4,omitempty"`
	Type         string `json:"type,omitempty"`
	ExpireMonth  int    `json:"expire_month,omitempty"`
	ExpireYear   int    `json:"expire_year,omitempty"`
}

func NewDefaultCreditCardToken() *CreditCardToken {
	return &CreditCardToken{}
}

func (cc *CreditCardToken) GetExpireMonth() int {
	return cc.ExpireMonth
}

func (cc *CreditCardToken) SetExpireMonth(value int) {
	cc.ExpireMonth = value
}

func (cc *CreditCardToken) GetExpireYear() int {
	return cc.ExpireYear
}

func (cc *CreditCardToken) SetExpireYear(value int) {
	cc.ExpireYear = value
}
func (fi *CreditCardToken) SetType(value string) {
	fi.Type = value
}
func (fi *CreditCardToken) GetType() string {
	return fi.Type
}

func (fi *CreditCardToken) SetLast4(value string) {
	fi.Last4 = value
}
func (fi *CreditCardToken) GetLast4() string {
	return fi.Last4
}

func (fi *CreditCardToken) SetPayerID(value string) {
	fi.PayerID = value
}
func (fi *CreditCardToken) GetPayerID() string {
	return fi.PayerID
}

func (fi *CreditCardToken) SetCreditCardID(value string) {
	fi.CreditCardID = value
}
func (fi *CreditCardToken) GetCreditCardID() string {
	return fi.CreditCardID
}
