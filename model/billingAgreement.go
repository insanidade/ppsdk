package model

type BillingAgreement struct {
	BAID string `json:"billing_agreement_id,omitempty"`
}

func NewDefaultBillingAgreement() *BillingAgreement {
	return &BillingAgreement{}
}
func (ba *BillingAgreement) GetBA() string {
	return ba.BAID
}
func (ba *BillingAgreement) SetBA(value string) {
	ba.BAID = value
}
