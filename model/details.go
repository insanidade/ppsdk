package model

//// TODO: create setters for all attributes

//Details defines a details json object
type Details struct {
	Subtotal         string `json:"subtotal,omitempty"`
	Shipping         string `json:"shipping,omitempty"`
	Tax              string `json:"tax,omitempty"`
	HandlingFee      string `json:"handling_fee,omitempty"`
	ShippingDiscount string `json:"shipping_discount,omitempty"`
	Insurance        string `json:"insurance,omitempty"`
	GiftWrap         string `json:"gift_wrap,omitempty"`
}

//NewDetails returns a new instance of a details json object
func NewDetails(subtotal string, shipping string) *Details {
	return &Details{
		Subtotal: subtotal,
		Shipping: shipping,
	}
}
