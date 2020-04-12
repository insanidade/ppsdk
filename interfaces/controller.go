package interfaces

type Controller interface {
	BodyFactory() BodyRoot
	HeaderFactory() Header
	Assemble(h Header, br BodyRoot)
	DoRequest() BodyRoot
}
