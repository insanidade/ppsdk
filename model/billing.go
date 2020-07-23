package model

type Billing struct {
	BillingAgreementID        string                     `json:"billing_agreement_id,omitempty"`
	SelectedInstallmentOption *SelectedInstallmentOption `json:"selected_installment_option,omitempty"`
}
