package model

//Order represents an order json object
type MerchantPreferences struct {
	CancelUrl                             string        `json:"cancel_url,omitempty"`
	ReturnUrl                             string        `json:"return_url,omitempty"`
	NotifyUrl                             string        `json:"notify_url,omitempty"`
	AcceptedPymtType                      string        `json:"accepted_pymt_type,omitempty"`
	SkipShippingAddress                   bool          `json:"skip_shipping_address,omitempty"`
	ImmutableShippingAddress              bool          `json:"immutable_shipping_address,omitempty"`
	ExperienceId                          string        `json:"experience_id,omitempty"`
	ExternalSelectedFundingInstrumentType string        `json:"external_selected_funding_instrument_type,omitempty"`
	AcceptedLegalCountryCodes             []CountryCode `json:"accepted_legal_country_codes,omitempty"`
}

// cancel_url
// return_url
// notify_url
// accepted_pymt_type
// skip_shipping_address
// immutable_shipping_address
// experience_id
// external_selected_funding_instrument_type
// accepted_legal_country_codes
