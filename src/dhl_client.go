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
	GetQuote(DCTFrom, DCTTo, BkgDetailsRequest, *DCTDutiable) (*DCTResponse, error)
	Tracking()
	Routing()
}

// NewDHLClient returns a DHLClient against DHL Production Environment
func NewDHLClient(siteID string, password string) (Client, error) {
	if siteID == "" {
		return nil, errors.New("siteID missing")
	}
	if password == "" {
		return nil, errors.New("password missing")
	}

	client := newDHLClient(siteID, password, "production")
	return client, nil
}

// NewDHLStagingClient returns a DHLClient against DHL Staging Environment
func NewDHLStagingClient(siteID string, password string) (Client, error) {
	if siteID == "" {
		return nil, errors.New("siteID missing")
	}
	if password == "" {
		return nil, errors.New("password missing")
	}

	client := newDHLClient(siteID, password, "staging")
	return client, nil
}

func newDHLClient(siteID string, password string, host string) Client {
	return &dhlClient{
		baseURL:    hosts[host],
		httpClient: &http.Client{},
		siteID:     siteID,
		password:   password,
	}
}

type dhlClient struct {
	baseURL    string
	httpClient *http.Client
	siteID     string
	password   string
}

func (c *dhlClient) GetQuote(from DCTFrom, to DCTTo, details BkgDetailsRequest, dutiable *DCTDutiable) (*DCTResponse, error) {
	data := c.buildQuoteRequest(from, to, details, dutiable)

	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s/%s", c.baseURL, "XMLShippingServlet")
	body := bytes.NewBuffer([]byte(xmlstring))

	fmt.Println(body)

	resp, err := c.httpClient.Post(url, "text/xml", body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", string(contents))

	var dctResponse DCTResponse
	if err := xml.Unmarshal(contents, &dctResponse); err != nil {
		return nil, err
	}

	return &dctResponse, nil
}

func (c *dhlClient) buildQuoteRequest(from DCTFrom, to DCTTo, details BkgDetailsRequest, dutiable *DCTDutiable) *DCTRequest {
	q := GetQuote{
		Request: Request{
			ServiceHeader: NewServiceHeader(c.siteID, c.password),
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

func (c *dhlClient) Tracking() {}

func (c *dhlClient) Routing() {}
