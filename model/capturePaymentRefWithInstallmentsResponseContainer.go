package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type CapturePaymentRefWithInstallmentsResponseContainer struct {
	Header iface.HeaderResponse
	Body   iface.ResponseBodyRoot
	Status string
	Code   int
}

func NewDefaultCapturePaymentRefWithInstallmentsResponseContainer() *CapturePaymentRefWithInstallmentsResponseContainer {
	var code int
	return &CapturePaymentRefWithInstallmentsResponseContainer{
		Header: NewHeaderForRESTResponse(),
		Body:   NewDefaultPaymentResponseRoot(),
		Status: "NO_STATUS",
		Code:   code,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *CapturePaymentRefWithInstallmentsResponseContainer) GetHeader() iface.HeaderResponse {
	return pc.Header
}

func (pc *CapturePaymentRefWithInstallmentsResponseContainer) GetBody() iface.ResponseBodyRoot {
	return pc.Body
}

func (pc *CapturePaymentRefWithInstallmentsResponseContainer) SetHeader(header iface.HeaderResponse) {
	pc.Header = header
}
func (pc *CapturePaymentRefWithInstallmentsResponseContainer) SetBody(body iface.ResponseBodyRoot) {
	pc.Body = body
}
func (pc *CapturePaymentRefWithInstallmentsResponseContainer) GetStatus() string {
	return pc.Status
}
func (pc *CapturePaymentRefWithInstallmentsResponseContainer) SetStatus(status string) {
	pc.Status = status
}
func (pc *CapturePaymentRefWithInstallmentsResponseContainer) GetCode() int {
	return pc.Code
}

func (pc *CapturePaymentRefWithInstallmentsResponseContainer) SetCode(code int) {
	pc.Code = code
}
