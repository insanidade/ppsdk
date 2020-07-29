package model

type SelectedInstallmentOption struct {
	Term               int             `json:"term,omitempty"`
	MonthlyPayment     *MonthlyPayment `json:"monthly_payment,omitempty"`
	DiscountPercentage string          `json:"discount_percentage,omitempty"`
	DiscountAmount     *DiscountAmount `json:"discount_amount,omitempty"`
}

func (sio *SelectedInstallmentOption) SetDiscountAmount(value *DiscountAmount) {
	sio.DiscountAmount = value
}

func (sio *SelectedInstallmentOption) GetDiscountAmount() *DiscountAmount {
	return sio.DiscountAmount
}

func (sio *SelectedInstallmentOption) SetDiscountPercentage(value string) {
	sio.DiscountPercentage = value
}

func (sio *SelectedInstallmentOption) GetDisountPercentage() string {
	return sio.DiscountPercentage
}

func (sio *SelectedInstallmentOption) SetMonthlyPayment(value *MonthlyPayment) {
	sio.MonthlyPayment = value
}
func (sio *SelectedInstallmentOption) GetMonthlyPaymente() *MonthlyPayment {
	return sio.MonthlyPayment
}

func (sio *SelectedInstallmentOption) SetTerm(value int) {
	sio.Term = value
}

func (sio *SelectedInstallmentOption) GetTerm() int {
	return sio.Term
}

func NewSelectedInstallmentOption() *SelectedInstallmentOption {
	return &SelectedInstallmentOption{}
}
