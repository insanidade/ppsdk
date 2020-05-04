package interfaces

type IFaceItemList interface {
	GetItems() []IFaceItem
	GetShippingAddress() *IFaceShippingAddress
	GetShippingPhoneNumber() string

	SetItems(i []IFaceItem)
	SetShippingAddress(sa *IFaceShippingAddress)
	SetShippingPhoneNumber(spn string)
}
