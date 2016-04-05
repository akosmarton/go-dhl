package models

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
