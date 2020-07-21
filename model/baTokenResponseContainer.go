package model

import (
	iface "github.com/insanidade/ppsdk/interfaces"
)

type BATokenResponseContainer struct {
	Header iface.HeaderResponse
	Body   iface.ResponseBodyRoot
	Status string
	Code   int
}

func NewBATokenResponseContainer() *BATokenResponseContainer {
	var code int
	return &BATokenResponseContainer{
		Header: NewHeaderForRESTResponse(),
		Body:   NewBATokenResponseRoot(),
		Status: "NO_STATUS",
		Code:   code,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *BATokenResponseContainer) GetHeader() iface.HeaderResponse {
	return pc.Header
}

func (pc *BATokenResponseContainer) GetBody() iface.ResponseBodyRoot {
	return pc.Body
}

func (pc *BATokenResponseContainer) SetHeader(header iface.HeaderResponse) {
	pc.Header = header
}
func (pc *BATokenResponseContainer) SetBody(body iface.ResponseBodyRoot) {
	pc.Body = body
}
func (pc *BATokenResponseContainer) GetStatus() string {
	return pc.Status
}
func (pc *BATokenResponseContainer) SetStatus(status string) {
	pc.Status = status
}
func (pc *BATokenResponseContainer) GetCode() int {
	return pc.Code
}

func (pc *BATokenResponseContainer) SetCode(code int) {
	pc.Code = code
}
