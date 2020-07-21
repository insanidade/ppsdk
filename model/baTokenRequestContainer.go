package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

// https://api.sandbox.paypal.com/v1/billing-agreements/agreement-tokens

const paypalSandboxAPIURLBA string = "https://api.sandbox.paypal.com"
const baTokenURL string = "/v1/billing-agreements/agreement-tokens"

type BATokenRequestContainer struct {
	Header iface.Header
	Body   iface.BodyRoot
	Method string
	URL    string
}

func NewBATokenRequestContainer(
	header iface.Header,
	body iface.BodyRoot) *BATokenRequestContainer {
	return &BATokenRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLBA + baTokenURL,
	}
}

func NewDefaultBATokenRequestContainer() *BATokenRequestContainer {
	header := NewHeaderForREST()
	body := NewBATokenRoot()
	return &BATokenRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLBA + baTokenURL,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *BATokenRequestContainer) GetHeader() iface.Header {
	return pc.Header
}

func (pc *BATokenRequestContainer) GetBody() iface.BodyRoot {
	return pc.Body
}
func (pc *BATokenRequestContainer) GetMethod() string {
	return pc.Method
}
func (pc *BATokenRequestContainer) GetURL() string {
	return pc.URL
}
func (pc *BATokenRequestContainer) SetHeader(header iface.Header) {
	pc.Header = header
}
func (pc *BATokenRequestContainer) SetBody(body iface.BodyRoot) {
	pc.Body = body
}
func (pc *BATokenRequestContainer) SetMethod(method string) {
	pc.Method = method
}
func (pc *BATokenRequestContainer) SetURL(url string) {
	pc.URL = url
}
