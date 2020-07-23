package model

//MerchantPreferences represents a merchant_preferences json object
type MerchantPreferences struct {
	CancelUrl                             string        `json:"cancel_url,omitempty"`
	ReturnUrl                             string        `json:"return_url,omitempty"`
	NotifyUrl                             string        `json:"notify_url,omitempty"`
	AcceptedPymtType                      string        `json:"accepted_pymt_type,omitempty"`
	ImmutableShippingAddress              bool          `json:"immutable_shipping_address,omitempty"`
	ExperienceId                          string        `json:"experience_id,omitempty"`
	ExternalSelectedFundingInstrumentType string        `json:"external_selected_funding_instrument_type,omitempty"`
	AcceptedLegalCountryCodes             []CountryCode `json:"accepted_legal_country_codes,omitempty"`
	SkipShippingAddress                   bool          `json:"skip_shipping_address,omitempty"`
}
