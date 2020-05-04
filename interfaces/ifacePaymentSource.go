package interfaces

type IfacePaymentSource interface{
GetToken() IfaceToken
SetToken(token IfaceToken)
}
