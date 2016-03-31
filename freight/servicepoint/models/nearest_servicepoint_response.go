package models

type GetNearestServicePointsResponse struct {
	XMLName       string         `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetNearestServicePointsResponse"`
	ServicePoints *ServicePoints `xml:"http://DHL.ServicePoint.DataContracts/2008/10 ServicePoints"`
}

type ServicePoints struct {
	XMLName            string               `xml:"http://DHL.ServicePoint.DataContracts/2008/10 ServicePoints"`
	NearbyServicePoint []NearbyServicePoint `xml:"http://DHL.ServicePoint.DataContracts/2008/10 NearbyServicePoint"`
}

type NearbyServicePoint struct {
	XMLName       string           `xml:"http://DHL.ServicePoint.DataContracts/2008/10 NearbyServicePoint"`
	Identity      *ServicePointRef `xml:"Identity,omitempty"`
	Distance      float64          `xml:"Distance,omitempty"`
	StreetName    string           `xml:"StreetName,omitempty"`
	PostCode      string           `xml:"PostCode,omitempty"`
	City          string           `xml:"City,omitempty"`
	FeatureCodes  *FeatureCodes    `xml:"FeatureCodes,omitempty"`
	RouteDistance float64          `xml:"RouteDistance,omitempty"`
}

type ServicePointRef struct {
	ID          string `xml:"Id,omitempty"`
	DisplayName string `xml:"DisplayName,omitempty"`
}

type SoapResponseEvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SOAPNameSpace string `xml:"xmlns:s,attr"`
	Body          *SoapResponseBody
}

type SoapResponseBody struct {
	XMLName  string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Response *GetNearestServicePointsResponse
	Fault    *SoapFault
}

type SoapFault struct {
	XMLName     string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}
