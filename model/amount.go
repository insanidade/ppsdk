package model

//Amount defines an amount json object
type Amount struct {
	Currency string   `json:"currency,omitempty"`
	Total    string   `json:"total,omitempty"`
	Details  *Details `json:"details,omitempty"`
}

//NewAmount returns an instance of an amount json object
func NewAmount(currency string, total string, details *Details) *Amount {
	return &Amount{
		Currency: currency,
		Total:    total,
		Details:  details}
}

//SetDetails sets a details json object for this
//amount object
func (a *Amount) SetDetails(d *Details) {
	a.Details = d
}
