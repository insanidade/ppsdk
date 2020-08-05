package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type CalculateFinancingOptionsResponseRoot struct {
	//Successful response parameters
	ConfigurationOwnerAccount string            `json:"configuration_owner_account,omitempty"`
	FinancingOptions          []FinancingOption `json:"financing_options,omitempty"`
	//Failure parameters
	Name             string             `json:"name,omitempty"`
	DebugID          string             `json:"debug_id,omitempty"`
	Error            string             `json:"error,omitempty"`
	ErrorDescription string             `json:"error_description,omitempty"`
	Message          string             `json:"message,omitempty"`
	InformationLink  string             `json:"information_link,omitempty"`
	Details          []DetailsOnFailure `json:"details,omitempty"`
	Links            []Link             `json:"links,omitempty"`
	//Non-json attributes. Not exported.
	valid bool
}

func NewCalculateFinancingOptionsResponseRoot() *CalculateFinancingOptionsResponseRoot {
	return &CalculateFinancingOptionsResponseRoot{}
}

func (prr *CalculateFinancingOptionsResponseRoot) GetFinancingOptions() []FinancingOption {
	return prr.FinancingOptions
}

// ##################################################################
// ###ResponseBodyRoot INTERFACE IMPLEMENTATIONS#####################
// ##################################################################
func (prr *CalculateFinancingOptionsResponseRoot) GetLinks() map[string]iface.Link {
	linkMap := make(map[string]iface.Link)
	for _, link := range prr.Links {
		newLink := NewLink(link.GetHref(), link.GetRel(), link.GetMethod())
		linkMap[link.GetRel()] = newLink
	}
	return linkMap
}
func (prr *CalculateFinancingOptionsResponseRoot) GetId() string {
	return ""
}

func (prr *CalculateFinancingOptionsResponseRoot) GetDebugID() string {
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
