package model

//Refund represents a sale json object
type Refund struct {
	ID            string  `json:"id,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
	State         string  `json:"state,omitempty"`
	Reason        string  `json:"reason,omitempty"`
	InvoiceNumber string  `json:"invoice_number,omitempty"`
	SaleID        string  `json:"sale_id,omitempty"`
	CaptureID     string  `json:"capture_id,omitempty"`
	ParentPayment string  `json:"parent_payment,omitempty"`
	Description   string  `json:"description,omitempty"`
	CreateTime    string  `json:"create_time,omitempty"`
	UpdateTime    string  `json:"update_time,omitempty"`
	Links         []Link  `json:"links,omitempty"`
}
