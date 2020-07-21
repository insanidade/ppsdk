package model

type CalculateFinancingOptionsRoot struct {
	FinancingContryCode string               `json:"financing_country_code,omitempty"`
	TransactionAmount   *TransactionAmount   `json:"transaction_amount,omitempty"`
	FundingInstrument   *BAFundingInstrument `json:"funding_instrument,omitempty"`
	// {
	// 	"financing_country_code":"BR",
	// 	"transaction_amount":{
	// 	   "value":"100.46",
	// 	   "currency_code":"BRL"
	// 	},
	// 	"funding_instrument":{
	// 	   "type":"BILLING_AGREEMENT",
	// 	   "billing_agreement":{
	// 		  "billing_agreement_id":"{{ba_id}}"
	// 	   }
	// 	}
	//  }

	//Not part of json so I don't want to export
	valid bool
}

func NewCalculateFinancingOptionsRoot() *CalculateFinancingOptionsRoot {
	// var emptyTransactions []Transaction
	return &CalculateFinancingOptionsRoot{
		FinancingContryCode: "",
		// Intent:       intent,
		// Transactions: emptyTransactions}
	}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//SetValid implements SetValid from BodyRoot interface
func (pr *CalculateFinancingOptionsRoot) SetValid(valid bool) {
	pr.valid = valid
}

//IsValid implements IsValid from BodyRoot interface
func (pr *CalculateFinancingOptionsRoot) IsValid() bool {
	return pr.valid
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
