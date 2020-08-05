package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type PaymentResponseContainer struct {
	Header iface.HeaderResponse
	Body   iface.ResponseBodyRoot
	Status string
	Code   int
}

func NewPaymentResponseContainer() *PaymentResponseContainer {
	var code int
	return &PaymentResponseContainer{
		Header: NewHeaderForRESTResponse(),
		Body:   NewDefaultPaymentResponseRoot(),
		Status: "NO_STATUS",
		Code:   code,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *PaymentResponseContainer) GetHeader() iface.HeaderResponse {
	return pc.Header
}

func (pc *PaymentResponseContainer) GetBody() iface.ResponseBodyRoot {
	return pc.Body
}

func (pc *PaymentResponseContainer) SetHeader(header iface.HeaderResponse) {
	pc.Header = header
}
func (pc *PaymentResponseContainer) SetBody(body iface.ResponseBodyRoot) {
	pc.Body = body
}
func (pc *PaymentResponseContainer) GetStatus() string {
	return pc.Status
}
func (pc *PaymentResponseContainer) SetStatus(status string) {
	pc.Status = status
}
func (pc *PaymentResponseContainer) GetCode() int {
	return pc.Code
}

func (pc *PaymentResponseContainer) SetCode(code int) {
	pc.Code = code
}
