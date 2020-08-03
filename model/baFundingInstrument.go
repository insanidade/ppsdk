package model

type BAFundingInstrument struct {
	Type string            `json:"type,omitempty"`
	BA   *BillingAgreement `json:"billing_agreement,omitempty"`
}

func NewDefaultBAFundingInstrument() *BAFundingInstrument {
	return &BAFundingInstrument{
		Type: "BILLING_AGREEMENT",
		BA:   NewDefaultBillingAgreement(),
	}
}
func (bfi *BAFundingInstrument) GetBA() *BillingAgreement {
	return bfi.BA
}
func (bfi *BAFundingInstrument) SetBA(value *BillingAgreement) {
	bfi.BA = value
}

func (bfi *BAFundingInstrument) GetType() string {
	return bfi.Type
}

func (bfi *BAFundingInstrument) SetType(value string) {
	bfi.Type = value
}
