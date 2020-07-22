package model

type CreditFinancing struct {
	Product          string `json:"product,omitempty"`
	CreditProduct    string `json:"credit_product_identifier,omitempty"`
	Apr              string `json:"apr,omitempty"`
	NominalRate      string `json:"nominal_rate,omitempty"`
	Term             int    `json:"term,omitempty"`
	Interval         string `json:"interval,omitempty"`
	IntervalDuration string `json:"interval_duration,omitempty"`
	CountryCode      string `json:"country_code,omitempty"`
	CreditType       string `json:"credit_type,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
}
