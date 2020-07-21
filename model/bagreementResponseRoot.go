package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type BAgreementResponseRoot struct {
	//Successful response parameters
	Id                 string    `json:"id,omitempty"`
	State              string    `json:"state,omitempty"`
	Description        string    `json:"description,omitempty"`
	Merchant           *Merchant `json:"merchant,omitempty"`
	Payer              *Payer    `json:"payer,omitempty"`
	Plan               *Plan     `json:"plan,omitempty"`
	MerchantCustomData string    `json:"merchant_custom_data,omitempty"`
	CreateTime         string    `json:"create_time,omitempty"`
	UpdateTime         string    `json:"update_time,omitempty"`
	Links              []Link    `json:"links,omitempty"`
	//Failure parameters
	Name             string             `json:"name,omitempty"`
	DebugID          string             `json:"debug_id,omitempty"`
	Error            string             `json:"error,omitempty"`
	ErrorDescription string             `json:"error_description,omitempty"`
	Message          string             `json:"message,omitempty"`
	InformationLink  string             `json:"information_link,omitempty"`
	Details          []DetailsOnFailure `json:"details,omitempty"`
	//Non-json attributes. Not exported.
	valid bool
}

func NewBAgreementResponseRoot() *BAgreementResponseRoot {
	return &BAgreementResponseRoot{}
}

// ##################################################################
// ###ResponseBodyRoot INTERFACE IMPLEMENTATIONS#####################
// ##################################################################
func (prr *BAgreementResponseRoot) GetLinks() map[string]iface.Link {
	linkMap := make(map[string]iface.Link)
	for _, link := range prr.Links {
		newLink := NewLink(link.GetHref(), link.GetRel(), link.GetMethod())
		linkMap[link.GetRel()] = newLink
	}
	return linkMap
}
func (prr *BAgreementResponseRoot) GetId() string {
	return prr.Id
}

func (prr *BAgreementResponseRoot) GetDebugID() string {
	return prr.DebugID
}

// //SetValid sets validity
// func (prr *PaymentResponseRoot) SetValid(valid bool) {
// 	prr.valid = valid
// }
//
// //IsValid returns validity
// func (prr *PaymentResponseRoot) IsValid() bool {
// 	return prr.valid
// }

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
