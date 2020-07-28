package model

//Order represents an order json object
type Plan struct {
	Type                string               `json:"type,omitempty"`
	MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
}

func NewPlan() *Plan {
	return &Plan{
		Type:                "MERCHANT_INITIATED_BILLING",
		MerchantPreferences: NewMerchantPreferences()}
}

func (pl *Plan) SetType(baType string) {
	pl.Type = baType
}

func (pl *Plan) GetType() string {
	return pl.Type
}

func (pl *Plan) SetMerchantPreferences(value *MerchantPreferences) {
	pl.MerchantPreferences = value
}

func (pl *Plan) GetMerchantPreferences() *MerchantPreferences {
	return pl.MerchantPreferences
}
