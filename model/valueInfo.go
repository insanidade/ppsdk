package model

type ValueInfo struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

func NewDefaultValueInfo() *ValueInfo {
	return &ValueInfo{}
}
func (vi *ValueInfo) GetValue() string {
	return vi.Value
}
func (vi *ValueInfo) SetValue(value string) {
	vi.Value = value
}
func (vi *ValueInfo) GetCurrency() string {
	return vi.Currency
}
func (vi *ValueInfo) SetCurrency(value string) {
	vi.Currency = value
}
