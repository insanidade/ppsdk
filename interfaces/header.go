package interfaces

type Header interface {
	AddCustomHeader(headerName string, headerValue string)
	GetHeaders() map[string]string
	SetNegativeTest(negativeTestValue string)
	RemoveNegativeTest()
	SetBearerToken(token string)
}
