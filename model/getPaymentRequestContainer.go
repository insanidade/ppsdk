package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

// const paypalSandboxAPIURL string = "https://api.sandbox.paypal.com"
const getPaymentURL string = "/v1/payments/payment"

type GetPaymentRequestContainer struct {
	Header iface.Header
	Body   iface.BodyRoot
	Method string
	URL    string
}

func NewGetPaymentRequestContainer(payid string) *GetPaymentRequestContainer {
	return &GetPaymentRequestContainer{
		Header: NewHeaderForREST(),
		Body:   NewEmptyBody(),
		Method: "GET",
		URL:    paypalSandboxAPIURL + getPaymentURL + "/" + payid,
	}
}

func (gprc *GetPaymentRequestContainer) GetHeader() iface.Header {
	return gprc.Header
}
func (gprc *GetPaymentRequestContainer) GetBody() iface.BodyRoot {
	return gprc.Body
}
func (gprc *GetPaymentRequestContainer) GetMethod() string {
	return gprc.Method
}
func (gprc *GetPaymentRequestContainer) GetURL() string {
	return gprc.URL
}
func (gprc *GetPaymentRequestContainer) SetHeader(header iface.Header) {
	gprc.Header = header
}
func (gprc *GetPaymentRequestContainer) SetBody(body iface.BodyRoot) {
	gprc.Body = body
}
func (gprc *GetPaymentRequestContainer) SetMethod(method string) {
	gprc.Method = method
}
func (gprc *GetPaymentRequestContainer) SetURL(url string) {
	gprc.URL = url
}
