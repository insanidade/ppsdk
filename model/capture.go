package model

//Capture represents a capture json object
type Capture struct {
	ID               string     `json:"id,omitempty"`
	Amount           *Amount    `json:"amount,omitempty"`
	IsFinalCapture   bool       `json:"is_final_capture,omitempty"`
	State            string     `json:"state,omitempty"`
	ReasonCode       string     `json:"reason_code,omitempty"`
	InvoiceNumber    string     `json:"invoice_number,omitempty"`
	TransactionFee   *ValueInfo `json:"transaction_fee,omitempty"`
	ReceivableAmount *ValueInfo `json:"receivable_amount,omitempty"`
	ExchangeRate     string     `json:"exchange_rate,omitempty"`
	NoteToPayer      string     `json:"note_to_payer,omitempty"`
	ParentPayment    string     `json:"parent_payment,omitempty"`
	CreateTime       string     `json:"create_time,omitempty"`
	UpdateTime       string     `json:"update_time,omitempty"`
	Links            []Link     `json:"links,omitempty"`
}
