package model

type Item struct {
	SKU         string `json:"sku,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity    int64  `json:"quantity,omitempty"`
	Price       string `json:"price,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Tax         string `json:"tax,omitempty"`
}

func NewItem(currency string,
	description string,
	name string,
	price string,
	quantity int64,
	sku string) *Item {
	return &Item{
		Currency:    currency,
		Description: description,
		Name:        name,
		Price:       price,
		Quantity:    quantity,
		SKU:         sku,
	}
}
