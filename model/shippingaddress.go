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

func NewShippingAddress(city string,
	countryCode string,
	line1 string,
	line2 string,
	normalizationStatus string,
	postalCode string,
	recipientName string,
	state string) *ShippingAddress {
	return &ShippingAddress{
		RecipientName:       recipientName,
		Line1:               line1,
		Line2:               line2,
		City:                city,
		PostalCode:          postalCode,
		State:               state,
		CountryCode:         countryCode,
		NormalizationStatus: normalizationStatus,
	}
}
