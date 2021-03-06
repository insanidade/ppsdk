package transactions

import (
	"log"
	"os"
	"runtime"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

//PaymentTransaction struct (testing access to private repo by go get)
type PaymentTransaction struct {
	// header            *model.HeaderForREST
	// bodyResponse      iface.ResponseBodyRoot
	// headerResponse    iface.Header
	invoice           iface.Invoice
	id                string
	links             map[string]iface.Link
	requestContainer  iface.RequestContainer
	responseContainer iface.ResponseContainer
}

//NewPaymentTransaction returns a new instance of a PaymentTransaction.
//The invoice private attribute is set as an empty value. It shall be set by
//invoking the SetInvoiceNumber() function.
func NewPaymentTransaction(reqContainer iface.RequestContainer, respContainer iface.ResponseContainer) *PaymentTransaction {
	// header := model.NewHeaderForREST()
	// body := model.NewPaymentRoot()
	setUpLoggingFile("payment.log")
	return &PaymentTransaction{
		// body:              body,
		// bodyResponse:      model.NewPaymentResponseRoot(),
		// header:            header,
		// headerResponse:    model.NewHeaderForRESTResponse(),
		invoice:           model.NewInvoiceNumber(), //empty invoice
		links:             make(map[string]iface.Link),
		requestContainer:  reqContainer,
		responseContainer: respContainer}
}

// func NewPaymentTransactionWithToken(token string) *PaymentTransaction {
// 	header := model.NewHeaderForREST()
// 	header.SetBearerToken(token)
// 	setUpLoggingFile("payment.log")
// 	return &PaymentTransaction{
// 		body:    model.NewPaymentRoot(),
// 		header:  header,
// 		invoice: model.NewInvoiceNumber(),
// 	}
// }

// func NewPaymentTransactionWithInvoiceNumber(token string, invoiceNumber iface.Invoice) *PaymentTransaction {
// 	header := model.NewHeaderForREST()
// 	header.SetBearerToken(token)
// 	setUpLoggingFile("payment.log")
// 	return &PaymentTransaction{
// 		body:    model.NewPaymentRoot(),
// 		header:  header,
// 		invoice: invoiceNumber,
// 	}
// }
//
// func NewPaymentTransactionWithNegativeTest(token string, negativeTestName string) *PaymentTransaction {
// 	header := model.NewHeaderForREST()
// 	header.SetBearerToken(token)
// 	header.SetNegativeTest(negativeTestName)
// 	setUpLoggingFile("payment.log")
// 	return &PaymentTransaction{
// 		body:    model.NewPaymentRoot(),
// 		header:  header,
// 		invoice: model.NewInvoiceNumber(),
// 	}
// }

// ##################################################################
// ##########Transaction INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pt *PaymentTransaction) GetRequestBody() iface.BodyRoot {
	return pt.requestContainer.GetBody()
}

func (pt *PaymentTransaction) GetResponseBody() iface.ResponseBodyRoot {
	return pt.responseContainer.GetBody()
}

func (pt *PaymentTransaction) GetRequestHeader() iface.Header {
	return pt.requestContainer.GetHeader()
}

func (pt *PaymentTransaction) GetResponseHeader() iface.HeaderResponse {
	return pt.responseContainer.GetHeader()
}

func (pt *PaymentTransaction) SetRequestBody(body iface.BodyRoot) {
	pt.requestContainer.SetBody(body)
}

func (pt *PaymentTransaction) SetResponseBody(body iface.ResponseBodyRoot) {
	pt.responseContainer.SetBody(body)
}
func (pt *PaymentTransaction) SetRequestHeader(header iface.Header) {
	pt.requestContainer.SetHeader(header)
}
func (pt *PaymentTransaction) SetResponseHeader(header iface.HeaderResponse) {
	pt.responseContainer.SetHeader(header)
}

func (pt *PaymentTransaction) GetResponseCode() int {
	return pt.responseContainer.GetCode()
}
func (pt *PaymentTransaction) SetResponseCode(code int) {
	pt.responseContainer.SetCode(code)
}

func (pt *PaymentTransaction) GetResponseStatus() string {
	return pt.responseContainer.GetStatus()
}
func (pt *PaymentTransaction) SetResponseStatus(status string) {
	pt.responseContainer.SetStatus(status)
}

func (pt *PaymentTransaction) GetRequestMethod() string {
	return pt.requestContainer.GetMethod()
}
func (pt *PaymentTransaction) GetRequestURL() string {
	return pt.requestContainer.GetURL()
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################

//BuildCreatePaymentContainers returns a container that shall be used for
//creating a new payment. That container is sent as a parameter to the
//httpHandler object, which shall perform the http tasks and return
//a response
// func (pt *PaymentTransaction) BuildCreatePaymentContainers() (*model.PaymentRequestContainer, *model.PaymentResponseContainer) {
// 	//Build the body.
// 	// pt.bodyFactory()
// 	return model.NewPaymentRequestContainer(pt.header,
// 		pt.body), model.NewPaymentResponseContainer()
// }

// func (pt *PaymentTransaction) BuildGetTransactionDetailsContainers() (iface.RequestContainer,
// 	iface.ResponseContainer) {
// 	trace()
// 	log.Printf("URL PARA GET PAYMENT DETAILS: %s\n", pt.links["self"].GetHref())
// 	return model.NewGetPaymentRequestContainer(pt.header, pt.links["self"].GetHref()),
// 		model.NewPaymentResponseContainer()
// }

// func (pt *PaymentTransaction) SetPaymentResponseInfo(responseContainer *model.PaymentResponseContainer) {
//
// 	log.Printf("SetPaymentResponseInfo: HTTP Status code: %d\n", responseContainer.GetCode())
//
// 	// pt.headerResponse = responseContainer.GetHeader()
// 	// pt.bodyResponse = responseContainer.GetBody()
// 	pt.links = responseContainer.GetBody().GetLinks() // pt.bodyResponse.GetLinks()
//
// }

func (pb *PaymentTransaction) GetLinks() map[string]iface.Link {
	return pb.links
}

func (pb *PaymentTransaction) SetInvoiceNumber(invoiceNumber iface.Invoice) {
	pb.invoice = invoiceNumber
}
func (pt *PaymentTransaction) SetBearerToken(token string) {
	pt.GetRequestHeader().SetBearerToken(token)
}

func (pt *PaymentTransaction) SetNegativeTest(negativeTestValue string) {
	pt.GetRequestHeader().SetNegativeTest(negativeTestValue)
}

func (pt *PaymentTransaction) RemoveNegativeTest() {
	pt.GetRequestHeader().RemoveNegativeTest()
}

//bodyFactory builds a PaymentRoot object and returns it.
// func (pc *PaymentTransaction) bodyFactory() {
// 	// fmt.Printf("bearer token: %s\n", cp.AuthToken)
//
// 	item := model.NewItem("BRL", "produto X", "Xzonator", "90", 2, "sku_123654")
// 	item2 := model.NewItem("BRL", "produto Y", "enhanced Y", "100", 3, "sku_9876")
// 	shippingAddress := model.NewShippingAddress("Campina Grannde",
// 		"BR",
// 		"Rua X,Bairro Y",
// 		"apto 1700",
// 		"NORMALIZED",
// 		"58108125",
// 		"Otávio Augusto",
// 		"PB")
//
// 	itemList := model.NewItemList("1133222222")
// 	itemList.AddItem(item)
// 	itemList.AddItem(item2)
// 	itemList.SetShippingAddress(shippingAddress)
//
// 	details := model.NewDetails("480", "20")
// 	amount := model.NewAmount("BRL", "500", details)
//
// 	transaction := model.NewTransaction("this is a description",
// 		"this is a note to the payee",
// 		"this is some custom message",
// 		pc.invoice.String(),
// 		"soft desc",
// 		"odefranca notify url")
// 	transaction.SetAmount(amount)
// 	transaction.SetItemList(itemList)
// 	fmt.Printf("Transaction: %+v", transaction)
//
// 	appContext := model.NewApplicationContext("odefranca brand name",
// 		"Login",
// 		"BR",
// 		"SET_PROVIDED_ADDRESS",
// 		"continue")
// 	fmt.Printf("App Context: %+v", appContext)
//
// 	payer := model.NewPayer("paypal")
// 	redirectUrls := model.NewRedirectURLS("https://localhost/go/return", "https://localhost/go/cancel")
//
// 	paymentBody := model.NewPaymentRoot()
// 	paymentBody.SetIntent("sale")
// 	paymentBody.SetApplicationContext(appContext)
// 	paymentBody.AddTransaction(transaction)
// 	paymentBody.SetPayer(payer)
// 	paymentBody.SetRedirectURLS(redirectUrls)
//
// 	pc.body = paymentBody
//
// }

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Printf("%s:%d %s\n", file, line, f.Name())
}

func setUpLoggingFile(logFileName string) {
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if nil != err {
		return
	}

	log.SetOutput(logFile)

}
