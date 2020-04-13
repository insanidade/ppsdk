package handler

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	model "github.com/insanidade/ppsdk/model"
)

const paypalSandboxAPIURL string = "https://api.sandbox.paypal.com"
const getKeyURL string = "/v1/oauth2/token"

type authControllerREST struct {
	ClientID string
	Secret   string
}

//NewAuthController returns a new instance of authControllerREST
func NewAuthController(clientID string, secret string) *authControllerREST {
	return &authControllerREST{
		ClientID: clientID,
		Secret:   secret}
}

//####################################
//DoRESTAuth builds an object with oauth data
func (authObj *authControllerREST) DoAuth() *model.OauthResponse {

	url := paypalSandboxAPIURL + getKeyURL

	payload := strings.NewReader("grant_type=client_credentials")

	request, _ := http.NewRequest("POST", url, payload)

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Accept-Language", "en_US")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	stringPair := authObj.ClientID + ":" + authObj.Secret
	b64encoded := b64.StdEncoding.EncodeToString([]byte(stringPair))
	// fmt.Printf("String Base64 encoded: %s\n", b64encoded)
	request.Header.Add("Authorization", "Basic "+b64encoded)

	response, er := http.DefaultClient.Do(request)

	if nil != er {
		log.Fatalf("Erro na requisição: %+v\n", er)
	}

	defer response.Body.Close()
	// io.Copy(os.Stdout, response.Body)
	decoder := json.NewDecoder(response.Body)

	oauthStruct := &model.OauthResponse{}

	err := decoder.Decode(oauthStruct)

	if nil != err {
		log.Fatalf("Erro  de leitura de JSON: %+v", err)
	}

	fmt.Printf("Recebeu: %+v\n", oauthStruct.AccessToken)
	fmt.Printf("Expira em %+v minutos\n", oauthStruct.GetMinutesUntilExpire())

	// if errEnc := encoder.Encode(oauthStruct); nil != errEnc {
	// 	log.Fatalf("Erro em ENCODE: %+v", errEnc)
	// }

	// for name, values := range response.Header {
	// 	fmt.Printf("Recebeu: %+v\n", name)
	// 	fmt.Printf("Recebeu: %+v\n", values)
	// 	wout.Header()[name] = values
	// }

	// wout.WriteHeader(response.StatusCode)
	// io.Copy(wout, response.Body)

	return oauthStruct

}
