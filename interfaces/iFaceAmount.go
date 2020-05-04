package interfaces

type IFaceAmount interface {
	GetCurrency() string
	GetTotal() string
	GetDetails() *IFaceDetails

	SetCurrency(c string)
	SetTotal(t string)
	SetDetails(d *IFaceDetails)
}
