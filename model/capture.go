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

func NewDefaultCapture() *Capture {
	var emptyLinks []Link
	return &Capture{
		Amount:           NewDefaultAmount(),
		TransactionFee:   NewDefaultValueInfo(),
		ReceivableAmount: NewDefaultValueInfo(),
		Links:            emptyLinks,
	}
}
func (r *Capture) AddLilnk(link *Link) {
	r.Links = append(r.Links, *link)
}

func (r *Capture) GetLinks() []Link {
	return r.Links
}
func (r *Capture) GetUpdateTime() string {
	return r.UpdateTime
}
func (r *Capture) SetUpdateTime(value string) {
	r.UpdateTime = value
}
func (r *Capture) GetCreateTime() string {
	return r.CreateTime
}

func (r *Capture) SetCreateTime(value string) {
	r.CreateTime = value
}

func (cap *Capture) GetParentPayment() string {
	return cap.ParentPayment
}
func (cap *Capture) SetParentPayment(value string) {
	cap.ParentPayment = value
}

func (cap *Capture) GetNoteToPayer() string {
	return cap.NoteToPayer
}
func (cap *Capture) SetNoteToPayer(value string) {
	cap.NoteToPayer = value
}
func (cap *Capture) GetExchangeRate() string {
	return cap.ExchangeRate
}
func (cap *Capture) SetExchangeRate(value string) {
	cap.ExchangeRate = value
}

func (cap *Capture) GetReceivableAmount() *ValueInfo {
	return cap.ReceivableAmount
}
func (cap *Capture) SetReceivableAmount(value *ValueInfo) {
	cap.ReceivableAmount = value
}

func (cap *Capture) GetTransactionFee() *ValueInfo {
	return cap.TransactionFee
}
func (cap *Capture) SetTransactionFee(value *ValueInfo) {
	cap.TransactionFee = value
}

func (cap *Capture) GetInvoiceNumber() string {
	return cap.InvoiceNumber
}
func (cap *Capture) SetInvoiceNumber(value string) {
	cap.InvoiceNumber = value
}
func (cap *Capture) GetReasonCode() string {
	return cap.ReasonCode
}
func (cap *Capture) SetReasonCode(value string) {
	cap.ReasonCode = value
}

func (cap *Capture) GetState() string {
	return cap.State
}
func (cap *Capture) SetState(value string) {
	cap.State = value
}
func (cap *Capture) GetIsFinalCapture() bool {
	return cap.IsFinalCapture
}
func (cap *Capture) SetIsFinalCapture(value bool) {
	cap.IsFinalCapture = value
}
func (cap *Capture) GetID() string {
	return cap.ID
}
func (cap *Capture) SetID(valut string) {
	cap.ID = valut
}
