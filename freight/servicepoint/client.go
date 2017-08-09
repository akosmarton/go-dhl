package servicepoint

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var hosts = map[string]string{
	"staging":    "https://www.dhltoolboxtest.se",
	"production": "https://www.dhltoolbox.se",
}

// Client interface for getting ServicePoints data
type Client interface {
	GetNearestServicePoints(NearestServicePointsQuery) (*GetNearestServicePointsResponse, error)
	GetServicePointDetail(ServicePointDetailQuery) (*GetServicePointDetailResponse, error)
}

// ClientConfig object used to configure dhlExpressClient
type ClientConfig struct {
	Host       string
	HTTPClient *http.Client
}

// NewClient used for creating a new client instance of ServicePoint
func NewClient(config ClientConfig) (Client, error) {
	host := "production"
	if config.Host != "" {
		host = config.Host
	}

	client := &http.Client{}
	if config.HTTPClient != nil {
		client = config.HTTPClient
	}

	return &dhlServicePointClient{
		baseURL:    hosts[host],
		httpClient: client,
	}, nil
}

type dhlServicePointClient struct {
	baseURL    string
	httpClient *http.Client
}

func (c *dhlServicePointClient) fetch(action string, data interface{}) (*[]byte, error) {
	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, "DHLServicePointLocatorWS/ServicePoint.svc")
	body := bytes.NewBuffer([]byte(xmlstring))

	log.Debugf("API header, SOAPAction: %s\n", action)
	log.Debugf("API Request URL: %s\n", url)
	log.Debugf("API Request Body: %s\n", body)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPAction", action)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	return &contents, err
}
