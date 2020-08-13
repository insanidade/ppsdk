package model

type DiscountAmount struct {
	CurreencyCode string `json:"currency_code,omitempty"`
	Curreency     string `json:"currency,omitempty"`
	Value         string `json:"value,omitempty"`
}

func (mp *DiscountAmount) SetValue(value string) {
	mp.Value = value
}
func (mp *DiscountAmount) GetValue() string {
	return mp.Value
}

func (mp *DiscountAmount) SetCurrency(value string) {
	mp.Curreency = value
}
func (mp *DiscountAmount) GetCurrency() string {
	return mp.Curreency
}

func (mp *DiscountAmount) SetCurrencyCode(value string) {
	mp.CurreencyCode = value
}
func (mp *DiscountAmount) GetCurrencyCode() string {
	return mp.CurreencyCode
}
func NewDefaultDiscountAmount() *DiscountAmount {
	return &DiscountAmount{}
}

// func NewDiscountAmount(currency string, value string) *DiscountAmount {
// 	return &DiscountAmount{
// 		CurreencyCode: currency,
// 		Value:         value,
// 	}
// }
