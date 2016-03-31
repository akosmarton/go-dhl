package models

type GetNearestServicePointsRequest struct {
	XMLName          string        `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetNearestServicePointsRequest"`
	CountryCode      string        `xml:"CountryCode,omitempty"`
	Street           string        `xml:"Street,omitempty"`
	PostCode         string        `xml:"PostCode,omitempty"`
	City             string        `xml:"City,omitempty"`
	FeatureCodes     *FeatureCodes `xml:"FeatureCodes,omitempty"`
	MaxNumberOfItems int           `xml:"MaxNumberOfItems,omitempty"`
	BitCatCodes      *BitCatCodes  `xml:"BitCatCodes,omitempty"`
}

type FeatureCodes struct {
	FeatureCode []string `xml:"http://DHL.ServicePoint.DataContracts/2008/10 FeatureCode,omitempty"`
}

type BitCatCodes struct {
	BitCatCode []string `xml:"http://DHL.ServicePoint.DataContracts/2008/10 BitCatCode,omitempty"`
}

type SoapRequestEnvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SOAPNameSpace string `xml:"xmlns:soap,attr"`
	Body          *SoapRequestBody
}

type SoapRequestBody struct {
	XMLName string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Request *GetNearestServicePointsRequest
}
