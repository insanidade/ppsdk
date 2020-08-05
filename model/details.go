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

func NewDefaultDetails() *Details {
	return &Details{}
}

//NewDetails returns a new instance of a details json object
func NewDetails(subtotal string, shipping string) *Details {
	return &Details{
		Subtotal: subtotal,
		Shipping: shipping,
	}
}
func (d *Details) GetGiftWrap() string {
	return d.GiftWrap
}
func (d *Details) SetGiftWrap(value string) {
	d.GiftWrap = value
}

func (d *Details) GetInsurance() string {
	return d.Insurance
}
func (d *Details) SetInsurance(value string) {
	d.Insurance = value
}

func (d *Details) GetShippingDiscount() string {
	return d.ShippingDiscount
}
func (d *Details) SetShippingDiscount(value string) {
	d.ShippingDiscount = value
}
func (d *Details) GetHandlingFee() string {
	return d.HandlingFee
}

func (d *Details) SetHandlingFee(value string) {
	d.HandlingFee = value
}
func (d *Details) GetTax() string {
	return d.Tax
}

func (d *Details) SetTax(value string) {
	d.Tax = value
}
func (d *Details) GetShipping() string {
	return d.Shipping
}
func (d *Details) SetShipping(value string) {
	d.Shipping = value
}
func (d *Details) GetSubtotal() string {
	return d.Subtotal
}
func (d *Details) SetSubtotal(value string) {
	d.Subtotal = value
}
