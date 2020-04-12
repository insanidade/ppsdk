package model

type CreditCard struct {
	Number         string          `json:"number,omitempty"`
	Type           string          `json:"type,omitempty"`
	ExpireMonth    int             `json:"expire_month,omitempty"`
	ExpireYear     int             `json:"expire_year,omitempty"`
	Cvv2           string          `json:"cvv2,omitempty"`
	FirstName      string          `json:"first_name,omitempty"`
	LastName       string          `json:"last_name,omitempty"`
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
	Links          []Link          `json:"links,omitempty"`
}

func (cc *CreditCard) SetBillingAddress(ba *BillingAddress) {
	cc.BillingAddress = ba
}

func (cc *CreditCard) AddLink(l *Link) {
	cc.Links = append(cc.Links, *l)
}
