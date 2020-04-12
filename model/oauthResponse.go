package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

//OauthResponse represents the json answer from the server
type OauthResponse struct {
	Scope       string `json:"scope,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	AppID       string `json:"app_id,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	Nonce       string `json:"nonce,omitempty"`
}

//SimpleElements implements the JSONObject SimpleElements function
// func (or *OauthResponse) SimpleElements() map[string]string {
// 	localMap := make(map[string]string)
//
// 	var buf bytes.Buffer
//
// 	encoder := json.NewEncoder(&buf)
//
// 	if errEnc := encoder.Encode(obj); nil != errEnc {
// 		log.Fatalf("Erro em ENCODE: %+v", errEnc)
// 	}
//
// 	fmt.Printf("buffer: %+v\n", &buf)
//
// 	localMap[or.Scope.]
// 	return nil
// }

//AddObject inserts another JSONObject to this JsonObject
func (or *OauthResponse) AddObject(obj iface.JSONObject) {}

//GetMinutesUntilExpire returns the amount of minutes remaining for the token to expire
func (or *OauthResponse) GetMinutesUntilExpire() int {
	return (or.ExpiresIn / 60)
}

func (or *OauthResponse) token() string {
	return or.AccessToken
}
func (or *OauthResponse) expiresIn() int {
	return or.ExpiresIn
}
func (or *OauthResponse) authType() string {
	return or.TokenType
}
