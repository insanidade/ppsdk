package interfaces

type IfaceFundingInstrument interface {
	GetCreditCard() *IfaceCreditCard
	GetCreditCardToken() *IfaceCreditCardToken
	SetCreditCard(cc *IfaceCreditCard)
	SetCreditCardToken(ccToken *IfaceCreditCardToken)
}
