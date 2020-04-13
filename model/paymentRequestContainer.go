package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type PaymentRequestContainer struct {
	Header iface.Header
	Body   iface.BodyRoot
	Method string
	URL    string
}

func NewPaymentRequestContainer(
	header *HeaderForREST,
	body *PaymentRoot,
	url string) *PaymentRequestContainer {
	return &PaymentRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    url,
	}
}

func (pc *PaymentRequestContainer) GetHeader() iface.Header {
	return pc.Header
}

func (pc *PaymentRequestContainer) GetBody() iface.BodyRoot {
	return pc.Body
}
func (pc *PaymentRequestContainer) GetMethod() string {
	return pc.Method
}
func (pc *PaymentRequestContainer) GetURL() string {
	return pc.URL
}
func (pc *PaymentRequestContainer) SetHeader(header iface.Header) {
	pc.Header = header
}
func (pc *PaymentRequestContainer) SetBody(body iface.BodyRoot) {
	pc.Body = body
}
func (pc *PaymentRequestContainer) SetMethod(method string) {
	pc.Method = method
}
func (pc *PaymentRequestContainer) SetURL(url string) {
	pc.URL = url
}
