package model

type Billing struct {
	BillingAgreementID        string                     `json:"billing_agreement_id,omitempty"`
	SelectedInstallmentOption *SelectedInstallmentOption `json:"selected_installment_option,omitempty"`
}

func (b *Billing) SetSelectedInstallmentOption(value *SelectedInstallmentOption) {
	b.SelectedInstallmentOption = value
}

func (b *Billing) GetSelectedInstallmentOption() *SelectedInstallmentOption {
	return b.SelectedInstallmentOption
}
func (b *Billing) SetBillingAgreementID(value string) {
	b.BillingAgreementID = value
}

func (b *Billing) GetBillingAdreementID() string {
	return b.BillingAgreementID
}

func NewDefaultBilling() *Billing {
	return &Billing{}
}
