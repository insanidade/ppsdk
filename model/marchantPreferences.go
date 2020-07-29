package model

//MerchantPreferences represents a merchant_preferences json object
type MerchantPreferences struct {
	SkipShippingAddress                   bool          `json:"skip_shipping_address"`
	ImmutableShippingAddress              bool          `json:"immutable_shipping_address"`
	CancelUrl                             string        `json:"cancel_url,omitempty"`
	ReturnUrl                             string        `json:"return_url,omitempty"`
	NotifyUrl                             string        `json:"notify_url,omitempty"`
	AcceptedPymtType                      string        `json:"accepted_pymt_type,omitempty"`
	ExperienceId                          string        `json:"experience_id,omitempty"`
	ExternalSelectedFundingInstrumentType string        `json:"external_selected_funding_instrument_type,omitempty"`
	AcceptedLegalCountryCodes             []CountryCode `json:"accepted_legal_country_codes,omitempty"`
}

func NewMerchantPreferences() *MerchantPreferences {
	return &MerchantPreferences{
		SkipShippingAddress: true,
		// ImmutableShippingAddress: true
	}
}

func (mp *MerchantPreferences) AddAcceptedLegalCountryCode(value *CountryCode) {
	mp.AcceptedLegalCountryCodes = append(mp.AcceptedLegalCountryCodes, *value)
}

func (mp *MerchantPreferences) GetAcceptedLegalCountryCode() []CountryCode {
	return mp.AcceptedLegalCountryCodes
}

func (mp *MerchantPreferences) GetSkipShippingAddress() bool {
	return mp.SkipShippingAddress
}

func (mp *MerchantPreferences) SetSkipShippingAddress(value bool) {
	mp.SkipShippingAddress = value
}

func (mp *MerchantPreferences) GetImmutableShippingAddress() bool {
	return mp.ImmutableShippingAddress
}

func (mp *MerchantPreferences) SetImmutableShippingAddress(value bool) {
	mp.ImmutableShippingAddress = value
}

func (mp *MerchantPreferences) GetCancelURL() string {
	return mp.CancelUrl
}

func (mp *MerchantPreferences) SetCancelURL(value string) {
	mp.CancelUrl = value
}

func (mp *MerchantPreferences) GetReturnURL() string {
	return mp.ReturnUrl
}

func (mp *MerchantPreferences) SetReturnURL(value string) {
	mp.ReturnUrl = value
}

func (mp *MerchantPreferences) GetNotifyURL() string {
	return mp.NotifyUrl
}

func (mp *MerchantPreferences) SetNotifyURL(value string) {
	mp.NotifyUrl = value
}

func (mp *MerchantPreferences) GetAcceptedPaymentType() string {
	return mp.AcceptedPymtType
}

func (mp *MerchantPreferences) SetAcceptedPaymentType(value string) {
	mp.AcceptedPymtType = value
}
func (mp *MerchantPreferences) GetExperienceID() string {
	return mp.ExperienceId
}
func (mp *MerchantPreferences) SetExperienceID(value string) {
	mp.ExperienceId = value
}

func (mp *MerchantPreferences) GetExternalSelectedFundingInstrumentType() string {
	return mp.ExternalSelectedFundingInstrumentType
}

func (mp *MerchantPreferences) SetExternalSelectedFundingInstrumentType(value string) {
	mp.ExternalSelectedFundingInstrumentType = value
}
