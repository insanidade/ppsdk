package model

//ApplicationContext struct
type ApplicationContext struct {
	BrandName   string `json:"brand_name,omitempty,omitempty"`
	LandingPage string `json:"landing_page,omitempty"`
	Locale      string `json:"locale,omitempty"`
	//Possible values for ShippingPreference are
	//NO_SHIPPING
	//GET_FROM_FILE
	//SET_PROVIDED_ADDRESS
	ShippingPreference string `json:"shipping_preference,omitempty"`
	//Possible values for UserAction are
	//commit
	//continue
	UserAction             string         `json:"user_action,omitempty"`
	PreferredPaymentSource *PaymentSource `json:"preferred_payment_source,omitempty"`
}

//NewApplicationContext constructor
func NewApplicationContext(brandName string,
	landingPage string,
	locale string,
	shippingPref string,
	userAction string) *ApplicationContext {

	return &ApplicationContext{
		BrandName:          brandName,
		LandingPage:        landingPage,
		Locale:             locale,
		ShippingPreference: shippingPref,
		UserAction:         userAction,
	}

}

func NewDefaultApplicationContext() *ApplicationContext {

	return &ApplicationContext{
		PreferredPaymentSource: NewDefaultPaymentSource(),
	}

}

func (ac *ApplicationContext) SetPreferredPaymentSource(pps *PaymentSource) {
	ac.PreferredPaymentSource = pps
}
