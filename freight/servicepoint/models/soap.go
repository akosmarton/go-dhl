package models

type SoapRequestEnvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SOAPNameSpace string `xml:"xmlns:soap,attr"`
	Body          *SoapRequestBody
}

type SoapRequestBody struct {
	XMLName string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Request interface{}
}

type SoapResponseEnvelopeNearestServicePoints struct {
	XMLName string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    *SoapResponseBodyNearestServicePoints
}

type SoapResponseBodyNearestServicePoints struct {
	XMLName  string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Response *GetNearestServicePointsResponse
	Fault    *SoapFault
}

type SoapResponseEnvelopeServicePointDetail struct {
	XMLName string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    *SoapResponseBodyServicePointDetail
}

type SoapResponseBodyServicePointDetail struct {
	XMLName  string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Response *GetServicePointDetailResponse
	Fault    *SoapFault
}

type SoapFault struct {
	XMLName     string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}
