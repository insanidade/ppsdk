package model

//PaymentRoot struct
type PaymentRoot struct {
	Intent             string              `json:"intent,omitempty"`
	NoteToPayer        string              `json:"note_to_payer,omitempty"`
	ApplicationContext *ApplicationContext `json:"application_context,omitempty"`
	Payer              *Payer              `json:"payer"`
	RedirectURLS       *RedirectURLS       `json:"redirect_urls,omitempty"`
	Transactions       []Transaction       `json:"transactions,omitempty"`
	//Not part of json so I don't want to export
	valid bool //`json:"-"` //`json:",omitempty"`
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//SetValid implements SetValid from BodyRoot interface
func (pr *PaymentRoot) SetValid(valid bool) {
	pr.valid = valid
}

//IsValid implements IsValid from BodyRoot interface
func (pr *PaymentRoot) IsValid() bool {
	return pr.valid
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################

//NewPaymentRoot constructor
func NewDefaultPaymentRoot() *PaymentRoot {
	var emptyTransactions []Transaction
	return &PaymentRoot{
		Payer:              NewDefaultPayer(),
		ApplicationContext: NewDefaultApplicationContext(),
		RedirectURLS:       NewDefaultRedirectURLS(),
		Transactions:       emptyTransactions,
	}
}

func (pr *PaymentRoot) SetIntent(intent string) {
	pr.Intent = intent
}

func (pr *PaymentRoot) SetNoteToPayer(ntp string) {
	pr.NoteToPayer = ntp
}

//SetRedirectURLS sets the return and cancel urls
func (pr *PaymentRoot) SetRedirectURLS(r *RedirectURLS) {
	pr.RedirectURLS = r
}

//SetPayer sets the payer
func (pr *PaymentRoot) SetPayer(p *Payer) {
	pr.Payer = p
}

//SetApplicationContext sets the application context object
func (pr *PaymentRoot) SetApplicationContext(ac *ApplicationContext) {
	pr.ApplicationContext = ac
}

//AddTransaction adds a Transaction instance to the list of
//transactions under this payment request.
func (pr *PaymentRoot) AddTransaction(tr *Transaction) {
	pr.Transactions = append(pr.Transactions, *tr)
}
