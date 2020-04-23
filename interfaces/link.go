package interfaces

type Link interface {
	GetHref() string
	GetRel() string
	GetMethod() string
}
