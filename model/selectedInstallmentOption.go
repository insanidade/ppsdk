package model

type SelectedInstallmentOption struct {
	Term               int             `json:"term,omitempty"`
	MonthlyPayment     *MonthlyPayment `json:"monthly_payment,omitempty"`
	DiscountPercentage string          `json:"discount_percentage,omitempty"`
	DiscountAmount     *DiscountAmount `json:"discount_amount,omitempty"`
}
