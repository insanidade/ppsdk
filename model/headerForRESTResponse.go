package model

//// TODO: add getters/setters for all default headers

//HeaderForRESTResponse class definition. It implements the Header interface
type HeaderForRESTResponse struct {
	Headers map[string]string
}

//NewHeaderForRESTResponse constructor
func NewHeaderForRESTResponse() *HeaderForRESTResponse {
	headersMap := make(map[string]string)
	return &HeaderForRESTResponse{
		Headers: headersMap,
	}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//AddCustomHeader adds custom header to the header
func (hfr *HeaderForRESTResponse) AddCustomHeader(headerName string, headerValue string) {
	hfr.Headers[headerName] = headerValue
}

//GetHeaders returns headers set so far
func (hfr *HeaderForRESTResponse) GetHeaders() map[string]string {
	return hfr.Headers
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################
