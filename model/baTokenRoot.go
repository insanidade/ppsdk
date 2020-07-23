package model

type BATokenRoot struct {
	Description        string           `json:"description,omitempty"`
	Payer              *Payer           `json:"payer,omitempty"`
	ShippingAddress    *ShippingAddress `json:"shipping_address,omitempty"`
	Plan               *Plan            `json:"plan,omitempty"`
	MerchantCustomData string           `json:"merchant_custom_data,omitempty"`
	//Not part of json so I don't want to export
	valid bool
}

func NewBATokenRoot() *BATokenRoot {
	// var emptyTransactions []Transaction
	return &BATokenRoot{}
}

// ##################################################################
// #############BodyRoot INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//SetValid implements SetValid from BodyRoot interface
func (pr *BATokenRoot) SetValid(valid bool) {
	pr.valid = valid
}

//IsValid implements IsValid from BodyRoot interface
func (pr *BATokenRoot) IsValid() bool {

	return pr.valid
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
