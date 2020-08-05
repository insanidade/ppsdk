package model

//Sale represents a sale json object
type Sale struct {
	ID                        string              `json:"id,omitempty"`
	Amount                    *Amount             `json:"amount,omitempty"`
	PaymentMode               string              `json:"payment_mode,omitempty"`
	State                     string              `json:"state,omitempty"`
	ReasonCode                string              `json:"reason_code,omitempty"`
	ProtectionEligibility     string              `json:"protection_eligibility,omitempty"`
	ProtectionEligibilityType string              `json:"protection_eligibility_type,omitempty"`
	ClearingTime              string              `json:"clearing_time,omitempty"`
	PaymentHoldStatus         string              `json:"payment_hold_status,omitempty"`
	PaymentHoldReasons        []PaymentHoldReason `json:"payment_hold_reasons,omitempty"`
	TransactionFee            *ValueInfo          `json:"transaction_fee,omitempty"`
	ReceivableAmount          *ValueInfo          `json:"receivable_amount,omitempty"`
	ExchangeRate              string              `json:"exchange_rate,omitempty"`
	FMFDetails                *FMFDetails         `json:"fmf_details,omitempty"`
	ReceiptID                 string              `json:"receipt_id,omitempty"`
	ParentPayment             string              `json:"parent_payment,omitempty"`
	ProcessorResponse         *ProcessorResponse  `json:"processor_response,omitempty"`
	BillingAgreementID        string              `json:"billing_agreement_id,omitempty"`
	CreateTime                string              `json:"create_time,omitempty"`
	UpdateTime                string              `json:"update_time,omitempty"`
	Links                     []Link              `json:"links,omitempty"`
}

func NewDefaultSale() *Sale {
	var emptyPaymentHoldReasons []PaymentHoldReason
	var emptyLinks []Link
	return &Sale{
		Amount:             NewDefaultAmount(),
		PaymentHoldReasons: emptyPaymentHoldReasons,
		TransactionFee:     NewDefaultValueInfo(),
		ReceivableAmount:   NewDefaultValueInfo(),
		FMFDetails:         NewDefaultFMFDetails(),
		ProcessorResponse:  NewDefaultProcessorResponse(),
		Links:              emptyLinks,
	}
}

func (r *Sale) AddPaymentHoldReason(value *PaymentHoldReason) {
	r.PaymentHoldReasons = append(r.PaymentHoldReasons, *value)
}

func (r *Sale) GetPaymentHoldReasons() []PaymentHoldReason {
	return r.PaymentHoldReasons
}

func (r *Sale) AddLink(link *Link) {
	r.Links = append(r.Links, *link)
}

func (r *Sale) GetLinks() []Link {
	return r.Links
}

func (r *Sale) GetUpdateTime() string {
	return r.UpdateTime
}
func (r *Sale) SetUpdateTime(value string) {
	r.UpdateTime = value
}
func (r *Sale) GetCreateTime() string {
	return r.CreateTime
}

func (r *Sale) SetCreateTime(value string) {
	r.CreateTime = value
}

func (a *Sale) GetProcessorResponse() *ProcessorResponse {
	return a.ProcessorResponse
}
func (a *Sale) SetProcessorResponse(value *ProcessorResponse) {
	a.ProcessorResponse = value
}

func (a *Sale) GetParentPayment() string {
	return a.ParentPayment
}
func (a *Sale) SetParentPayment(value string) {
	a.ParentPayment = value
}

func (a *Sale) GetReceiptID() string {
	return a.ReceiptID
}
func (a *Sale) SetReceiptID(value string) {
	a.ReceiptID = value
}

func (a *Sale) GetFMFDetails() *FMFDetails {
	return a.FMFDetails
}
func (a *Sale) SetFMFDetails(value *FMFDetails) {
	a.FMFDetails = value
}

func (a *Sale) GetProtectionEligibilityType() string {
	return a.ProtectionEligibilityType
}
func (a *Sale) SetProtectionEligibilityType(value string) {
	a.ProtectionEligibilityType = value
}

func (a *Sale) GetProtectionEligibility() string {
	return a.ProtectionEligibility
}
func (a *Sale) SetProtectionEligibility(value string) {
	a.ProtectionEligibility = value
}
func (a *Sale) GetReasonCode() string {
	return a.ReasonCode
}
func (a *Sale) SetReasonCode(value string) {
	a.ReasonCode = value
}
func (a *Sale) GetState() string {
	return a.State
}
func (a *Sale) SetState(value string) {
	a.State = value
}
func (a *Sale) GetPaymentMode() string {
	return a.PaymentMode
}
func (a *Sale) SetPaymentMode(value string) {
	a.PaymentMode = value
}

func (a *Sale) GetAmount() *Amount {
	return a.Amount
}
func (a *Sale) SetAmount(value *Amount) {
	a.Amount = value
}
func (a *Sale) GetID() string {
	return a.ID
}
func (a *Sale) SetID(value string) {
	a.ID = value
}
