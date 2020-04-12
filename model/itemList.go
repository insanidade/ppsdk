package model

type ItemList struct {
	Items               []Item           `json:"items,omitempty"`
	ShippingAddress     *ShippingAddress `json:"shipping_address,omitempty"`
	ShippingPhoneNumber string           `json:"shipping_phone_number,omitempty"`
}

func NewItemList(shippingPhonenumber string) *ItemList {
	var emptyItemSlice []Item
	return &ItemList{ShippingPhoneNumber: shippingPhonenumber,
		Items: emptyItemSlice}
}

func (il *ItemList) AddItem(i *Item) {
	il.Items = append(il.Items, *i)
}

func (il *ItemList) SetShippingAddress(sa *ShippingAddress) {
	il.ShippingAddress = sa
}
