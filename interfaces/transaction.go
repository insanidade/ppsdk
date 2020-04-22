package interfaces

//Transaction is the interface that represents a transaction in the session.
//It provides services related to the transaction status, headers and other
//important information a user might need.
type Transaction interface {
	// BuildRequestContainer() RequestContainer
	// BuildResponseContainer() ResponseContainer

	GetRequestBody() BodyRoot
	GetResponseBody() BodyRoot
	GetRequestHeader() Header
	GetResponseHeader() Header

	// BodyFactory() BodyRoot
	// HeaderFactory() Header
	// Assemble(h Header, br BodyRoot)
	// DoRequest() BodyRoot
}
