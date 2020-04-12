package model

//Payee represents a payee json object
type Payee struct {
	Email      string `json:"email,omitempty"`
	MerchantID string `json:"merchant_id,omitempty"`
}

//NewPayee returns a new Payee instance
func NewPayee(email string, merchantID string) *Payee {
	return &Payee{Email: email,
		MerchantID: merchantID,
	}
}
