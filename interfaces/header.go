package interfaces

type Header interface {
	AddCustomHeader(headerName string, headerValue string)
	GetHeaders() map[string]string
}
