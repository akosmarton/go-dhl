package express

import (
	"encoding/xml"
	"errors"

	log "github.com/sirupsen/logrus"
)

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

// TrackingResponse object
type TrackingResponse struct {
	Response     *Response `xml:"Response,omitempty"`
	AWBInfo      []AWBInfo `xml:"AWBInfo,omitempty"`
	Fault        *Fault    `xml:"Fault,omitempty"`
	LanguageCode string    `xml:"LanguageCode,omitempty"`
}

// AWBInfo response object
type AWBInfo struct {
	AWBNumber    AWBNumber       `xml:"AWBNumber,omitempty"`
	Status       *Status         `xml:"Status,omitempty"`
	ShipmentInfo *ShipmentInfo   `xml:"ShipmentInfo,omitempty"`
	Pieces       *TrackingPieces `xml:"TrackingPieces,omitempty"`
}

// Status response object
type Status struct {
	ActionStatus string     `xml:"ActionStatus,omitempty"`
	Condition    *Condition `xml:"Condition,omitempty"`
}

// Fault response object
type Fault struct {
	PieceFault *PieceFault `xml:"PieceFault,omitempty"`
}

// ShipmentInfo response object
type ShipmentInfo struct {
	OriginServiceArea      *OriginServiceArea      `xml:"OriginServiceArea,omitempty"`
	DestinationServiceArea *DestinationServiceArea `xml:"DestinationServiceArea,omitempty"`
	ShipperName            string                  `xml:"ShipperName,omitempty"`          // max 35 characters
	ShipperAccountNumber   int                     `xml:"ShipperAccountNumber,omitempty"` // positiv integer
	ConsigneeName          string                  `xml:"ConsigneeName,omitempty"`        // max 35 characters
	ShipmentDate           string                  `xml:"ShipmentDate,omitempty"`         // datetime string
	Pieces                 *TrackingPieces         `xml:"Pieces,omitempty"`
	Weight                 string                  `xml:"Weight,omitempty"`
	WeightUnit             string                  `xml:"WeightUnit,omitempty"`     // L K or G
	EstDlvyDate            string                  `xml:"EstDlvyDate,omitempty"`    // datetime string
	EstDlvyDateUTC         string                  `xml:"EstDlvyDateUTC,omitempty"` // datetime string UTC
	GlobalProductCode      string                  `xml:"GlobalProductCode,omitempty"`
	ShipmentDesc           string                  `xml:"ShipmentDesc,omitempty"`
	DlvyNotificationFlag   string                  `xml:"DlvyNotificationFlag,omitempty"` // Y or N
	Shipper                *Shipper                `xml:"Shipper,omitempty"`
	Consignee              *Consignee              `xml:"Consignee,omitempty"`
	ShipmentEvent          []ShipmentEvent         `xml:"ShipmentEvent,omitempty"`
	ShipperReference       Reference               `xml:"ShipperReference,omitempty"`
}

// OriginServiceArea response object
type OriginServiceArea struct {
	ServiceAreaCode  string `xml:"ServiceAreaCode,omitempty"`
	Description      string `xml:"Description,omitempty"`
	OutboundSortCode string `xml:"OutboundSortCode,omitempty"`
}

// DestinationServiceArea response object
type DestinationServiceArea struct {
	ServiceAreaCode string `xml:"ServiceAreaCode,omitempty"`
	Description     string `xml:"Description,omitempty"`
	FacilityCode    string `xml:"FacilityCode,omitempty"`
	InboundSortCode string `xml:"InboundSortCode,omitempty"`
}

// Shipper response object
type Shipper struct {
	ShipperID             string   `xml:"ShipperID,omitempty"` // max 30 characters
	CompanyName           string   `xml:"CompanyName,omitempty"`
	RegisteredAccount     int      `xml:"RegisteredAccount,omitempty"` // positive integer
	AddressLines          []string `xml:"AddressLine,omitempty"`
	City                  string   `xml:"City,omitempty"`
	Division              string   `xml:"Division,omitempty"`
	DivisionCode          string   `xml:"DivisionCode,omitempty"`
	PostalCode            string   `xml:"PostalCode,omitempty"`
	OriginServiceAreaCode string   `xml:"OriginServiceAreaCode,omitempty"`
	OriginFacilityCode    string   `xml:"OriginFacilityCode,omitempty"`
	CountryCode           string   `xml:"CountryCode,omitempty"`
	CountryName           string   `xml:"CountryName,omitempty"`
	FederalTaxID          string   `xml:"FederalTaxID,omitempty"`
	StateTaxID            string   `xml:"StateTaxId,omitempty"`
	Contact               *Contact `xml:"Contact,omitempty"`
}

// Consignee response object
type Consignee struct {
	CompanyName       string   `xml:"CompanyName,omitempty"`
	RegisteredAccount int      `xml:"RegisteredAccount,omitempty"` // positive integer
	AddressLines      []string `xml:"AddressLine,omitempty"`
	City              string   `xml:"City,omitempty"`
	Division          string   `xml:"Division,omitempty"`
	DivisionCode      string   `xml:"DivisionCode,omitempty"`
	PostalCode        string   `xml:"PostalCode,omitempty"`
	CountryCode       string   `xml:"CountryCode,omitempty"`
	CountryName       string   `xml:"CountryName,omitempty"`
	FederalTaxID      string   `xml:"FederalTaxID,omitempty"`
	StateTaxID        string   `xml:"StateTaxId,omitempty"`
	Contact           *Contact `xml:"Contact,omitempty"`
}

// Contact response object
type Contact struct {
	PersonName        string `xml:"PersonName,omitempty"`
	PhoneNumber       string `xml:"PhoneNumber,omitempty"`
	PhoneExtension    string `xml:"PhoneExtension,omitempty"`
	FaxNumber         string `xml:"FaxNumber,omitempty"`
	Telex             string `xml:"Telex,omitempty"`
	Email             *Email `xml:"Email,omitempty"`
	MobilePhoneNumber int    `xml:"MobilePhoneNumber,omitempty"`
}

// Email response object
type Email struct {
	From    string   `xml:"From,omitempty"`
	To      string   `xml:"To,omitempty"`
	CC      []string `xml:"cc,omitempty"`
	Subject string   `xml:"Subject,omitempty"`
	ReplyTo string   `xml:"ReplyTo,omitempty"`
	Body    string   `xml:"Body,omitempty"`
}

// ShipmentEvent response object
type ShipmentEvent struct {
	Date         string        `xml:"Date,omitempty"` // date string
	Time         string        `xml:"Time,omitempty"` // Time
	ServiceEvent *ServiceEvent `xml:"ServiceEvent,omitempty"`
	Signatory    string        `xml:"Signatory,omitempty"`
	EventRemarks *EventRemarks `xml:"EventRemarks,omitempty"`
	ServiceArea  string        `xml:"ServiceArea,omitempty"`
}

// EventRemarks response object
type EventRemarks struct {
	FurtherDetails string `xml:"FurtherDetails,omitempty"`
	NextSteps      string `xml:"NextSteps,omitempty"`
}

// PieceFault response object
type PieceFault struct {
	PieceID       TrackingPieceID `xml:"AWBNumber,omitempty"`
	ConditionCode string          `xml:"ConditionCode,omitempty"`
	ConditionData string          `xml:"ConditionData,omitempty"`
}

// TrackingPieces response object
type TrackingPieces struct {
	PieceInfo []PieceInfo `xml:"PieceInfo,omitempty"`
}

// PieceInfo response object
type PieceInfo struct {
	PieceDetails PieceDetails `xml:"PieceDetails,omitempty"`
	PieceEvent   []PieceEvent `xml:"PieceEvent,omitempty"`
}

// PieceDetails response object
type PieceDetails struct {
	AWBNumber     AWBNumber       `xml:"AWBNumber,omitempty"`
	LicensePlate  TrackingPieceID `xml:"LicensePlate,omitempty"`
	PieceNumber   string          `xml:"PieceNumber,omitempty"`
	ActualDepth   string          `xml:"ActualDepth,omitempty"`
	ActualWidth   string          `xml:"ActualWidth,omitempty"`
	ActualHeight  string          `xml:"ActualHeight,omitempty"`
	ActualWeight  string          `xml:"ActualWeight,omitempty"`
	Depth         string          `xml:"Depth,omitempty"`
	Width         string          `xml:"Width,omitempty"`
	Height        string          `xml:"Height,omitempty"`
	Weight        string          `xml:"Weight,omitempty"`
	PackageType   string          `xml:"PackageType,omitempty"` // EE OD or CP
	DimWeight     string          `xml:"DimWeight,omitempty"`
	WeightUnit    string          `xml:"WeightUnit,omitempty"`
	PieceContents string          `xml:"PieceContents,omitempty"`
}

// PieceEvent response object
type PieceEvent struct {
	Date         string        `xml:"Date,omitempty"` // date string
	Time         string        `xml:"Time,omitempty"` // Time
	ServiceEvent *ServiceEvent `xml:"ServiceEvent,omitempty"`
	Signatory    string        `xml:"Signatory,omitempty"`
	ServiceArea  string        `xml:"ServiceArea,omitempty"`
}

// ServiceEvent response object
type ServiceEvent struct {
	EventCode   string `xml:"EventCode,omitempty"`
	Description string `xml:"Description,omitempty"`
}

// TrackingQuery contains parameters to be used for making a Known or Unknown Tracking request
type TrackingQuery struct {
	LanguageCode     string
	AWBNumbers       []AWBNumber
	LPNumbers        []TrackingPieceID
	LevelOfDetails   string // LAST_CHECK_POINT_ONLY or ALL_CHECK_POINTS
	PiecesEnabled    string // S, B or P
	CountryCode      string // 2 character country id
	AccountNumber    int
	ShipmentDate     *ShipmentDate
	ShipperReference *Reference
}

func (c *dhlExpressClient) determineRequestType(query TrackingQuery) (interface{}, error) {
	var data interface{}

	if len(query.LPNumbers) != 0 || len(query.AWBNumbers) != 0 {
		log.Debugf("Building KnownTrackingRequest")
		data = c.buildKnownTrackingRequest(query)
		return data, nil
	}
	if query.AccountNumber != 0 {
		log.Debugf("Building UnknownTrackingRequest")
		data = c.buildUnknownTrackingRequest(query)
		return data, nil
	}
	return nil, errors.New("Missing LPNumber (KnownTrackingRequest) or AccountNumber (UnknowTrackingRequest)")
}

// Tracking takes a TrackingQuery and makes a Known or Unknown Tracking Request
//
// AWBNumbers, LPNumbers, LevelOfDetails and PiecesEnabled are KnownTrackingRequest specific
// AccountNumber, ShipmentDate and ShipperReference are UnknownTrackingRequest specific
func (c *dhlExpressClient) Tracking(query TrackingQuery) (*TrackingResponse, error) {
	data, err := c.determineRequestType(query)
	if err != nil {
		return nil, err
	}
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}
	log.Debugf("Tracking Response Body: %s\n", string(*contents))

	var trackingResponse TrackingResponse
	if err := xml.Unmarshal(*contents, &trackingResponse); err != nil {
		return nil, err
	}

	return &trackingResponse, nil
}

func (c *dhlExpressClient) buildKnownTrackingRequest(query TrackingQuery) *KnownTrackingRequest {
	sh := NewServiceHeader(c.siteID, c.password)
	return &KnownTrackingRequest{
		XReqNameSpace:   "http://www.dhl.com",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com TrackingRequestKnown.xsd",
		Request: &Request{
			ServiceHeader: &sh,
		},
		LanguageCode:   query.LanguageCode,
		AWBNumber:      query.AWBNumbers,
		LPNumber:       query.LPNumbers,
		LevelOfDetails: query.LevelOfDetails,
		PiecesEnabled:  query.PiecesEnabled,
		CountryCode:    query.CountryCode,
	}
}

func (c *dhlExpressClient) buildUnknownTrackingRequest(query TrackingQuery) *UnknownTrackingRequest {
	sh := NewServiceHeader(c.siteID, c.password)
	return &UnknownTrackingRequest{
		XReqNameSpace:   "http://www.dhl.com",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com TrackingRequestKnown.xsd",
		Request: &Request{
			ServiceHeader: &sh,
		},
		LanguageCode:     query.LanguageCode,
		AccountNumber:    query.AccountNumber,
		ShipperReference: query.ShipperReference,
		CountryCode:      query.CountryCode,
	}
}
