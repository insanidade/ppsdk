package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type BAgreementResponseContainer struct {
	Header iface.HeaderResponse
	Body   iface.ResponseBodyRoot
	Status string
	Code   int
}

func NewBAgreementResponseContainer() *BAgreementResponseContainer {
	var code int
	return &BAgreementResponseContainer{
		Header: NewHeaderForRESTResponse(),
		Body:   NewBAgreementResponseRoot(),
		Status: "NO_STATUS",
		Code:   code,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *BAgreementResponseContainer) GetHeader() iface.HeaderResponse {
	return pc.Header
}

func (pc *BAgreementResponseContainer) GetBody() iface.ResponseBodyRoot {
	return pc.Body
}

func (pc *BAgreementResponseContainer) SetHeader(header iface.HeaderResponse) {
	pc.Header = header
}
func (pc *BAgreementResponseContainer) SetBody(body iface.ResponseBodyRoot) {
	pc.Body = body
}
func (pc *BAgreementResponseContainer) GetStatus() string {
	return pc.Status
}
func (pc *BAgreementResponseContainer) SetStatus(status string) {
	pc.Status = status
}
func (pc *BAgreementResponseContainer) GetCode() int {
	return pc.Code
}

func (pc *BAgreementResponseContainer) SetCode(code int) {
	pc.Code = code
}
