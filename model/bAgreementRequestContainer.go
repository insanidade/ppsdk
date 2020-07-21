package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

// https://api.sandbox.paypal.com/v1/billing-agreements/agreement-tokens

const paypalSandboxAPIURLBAgreement string = "https://api.sandbox.paypal.com"
const baAgreementURL string = "/v1/billing-agreements/agreements"

type BAgreementRequestContainer struct {
	Header iface.Header
	Body   iface.BodyRoot
	Method string
	URL    string
}

func NewBAgreementRequestContainer(
	header iface.Header,
	body iface.BodyRoot) *BAgreementRequestContainer {
	return &BAgreementRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLBAgreement + baAgreementURL,
	}
}

func NewDefaultBAgreementRequestContainer() *BAgreementRequestContainer {
	header := NewHeaderForREST()
	body := NewBAgreementRoot()
	return &BAgreementRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLBAgreement + baAgreementURL,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *BAgreementRequestContainer) GetHeader() iface.Header {
	return pc.Header
}

func (pc *BAgreementRequestContainer) GetBody() iface.BodyRoot {
	return pc.Body
}
func (pc *BAgreementRequestContainer) GetMethod() string {
	return pc.Method
}
func (pc *BAgreementRequestContainer) GetURL() string {
	return pc.URL
}
func (pc *BAgreementRequestContainer) SetHeader(header iface.Header) {
	pc.Header = header
}
func (pc *BAgreementRequestContainer) SetBody(body iface.BodyRoot) {
	pc.Body = body
}
func (pc *BAgreementRequestContainer) SetMethod(method string) {
	pc.Method = method
}
func (pc *BAgreementRequestContainer) SetURL(url string) {
	pc.URL = url
}
