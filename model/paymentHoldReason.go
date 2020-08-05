package model

type PaymentHoldReason struct {
	PaymentHoldReason string `json:"payment_hold_reason,omitempty"`
}

func NewDefaultPaymentHoldReason() *PaymentHoldReason {
	return &PaymentHoldReason{}
}

func (phr *PaymentHoldReason) GetPaymentHoldReason() string {
	return phr.PaymentHoldReason
}

func (phr *PaymentHoldReason) SetPaymentHoldReason(value string) {
	phr.PaymentHoldReason = value
}
