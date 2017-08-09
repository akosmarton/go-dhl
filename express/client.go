package express

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
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
	Routing(RouteQuery) (*RouteResponse, error)
}

// ClientConfig object used to configure dhlExpressClient
type ClientConfig struct {
	Host       string
	HTTPClient *http.Client
}

type dhlExpressClient struct {
	baseURL    string
	httpClient *http.Client
	siteID     string
	password   string
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
	if config.HTTPClient != nil {
		client = config.HTTPClient
	}

	return &dhlExpressClient{
		baseURL:    hosts[host],
		httpClient: client,
		siteID:     siteID,
		password:   password,
	}, nil
}

func (c *dhlExpressClient) fetch(data interface{}) (*[]byte, error) {
	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer([]byte(xmlstring))
	log.Debugf("API Request Body: %s\n", body)

	url := fmt.Sprintf("http://%s/%s", c.baseURL, "XMLShippingServlet")
	resp, err := c.httpClient.Post(url, "text/xml", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	return &contents, err
}
