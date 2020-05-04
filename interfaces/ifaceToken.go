package interfaces

type IfaceToken interface {
	GetId() string
	GetType() string

	SetId(id string)
	SetType(tokenType string)
}
