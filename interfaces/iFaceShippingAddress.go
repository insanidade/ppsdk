package interfaces

type IFaceShippingAddress interface {
	GetLine1() string
	GetLine2() string
	GetCity() string
	GetCountryCode() string
	GetPostalCode() string
	GetState() string
	GetPhone() string
	GetNormalizationStatus() string
	GetType() string
	GetRecipientName() string

	SetLine1(l1 string)
	SetLine2(l2 string)
	SetCity(city string)
	SetCountryCode(ccode string)
	SetPostalCode(pcode string)
	SetState(state string)
	SetPhone(phone string)
	SetNormalizationStatus(normalStatus string)
	SetType(shippingType string)
	SetRecipientName(rname string)
}
