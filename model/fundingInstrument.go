package model

//FundingInstrument object
type FundingInstrument struct {
	CreditCard      *CreditCard      `json:"credit_card,omitempty"`
	CreditCardToken *CreditCardToken `json:"credit_card_token,omitempty"`
	Billing         *Billing         `json:"billing,omitempty"`
}

//SetCreditCard sets the credit card object
func (fi *FundingInstrument) SetCreditCard(cc *CreditCard) {
	fi.CreditCard = cc
}

//SetCreditCardToken sets the credit card token object
func (fi *FundingInstrument) SetCreditCardToken(cct *CreditCardToken) {
	fi.CreditCardToken = cct
}

func (fi *FundingInstrument) SetBilling(value *Billing) {
	fi.Billing = value
}

func (fi *FundingInstrument) GetCreditCard() *CreditCard {
	return fi.CreditCard
}

func (fi *FundingInstrument) GetCretidCardToken() *CreditCardToken {
	return fi.CreditCardToken
}

func (fi *FundingInstrument) GetBilling() *Billing {
	return fi.Billing
}
