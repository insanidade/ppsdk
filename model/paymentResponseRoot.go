package model

type PaymentResponseRoot struct {
	//Successful response parameters
	ID                 string              `json:"id,omitempty"`
	Intent             string              `json:"intent,omitempty"`
	Payer              *Payer              `json:"payer,omitempty"`
	ApplicationContext *ApplicationContext `json:"application_context,omitempty"`
	Transactions       []Transaction       `json:"transactions,omitempty"`
	State              string              `json:"string,omitempty"`
	NoteToPayer        string              `json:"note_to_payer,omitempty"`
	RedirectURLS       *RedirectURLS       `json:"redirect_urls,omitempty"`
	FailureReason      string              `json:"failure_reason,omitempty"`
	CreateTime         string              `json:"create_time,omitempty"`
	UpdateTime         string              `json:"update_time,omitempty"`
	Links              []Link              `json:"links,omitempty"`
	//Failure parameters
	Name             string             `json:"name,omitempty"`
	DebugID          string             `json:"debug_id,omitempty"`
	Error            string             `json:"error,omitempty"`
	ErrorDescription string             `json:"error_description,omitempty"`
	Message          string             `json:"message,omitempty"`
	InformationLink  string             `json:"information_link,omitempty"`
	Details          []DetailsOnFailure `json:"details,omitempty"`

	valid bool
}

// ##################################################################
// ##########BodyRoot INTERFACE IMPLEMENTATIONS######################
// ##################################################################

//SetValid sets validity
func (prr *PaymentResponseRoot) SetValid(valid bool) {
	prr.valid = valid
}

//IsValid returns validity
func (prr *PaymentResponseRoot) IsValid() bool {
	return prr.valid
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
