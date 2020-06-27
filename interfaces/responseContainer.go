package interfaces

type ResponseContainer interface {
	GetHeader() HeaderResponse
	GetBody() ResponseBodyRoot
	SetHeader(header HeaderResponse)
	SetBody(body ResponseBodyRoot)
	GetStatus() string
	SetStatus(status string)
	GetCode() int
	SetCode(code int)
}
