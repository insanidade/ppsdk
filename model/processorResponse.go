package model

type ProcessorResponse struct {
	ResponseCode string `json:"response_code,omitempty"`
	AVSCode      string `json:"avs_code,omitempty"`
	CVVCode      string `json:"cvv_code,omitempty"`
	AdviceCode   string `json:"advice_code,omitempty"`
	ECISubmitted string `json:"eci_submitted,omitempty"`
	VPAS         string `json:"vpas,omitempty"`
}
