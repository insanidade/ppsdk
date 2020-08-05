package model

//Amount defines an amount json object
type Amount struct {
	Currency string   `json:"currency,omitempty"`
	Total    string   `json:"total,omitempty"`
	Details  *Details `json:"details,omitempty"`
}

func NewDefaultAmount() *Amount {
	return &Amount{
		Details: NewDefaultDetails()}
}

//NewAmount returns an instance of an amount json object
func NewAmount(currency string, total string, details *Details) *Amount {
	return &Amount{
		Currency: currency,
		Total:    total,
		Details:  details}
}

func (a *Amount) GetTotal() string {
	return a.Total
}

func (a *Amount) SetTotal(value string) {
	a.Total = value
}

func (a *Amount) GetCurrency() string {
	return a.Currency
}

func (a *Amount) SetCurrency(value string) {
	a.Currency = value
}

func (a *Amount) GetDetails() *Details {
	return a.Details
}

//SetDetails sets a details json object for this
//amount object
func (a *Amount) SetDetails(d *Details) {
	a.Details = d
}
