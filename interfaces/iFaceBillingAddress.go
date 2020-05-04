package interfaces

type IFaceBillingAddress interface {
	GetLine1() string
	GetLine2() string
	GetCity() string
	GetCountryCode() string
	GetPostalCode() string
	GetState() string
	GetPhone() string
	GetNormalizationStatus() string
	GetType() string

	SetLine1(line1 string)
	SetLine2(line2 string)
	SetCity(city string)
	SetCountryCode(cc string)
	SetPostalCode(pc string)
	SetState(state string)
	SetPhone(phone string)
	SetNormalizationStatus(ns string)
	SetType(t string)
}
