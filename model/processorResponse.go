package model

type ProcessorResponse struct {
	ResponseCode string `json:"response_code,omitempty"`
	AVSCode      string `json:"avs_code,omitempty"`
	CVVCode      string `json:"cvv_code,omitempty"`
	AdviceCode   string `json:"advice_code,omitempty"`
	ECISubmitted string `json:"eci_submitted,omitempty"`
	VPAS         string `json:"vpas,omitempty"`
}

func NewDefaultProcessorResponse() *ProcessorResponse {
	return &ProcessorResponse{}
}
func (pr *ProcessorResponse) GetVPAS() string {
	return pr.VPAS
}
func (pr *ProcessorResponse) SetVPAS(value string) {
	pr.VPAS = value
}
func (pr *ProcessorResponse) GetECISubmitted() string {
	return pr.ECISubmitted
}

func (pr *ProcessorResponse) SetECISubmitted(value string) {
	pr.ECISubmitted = value
}
func (pr *ProcessorResponse) GetAdviceCode() string {
	return pr.AdviceCode
}
func (pr *ProcessorResponse) SetAdviceCode(value string) {
	pr.AdviceCode = value
}

func (pr *ProcessorResponse) GetCVVCode() string {
	return pr.CVVCode
}
func (pr *ProcessorResponse) SetCVVCode(value string) {
	pr.CVVCode = value
}
func (pr *ProcessorResponse) GetAVSCode() string {
	return pr.AVSCode
}
func (pr *ProcessorResponse) SetAVSCode(value string) {
	pr.AVSCode = value
}

func (pr *ProcessorResponse) GetResponseCode() string {
	return pr.ResponseCode
}
func (pr *ProcessorResponse) SetResponseCode(value string) {
	pr.ResponseCode = value
}
