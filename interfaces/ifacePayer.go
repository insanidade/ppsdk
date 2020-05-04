package interfaces

type IFacePayer interface {
	GetPaymentMethod() string
	GetStatus() string
	GetTFundingInstruments() []IfaceFundingInstrument
	GetPayerInfo() *IFacePayerInfo
}
