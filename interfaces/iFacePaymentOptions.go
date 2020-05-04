package interfaces

type IFacePaymentOptions interface {
	GetAllowedPaymentMethod() string
	SetAllowedPaymentMethod(apm string)
}
