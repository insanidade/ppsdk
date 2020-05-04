package interfaces

type IFaceRedirectURLS interface {
	GetReturnURL() string
	GetCancelURL() string

	SetReturnURL(r string)
	SetCancelURL(c string)
}
