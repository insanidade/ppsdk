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

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}

func (cc *CreditCard) GetLinks() []Link {
	return cc.Links
}

func (cc *CreditCard) GetBillingAddress() *BillingAddress {
	return cc.BillingAddress
}

func (cc *CreditCard) SetLastName(value string) {
	cc.LastName = value
}
func (cc *CreditCard) GetLastName() string {
	return cc.LastName
}

func (cc *CreditCard) SetFirstName(value string) {
	cc.FirstName = value
}

func (cc *CreditCard) GetFirstName() string {
	return cc.FirstName
}

func (cc *CreditCard) SetBillingAddress(ba *BillingAddress) {
	cc.BillingAddress = ba
}

func (cc *CreditCard) AddLink(l *Link) {
	cc.Links = append(cc.Links, *l)
}

func (cc *CreditCard) GtNumber() string {
	return cc.Number
}

func (cc *CreditCard) SetNumber(value string) {
	cc.Number = value
}

func (cc *CreditCard) GetType() string {
	return cc.Type
}

func (cc *CreditCard) SetType(value string) {
	cc.Type = value
}

func (cc *CreditCard) GetExpireMonth() int {
	return cc.ExpireMonth
}

func (cc *CreditCard) SetExpireMonth(value int) {
	cc.ExpireMonth = value
}

func (cc *CreditCard) GetExpireYear() int {
	return cc.ExpireYear
}

func (cc *CreditCard) SetExpireYear(value int) {
	cc.ExpireYear = value
}

func (cc *CreditCard) GetCVV2() string {
	return cc.Cvv2
}

func (cc *CreditCard) SetCVV2(value string) {
	cc.Cvv2 = value
}
