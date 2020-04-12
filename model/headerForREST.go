package model

//// TODO: add getters/setters for all default headers

//HeaderForREST class definition. It implements the Header interface
type HeaderForREST struct {
	Headers           map[string]string
	BearerToken       string
	NegativeTestValue string
}

//NewHeaderForREST constructor
func NewHeaderForREST() *HeaderForREST {
	headersMap := make(map[string]string)
	headersMap["Accept"] = "application/json"
	headersMap["Accept-Language"] = "en_US"
	headersMap["Content-Type"] = "application/json"
	return &HeaderForREST{
		Headers: headersMap,
	}
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

//AddCustomHeader adds custom header to the header
func (hfr *HeaderForREST) AddCustomHeader(headerName string, headerValue string) {
	hfr.Headers[headerName] = headerValue
}

//GetHeaders returns headers set so far
func (hfr *HeaderForREST) GetHeaders() map[string]string {
	return hfr.Headers
}

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################

//SetNegativeTest adds values for negative test simulations.
//Example of valid value: {"mock_application_codes": "INSTRUMENT_DECLINED"}
func (hfr *HeaderForREST) SetNegativeTest(negativeTestValue string) {
	hfr.Headers["PayPal-Mock-Response"] = negativeTestValue
	hfr.NegativeTestValue = negativeTestValue
}

func (hfr *HeaderForREST) RemoveNegativeTest() {
	delete(hfr.Headers, "PayPal-Mock-Response")
}

func (hfr *HeaderForREST) SetBearerToken(token string) {
	hfr.Headers["Authorization"] = "Bearer " + token
	hfr.BearerToken = token
}
