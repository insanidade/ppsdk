package model

type PayerInfo struct {
	Email           string           `json:"email,omitempty"`
	Salutation      string           `json:"salutation,omitempty"`
	FirstName       string           `json:"first_name,omitempty"`
	MiddleName      string           `json:"middle_name,omitempty"`
	LastName        string           `json:"last_name,omitempty"`
	Suffix          string           `json:"suffix,omitempty"`
	PayerID         string           `json:"payer_id,omitempty"`
	BirthDate       string           `json:"birth_date,omitempty"`
	TaxID           string           `json:"tax_id,omitempty"`
	TaxIDType       string           `json:"tax_id_type,omitempty"`
	BillingAddress  *BillingAddress  `json:"billing_address,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

func (pi *PayerInfo) SetBillingAddress(ba *BillingAddress) {
	pi.BillingAddress = ba
}

func (pi PayerInfo) SetShippingAddress(sa *ShippingAddress) {
	pi.ShippingAddress = sa
}
