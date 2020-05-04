package interfaces

type IFaceItem interface {
	GetSKU() string
	GetName() string
	GetDescription() string
	GetQuantity() int
	GetPrice() string
	GetCurrency() string
	GetTax() string

	SetSKU(sku string)
	SetName(n string)
	SetDescription(d string)
	SetQuantity(q int)
	SetPrice(p string)
	SetCurrency(c string)
	SetTax(t string)
}
