package model

type BAFundingInstrument struct {
	Type string            `json:"type,omitempty"`
	BA   *BillingAgreement `json:"billing_agreement,omitempty"`
}

func NewBAFundingInstrument() *BAFundingInstrument {
	return &BAFundingInstrument{}
}
