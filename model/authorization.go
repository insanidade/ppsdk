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
