package model

type InvoiceNumber struct {
	prefix    string
	value     string
	suffix    string
	separator string
}

func NewInvoiceNumber() *InvoiceNumber {
	return &InvoiceNumber{
		prefix:    *new(string),
		value:     *new(string),
		suffix:    *new(string),
		separator: *new(string),
	}
}

func NewInvoiceNumberWithValuePreAndSuf(p string, v string, s string) *InvoiceNumber {
	return &InvoiceNumber{
		prefix:    p,
		value:     v,
		suffix:    s,
		separator: "_",
	}
}

func NewInvoiceNumberWithValue(v string) *InvoiceNumber {
	return &InvoiceNumber{
		value: v}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################
func (in *InvoiceNumber) String() string {
	return in.prefix + in.separator + in.value + in.separator + in.suffix
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################

func (in *InvoiceNumber) SetSeparator(s string) {
	in.separator = s
}

func (in *InvoiceNumber) SetPrefix(p string) {
	in.prefix = p
}

func (in *InvoiceNumber) SetSuffix(s string) {
	in.suffix = s
}
