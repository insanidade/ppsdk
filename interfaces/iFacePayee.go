package interfaces

type IFacePayee interface {
	GetEmail() string
	GetMerchantID() string

	SetEmail(e string)
	SetMerchantID(mID string)
}
