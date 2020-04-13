package interfaces

type Builder interface {
	BuildRequestContainer() RequestContainer
	BuildResponseContainer() ResponseContainer

	// BodyFactory() BodyRoot
	// HeaderFactory() Header
	// Assemble(h Header, br BodyRoot)
	// DoRequest() BodyRoot
}
