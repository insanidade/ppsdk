package interfaces

type IfaceApplicationContext interface {
	GetBrandName() string
	SetBrandName(bn string)

	GetLandingPage() string
	SetLandingPage(lp string)

	GetLocale() string
	SetLocale(locale string)

	GetPreferredPaymentSource() *IfacePaymentSource
	SetPreferredPaymentSource(ps *IfacePaymentSource)

	//Possible values for ShippingPreference are
	//NO_SHIPPING
	//GET_FROM_FILE
	//SET_PROVIDED_ADDRESS
	GetShippingPreference() string
	SetShippingPreference(sp string)
	//Possible values for UserAction are
	//commit
	//continue
	GetUserAction() string
	SetUserAction(ua string)
}
