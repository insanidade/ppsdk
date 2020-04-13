package interfaces

type RequestContainer interface {
	GetHeader() Header
	GetBody() BodyRoot
	GetMethod() string
	GetURL() string
	SetHeader(header Header)
	SetBody(body BodyRoot)
	SetMethod(string)
	SetURL(string)
}
