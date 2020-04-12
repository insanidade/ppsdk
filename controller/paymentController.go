package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

// const paypalSandboxAPIURL string = "https://api.sandbox.paypal.com"
const createPaymentURL string = "/v1/payments/payment"

type PaymentController struct {
	Body    *model.PaymentRoot
	Header  *model.HeaderForREST
	request *http.Request
}

func NewPaymentController(token string) *PaymentController {
	header := model.NewHeaderForREST()
	header.SetBearerToken(token)
	setUpLoggingFile("payment.log")
	return &PaymentController{
		Body:   model.NewPaymentRoot(),
		Header: header,
	}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//DoRequest performs the request call.
func (pc *PaymentController) DoRequest() iface.BodyRoot {
	log.Printf("Chamando DoRequest para o request: %+v\n", pc.request.Body)
	log.Println("#####################################")

	response, er := http.DefaultClient.Do(pc.request)

	if nil != er {
		log.Fatalf("Erro na requisição: %+v\n", er)
	}
	defer response.Body.Close()

	log.Printf("Recebeu HTTP Status code: %d\n", response.StatusCode)
	log.Println("#####################################")
	for headerName, headerValue := range response.Header {
		log.Printf("Header na resposta[%s : %s]\n", headerName, strings.Join(headerValue, ","))
		log.Println("#####################################")
	}

	// bodyBytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// log.Printf("Recebeu response.Body: %s\n", bodyString)
	// log.Println("#####################################")

	decoder := json.NewDecoder(response.Body)

	// log.Println("######### T R A I L E R S ###############")
	// for trailerName, trailerValue := range response.Trailer {
	// 	log.Printf("Trailer na resposta[%s : %s]\n", trailerName, strings.Join(trailerValue, ","))
	// 	log.Println("#####################################")
	// }

	responseBody := &model.PaymentResponseRoot{}
	errr := decoder.Decode(responseBody)

	if nil != errr {
		log.Fatalf("Erro  de leitura de JSON: %s", errr.Error())
	}

	log.Printf("Recebeu: %+v\n", responseBody)
	log.Println("#####################################")

	return responseBody
}

//Assemble assembles the request. Implementation from interface Controller
func (pc *PaymentController) Assemble(h iface.Header, br iface.BodyRoot) {
	url := paypalSandboxAPIURL + createPaymentURL

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if errEnc := encoder.Encode(br); nil != errEnc {
		log.Fatalf("Erro em ENCODE no Assemble: %+v", errEnc)
	}
	log.Printf("buffer em Assemble: %+v\n", &buf)
	log.Println("#####################################")

	pc.request, _ = http.NewRequest("POST", url, &buf)

	for headerName, headerValue := range h.GetHeaders() {
		log.Printf("Adicionando header[%s : %s]\n", headerName, headerValue)
		log.Println("#####################################")
		pc.request.Header.Add(headerName, headerValue)
	}

}

//HeaderFactory is almost useless here (I agree) but maybe
//we will see other REST specialized headers in the future and
//that's why I still keep this function in the interface
func (pc *PaymentController) HeaderFactory() *model.HeaderForREST {
	for headerName, headerValue := range pc.Header.GetHeaders() {
		log.Printf("Listando headers:%s:%s\n", headerName, headerValue)
		log.Println("#####################################")
	}
	return pc.Header
}

//BodyFactory builds a PaymentRoot object and returns it.
//#################################################################
func (pc *PaymentController) BodyFactory() *model.PaymentRoot {
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
		"odefranca invoice - SHALL BE RANDOM",
		"soft desc",
		"odefranca notify url")
	transaction.SetAmount(amount)
	transaction.SetItemList(itemList)
	fmt.Printf("Transaction: %+v", transaction)

	appContext := model.NewApplicationContext("odefranca brand name",
		"Login",
		"odefranca locale",
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

	return paymentBody

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
