package model

type EmptyBody struct {
	Void  string `json:"void,omitempty"`
	valid bool
}

func NewEmptyBody() *EmptyBody {
	return &EmptyBody{
		valid: true}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//SetValid implements SetValid from BodyRoot interface
func (pr *EmptyBody) SetValid(valid bool) {
	pr.valid = valid
}

//IsValid implements IsValid from BodyRoot interface
func (pr *EmptyBody) IsValid() bool {
	return pr.valid
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
