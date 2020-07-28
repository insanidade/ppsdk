package model

//Payer object. PaymentMethod must be as follows:
// Possible values foar PaymentMethod are:
// credit_card
// paypal
// pay_upon_invoice
// carrier
// alternate_payment
// bank
type Payer struct {
	PaymentMethod      string              `json:"payment_method,omitempty"`
	Status             string              `json:"status,omitempty"`
	FundingInstruments []FundingInstrument `json:"funding_instruments,omitempty"`
	PayerInfo          *PayerInfo          `json:"payer_info,omitempty"`
}

func NewPayer(paymentMethod string) *Payer {
	return &Payer{
		PaymentMethod: paymentMethod,
	}
}

func NewDefaultPayer() *Payer {
	return &Payer{
		PaymentMethod: "PAYPAL",
	}
}

func (p *Payer) SetPaymentMethod(value string) {
	p.PaymentMethod = value
}

func (p *Payer) GetPaymentMethod() string {
	return p.PaymentMethod
}

func (p *Payer) SetPayerInfo(pi *PayerInfo) {
	p.PayerInfo = pi
}
func (p *Payer) GetPayerInfo() *PayerInfo {
	return p.PayerInfo
}

func (p *Payer) SetStatus(status string) {
	p.Status = status
}
func (p *Payer) GetStatus() string {
	return p.Status
}

func (p *Payer) AddFundingInstrument(fi *FundingInstrument) {
	p.FundingInstruments = append(p.FundingInstruments, *fi)
}

func (p *Payer) GetFundingInstruments() []FundingInstrument {
	return p.FundingInstruments
}
