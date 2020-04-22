package interfaces

type ResponseContainer interface {
	GetHeader() Header
	GetBody() BodyRoot
	SetHeader(header Header)
	SetBody(body BodyRoot)
	GetStatus() string
	SetStatus(status string)
	GetCode() int
	SetCode(code int)
}
