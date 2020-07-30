package model

//// TODO: setters for all attributes. Add to the constructor
//only for the requried ones.

type ShippingAddress struct {
	Line1               string `json:"line1,omitempty"`
	Line2               string `json:"line2,omitempty"`
	City                string `json:"city,omitempty"`
	CountryCode         string `json:"country_code,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	State               string `json:"state,omitempty"`
	Phone               string `json:"phone,omitempty"`
	NormalizationStatus string `json:"normalization_status,omitempty"`
	Type                string `json:"type,omitempty"`
	RecipientName       string `json:"recipient_name,omitempty"`
}

func (ba *ShippingAddress) SetRecipientName(value string) {
	ba.RecipientName = value
}
func (ba *ShippingAddress) GetRecipientName() string {
	return ba.RecipientName
}

func (ba *ShippingAddress) SetType(value string) {
	ba.Type = value
}
func (ba *ShippingAddress) GetType() string {
	return ba.Type
}
func (ba *ShippingAddress) SetNormalizationStatus(value string) {
	ba.NormalizationStatus = value
}

func (ba *ShippingAddress) GetNormalizationStatus() string {
	return ba.NormalizationStatus
}
func (ba *ShippingAddress) SetPhone(value string) {
	ba.Phone = value
}
func (ba *ShippingAddress) GetPhone() string {
	return ba.Phone
}

func (ba *ShippingAddress) SetState(value string) {
	ba.State = value
}
func (ba *ShippingAddress) GetState() string {
	return ba.State
}

func (ba *ShippingAddress) SetPostalCode(value string) {
	ba.PostalCode = value
}

func (ba *ShippingAddress) GetPostalCode() string {
	return ba.PostalCode
}

func (ba *ShippingAddress) SetCountryCode(value string) {
	ba.CountryCode = value
}

func (ba *ShippingAddress) GetCountryCode() string {
	return ba.CountryCode
}

func (ba *ShippingAddress) SetCity(value string) {
	ba.City = value
}

func (ba *ShippingAddress) GetCity() string {
	return ba.City
}

func (ba *ShippingAddress) SetLine2(value string) {
	ba.Line2 = value
}

func (ba *ShippingAddress) GetLine2() string {
	return ba.Line2
}

func (ba *ShippingAddress) SetLine1(value string) {
	ba.Line1 = value
}

func (ba *ShippingAddress) GetLine1() string {
	return ba.Line1
}
func NewShippingAddress() *ShippingAddress {
	return &ShippingAddress{}
}

// func NewShippingAddress(city string,
// 	countryCode string,
// 	line1 string,
// 	line2 string,
// 	normalizationStatus string,
// 	postalCode string,
// 	recipientName string,
// 	state string) *ShippingAddress {
// 	return &ShippingAddress{
// 		RecipientName:       recipientName,
// 		Line1:               line1,
// 		Line2:               line2,
// 		City:                city,
// 		PostalCode:          postalCode,
// 		State:               state,
// 		CountryCode:         countryCode,
// 		NormalizationStatus: normalizationStatus,
// 	}
// }
