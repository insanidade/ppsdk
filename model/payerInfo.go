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

func (pi PayerInfo) SetTaxIDType(value string) {
	pi.TaxIDType = value
}
func (pi PayerInfo) GetTaxIDType() string {
	return pi.TaxIDType
}
func (pi PayerInfo) SetTaxID(value string) {
	pi.TaxID = value
}
func (pi PayerInfo) GetTaxID() string {
	return pi.TaxID
}

func (pi PayerInfo) SetBirthDate(value string) {
	pi.BirthDate = value
}
func (pi PayerInfo) GetBirthDate() string {
	return pi.BirthDate
}

func (pi PayerInfo) SetPayerID(value string) {
	pi.PayerID = value
}
func (pi PayerInfo) GetPayerID() string {
	return pi.PayerID
}

func (pi PayerInfo) SetSuffix(value string) {
	pi.Suffix = value
}
func (pi PayerInfo) GetSuffix() string {
	return pi.Suffix
}
func (pi PayerInfo) SetLastName(value string) {
	pi.LastName = value
}
func (pi PayerInfo) GetLastName() string {
	return pi.LastName
}

func (pi PayerInfo) SetMiddleName(value string) {
	pi.MiddleName = value
}
func (pi PayerInfo) GetMiddleName() string {
	return pi.MiddleName
}
func (pi PayerInfo) SetFirstName(value string) {
	pi.FirstName = value
}
func (pi PayerInfo) GetFirstName() string {
	return pi.FirstName
}
func (pi PayerInfo) SetSalutation(value string) {
	pi.Salutation = value
}
func (pi PayerInfo) GetSalutation() string {
	return pi.Salutation
}
func (pi PayerInfo) SetEmail(value string) {
	pi.Email = value
}

func (pi PayerInfo) GetEmail() string {
	return pi.Email
}

func (pi PayerInfo) GetBillingAddress() *BillingAddress {
	return pi.BillingAddress
}

func (pi *PayerInfo) SetBillingAddress(ba *BillingAddress) {
	pi.BillingAddress = ba
}

func (pi PayerInfo) GetShippingAddress() *ShippingAddress {
	return pi.ShippingAddress
}

func (pi PayerInfo) SetShippingAddress(sa *ShippingAddress) {
	pi.ShippingAddress = sa
}

func NewDefaultPayerInfo() *PayerInfo {
	return &PayerInfo{}
}
