package model

//Authorization represents an authorization json object
type Authorization struct {
	ID                        string             `json:"id,omitempty"`
	Amount                    *Amount            `json:"amount,omitempty"`
	PaymentMode               string             `json:"payment_mode,omitempty"`
	State                     string             `json:"state,omitempty"`
	ReasonCode                string             `json:"reason_code,omitempty"`
	ProtectionEligibility     string             `json:"protection_eligibility,omitempty"`
	ProtectionEligibilityType string             `json:"protection_eligibility_type,omitempty"`
	FMFDetails                *FMFDetails        `json:"fmf_details,omitempty"`
	ReceiptID                 string             `json:"receipt_id,omitempty"`
	ParentPayment             string             `json:"parent_payment,omitempty"`
	ProcessorResponse         *ProcessorResponse `json:"processor_response,omitempty"`
	ValidUntil                string             `json:"valid_until,omitempty"`
	CreateTime                string             `json:"create_time,omitempty"`
	UpdateTime                string             `json:"update_time,omitempty"`
	Links                     []Link             `json:"links,omitempty"`
}

func NewDefaultAuthorization() *Authorization {
	var emptyLinks []Link
	return &Authorization{
		Amount:            NewDefaultAmount(),
		FMFDetails:        NewDefaultFMFDetails(),
		ProcessorResponse: NewDefaultProcessorResponse(),
		Links:             emptyLinks,
	}
}

func (r *Authorization) AddLilnk(link *Link) {
	r.Links = append(r.Links, *link)
}

func (r *Authorization) GetLinks() []Link {
	return r.Links
}

func (r *Authorization) GetUpdateTime() string {
	return r.UpdateTime
}
func (r *Authorization) SetUpdateTime(value string) {
	r.UpdateTime = value
}
func (r *Authorization) GetCreateTime() string {
	return r.CreateTime
}

func (r *Authorization) SetCreateTime(value string) {
	r.CreateTime = value
}

func (a *Authorization) GetValidUntil() string {
	return a.ValidUntil
}
func (a *Authorization) SetValidUntil(value string) {
	a.ValidUntil = value
}
func (a *Authorization) GetProcessorResponse() *ProcessorResponse {
	return a.ProcessorResponse
}
func (a *Authorization) SetProcessorResponse(value *ProcessorResponse) {
	a.ProcessorResponse = value
}

func (a *Authorization) GetParentPayment() string {
	return a.ParentPayment
}
func (a *Authorization) SetParentPayment(value string) {
	a.ParentPayment = value
}

func (a *Authorization) GetReceiptID() string {
	return a.ReceiptID
}
func (a *Authorization) SetReceiptID(value string) {
	a.ReceiptID = value
}

func (a *Authorization) GetFMFDetails() *FMFDetails {
	return a.FMFDetails
}
func (a *Authorization) SetFMFDetails(value *FMFDetails) {
	a.FMFDetails = value
}

func (a *Authorization) GetProtectionEligibilityType() string {
	return a.ProtectionEligibilityType
}
func (a *Authorization) SetProtectionEligibilityType(value string) {
	a.ProtectionEligibilityType = value
}

func (a *Authorization) GetProtectionEligibility() string {
	return a.ProtectionEligibility
}
func (a *Authorization) SetProtectionEligibility(value string) {
	a.ProtectionEligibility = value
}
func (a *Authorization) GetReasonCode() string {
	return a.ReasonCode
}
func (a *Authorization) SetReasonCode(value string) {
	a.ReasonCode = value
}
func (a *Authorization) GetState() string {
	return a.State
}
func (a *Authorization) SetState(value string) {
	a.State = value
}
func (a *Authorization) GetPaymentMode() string {
	return a.PaymentMode
}
func (a *Authorization) SetPaymentMode(value string) {
	a.PaymentMode = value
}

func (a *Authorization) GetAmount() *Amount {
	return a.Amount
}
func (a *Authorization) SetAmount(value *Amount) {
	a.Amount = value
}
func (a *Authorization) GetID() string {
	return a.ID
}
func (a *Authorization) SetID(value string) {
	a.ID = value
}
