package client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/shipwallet/go-dhl/freight/timetable/models"
)

var hosts = map[string]string{
	"staging":    "https://www.dhltoolboxtest.se",
	"production": "https://www.dhltoolbox.se",
}

// TimeTableClient interface
type TimeTableClient interface {
	GetTimeTable(TimeTableQuery) (*models.GetTimeTableResponse, error)
	//GettimeTableAndPriceTD not implemented
}

// ClientConfig object used to configure dhlExpressClient
type ClientConfig struct {
	Debug      bool
	Host       string
	HttpClient *http.Client
}

// NewTimeTableClient used for creating a new TimeTableClient instance
func NewTimeTableClient(username string, password string, config ClientConfig) (TimeTableClient, error) {

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
	if config.HttpClient != nil {
		client = config.HttpClient
	}

	return &dhlTimeTableClient{
		baseURL:    hosts[host],
		debug:      config.Debug,
		httpClient: client,
		username:   username,
		password:   password,
	}, nil
}

type dhlTimeTableClient struct {
	baseURL    string
	debug      bool
	httpClient *http.Client
	username   string
	password   string
}

var (
	dhlErrIDPattern  = regexp.MustCompile(`<DHLErrID>(.*)</DHLErrID>`)
	dhlErrMsgPattern = regexp.MustCompile(`<DHLErrMessage>(.*)</DHLErrMessage>`)
)

type DHLError struct {
	DHLErrID      string
	DHLErrMessage string
}

func (d DHLError) Error() string {
	return fmt.Sprintf("DHLError id=%s msg=%s", d.DHLErrID, d.DHLErrMessage)
}

func (c *dhlTimeTableClient) GetTimeTable(query TimeTableQuery) (*models.GetTimeTableResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	data := c.buildGetTimeTableRequest(query)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("GetTimeTable Response Body: %s\n", string(*contents))
	}

	var response models.SoapResponseEvelope

	if err := xml.Unmarshal(*contents, &response); err != nil {
		return nil, err
	}

	if response.Body.Fault != nil {
		dhlErrID := dhlErrIDPattern.FindAllStringSubmatch(response.Body.Fault.FaultString, -1)
		dhlErrMsg := dhlErrMsgPattern.FindAllStringSubmatch(response.Body.Fault.FaultString, -1)

		if len(dhlErrID) == 1 && len(dhlErrID[0]) == 2 && len(dhlErrMsg) == 1 && len(dhlErrMsg[0]) == 2 {
			dhlError := DHLError{dhlErrID[0][1], dhlErrMsg[0][1]}
			return nil, dhlError
		}

		return nil, errors.New(response.Body.Fault.FaultString)
	}

	return response.Body.Response, nil
}

func (c *dhlTimeTableClient) buildGetTimeTableRequest(q TimeTableQuery) *models.SoapRequestEnvelope {

	timetable := models.GetTimeTableRequest{
		User:                c.username,
		Password:            c.password,
		Type:                1,
		CountryCode:         q.OriginCountryCode,
		PickupPostCode:      q.OriginPostCode,
		PickupPlace:         q.OriginPlace,
		DeliveryCountryCode: q.DestinationCountryCode,
		DeliveryPostCode:    q.DestinationPostCode,
		DeliveryPlace:       q.DestinationPlace,
		EarliestSent:        q.EarliestSent,
		Date:                q.Date,
		HolidayCheck:        q.HolidayCheck,
	}

	return &models.SoapRequestEnvelope{
		XSINameSpace:  "http://www.w3.org/2001/XMLSchema-instance",
		XSDNameSpace:  "http://www.w3.org/2001/XMLSchema",
		SOAPNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &models.SoapRequestBody{
			Request: &timetable,
		},
	}
}

func (c *dhlTimeTableClient) fetch(data interface{}) (*[]byte, error) {
	xmlstring, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, "DHLTimeTableWS/TimeTable.asmx")
	body := bytes.NewBuffer([]byte(xmlstring))

	action := "http://dhltimetable.dhl.com/GetTimeTable"
	if c.debug {
		fmt.Printf("API header, SOAPAction: %s\n", action)
		fmt.Printf("API Request URL: %s\n", url)
		fmt.Printf("API Request Body: %s\n", body)
	}

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
