package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
)

type hTTPRESTHandler struct {
	requestContainer  iface.RequestContainer
	responseContainer iface.ResponseContainer
	statusCode        int
	request           *http.Request
}

//NewHTTPRESTHandler returns a new instance of the http REST request handler
func NewHTTPRESTHandler(requestContainer iface.RequestContainer) *hTTPRESTHandler {
	return &hTTPRESTHandler{
		requestContainer: requestContainer,
	}
}

//DoRequest performs the request call.
func (pc *hTTPRESTHandler) DoRequest() iface.BodyRoot {

	pc.assemble()
	// log.Printf("Chamando DoRequest para o request: %+v\n", pc.request.Body)
	// log.Println("#####################################")

	response, er := http.DefaultClient.Do(pc.request)

	if nil != er {
		trace()
		log.Fatalf("Erro na requisição HTTP: %+v\n", er)
	}
	defer response.Body.Close()

	pc.statusCode = response.StatusCode

	log.Printf("Recebeu HTTP Status code: %d\n", response.StatusCode)
	log.Println("#####################################")
	for headerName, headerValue := range response.Header {
		log.Printf("Header na resposta[%s : %s]\n", headerName, strings.Join(headerValue, ","))
		log.Println("#####################################")
	}
	//#############################################

	decoder := json.NewDecoder(response.Body)

	responseBody := &model.PaymentResponseRoot{}
	errr := decoder.Decode(responseBody)

	if nil != errr {
		trace()
		log.Fatalf("Erro  de leitura de JSON: %s", errr.Error())
	}

	log.Printf("Recebeu: %+v\n", *responseBody)
	log.Println("#####################################")

	return responseBody
}

//assemble assembles the request. Implementation from interface Controller
func (pc *hTTPRESTHandler) assemble() {

	log.Printf("##### ASSEMBLE CONTAINER #####:\n")

	header := pc.requestContainer.GetHeader()
	body := pc.requestContainer.GetBody()
	method := pc.requestContainer.GetMethod()
	url := pc.requestContainer.GetURL()

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if errEnc := encoder.Encode(body); nil != errEnc {
		log.Fatalf("Erro em ENCODE no Assemble: %+v", errEnc)
	}
	log.Printf("buffer em Assemble: %+v\n", &buf)
	log.Println("#####################################")

	pc.request, _ = http.NewRequest(method, url, &buf)

	for headerName, headerValue := range header.GetHeaders() {
		log.Printf("Adicionando header[%s : %s]\n", headerName, headerValue)
		log.Println("#####################################")
		pc.request.Header.Add(headerName, headerValue)
	}

}

func (hrh *hTTPRESTHandler) GetRequestContainer() iface.RequestContainer {
	return hrh.requestContainer
}

func (hrh *hTTPRESTHandler) GetStatusCode() int {
	return hrh.statusCode
}

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
