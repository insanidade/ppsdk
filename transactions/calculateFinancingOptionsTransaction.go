package transactions

import (
	"fmt"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type CalculateFinancingOptionsTransaction struct {
	requestContainer  *model.CalculateFinancingOptionsRequestContainer
	responseContainer *model.CalculateFinancingOptionsResponseContainer
	baID              string
}

func NewCalculateFinancingOptionsTransaction() *CalculateFinancingOptionsTransaction {
	return &CalculateFinancingOptionsTransaction{
		requestContainer:  model.NewDefaultCalculateFinancingOptionsRequestContainer(),
		responseContainer: model.NewCalculateFinancingOptionsResponseContainer(),
	}
}

// ##################################################################
// ##########Transaction INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pt *CalculateFinancingOptionsTransaction) GetRequestBody() iface.BodyRoot {
	return pt.requestContainer.GetBody()
}

func (pt *CalculateFinancingOptionsTransaction) GetResponseBody() iface.ResponseBodyRoot {
	return pt.responseContainer.GetBody()
}

func (pt *CalculateFinancingOptionsTransaction) GetRequestHeader() iface.Header {
	return pt.requestContainer.GetHeader()
}

func (pt *CalculateFinancingOptionsTransaction) GetResponseHeader() iface.HeaderResponse {
	return pt.responseContainer.GetHeader()
}

func (pt *CalculateFinancingOptionsTransaction) SetRequestBody(body iface.BodyRoot) {
	pt.requestContainer.SetBody(body)
}

func (pt *CalculateFinancingOptionsTransaction) SetResponseBody(body iface.ResponseBodyRoot) {
	pt.responseContainer.SetBody(body)
}
func (pt *CalculateFinancingOptionsTransaction) SetRequestHeader(header iface.Header) {
	pt.requestContainer.SetHeader(header)
}
func (pt *CalculateFinancingOptionsTransaction) SetResponseHeader(header iface.HeaderResponse) {
	pt.responseContainer.SetHeader(header)
}

func (pt *CalculateFinancingOptionsTransaction) GetResponseCode() int {
	return pt.responseContainer.GetCode()
}
func (pt *CalculateFinancingOptionsTransaction) SetResponseCode(code int) {
	pt.responseContainer.SetCode(code)
}

func (pt *CalculateFinancingOptionsTransaction) GetResponseStatus() string {
	return pt.responseContainer.GetStatus()
}
func (pt *CalculateFinancingOptionsTransaction) SetResponseStatus(status string) {
	pt.responseContainer.SetStatus(status)
}

func (pt *CalculateFinancingOptionsTransaction) GetRequestMethod() string {
	return pt.requestContainer.GetMethod()
}
func (pt *CalculateFinancingOptionsTransaction) GetRequestURL() string {
	return pt.requestContainer.GetURL()
}

func (pt *CalculateFinancingOptionsTransaction) SetBearerToken(token string) {
	pt.GetRequestHeader().SetBearerToken(token)
}

func (pt *CalculateFinancingOptionsTransaction) AssembleRequestBody() {
	//montar objeto para, depois de pronto, fazer pc.requestcontainer.setbody(obj)
	fmt.Println("CHAMANDO TRANSACTION ASSEMBLE PARA CalculateFinancingOptionsTransaction")

	root := model.NewCalculateFinancingOptionsRoot()
	root.SetFinancingCountryCode("BR")

	txnAmount := root.GetTransactionAmount()
	txnAmount.SetValue("251")
	txnAmount.SetCurrency("BRL")

	fundingInst := root.GetFundingInstrument()
	fundingInst.SetType("BILLING_AGREEMENT")
	ba := fundingInst.GetBA()
	ba.SetBA(pt.baID)

	fundingInst.SetBA(ba)

	root.SetFundingInstrument(fundingInst)
	root.SetTransactionAmount(txnAmount)

	pt.SetRequestBody(root)

}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
func (pt *CalculateFinancingOptionsTransaction) SetBAID(id string) {

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
