package builder

import (
	"math/rand"
	"strconv"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type DefaultInvoiceBuilder struct {
	invoice *model.InvoiceNumber
}

func NewDefaultInvoiceBuilder() *DefaultInvoiceBuilder {
	randomVal := invoiceFactory()
	return &DefaultInvoiceBuilder{
		invoice: model.NewInvoiceNumberWithValue(randomVal),
	}
}

func (dib *DefaultInvoiceBuilder) GetInvoice() iface.Invoice {
	return dib.invoice
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func invoiceFactory() string {
	return strconv.FormatInt(randomInt(1000000000, 9999999999), 10)
}
