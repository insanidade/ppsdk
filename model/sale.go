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
