package interfaces

type IFaceValueInfo interface {
	GetCurrency() string
	GetValue() string

	SetCurrency(c string)
	SetValue(v string)
}
