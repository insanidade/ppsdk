package transactions

import (
	"fmt"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type CreateBATokenTransaction struct {
	requestContainer  *model.BATokenRequestContainer
	responseContainer *model.BATokenResponseContainer
	baType            string
}

func NewCreateBATokenTransaction() *CreateBATokenTransaction {
	return &CreateBATokenTransaction{
		requestContainer:  model.NewDefaultBATokenRequestContainer(),
		responseContainer: model.NewBATokenResponseContainer(),
	}
}

// ##################################################################
// ##########Transaction INTERFACE IMPLEMENTATIONS###################
// ##################################################################

func (pt *CreateBATokenTransaction) GetRequestBody() iface.BodyRoot {
	return pt.requestContainer.GetBody()
}

func (pt *CreateBATokenTransaction) GetResponseBody() iface.ResponseBodyRoot {
	return pt.responseContainer.GetBody()
}

func (pt *CreateBATokenTransaction) GetRequestHeader() iface.Header {
	return pt.requestContainer.GetHeader()
}

func (pt *CreateBATokenTransaction) GetResponseHeader() iface.HeaderResponse {
	return pt.responseContainer.GetHeader()
}

func (pt *CreateBATokenTransaction) SetRequestBody(body iface.BodyRoot) {
	pt.requestContainer.SetBody(body)
}

func (pt *CreateBATokenTransaction) SetResponseBody(body iface.ResponseBodyRoot) {
	pt.responseContainer.SetBody(body)
}
func (pt *CreateBATokenTransaction) SetRequestHeader(header iface.Header) {
	pt.requestContainer.SetHeader(header)
}
func (pt *CreateBATokenTransaction) SetResponseHeader(header iface.HeaderResponse) {
	pt.responseContainer.SetHeader(header)
}

func (pt *CreateBATokenTransaction) GetResponseCode() int {
	return pt.responseContainer.GetCode()
}
func (pt *CreateBATokenTransaction) SetResponseCode(code int) {
	pt.responseContainer.SetCode(code)
}

func (pt *CreateBATokenTransaction) GetResponseStatus() string {
	return pt.responseContainer.GetStatus()
}
func (pt *CreateBATokenTransaction) SetResponseStatus(status string) {
	pt.responseContainer.SetStatus(status)
}

func (pt *CreateBATokenTransaction) GetRequestMethod() string {
	return pt.requestContainer.GetMethod()
}
func (pt *CreateBATokenTransaction) GetRequestURL() string {
	return pt.requestContainer.GetURL()
}

func (pt *CreateBATokenTransaction) SetBearerToken(token string) {
	pt.GetRequestHeader().SetBearerToken(token)
}

func (pt *CreateBATokenTransaction) AssembleRequestBody() {
	//montar objeto para, depois de pronto, fazer pc.requestcontainer.setbody(obj)
	fmt.Println("CHAMANDO TRANSACTION ASSEMBLE")
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
func (pt *CreateBATokenTransaction) SetBAType(baType string) {

	pt.baType = baType

	fmt.Printf("VALOR DE BA TYPE ESTA DEFINIDO: %s\n", pt.baType)
}

// func setUpLoggingFile(logFileName string) {
// 	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

// 	if nil != err {
// 		return
// 	}

// 	log.SetOutput(logFile)

// }
