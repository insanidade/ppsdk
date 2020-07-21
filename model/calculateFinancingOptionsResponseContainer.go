package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type CalculateFinancingOptionsResponseContainer struct {
	Header iface.HeaderResponse
	Body   iface.ResponseBodyRoot
	Status string
	Code   int
}

func NewCalculateFinancingOptionsResponseContainer() *CalculateFinancingOptionsResponseContainer {
	var code int
	return &CalculateFinancingOptionsResponseContainer{
		Header: NewHeaderForRESTResponse(),
		Body:   NewCalculateFinancingOptionsResponseRoot(),
		Status: "NO_STATUS",
		Code:   code,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *CalculateFinancingOptionsResponseContainer) GetHeader() iface.HeaderResponse {
	return pc.Header
}

func (pc *CalculateFinancingOptionsResponseContainer) GetBody() iface.ResponseBodyRoot {
	return pc.Body
}

func (pc *CalculateFinancingOptionsResponseContainer) SetHeader(header iface.HeaderResponse) {
	pc.Header = header
}
func (pc *CalculateFinancingOptionsResponseContainer) SetBody(body iface.ResponseBodyRoot) {
	pc.Body = body
}
func (pc *CalculateFinancingOptionsResponseContainer) GetStatus() string {
	return pc.Status
}
func (pc *CalculateFinancingOptionsResponseContainer) SetStatus(status string) {
	pc.Status = status
}
func (pc *CalculateFinancingOptionsResponseContainer) GetCode() int {
	return pc.Code
}

func (pc *CalculateFinancingOptionsResponseContainer) SetCode(code int) {
	pc.Code = code
}
