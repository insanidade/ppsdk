package model

// model "github.com/insanidade/ppsdk/model"

type PaymentResponseContainer struct {
	Header *HeaderForRESTResponse
	Body   *PaymentResponseRoot //iface.BodyRoot
	Status string
	Code   int
}

func NewPaymentResponseContainer() *PaymentResponseContainer {
	var code int
	return &PaymentResponseContainer{
		Header: NewHeaderForRESTResponse(),
		Body:   NewPaymentResponseRoot(),
		Status: "NO_STATUS",
		Code:   code,
	}
}

//GetHeader returns the header that is set for this container.
func (pc *PaymentResponseContainer) GetHeader() *HeaderForRESTResponse {
	return pc.Header
}

func (pc *PaymentResponseContainer) GetBody() *PaymentResponseRoot {
	return pc.Body
}

func (pc *PaymentResponseContainer) SetHeader(header *HeaderForRESTResponse) {
	pc.Header = header
}
func (pc *PaymentResponseContainer) SetBody(body *PaymentResponseRoot) {
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
