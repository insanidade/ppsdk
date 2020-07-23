package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"

	iface "github.com/insanidade/ppsdk/interfaces"
)

type hTTPRESTHandler struct {
	transaction iface.Transaction
	statusCode  int
	request     *http.Request
}

//NewHTTPRESTHandler returns a new instance of the http REST request handler
func NewHTTPRESTHandler(transaction iface.Transaction) *hTTPRESTHandler {
	return &hTTPRESTHandler{
		transaction: transaction}
}

//DoRequest performs the request call.
//iface.ResponseContainer
func (pc *hTTPRESTHandler) DoRequest() {

	pc.configureRequest()

	response, er := http.DefaultClient.Do(pc.request)

	if nil != er {
		trace()
		log.Fatalf("Erro na requisição HTTP: %+v\n", er)
	}
	defer response.Body.Close()

	pc.statusCode = response.StatusCode
	pc.transaction.SetResponseStatus(response.Status)
	pc.transaction.SetResponseCode(response.StatusCode)

	log.Printf("Recebeu HTTP Status code: %d\n", pc.transaction.GetResponseCode())
	log.Println("#####################################")

	responseHeader := pc.transaction.GetResponseHeader()
	for headerName, headerValue := range response.Header {
		responseHeader.AddCustomHeader(headerName, strings.Join(headerValue, ","))
		log.Printf("Header na resposta[%s : %s]\n", headerName, strings.Join(headerValue, ","))
		log.Println("#####################################")
	}
	pc.transaction.SetResponseHeader(responseHeader)
	//#############################################

	decoder := json.NewDecoder(response.Body)
	//
	responseBody := pc.transaction.GetResponseBody()
	errr := decoder.Decode(responseBody)
	//
	if nil != errr {
		trace()
		log.Fatalf("Erro  de leitura de JSON: %s", errr.Error())
	}
	//
	log.Printf("Recebeu: %+v\n", responseBody)
	log.Println("#####################################")
	//
	pc.transaction.SetResponseBody(responseBody)

	// return pc.responseContainer

}

//configureRequest assembles the request. Implementation from interface Controller
func (pc *hTTPRESTHandler) configureRequest() {

	log.Printf("##### ASSEMBLE CONTAINER #####:\n")

	header := pc.transaction.GetRequestHeader()
	body := pc.transaction.GetRequestBody()
	method := pc.transaction.GetRequestMethod()
	url := pc.transaction.GetRequestURL()

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

	//#########################################
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(pc.request, true)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("############### REQUEST QUE FORMOU: %s\n", string(requestDump))
	fmt.Println(string(requestDump))
	//#########################################

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
