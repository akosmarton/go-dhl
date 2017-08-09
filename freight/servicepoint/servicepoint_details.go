package servicepoint

import (
	"encoding/xml"
	"errors"

	log "github.com/sirupsen/logrus"
)

type ServicePointDetail struct {
	XMLName        string           `xml:"ServicePointDetail"`
	Identity       *ServicePointRef `xml:"http://DHL.ServicePoint.DataContracts/2008/10 Identity,omitempty"`
	ServiceAddress *ServiceAddress  `xml:"http://DHL.ServicePoint.DataContracts/2008/10 ServiceAddress,omitempty"`
	FeatureCodes   *FeatureCodes    `xml:"http://DHL.ServicePoint.DataContracts/2008/10 FeatureCodes,omitempty"`
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

type ServiceAddress struct {
	Address           *PhysicalAddress `xml:"Address,omitempty"`
	ContactPersonName string           `xml:"ContactPersonName,omitempty"`
	Telecom           *TelecomInfo     `xml:"Telecom,omitempty"`
}

type PhysicalAddress struct {
	AddresseeName string `xml:"AddresseeName,omitempty"`
	City          string `xml:"City,omitempty"`
	CountryCode   string `xml:"CountryCode,omitempty"`
	PostCode      string `xml:"PostCode,omitempty"`
	Street1       string `xml:"Street1,omitempty"`
	Street2       string `xml:"Street2,omitempty"`
}

type TelecomInfo struct {
	Email string `xml:"Email,omitempty"`
	Fax   string `xml:"Fax,omitempty"`
	Phone string `xml:"Phone,omitempty"`
}

type ServicePointRef struct {
	ID          string `xml:"Id,omitempty"`
	DisplayName string `xml:"DisplayName,omitempty"`
}

type ServicePoints struct {
	XMLName            string               `xml:"ServicePoints"`
	NearbyServicePoint []NearbyServicePoint `xml:"http://DHL.ServicePoint.DataContracts/2008/10 NearbyServicePoint"`
}

type GetServicePointDetailRequest struct {
	XMLName         string           `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetServicePointDetailRequest"`
	ServicePointRef *ServicePointRef `xml:"http://DHL.ServicePoint.DataContracts/2008/10 ServicePointRef"`
}

type GetServicePointDetailResponse struct {
	XMLName            string `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetServicePointDetailResponse"`
	ServicePointDetail *ServicePointDetail
}

func (c *dhlServicePointClient) GetServicePointDetail(query ServicePointDetailQuery) (*GetServicePointDetailResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	data := c.buildGetServicePointDetailRequest(query)
	contents, err := c.fetch("GetServicePointDetail", data)
	if err != nil {
		return nil, err
	}
	log.Debugf("GetServicePointDetail Response Body: %s\n", string(*contents))

	var response SoapResponseEnvelopeServicePointDetail
	if err := xml.Unmarshal(*contents, &response); err != nil {
		return nil, err
	}
	if response.Body.Fault != nil {
		return nil, errors.New(response.Body.Fault.FaultString)
	}

	return response.Body.Response, nil
}

func (c *dhlServicePointClient) buildGetServicePointDetailRequest(q ServicePointDetailQuery) *SoapRequestEnvelope {
	servicepoint := GetServicePointDetailRequest{
		ServicePointRef: &ServicePointRef{
			ID:          q.ID,
			DisplayName: q.DisplayName,
		},
	}

	return &SoapRequestEnvelope{
		SOAPNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &SoapRequestBody{
			Request: &servicepoint,
		},
	}
}

// NearestServicePointsQuery used for doing GetNearestServicePoints requests
type NearestServicePointsQuery struct {
	CountryCode      string
	Street           string
	PostCode         string
	City             string
	FeatureCodes     []string
	MaxNumberOfItems int
	BitCatCodes      []string
}

// Validate is used to validate query input
func (q *NearestServicePointsQuery) Validate() error {
	if q.CountryCode == "" {
		return errors.New("CountryCode is required")
	}

	if q.PostCode == "" {
		return errors.New("PostCode is required")
	}

	return nil
}

// ServicePointDetailQuery used for doing GetServicePointDetail requests
type ServicePointDetailQuery struct {
	ID          string
	DisplayName string
}

// Validate is used to validate query input
func (q *ServicePointDetailQuery) Validate() error {
	if q.ID == "" {
		return errors.New("ID is required")
	}

	return nil
}
