package transactions

import (
	"fmt"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type CapturePaymentRefWithInstallmentsTransaction struct {
	requestContainer    *model.CapturePaymentRefWithInstallmentsRequestContainer
	responseContainer   *model.CapturePaymentRefWithInstallmentsResponseContainer
	BAID                string
	QualifyingFinancing *model.QualifyingFinancingOption
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
	billing.SetBillingAgreementID(pt.BAID)
	installmentOption := billing.GetSelectedInstallmentOption()
	installmentOption.SetTerm(pt.QualifyingFinancing.CreditFinancing.Term)
	monthlyPayment := installmentOption.GetMonthlyPaymente()
	monthlyPayment.SetCurrency(pt.QualifyingFinancing.MonthlyPayment.CurreencyCode)
	monthlyPayment.SetValue(pt.QualifyingFinancing.MonthlyPayment.Value)
	//### check for the existence of discount ###
	if 1 == pt.QualifyingFinancing.CreditFinancing.Term {
		installmentOption.SetDiscountPercentage(pt.QualifyingFinancing.DiscountPercentage)
		discountAmount := installmentOption.GetDiscountAmount()
		discountAmount.SetCurrency(pt.QualifyingFinancing.DiscountAmount.CurreencyCode)
		discountAmount.SetValue(pt.QualifyingFinancing.DiscountAmount.Value)
		installmentOption.SetDiscountAmount(discountAmount)
	}
	//############################
	installmentOption.SetMonthlyPayment(monthlyPayment)

	billing.SetSelectedInstallmentOption(installmentOption)

	fi.SetBilling(billing)

	payer.AddFundingInstrument(fi)
	//TODO: IMPLEMENT TRANSACTIONS OBJECT
	pt.SetRequestBody(root)

}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
func (pt *CapturePaymentRefWithInstallmentsTransaction) SetBAID(value string) {

	pt.BAID = value
	// fmt.Printf("VALOR DE BA TYPE ESTA DEFINIDO: %s\n", pt.baType)
}

func (pt *CapturePaymentRefWithInstallmentsTransaction) SetIntallmentPlan(value *model.QualifyingFinancingOption) {

	pt.QualifyingFinancing = value

	// fmt.Printf("VALOR DE BA TYPE ESTA DEFINIDO: %s\n", pt.baType)
}
