package express

import (
	"encoding/xml"
)

// RouteRequest object
type RouteRequest struct {
	XMLName           string   `xml:"ns1:RouteRequest"`
	XNS1NameSpace     string   `xml:"xmlns:ns1,attr"`
	XXSINameSpace     string   `xml:"xmlns:xsi,attr"`
	XSchemaLocation   string   `xml:"xsi:schemaLocation,attr"`
	Request           *Request `xml:"Request,omitempty"`
	RegionCode        string   `xml:"RegionCode,omitempty"`  // AP EU AM
	RequestType       string   `xml:"RequestType,omitempty"` // O D
	Address1          string   `xml:"Address1,omitempty"`
	Address2          string   `xml:"Address2,omitempty"`
	Address3          string   `xml:"Address3,omitempty"`
	PostalCode        string   `xml:"PostalCode,omitempty"`
	City              string   `xml:"City,omitempty"`
	Division          string   `xml:"Division,omitempty"`
	CountryCode       string   `xml:"CountryCode,omitempty"`
	CountryName       string   `xml:"CountryName,omitempty"`
	OriginCountryCode string   `xml:"OriginCountryCode,omitempty"`
	SchemaVersion     float32  `xml:"schemaVersion,attr,omitempty"` // 1.0
}

type RouteResponse struct {
	Response             *Response
	Note                 *Note
	RegionCode           string
	GMTNegativeIndicator string
	GMTOffset            string
	ServiceArea          *ServiceArea
}

type Note struct {
	ActionNote string
	Condition  *Condition
}

type ServiceArea struct {
	ServiceAreaCode string
	Description     string
}

// RouteQuery contains parameters to be used for making a Routing request
type RouteQuery struct {
	RegionCode        string // AM EU AP
	RequestType       string // O D
	Address1          string
	Address2          string
	Address3          string
	PostalCode        string
	City              string
	CountryCode       string
	OriginCountryCode string
	CountryName       string
	Division          string
}

func (c *dhlExpressClient) Routing(query RouteQuery) (*RouteResponse, error) {
	data := c.buildRoutingRequest(query)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	var routeResponse RouteResponse
	if err := xml.Unmarshal(*contents, &routeResponse); err != nil {
		return nil, err
	}

	return &routeResponse, nil
}

func (c *dhlExpressClient) buildRoutingRequest(query RouteQuery) *RouteRequest {
	sh := NewServiceHeader(c.siteID, c.password)
	return &RouteRequest{
		XNS1NameSpace:   "http://www.dhl.com",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com routing-req.xsd",
		SchemaVersion:   1.0,
		Request: &Request{
			ServiceHeader: &sh,
		},
		RegionCode:        query.RegionCode,
		RequestType:       query.RequestType,
		Address1:          query.Address1,
		Address2:          query.Address2,
		Address3:          query.Address3,
		PostalCode:        query.PostalCode,
		City:              query.City,
		CountryCode:       query.CountryCode,
		OriginCountryCode: query.OriginCountryCode,
		CountryName:       query.CountryName,
		Division:          query.Division,
	}
}
