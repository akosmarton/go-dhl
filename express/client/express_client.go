package client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shipwallet/go-dhl/express/models"
)

var hosts = map[string]string{
	"staging":    "xmlpitest-ea.dhl.com",
	"production": "xmlpi-ea.dhl.com",
}

// Client interface
type Client interface {
	GetQuote(*models.DCTFrom, *models.DCTTo, *models.BkgDetailsRequest, *models.DCTDutiable) (*models.DCTResponse, error)
	GetCapability(*models.DCTFrom, *models.DCTTo, *models.BkgDetailsRequest, *models.DCTDutiable) (*models.DCTResponse, error)
	Tracking(TrackingQuery) (*models.TrackingResponse, error)
	Routing(RouteQuery) (*models.RouteResponse, error)
}

// ClientConfig object used to configure dhlExpressClient
type ClientConfig struct {
	Debug      bool
	Host       string
	HttpClient *http.Client
}

// TrackingQuery contains parameters to be used for making a Known or Unknown Tracking request
type TrackingQuery struct {
	LanguageCode     string
	AWBNumbers       []models.AWBNumber
	LPNumbers        []models.TrackingPieceID
	LevelOfDetails   string // LAST_CHECK_POINT_ONLY or ALL_CHECK_POINTS
	PiecesEnabled    string // S, B or P
	CountryCode      string // 2 character country id
	AccountNumber    int
	ShipmentDate     *models.ShipmentDate
	ShipperReference *models.Reference
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

// NewDHLExpressClient returns a DHLExpressClient against DHL Production Environment
func NewDHLExpressClient(siteID string, password string, config ClientConfig) (Client, error) {
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

	client := &http.Client{}
	if config.HttpClient != nil {
		client = config.HttpClient
	}

	return &dhlExpressClient{
		baseURL:    hosts[host],
		httpClient: client,
		debug:      config.Debug,
		siteID:     siteID,
		password:   password,
	}, nil
}

type dhlExpressClient struct {
	baseURL    string
	httpClient *http.Client
	debug      bool
	siteID     string
	password   string
}

// GetQuote takes DCTFrom, DCTTo, BkgDetailsRequest and DCTDutiable and makes DCTRequest with GetQuote.
// Returns a DCTResponse with GetQuoteResponse
func (c *dhlExpressClient) GetQuote(from *models.DCTFrom, to *models.DCTTo, details *models.BkgDetailsRequest, dutiable *models.DCTDutiable) (*models.DCTResponse, error) {
	data := c.buildQuoteRequest(from, to, details, dutiable)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetQuote Response Body: %s\n", string(*contents))
	}

	var dctResponse models.DCTResponse
	if err := xml.Unmarshal(*contents, &dctResponse); err != nil {
		return nil, err
	}

	return &dctResponse, nil
}

func (c *dhlExpressClient) buildQuoteRequest(from *models.DCTFrom, to *models.DCTTo, details *models.BkgDetailsRequest, dutiable *models.DCTDutiable) *models.DCTRequest {
	sh := models.NewServiceHeader(c.siteID, c.password)
	q := models.GetQuote{
		Request: &models.Request{
			ServiceHeader: &sh,
		},
		From:       from,
		To:         to,
		BkgDetails: details,
		Dutiable:   dutiable,
	}
	return &models.DCTRequest{
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
func (c *dhlExpressClient) GetCapability(from *models.DCTFrom, to *models.DCTTo, details *models.BkgDetailsRequest, dutiable *models.DCTDutiable) (*models.DCTResponse, error) {
	data := c.buildCapabilityRequest(from, to, details, dutiable)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetCapability Response Body: %s\n", string(*contents))
	}

	var dctResponse models.DCTResponse
	if err := xml.Unmarshal(*contents, &dctResponse); err != nil {
		return nil, err
	}

	return &dctResponse, nil
}

func (c *dhlExpressClient) buildCapabilityRequest(from *models.DCTFrom, to *models.DCTTo, details *models.BkgDetailsRequest, dutiable *models.DCTDutiable) *models.DCTRequest {
	sh := models.NewServiceHeader(c.siteID, c.password)
	q := models.GetCapability{
		Request: &models.Request{
			ServiceHeader: &sh,
		},
		From:       from,
		To:         to,
		BkgDetails: details,
		Dutiable:   dutiable,
	}
	return &models.DCTRequest{
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
func (c *dhlExpressClient) Tracking(query TrackingQuery) (*models.TrackingResponse, error) {
	var data interface{}

	if len(query.LPNumbers) != 0 || len(query.AWBNumbers) != 0 {
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

	var trackingResponse models.TrackingResponse
	if err := xml.Unmarshal(*contents, &trackingResponse); err != nil {
		return nil, err
	}

	return &trackingResponse, nil
}

func (c *dhlExpressClient) buildKnownTrackingRequest(query TrackingQuery) *models.KnownTrackingRequest {
	sh := models.NewServiceHeader(c.siteID, c.password)
	return &models.KnownTrackingRequest{
		XReqNameSpace:   "http://www.dhl.com",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com TrackingRequestKnown.xsd",
		Request: &models.Request{
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

func (c *dhlExpressClient) buildUnknownTrackingRequest(query TrackingQuery) *models.UnknownTrackingRequest {
	sh := models.NewServiceHeader(c.siteID, c.password)
	return &models.UnknownTrackingRequest{
		XReqNameSpace:   "http://www.dhl.com",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com TrackingRequestKnown.xsd",
		Request: &models.Request{
			ServiceHeader: &sh,
		},
		LanguageCode:     query.LanguageCode,
		AccountNumber:    query.AccountNumber,
		ShipperReference: query.ShipperReference,
		CountryCode:      query.CountryCode,
	}
}

func (c *dhlExpressClient) Routing(query RouteQuery) (*models.RouteResponse, error) {
	data := c.buildRoutingRequest(query)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("Routing Response Body: %s\n", string(*contents))
	}

	var routeResponse models.RouteResponse
	if err := xml.Unmarshal(*contents, &routeResponse); err != nil {
		return nil, err
	}

	return &routeResponse, nil
}

func (c *dhlExpressClient) buildRoutingRequest(query RouteQuery) *models.RouteRequest {
	sh := models.NewServiceHeader(c.siteID, c.password)
	return &models.RouteRequest{
		XNS1NameSpace:   "http://www.dhl.com",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com routing-req.xsd",
		SchemaVersion:   1.0,
		Request: &models.Request{
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

func (c *dhlExpressClient) fetch(data interface{}) (*[]byte, error) {
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
