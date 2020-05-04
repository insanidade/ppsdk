package interfaces

type IFaceTransaction interface {
	GetDescription() string
	GetNoteToPayee() string
	GetCustom() string
	GetInvoiceNumber() string
	GetSoftDescriptor() string
	GetNotifyURL() string
	GetAmount() *IFaceAmount
	GetPayee() *IFacePayee
	GetPaymentOptions() *IFacePaymentOptions
	GetItemList() *IFaceItemList
	GetRelatedResources() []IFaceRelatedResource

	SetDescription(desc string)
	SetNoteToPayee(note string)
	SetCustom(custom string)
	SetInvoiceNumber(in string)
	SetSoftDescriptor(sdesc string)
	SetNotifyURL(nurl string)
	SetAmount(amount *IFaceAmount)
	SetPayee(payee *IFacePayee)
	SetPaymentOptions(po *IFacePaymentOptions)
	SetItemList(il *IFaceItemList)
	SetRelatedResources(rr []IFaceRelatedResource)
}
