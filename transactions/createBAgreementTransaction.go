package transactions

import (
	"fmt"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type CreateBAgreementTransaction struct {
	requestContainer  *model.BAgreementRequestContainer
	responseContainer *model.BAgreementResponseContainer
	baToken           string
}

func NewCreateBAgreementTransaction() *CreateBAgreementTransaction {
	return &CreateBAgreementTransaction{
		requestContainer:  model.NewDefaultBAgreementRequestContainer(),
		responseContainer: model.NewBAgreementResponseContainer(),
	}
}

// ##################################################################
// ##########Transaction INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pt *CreateBAgreementTransaction) GetRequestBody() iface.BodyRoot {
	return pt.requestContainer.GetBody()
}

func (pt *CreateBAgreementTransaction) GetResponseBody() iface.ResponseBodyRoot {
	return pt.responseContainer.GetBody()
}

func (pt *CreateBAgreementTransaction) GetRequestHeader() iface.Header {
	return pt.requestContainer.GetHeader()
}

func (pt *CreateBAgreementTransaction) GetResponseHeader() iface.HeaderResponse {
	return pt.responseContainer.GetHeader()
}

func (pt *CreateBAgreementTransaction) SetRequestBody(body iface.BodyRoot) {
	pt.requestContainer.SetBody(body)
}

func (pt *CreateBAgreementTransaction) SetResponseBody(body iface.ResponseBodyRoot) {
	pt.responseContainer.SetBody(body)
}
func (pt *CreateBAgreementTransaction) SetRequestHeader(header iface.Header) {
	pt.requestContainer.SetHeader(header)
}
func (pt *CreateBAgreementTransaction) SetResponseHeader(header iface.HeaderResponse) {
	pt.responseContainer.SetHeader(header)
}

func (pt *CreateBAgreementTransaction) GetResponseCode() int {
	return pt.responseContainer.GetCode()
}
func (pt *CreateBAgreementTransaction) SetResponseCode(code int) {
	pt.responseContainer.SetCode(code)
}

func (pt *CreateBAgreementTransaction) GetResponseStatus() string {
	return pt.responseContainer.GetStatus()
}
func (pt *CreateBAgreementTransaction) SetResponseStatus(status string) {
	pt.responseContainer.SetStatus(status)
}

func (pt *CreateBAgreementTransaction) GetRequestMethod() string {
	return pt.requestContainer.GetMethod()
}
func (pt *CreateBAgreementTransaction) GetRequestURL() string {
	return pt.requestContainer.GetURL()
}

func (pt *CreateBAgreementTransaction) SetBearerToken(token string) {
	pt.GetRequestHeader().SetBearerToken(token)
}

func (pt *CreateBAgreementTransaction) AssembleRequestBody() {
	//montar objeto para, depois de pronto, fazer pc.requestcontainer.setbody(obj)
	fmt.Println("CHAMANDO TRANSACTION ASSEMBLE PARA CreateBAgreementTransaction")

	root := model.NewBAgreementRoot()

	root.SetTokenID(pt.baToken)

	pt.SetRequestBody(root)

}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
func (pt *CreateBAgreementTransaction) SetBAToken(baToken string) {

	pt.baToken = baToken

	// fmt.Printf("VALOR DE BA TYPE ESTA DEFINIDO: %s\n", pt.baType)
}

func (pt *CreateBAgreementTransaction) GetBAID() string {

	obj := pt.GetResponseBody().(*model.BAgreementResponseRoot)

	fmt.Printf("VALOR DE BA ID VINDO DA RESPOSTA NA TRANSACAO: %s\n", obj.GetId())

	return obj.GetId()
}
