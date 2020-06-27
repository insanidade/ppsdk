package interfaces

//Transaction is the interface that represents a transaction in the session.
//It provides services related to the transaction status, headers and other
//important information a user might need.
type Transaction interface {
	// BuildRequestContainer() RequestContainer
	// BuildResponseContainer() ResponseContainer

	GetRequestBody() BodyRoot
	GetResponseBody() ResponseBodyRoot
	GetRequestHeader() Header
	GetResponseHeader() HeaderResponse
	// GetStatus() string
	GetRequestMethod() string
	GetRequestURL() string
	GetResponseStatus() string
	GetResponseCode() int

	SetRequestBody(BodyRoot)
	SetResponseBody(ResponseBodyRoot)
	SetRequestHeader(Header)
	SetResponseHeader(HeaderResponse)
	SetResponseCode(code int)
	SetResponseStatus(status string)
	// SetStatus(status string)

	// BodyFactory() BodyRoot
	// HeaderFactory() Header
	// Assemble(h Header, br BodyRoot)
	// DoRequest() BodyRoot
}
