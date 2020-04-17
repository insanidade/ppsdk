package interfaces

type Builder interface {
	BuildRequestContainer() RequestContainer
	BuildResponseContainer() ResponseContainer
	// GetBody() BodyRoot
	// GetHeader() Header

	// BodyFactory() BodyRoot
	// HeaderFactory() Header
	// Assemble(h Header, br BodyRoot)
	// DoRequest() BodyRoot
}
