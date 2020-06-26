package interfaces

//IFaceCapture interface
type IFaceCapture interface {
	GetID() string
	GetAmount() *IFaceAmount
	GetIsFinalCapture() bool
	GetState() string
	GetTransactionFee() *IFaceValueInfo
	GetReasonCode() string
	GetInvoiceNumber() string
	GetReceivableAmount() *IFaceValueInfo
	GetExchangeRate() string
	GetNoteToPayer() string
	GetCreateTime() string
	GetParentPayment() string
	GetUpdateTime() string
	GetLinks() []IFaceLink
}
