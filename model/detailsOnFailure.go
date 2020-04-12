package model

//// TODO: create setters for all attributes

//DetailsOnFailure defines a details json object
//for unsuccessful requests
type DetailsOnFailure struct {
	Location string `json:"location,omitempty"`
	Field    string `json:"field,omitempty"`
	Issue    string `json:"issue,omitempty"`
}

//NewDetailsOnFailure returns a new instance of a details json object
//for unsuccessful requests
func NewDetailsOnFailure(issue string) *DetailsOnFailure {
	return &DetailsOnFailure{
		Issue: issue,
	}
}
