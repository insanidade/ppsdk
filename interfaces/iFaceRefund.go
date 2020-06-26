package interfaces

type IFaceRefund interface {
	GetID() string
	GetAmount() *IFaceAmount
	GetState() string
	GetInvoiceNumber() string
	GetReason() string
	GetSaleID() string
	GetCaptureID() string
	GetParentPayment() string
	GetDescription() string
	GetCreateTime() string
	GetUpdateTime() string
	GetLinks() []IFaceLink
}
