package model

//Order represents an order json object
type Plan struct {
	Type                string               `json:"type,omitempty"`
	MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
}
