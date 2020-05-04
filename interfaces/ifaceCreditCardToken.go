package interfaces

type IfaceCreditCardToken interface {
	GetCreditCardID() string
	GetPayerID() string
	GetLast4() string
	GetType() string
	GetExpireMonth() int
	GetExpireYear() int

	SetCreditCardID(ccID string)
	SetPayerID(payerID string)
	SetLast4(last4 string)
	SetType(tokenType string)
	SetExpireMonth(em int)
	SetExpireYear(ey int)
}
