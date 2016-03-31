package client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shipwallet/go-dhl/freight/servicepoint/models"
)

var hosts = map[string]string{
	"staging":    "https://www.dhltoolboxtest.se",
	"production": "http://164.9.104.199",
}

// ServicePointClient interface
type ServicePointClient interface {
	GetNearestServicePoints(NearestServicePointsQuery) (*models.GetNearestServicePointsResponse, error)
	GetServicePointDetail()
}

// ClientConfig object used to configure dhlExpressClient
type ClientConfig struct {
	Debug bool
	Host  string
}

// NewServicePointClient used for creating a new TimeTableClient instance
func NewServicePointClient(config ClientConfig) (ServicePointClient, error) {

	host := "production"
	if config.Host != "" {
		host = config.Host
	}

	return &dhlServicePointClient{
		baseURL:    hosts[host],
		debug:      config.Debug,
		httpClient: &http.Client{},
	}, nil
}

type dhlServicePointClient struct {
	baseURL    string
	debug      bool
	httpClient *http.Client
}

func (c *dhlServicePointClient) GetNearestServicePoints(query NearestServicePointsQuery) (*models.GetNearestServicePointsResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	data := c.buildGetNearestServicePointsRequest(query)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetNearestServicePoints Response Body: %s\n", string(*contents))
	}

	var response models.SoapResponseEvelope

	if err := xml.Unmarshal(*contents, &response); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", response)

	if response.Body.Fault != nil {
		return nil, errors.New(response.Body.Fault.FaultString)
	}

	return response.Body.Response, nil
}

func (c *dhlServicePointClient) buildGetNearestServicePointsRequest(q NearestServicePointsQuery) *models.SoapRequestEnvelope {

	servicepoint := models.GetNearestServicePointsRequest{
		CountryCode: q.CountryCode,
		Street:      q.Street,
		PostCode:    q.PostCode,
		City:        q.City,
		FeatureCodes: &models.FeatureCodes{
			FeatureCode: q.FeatureCodes,
		},
		MaxNumberOfItems: q.MaxNumberOfItems,
		BitCatCodes: &models.BitCatCodes{
			BitCatCode: q.BitCatCodes,
		},
	}

	return &models.SoapRequestEnvelope{
		SOAPNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &models.SoapRequestBody{
			Request: &servicepoint,
		},
	}
}

func (c *dhlServicePointClient) GetServicePointDetail() {

}

func (c *dhlServicePointClient) fetch(data interface{}) (*[]byte, error) {
	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, "DHLServicePointLocatorWS/ServicePoint.svc")
	body := bytes.NewBuffer([]byte(xmlstring))

	if c.debug {
		fmt.Printf("API Request Body: %s\n", body)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("SOAPAction", "GetNearestServicePoints")
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	return &contents, err
}
