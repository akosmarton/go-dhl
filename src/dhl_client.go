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
	Tracking()
	Routing()
}

// DHLClientConfig object used to configure DHLClient
type DHLClientConfig struct {
	Debug bool
	Host  string
}

// NewDHLClient returns a DHLClient against DHL Production Environment
func NewDHLClient(siteID string, password string, config DHLClientConfig) (Client, error) {
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

func (c *dhlClient) GetQuote(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) (*DCTResponse, error) {
	data := c.buildQuoteRequest(from, to, details, dutiable)

	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s/%s", c.baseURL, "XMLShippingServlet")
	body := bytes.NewBuffer([]byte(xmlstring))

	if c.debug {
		fmt.Printf("GetQuote Request Body: %s\n", body)
	}

	resp, err := c.httpClient.Post(url, "text/xml", body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetQuote Response Body: %s\n", string(contents))
	}

	var dctResponse DCTResponse
	if err := xml.Unmarshal(contents, &dctResponse); err != nil {
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

func (c *dhlClient) Tracking() {}

func (c *dhlClient) Routing() {}
