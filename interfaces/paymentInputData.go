package interfaces

type PaymentInputData interface {
	GetIntent() string
	GetNoteToPayer() string
	GetApplicationContext() *IfaceApplicationContext
	GetPayer() *IFacePayer
	GetRedirectURLS() *IFaceRedirectURLS
	GetTransactions() []IFaceTransaction

	SetIntent(intent string)
	SetNoteToPayer(note string)
	SetApplicationContext(appContext *IfaceApplicationContext)
	SetPayer(payer *IFacePayer)
	SetRedirectURLS(urls *IFaceRedirectURLS)
	SetTransactions(transactions []IFaceTransaction)
}
