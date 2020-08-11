package model

type QualifyingFinancingOption struct {
	CreditFinancing       *CreditFinancing `json:"credit_financing,omitempty"`
	MonthlyPercentageRate string           `json:"monthly_percentage_rate,omitempty"`
	DiscountPercentage    string           `json:"discount_percentage,omitempty"`
	PayPalSubsidy         bool             `json:"paypal_subsidy,omitempty"`
	MinAmount             *MinAmount       `json:"min_amount,omitempty"`
	PeriodicPayment       *PeriodicPayment `json:"periodic_payment,omitempty"`
	MonthlyPayment        *MonthlyPayment  `json:"monthly_payment,omitempty"`
	DiscountAmount        *DiscountAmount  `json:"discount_amount,omitempty"`
	TotalInterest         *TotalInterest   `json:"total_interest,omitempty"`
	TotalCost             *TotalCost       `json:"total_cost,omitempty"`
}

func NewDefaultQualifyingFinancingOption() *QualifyingFinancingOption {
	return &QualifyingFinancingOption{}
}
