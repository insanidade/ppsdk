package interfaces

type IFaceDetails interface {
	GetSubtotal() string
	GetShipping() string
	GetTax() string
	GetHandlingFee() string
	GetShippingDiscount() string
	GetInsurance() string
	GetGiftWrap() string

	SetSubtotal(s string)
	SetShipping(sh string)
	SetTax(tax string)
	SetHandlingFee(hf string)
	SetShippingDiscount(sd string)
	SetInsurance(in string)
	SetGiftWrap(gw string)
}
