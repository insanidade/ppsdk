package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type BATokenResponseRoot struct {
	//Successful response parameters
	TokenId string `json:"token_id,omitempty"`
	Links   []Link `json:"links,omitempty"`
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

func NewBATokenResponseRoot() *BATokenResponseRoot {
	return &BATokenResponseRoot{}
}

// ##################################################################
// ###ResponseBodyRoot INTERFACE IMPLEMENTATIONS#####################
// ##################################################################
func (prr *BATokenResponseRoot) GetLinks() map[string]iface.Link {
	linkMap := make(map[string]iface.Link)
	for _, link := range prr.Links {
		newLink := NewLink(link.GetHref(), link.GetRel(), link.GetMethod())
		linkMap[link.GetRel()] = newLink
	}
	return linkMap
}
func (prr *BATokenResponseRoot) GetId() string {
	return prr.TokenId
}

func (prr *BATokenResponseRoot) GetDebugID() string {
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
