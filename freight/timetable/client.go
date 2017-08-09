package timetable

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
	"staging":    "https://www.dhltoolboxtest.se",
	"production": "https://www.dhltoolbox.se",
}

// Client interface for getting timetable
type Client interface {
	GetTimeTable(TimeTableQuery) (*GetTimeTableResponse, error)
	//GettimeTableAndPriceTD not implemented
}

// Config object used to configure dhlExpressClient
type Config struct {
	Host       string
	HTTPClient *http.Client
}

// NewClient used for creating a new TimeTable Client instance
func NewClient(username string, password string, config Config) (Client, error) {
	if username == "" {
		return nil, errors.New("Username missing")
	}
	if password == "" {
		return nil, errors.New("Password missing")
	}

	host := "production"
	if config.Host != "" {
		host = config.Host
	}

	client := &http.Client{}
	if config.HTTPClient != nil {
		client = config.HTTPClient
	}

	return &dhlTimeTableClient{
		baseURL:    hosts[host],
		httpClient: client,
		username:   username,
		password:   password,
	}, nil
}

type dhlTimeTableClient struct {
	baseURL    string
	httpClient *http.Client
	username   string
	password   string
}

func (c *dhlTimeTableClient) fetch(data interface{}) (*[]byte, error) {
	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, "DHLTimeTableWS/TimeTable.asmx")
	body := bytes.NewBuffer([]byte(xmlstring))

	action := "http://dhltimetable.dhl.com/GetTimeTable"
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
