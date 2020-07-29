package model

type BillingAddress struct {
	Line1               string `json:"line1,omitempty"`
	Line2               string `json:"line2,omitempty"`
	City                string `json:"city,omitempty"`
	CountryCode         string `json:"country_code,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	State               string `json:"state,omitempty"`
	Phone               string `json:"phone,omitempty"`
	NormalizationStatus string `json:"normalization_status,omitempty"`
	Type                string `json:"type,omitempty"`
}

func (ba *BillingAddress) SetType(value string) {
	ba.Type = value
}
func (ba *BillingAddress) GetType() string {
	return ba.Type
}
func (ba *BillingAddress) SetNormalizationStatus(value string) {
	ba.NormalizationStatus = value
}

func (ba *BillingAddress) GetNormalizationStatus() string {
	return ba.NormalizationStatus
}
func (ba *BillingAddress) SetPhone(value string) {
	ba.Phone = value
}
func (ba *BillingAddress) GetPhone() string {
	return ba.Phone
}

func (ba *BillingAddress) SetState(value string) {
	ba.State = value
}
func (ba *BillingAddress) GetState() string {
	return ba.State
}

func (ba *BillingAddress) SetPostalCode(value string) {
	ba.PostalCode = value
}

func (ba *BillingAddress) GetPostalCode() string {
	return ba.PostalCode
}

func (ba *BillingAddress) SetCountryCode(value string) {
	ba.CountryCode = value
}

func (ba *BillingAddress) GetCountryCode() string {
	return ba.CountryCode
}

func (ba *BillingAddress) SetCity(value string) {
	ba.City = value
}

func (ba *BillingAddress) GetCity() string {
	return ba.City
}

func (ba *BillingAddress) SetLine2(value string) {
	ba.Line2 = value
}

func (ba *BillingAddress) GetLine2() string {
	return ba.Line2
}

func (ba *BillingAddress) SetLine1(value string) {
	ba.Line1 = value
}

func (ba *BillingAddress) GetLine1() string {
	return ba.Line1
}

func NewBillingAddress() *BillingAddress {
	return &BillingAddress{}
}
