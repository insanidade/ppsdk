package interfaces

type IFaceLink interface {
	GetHref() string
	GetRel() string
	GetMethod() string

	SetHref(href string)
	SetRel(rel string)
	SetMethod(meth string)
}
