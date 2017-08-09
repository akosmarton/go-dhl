package servicepoint

import (
	"encoding/xml"
	"errors"

	log "github.com/sirupsen/logrus"
)

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

type GetNearestServicePointsResponse struct {
	XMLName       string `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetNearestServicePointsResponse"`
	ServicePoints *ServicePoints
}

func (c *dhlServicePointClient) GetNearestServicePoints(query NearestServicePointsQuery) (*GetNearestServicePointsResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	data := c.buildGetNearestServicePointsRequest(query)
	contents, err := c.fetch("GetNearestServicePoints", data)
	if err != nil {
		return nil, err
	}
	log.Debugf("GetNearestServicePoints Response Body: %s\n", string(*contents))

	var response SoapResponseEnvelopeNearestServicePoints
	if err := xml.Unmarshal(*contents, &response); err != nil {
		return nil, err
	}
	if response.Body.Fault != nil {
		return nil, errors.New(response.Body.Fault.FaultString)
	}

	return response.Body.Response, nil
}

func (c *dhlServicePointClient) buildGetNearestServicePointsRequest(q NearestServicePointsQuery) *SoapRequestEnvelope {
	servicepoint := GetNearestServicePointsRequest{
		CountryCode: q.CountryCode,
		Street:      q.Street,
		PostCode:    q.PostCode,
		City:        q.City,
		FeatureCodes: &FeatureCodes{
			FeatureCode: q.FeatureCodes,
		},
		MaxNumberOfItems: q.MaxNumberOfItems,
		BitCatCodes: &BitCatCodes{
			BitCatCode: q.BitCatCodes,
		},
	}

	return &SoapRequestEnvelope{
		SOAPNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &SoapRequestBody{
			Request: &servicepoint,
		},
	}
}
