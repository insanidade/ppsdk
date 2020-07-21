package model

type BillingAgreement struct {
	BAID string `json:"billing_agreement_id,omitempty"`
}

func NewBillingAgreement() *BillingAgreement {
	return &BillingAgreement{}
}
