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
	return &BATokenRoot{
		Payer:              NewPayer(),
		ShippingAddress:    NewShippingAddress(),
		Plan:               NewPlan(),
		MerchantCustomData: "custom infor from merchant"}
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
func (bat *BATokenRoot) SetDescription(value string) {
	bat.Description = value
}
func (bat *BATokenRoot) GetDescription() string {
	return bat.Description
}

func (bat *BATokenRoot) SetPayer(thePayer *Payer) {
	bat.Payer = thePayer
}
func (bat *BATokenRoot) GetPayer() *Payer {
	return bat.Payer
}

func (bat *BATokenRoot) SetShippingAddress(value *ShippingAddress) {
	bat.ShippingAddress = value
}
func (bat *BATokenRoot) GetShippingAddress() *ShippingAddress {
	return bat.ShippingAddress
}

func (bat *BATokenRoot) GetPlan() *Plan {
	return bat.Plan
}

func (bat *BATokenRoot) SetPlan(thePlan *Plan) {
	bat.Plan = thePlan
}
func (bat *BATokenRoot) GetMerchantCustomData() string {
	return bat.MerchantCustomData
}
func (bat *BATokenRoot) SetMerchantCustomData(value string) {
	bat.MerchantCustomData = value
}
