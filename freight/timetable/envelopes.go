package timetable

// SoapRequestEnvelope is SOAP envelope for SoapRequestBody
type SoapRequestEnvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	XSINameSpace  string `xml:"xmlns:xsi,attr"`
	XSDNameSpace  string `xml:"xmlns:xsd,attr"`
	SOAPNameSpace string `xml:"xmlns:soap,attr"`
	Body          *SoapRequestBody
}

// SoapRequestBody is SOAP envelope for RequestBody
type SoapRequestBody struct {
	XMLName string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Request *GetTimeTableRequest
}

// SoapResponseEvelope is SOAP envelope for SoapResponseBody
type SoapResponseEvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	XSINameSpace  string `xml:"xmlns:xsi,attr"`
	XSDNameSpace  string `xml:"xmlns:xsd,attr"`
	SOAPNameSpace string `xml:"xmlns:soap,attr"`
	Body          *SoapResponseBody
}

// SoapResponseBody is SOAP envelope for ResponseBody
type SoapResponseBody struct {
	XMLName  string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Response *GetTimeTableResponse
	Fault    *SoapFault
}

// SoapFault is SOAP envelope for Fault
type SoapFault struct {
	XMLName     string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}
