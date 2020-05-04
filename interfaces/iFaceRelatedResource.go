package interfaces

type IFaceRelatedResource interface {
	GetSale() *IFaceSale
	GetAuthorization() *IFaceAuthorization
	GetOrder() *IFaceOrder
	GetCapture() *IFaceCapture
	GetRefund() *IFaceRefund
}
