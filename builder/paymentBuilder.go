package builder

import (
	"fmt"
	"log"
	"os"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

const paypalSandboxAPIURL string = "https://api.sandbox.paypal.com"
const createPaymentURL string = "/v1/payments/payment"

//PaymentBuilder struct (testing access to private repo by go get)
type PaymentBuilder struct {
	body    *model.PaymentRoot
	header  *model.HeaderForREST
	invoice iface.Invoice
}

func NewPaymentBuilder(token string) *PaymentBuilder {
	header := model.NewHeaderForREST()
	header.SetBearerToken(token)
	setUpLoggingFile("payment.log")
	return &PaymentBuilder{
		body:    model.NewPaymentRoot(),
		header:  header,
		invoice: model.NewInvoiceNumber(),
	}
}

func NewPaymentBuilderWithInvoiceNumber(token string, invoiceNumber iface.Invoice) *PaymentBuilder {
	header := model.NewHeaderForREST()
	header.SetBearerToken(token)
	setUpLoggingFile("payment.log")
	return &PaymentBuilder{
		body:    model.NewPaymentRoot(),
		header:  header,
		invoice: invoiceNumber,
	}
}

func NewPaymentBuilderWithNegativeTest(token string, negativeTestName string) *PaymentBuilder {
	header := model.NewHeaderForREST()
	header.SetBearerToken(token)
	header.SetNegativeTest(negativeTestName)
	setUpLoggingFile("payment.log")
	return &PaymentBuilder{
		body:    model.NewPaymentRoot(),
		header:  header,
		invoice: model.NewInvoiceNumber(),
	}
}

// SetNegativeTest

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pb *PaymentBuilder) BuildRequestContainer() iface.RequestContainer {
	pb.bodyFactory()
	return model.NewPaymentRequestContainer(pb.header,
		pb.body,
		paypalSandboxAPIURL+createPaymentURL)
}
func (pb *PaymentBuilder) BuildResponseContainer() iface.ResponseContainer {
	return nil
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################

func (pb *PaymentBuilder) SetInvoiceNumber(invoiceNumber iface.Invoice) {
	pb.invoice = invoiceNumber
}

//BodyFactory builds a PaymentRoot object and returns it.
//#################################################################
func (pc *PaymentBuilder) bodyFactory() {
	// fmt.Printf("bearer token: %s\n", cp.AuthToken)

	item := model.NewItem("BRL", "produto X", "Xzonator", "90", 2, "sku_123654")
	item2 := model.NewItem("BRL", "produto Y", "enhanced Y", "100", 3, "sku_9876")
	shippingAddress := model.NewShippingAddress("Campina Grannde",
		"BR",
		"Rua X,Bairro Y",
		"apto 1700",
		"NORMALIZED",
		"58108125",
		"Otávio Augusto",
		"PB")

	itemList := model.NewItemList("1133222222")
	itemList.AddItem(item)
	itemList.AddItem(item2)
	itemList.SetShippingAddress(shippingAddress)

	details := model.NewDetails("480", "20")
	amount := model.NewAmount("BRL", "500", details)

	transaction := model.NewTransaction("this is a description",
		"this is a note to the payee",
		"this is some custom message",
		pc.invoice.String(),
		"soft desc",
		"odefranca notify url")
	transaction.SetAmount(amount)
	transaction.SetItemList(itemList)
	fmt.Printf("Transaction: %+v", transaction)

	appContext := model.NewApplicationContext("odefranca brand name",
		"Login",
		"BR",
		"SET_PROVIDED_ADDRESS",
		"continue")
	fmt.Printf("App Context: %+v", appContext)

	payer := model.NewPayer("paypal")
	redirectUrls := model.NewRedirectURLS("https://localhost/go/return", "https://localhost/go/cancel")

	paymentBody := model.NewPaymentRoot()
	paymentBody.SetIntent("sale")
	paymentBody.SetApplicationContext(appContext)
	paymentBody.AddTransaction(transaction)
	paymentBody.SetPayer(payer)
	paymentBody.SetRedirectURLS(redirectUrls)

	pc.body = paymentBody

}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
func setUpLoggingFile(logFileName string) {
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if nil != err {
		return
	}

	log.SetOutput(logFile)

}
