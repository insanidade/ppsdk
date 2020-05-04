package interfaces

type IfaceCreditCard interface {
	GetNumber() string
	GetType() string
	GetExpireMonth() int
	GetExpireYear() int
	GetCvv2() string
	GetFirstName() string
	GetLastName() string
	GetBillingAddress() *IFaceBillingAddress
	GetLinks() []IFaceLink

	SetNumber(n string)
	SetType(t string)
	SetExpireMonth(em int)
	SetExpireYear(ey int)
	SetCvv2(cvv2 string)
	SetFirstName(fn string)
	SetLastName(ln string)
	SetBillingAddress(ba *IFaceBillingAddress)
	SetLinks(l []IFaceLink)
}
