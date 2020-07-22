package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

// https://api.sandbox.paypal.com/v1/billing-agreements/agreement-tokens

const paypalSandboxAPIURLCalcFinancingOptions string = "https://api.sandbox.paypal.com"
const CalcFinancingOptionsURL string = "/v1/credit/calculated-financing-options"

type CalculateFinancingOptionsRequestContainer struct {
	Header iface.Header
	Body   iface.BodyRoot
	Method string
	URL    string
}

func NewCalculateFinancingOptionsRequestContainer(
	header iface.Header,
	body iface.BodyRoot) *CalculateFinancingOptionsRequestContainer {
	return &CalculateFinancingOptionsRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLCalcFinancingOptions + CalcFinancingOptionsURL,
	}
}

func NewDefaultCalculateFinancingOptionsRequestContainer() *CalculateFinancingOptionsRequestContainer {
	header := NewHeaderForREST()
	body := NewCalculateFinancingOptionsRoot()
	return &CalculateFinancingOptionsRequestContainer{
		Header: header,
		Body:   body,
		Method: "POST",
		URL:    paypalSandboxAPIURLCalcFinancingOptions + CalcFinancingOptionsURL,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *CalculateFinancingOptionsRequestContainer) GetHeader() iface.Header {
	return pc.Header
}

func (pc *CalculateFinancingOptionsRequestContainer) GetBody() iface.BodyRoot {
	return pc.Body
}
func (pc *CalculateFinancingOptionsRequestContainer) GetMethod() string {
	return pc.Method
}
func (pc *CalculateFinancingOptionsRequestContainer) GetURL() string {
	return pc.URL
}
func (pc *CalculateFinancingOptionsRequestContainer) SetHeader(header iface.Header) {
	pc.Header = header
}
func (pc *CalculateFinancingOptionsRequestContainer) SetBody(body iface.BodyRoot) {
	pc.Body = body
}
func (pc *CalculateFinancingOptionsRequestContainer) SetMethod(method string) {
	pc.Method = method
}
func (pc *CalculateFinancingOptionsRequestContainer) SetURL(url string) {
	pc.URL = url
}
