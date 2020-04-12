package model

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method,omitempty"`
}

func NewPaymentOptions() *PaymentOptions {
	return &PaymentOptions{
		AllowedPaymentMethod: "UNRESTRICTED",
	}
}

func (po *PaymentOptions) SetUnrestricted() {
	po.AllowedPaymentMethod = "UNRESTRICTED"
}

func (po *PaymentOptions) SetInstantFunding() {
	po.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
}

func (po *PaymentOptions) SetImmediatePay() {
	po.AllowedPaymentMethod = "IMMEDIATE_PAY"
}
