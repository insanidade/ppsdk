package interfaces

type HeaderResponse interface {
	AddCustomHeader(headerName string, headerValue string)
	GetHeaders() map[string]string
}
