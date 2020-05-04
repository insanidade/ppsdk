package interfaces

type IFacePayerInfo interface {
	GetEmail() string
	GetSalutation() string
	GetFirstName() string
	GetMiddleName() string
	GetLastName() string
	GetSuffix() string
	GetPayerID() string
	GetBirthDate() string
	GetTaxID() string
	GetTaxIDType() string
	GetBillingAddress() *IFaceBillingAddress
	GetShippingAddress() *IFaceShippingAddress

	SetEmail(email string)
	SetSalutation(salut string)
	SetFirstName(fname string)
	SetMiddleName(midname string)
	SetLastName(lname string)
	SetSuffix(suffix string)
	SetPayerID(payerID string)
	SetBirthDate(bdate string)
	SetTaxID(taxID string)
	SetTaxIDType(taxIDType string)
	SetBillingAddress(ba *IFaceBillingAddress)
	SetShippingAddress(sa *IFaceShippingAddress)
}
