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

func NewDefaultRefund() *Refund {
	var emptyLinks []Link
	return &Refund{
		Amount: NewDefaultAmount(),
		Links:  emptyLinks,
	}
}

func (r *Refund) AddLilnk(link *Link) {
	r.Links = append(r.Links, *link)
}

func (r *Refund) GetLinks() []Link {
	return r.Links
}

func (r *Refund) GetUpdateTime() string {
	return r.UpdateTime
}
func (r *Refund) SetUpdateTime(value string) {
	r.UpdateTime = value
}
func (r *Refund) GetCreateTime() string {
	return r.CreateTime
}

func (r *Refund) SetCreateTime(value string) {
	r.CreateTime = value
}
func (r *Refund) GetDescription() string {
	return r.Description
}
func (r *Refund) SetDescription(value string) {
	r.Description = value
}
func (r *Refund) GetParentPayment() string {
	return r.ParentPayment
}

func (r *Refund) SetParentPayment(value string) {
	r.ParentPayment = value
}

func (r *Refund) GetCaptureID() string {
	return r.CaptureID
}

func (r *Refund) SetCaptureID(value string) {
	r.CaptureID = value
}

func (r *Refund) GetSaleID() string {
	return r.SaleID
}
func (r *Refund) SetSaleID(value string) {
	r.SaleID = value
}

func (r *Refund) GetInvoiceNumber() string {
	return r.InvoiceNumber
}
func (r *Refund) SetInvoiceNumber(value string) {
	r.InvoiceNumber = value
}

func (r *Refund) GetReason() string {
	return r.Reason
}

func (r *Refund) SetReason(value string) {
	r.Reason = value
}
func (r *Refund) GetState() string {
	return r.State
}

func (r *Refund) SetState(value string) {
	r.State = value
}
func (r *Refund) GetID() string {
	return r.ID
}
func (r *Refund) SetID(value string) {
	r.ID = value
}
