package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

// https://api.sandbox.paypal.com/v1/billing-agreements/agreement-tokens

const paypalSandboxAPIURLCapturePaymentRefWithInstallments string = "https://api.sandbox.paypal.com"
const CapturePaymentRefWithInstallmentsURL string = "/v1/payments/payment"

type CapturePaymentRefWithInstallmentsRequestContainer struct {
	Header iface.Header
	Body   iface.BodyRoot
	Method string
	URL    string
}

func NewCapturePaymentRefWithInstallmentsRequestContainer(
	header iface.Header,
	body iface.BodyRoot) *CapturePaymentRefWithInstallmentsRequestContainer {
	return &CapturePaymentRefWithInstallmentsRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLCapturePaymentRefWithInstallments + CapturePaymentRefWithInstallmentsURL,
	}
}

func NewDefaultCapturePaymentRefWithInstallmentsRequestContainer() *CapturePaymentRefWithInstallmentsRequestContainer {
	header := NewHeaderForREST()
	body := NewDefaultPaymentRoot()
	return &CapturePaymentRefWithInstallmentsRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLCapturePaymentRefWithInstallments + CapturePaymentRefWithInstallmentsURL,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) GetHeader() iface.Header {
	return pc.Header
}

func (pc *CapturePaymentRefWithInstallmentsRequestContainer) GetBody() iface.BodyRoot {
	return pc.Body
}
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) GetMethod() string {
	return pc.Method
}
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) GetURL() string {
	return pc.URL
}
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) SetHeader(header iface.Header) {
	pc.Header = header
}
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) SetBody(body iface.BodyRoot) {
	pc.Body = body
}
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) SetMethod(method string) {
	pc.Method = method
}
func (pc *CapturePaymentRefWithInstallmentsRequestContainer) SetURL(url string) {
	pc.URL = url
}
