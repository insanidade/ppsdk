package model

type FinancingOption struct {
	Product                    string                      `json:"product,omitempty"`
	CreditProductIdentifier    string                      `json:"credit_product_identifier,omitempty"`
	OptionSetID                string                      `json:"option_set_id,omitempty"`
	QualifyingFinancingOptions []QualifyingFinancingOption `json:"qualifying_financing_options,omitempty"`
}
