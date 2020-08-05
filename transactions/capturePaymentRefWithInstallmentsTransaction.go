package transactions

import (
	"fmt"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type CapturePaymentRefWithInstallmentsTransaction struct {
	requestContainer  *model.CapturePaymentRefWithInstallmentsRequestContainer
	responseContainer *model.CapturePaymentRefWithInstallmentsResponseContainer
	baID              string
}

func NewCapturePaymentRefWithInstallmentsTransaction() *CapturePaymentRefWithInstallmentsTransaction {
	return &CapturePaymentRefWithInstallmentsTransaction{
		requestContainer:  model.NewDefaultCapturePaymentRefWithInstallmentsRequestContainer(),
		responseContainer: model.NewDefaultCapturePaymentRefWithInstallmentsResponseContainer(),
	}
}

// ##################################################################
// ##########Transaction INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetRequestBody() iface.BodyRoot {
	return pt.requestContainer.GetBody()
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetResponseBody() iface.ResponseBodyRoot {
	return pt.responseContainer.GetBody()
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetRequestHeader() iface.Header {
	return pt.requestContainer.GetHeader()
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetResponseHeader() iface.HeaderResponse {
	return pt.responseContainer.GetHeader()
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) SetRequestBody(body iface.BodyRoot) {
	pt.requestContainer.SetBody(body)
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) SetResponseBody(body iface.ResponseBodyRoot) {
	pt.responseContainer.SetBody(body)
}
func (pt *CapturePaymentRefWithInstallmentsTransaction) SetRequestHeader(header iface.Header) {
	pt.requestContainer.SetHeader(header)
}
func (pt *CapturePaymentRefWithInstallmentsTransaction) SetResponseHeader(header iface.HeaderResponse) {
	pt.responseContainer.SetHeader(header)
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetResponseCode() int {
	return pt.responseContainer.GetCode()
}
func (pt *CapturePaymentRefWithInstallmentsTransaction) SetResponseCode(code int) {
	pt.responseContainer.SetCode(code)
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetResponseStatus() string {
	return pt.responseContainer.GetStatus()
}
func (pt *CapturePaymentRefWithInstallmentsTransaction) SetResponseStatus(status string) {
	pt.responseContainer.SetStatus(status)
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) GetRequestMethod() string {
	return pt.requestContainer.GetMethod()
}
func (pt *CapturePaymentRefWithInstallmentsTransaction) GetRequestURL() string {
	return pt.requestContainer.GetURL()
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) SetBearerToken(token string) {
	pt.GetRequestHeader().SetBearerToken(token)
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) AssembleRequestBody() {
	//montar objeto para, depois de pronto, fazer pc.requestcontainer.setbody(obj)
	fmt.Println("CHAMANDO TRANSACTION ASSEMBLE PARA CapturePaymentRefWithInstallmentsTransaction")

	root := model.NewDefaultPaymentRoot()
	root.SetIntent("sale")

	payer := root.Payer
	payer.SetPaymentMethod("paypal")

	fi := model.NewDefaultFundingInstrument()

	billing := fi.Billing
	billing.SetBillingAgreementID(pt.baID)
	installmentOption := billing.GetSelectedInstallmentOption()
	installmentOption.SetTerm(1)
	monthlyPayment := installmentOption.GetMonthlyPaymente()
	monthlyPayment.SetCurrency("BRL")
	monthlyPayment.SetValue("251")
	installmentOption.SetMonthlyPayment(monthlyPayment)

	pt.SetRequestBody(root)

}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
func (pt *CapturePaymentRefWithInstallmentsTransaction) SetBAID(id string) {

	pt.baID = id

	// fmt.Printf("VALOR DE BA TYPE ESTA DEFINIDO: %s\n", pt.baType)
}

// func setUpLoggingFile(logFileName string) {
// 	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

// 	if nil != err {
// 		return
// 	}

// 	log.SetOutput(logFile)

// }
