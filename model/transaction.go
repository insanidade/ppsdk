package model

//Transaction defines a transaction json object
type Transaction struct {
	Description      string            `json:"description,omitempty"`
	NoteToPayee      string            `json:"note_to_payee,omitempty"`
	Custom           string            `json:"custom,omitempty"`
	InvoiceNumber    string            `json:"invoice_number,omitempty"`
	SoftDescriptor   string            `json:"soft_descriptor,omitempty"`
	NotifyURL        string            `json:"notify_url,omitempty"`
	Amount           *Amount           `json:"amount,omitempty"`
	Payee            *Payee            `json:"payee,omitempty"`
	PaymentOptions   *PaymentOptions   `json:"payment_options,omitempty"`
	ItemList         *ItemList         `json:"item_list,omitempty"`
	RelatedResources []RelatedResource `json:"related_resources,omitempty"`
}

func NewDefaultTransaction() *Transaction {
	var emptyRelatedResource []RelatedResource
	return &Transaction{
		PaymentOptions:   NewPaymentOptions(),
		Amount:           NewDefaultAmount(),
		RelatedResources: emptyRelatedResource,
	}
}

//NewTransaction constructor
func NewTransaction(desc string,
	note string,
	custom string,
	inv string,
	soft string,
	notifURL string) *Transaction {
	var emptyRelatedResource []RelatedResource
	return &Transaction{
		Description:      desc,
		NoteToPayee:      note,
		Custom:           custom,
		InvoiceNumber:    inv,
		SoftDescriptor:   soft,
		NotifyURL:        notifURL,
		PaymentOptions:   NewPaymentOptions(),
		RelatedResources: emptyRelatedResource,
	}
}

//SetItemList sets the item_list for this transaction object
func (t *Transaction) SetItemList(il *ItemList) {
	t.ItemList = il
}

//SetAmount sets the Amount object for this transaction
//######################################################
func (t *Transaction) SetAmount(a *Amount) {
	t.Amount = a
}

//AddRelatedResource adds a RelatedResource object to the
//RelatedResources slice
func (t *Transaction) AddRelatedResource(rs *RelatedResource) {
	t.RelatedResources = append(t.RelatedResources, *rs)
}
