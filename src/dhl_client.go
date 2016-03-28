package dhl

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var hosts = map[string]string{
	"staging":    "xmlpitest-ea.dhl.com",
	"production": "xmlpi-ea.dhl.com",
}

// Client interface
type Client interface {
	GetQuote(*DCTFrom, *DCTTo, *BkgDetailsRequest, *DCTDutiable) (*DCTResponse, error)
	GetCapability(*DCTFrom, *DCTTo, *BkgDetailsRequest, *DCTDutiable) (*DCTResponse, error)
	Tracking(TrackingQuery) (*TrackingResponse, error)
	Routing()
}

// ClientConfig object used to configure DHLClient
type ClientConfig struct {
	Debug bool
	Host  string
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

// NewDHLClient returns a DHLClient against DHL Production Environment
func NewDHLClient(siteID string, password string, config ClientConfig) (Client, error) {
	if siteID == "" {
		return nil, errors.New("siteID missing")
	}
	if password == "" {
		return nil, errors.New("password missing")
	}

	host := "production"
	if config.Host != "" {
		host = config.Host
	}

	return &dhlClient{
		baseURL:    hosts[host],
		httpClient: &http.Client{},
		debug:      config.Debug,
		siteID:     siteID,
		password:   password,
	}, nil
}

type dhlClient struct {
	baseURL    string
	httpClient *http.Client
	debug      bool
	siteID     string
	password   string
}

// GetQuote takes DCTFrom, DCTTo, BkgDetailsRequest and DCTDutiable and makes DCTRequest with GetQuote.
// Returns a DCTResponse with GetQuoteResponse
func (c *dhlClient) GetQuote(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) (*DCTResponse, error) {
	data := c.buildQuoteRequest(from, to, details, dutiable)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetQuote Response Body: %s\n", string(*contents))
	}

	var dctResponse DCTResponse
	if err := xml.Unmarshal(*contents, &dctResponse); err != nil {
		return nil, err
	}

	return &dctResponse, nil
}

func (c *dhlClient) buildQuoteRequest(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) *DCTRequest {
	sh := NewServiceHeader(c.siteID, c.password)
	q := GetQuote{
		Request: &Request{
			ServiceHeader: &sh,
		},
		From:       from,
		To:         to,
		BkgDetails: details,
		Dutiable:   dutiable,
	}
	return &DCTRequest{
		XPNameSpace:     "http://www.dhl.com",
		XP1NameSpace:    "http://www.dhl.com/datatypes",
		XP2NameSpace:    "http://www.dhl.com/DCTRequestdatatypes",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com DCT-req.xsd ",
		GetQuote:        &q,
	}
}

// GetCapability takes DCTFrom, DCTTo, BkgDetailsRequest and DCTDutiable and makes DCTRequest with GetQuote.
// Returns a DCTResponse with GetCapabilityResponse
func (c *dhlClient) GetCapability(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) (*DCTResponse, error) {
	data := c.buildCapabilityRequest(from, to, details, dutiable)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetCapability Response Body: %s\n", string(*contents))
	}

	var dctResponse DCTResponse
	if err := xml.Unmarshal(*contents, &dctResponse); err != nil {
		return nil, err
	}

	return &dctResponse, nil
}

func (c *dhlClient) buildCapabilityRequest(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) *DCTRequest {
	sh := NewServiceHeader(c.siteID, c.password)
	q := GetCapability{
		Request: &Request{
			ServiceHeader: &sh,
		},
		From:       from,
		To:         to,
		BkgDetails: details,
		Dutiable:   dutiable,
	}
	return &DCTRequest{
		XPNameSpace:     "http://www.dhl.com",
		XP1NameSpace:    "http://www.dhl.com/datatypes",
		XP2NameSpace:    "http://www.dhl.com/DCTRequestdatatypes",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com DCT-req.xsd ",
		GetCapability:   &q,
	}
}

// Tracking takes a TrackingQuery and makes a Known or Unknown Tracking Request
//
// AWBNumbers, LPNumbers, LevelOfDetails and PiecesEnabled are KnownTrackingRequest specific
// AccountNumber, ShipmentDate and ShipperReference are UnknownTrackingRequest specific
func (c *dhlClient) Tracking(query TrackingQuery) (*TrackingResponse, error) {
	var data interface{}

	if len(query.LPNumbers) != 0 {
		if c.debug {
			fmt.Println("Build KnownTrackingRequest")
		}
		data = c.buildKnownTrackingRequest(query)
	} else if query.AccountNumber != 0 {
		if c.debug {
			fmt.Println("Build UnknownTrackingRequest")
		}
		data = c.buildUnknownTrackingRequest(query)
	} else {
		return nil, errors.New("Missing LPNumber (KnownTrackingRequest) or AccountNumber (UnknowTrackingRequest)")
	}

	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("Tracking Response Body: %s\n", string(*contents))
	}

	var trackingResponse TrackingResponse
	if err := xml.Unmarshal(*contents, &trackingResponse); err != nil {
		return nil, err
	}

	return &trackingResponse, nil
}

func (c *dhlClient) buildKnownTrackingRequest(query TrackingQuery) *KnownTrackingRequest {
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

func (c *dhlClient) buildUnknownTrackingRequest(query TrackingQuery) *UnknownTrackingRequest {
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

func (c *dhlClient) fetch(data interface{}) (*[]byte, error) {
	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s/%s", c.baseURL, "XMLShippingServlet")
	body := bytes.NewBuffer([]byte(xmlstring))

	if c.debug {
		fmt.Printf("API Request Body: %s\n", body)
	}

	resp, err := c.httpClient.Post(url, "text/xml", body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	return &contents, err
}

func (c *dhlClient) Routing() {}
