package dhl

// KnownTrackingRequest request object
type KnownTrackingRequest struct {
	XMLName         string            `xml:"req:KnownTrackingRequest"`
	XReqNameSpace   string            `xml:"xmlns:req,attr"`
	XXSINameSpace   string            `xml:"xmlns:xsi,attr"`
	XSchemaLocation string            `xml:"xsi:schemaLocation,attr"`
	Request         *Request          `xml:"Request,omitempty"`
	LanguageCode    string            `xml:"LanguageCode,omitempty"`
	AWBNumber       []AWBNumber       `xml:"AWBNumber,omitempty"`
	LPNumber        []TrackingPieceID `xml:"LPNumber,omitempty"`
	LevelOfDetails  string            `xml:"LevelOfDetails,omitempty"` // LAST_CHECK_POINT_ONLY or ALL_CHECK_POINTS
	PiecesEnabled   string            `xml:"PiecesEnabled,omitempty"`  // S, B or P
	CountryCode     string            `xml:"CountryCode,omitempty"`
}

// UnknownTrackingRequest request object
type UnknownTrackingRequest struct {
	XMLName          string     `xml:"req:UnknownTrackingRequest"`
	XReqNameSpace    string     `xml:"xmlns:req,attr"`
	XXSINameSpace    string     `xml:"xmlns:xsi,attr"`
	XSchemaLocation  string     `xml:"xsi:schemaLocation,attr"`
	Request          *Request   `xml:"Request,omitempty"`
	LanguageCode     string     `xml:"LanguageCode,omitempty"`
	AccountNumber    int        `xml:"AccountNumber,omitempty"`
	ShipperReference *Reference `xml:"ShipperReference,omitempty"`
	CountryCode      string     `xml:"CountryCode,omitempty"`
}

//AWBNumber or Airway bill number is 11 character long string used for tracking
type AWBNumber string

//TrackingPieceID 1 to 35 character long string used for tracking
type TrackingPieceID string

// Reference request object
type Reference struct {
	ReferenceID   string `xml:"ReferenceID,omitempty"`
	ReferenceType string `xml:"ReferenceType,omitempty"` // 2-3 characters
}

// ShipmentDate request object
type ShipmentDate struct {
	ShipmentDateFrom string `xml:"ShipmentDateFrom,omitempty"`
	ShipmentDateTo   string `xml:"ShipmentDateFrom,omitempty"`
}
