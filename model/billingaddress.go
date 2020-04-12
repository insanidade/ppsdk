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
