package interfaces

type ResponseContainer interface {
	GetHeader() Header
	GetBody() ResponseBodyRoot
	SetHeader(header Header)
	SetBody(body ResponseBodyRoot)
	GetStatus() string
	SetStatus(status string)
	GetCode() int
	SetCode(code int)
}
