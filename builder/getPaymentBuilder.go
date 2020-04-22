package builder

import (
	"log"
	"os"
	"runtime"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

// const paypalSandboxAPIURL string = "https://api.sandbox.paypal.com"
// const createPaymentURL string = "/v1/payments/payment"

type GetPaymentBuilder struct {
	body   *model.EmptyBody
	header *model.HeaderForREST
	url    string
}

func NewGetPaymentBuilder(token string, url string) *GetPaymentBuilder {
	header := model.NewHeaderForREST()
	header.SetBearerToken(token)
	// teste := NewPaymentBuilder(token)
	// teste.header.SetNegativeTest("negativeTestName")
	setUpLoggingFile("payment.log")
	return &GetPaymentBuilder{
		body:   model.NewEmptyBody(),
		header: header,
		url:    url}
}

func NewGetPaymentBuilderWithNegativeTest(token string, url string, negativeTest string) *GetPaymentBuilder {
	header := model.NewHeaderForREST()
	header.SetBearerToken(token)
	header.SetNegativeTest(negativeTest)
	// teste := NewPaymentBuilder(token)
	// teste.header.SetNegativeTest("negativeTestName")
	setUpLoggingFile("get_payment.log")
	return &GetPaymentBuilder{
		body:   model.NewEmptyBody(),
		header: header,
		url:    url}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pb *GetPaymentBuilder) BuildRequestContainer() iface.RequestContainer {

	return model.NewGetPaymentRequestContainer(pb.header, pb.url)

}
func (pb *GetPaymentBuilder) BuildResponseContainer() iface.ResponseContainer {
	return nil
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
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
