package dhl

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
