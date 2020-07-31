package model

type BAgreementRoot struct {
	TokenID string `json:"token_id,omitempty"`

	//Not part of json so I don't want to export
	valid bool
}

func NewBAgreementRoot() *BAgreementRoot {
	// var emptyTransactions []Transaction
	return &BAgreementRoot{
		TokenID: "EMPTY TOKEN ID",
		// Intent:       intent,
		// Transactions: emptyTransactions}
	}
}
func (bagreement *BAgreementRoot) SetTokenID(value string) {
	bagreement.TokenID = value
}

func (bagreement *BAgreementRoot) GetTokenID() string {
	return bagreement.TokenID
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//SetValid implements SetValid from BodyRoot interface
func (pr *BAgreementRoot) SetValid(valid bool) {
	pr.valid = valid
}

//IsValid implements IsValid from BodyRoot interface
func (pr *BAgreementRoot) IsValid() bool {
	return pr.valid
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
