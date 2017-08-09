package timetable

import (
	"encoding/xml"
	"errors"

	log "github.com/sirupsen/logrus"
)

// GetTimeTableRequest represents a request for TimeTable
type GetTimeTableRequest struct {
	XMLName             string `xml:"http://dhltimetable.dhl.com/ GetTimeTable"`
	User                string `xml:"strUser"`
	Password            string `xml:"strPwd"`
	Type                int    `xml:"intType"` // 1 = styckesgods
	CountryCode         string `xml:"strCountryCode"`
	PickupPostCode      string `xml:"strPickupPostCode"`
	PickupPlace         string `xml:"strPickupPlace"`
	DeliveryCountryCode string `xml:"strDeliveryCountryCode"` //SE or NO
	DeliveryPostCode    string `xml:"strDeliveryPostCode"`
	DeliveryPlace       string `xml:"strDeliveryPlace"`
	EarliestSent        bool   `xml:"blnEarliestSent"`
	Date                string `xml:"dtDate"` // Datetime string YYYY-MM-DD
	HolidayCheck        bool   `xml:"blnHolidayCheck"`
}

// GetTimeTableResponse represents a response with TimeTable
type GetTimeTableResponse struct {
	XMLName            string              `xml:"http://dhltimetable.dhl.com/ GetTimeTableResponse"`
	GetTimeTableResult *GetTimeTableResult `xml:"GetTimeTableResult,omitempty"`
}

// GetTimeTableResult represents a result with TimeTable
type GetTimeTableResult struct {
	SCHTimeTable *SCHTimeTable `xml:"http://dhltimetable.dhl.com/schTimeTable.xsd schTimeTable,omitempty"`
}

// SCHTimeTable ...
type SCHTimeTable struct {
	XMLName      string         `xml:"http://dhltimetable.dhl.com/schTimeTable.xsd schTimeTable"`
	Information  *Information   `xml:"Information,omitempty"`
	TimeTable    []TimeTable    `xml:"TimeTable,omitempty"`
	TimeTableOrg []TimeTableOrg `xml:"TimeTableOrg,omitempty"`
}

// Information represents pickup and delivery data
type Information struct {
	PickupCountryCode   string `xml:"PickupCountryCode,omitempty"`
	PickupPostCode      string `xml:"PickupPostCode,omitempty"`
	PickupPlace         string `xml:"PickupPlace,omitempty"`
	DeliveryCountryCode string `xml:"DeliveryCountryCode,omitempty"`
	DeliveryPostCode    string `xml:"DeliveryPostCode,omitempty"`
	DeliveryPlace       string `xml:"DeliveryPlace,omitempty"`
}

// TimeTable represents pickup and delivery time
type TimeTable struct {
	Product            string `xml:"Product,omitempty"`
	PickupDateString   string `xml:"PickupDateString,omitempty"`
	DeliveryDateString string `xml:"DeliveryDateString,omitempty"`
	Options            string `xml:"Options,omitempty"`
	PickupDate         string `xml:"PickupDate,omitempty"`   // Datetime string
	DeliveryDate       string `xml:"DeliveryDate,omitempty"` // Datetime string
}

// TimeTableOrg ...
type TimeTableOrg struct {
	SAvgTerminal string `xml:"sAvgTerminal,omitempty"`
	OPickup      string `xml:"oPickup,omitempty"`
	OAvgTerm     string `xml:"oAvgTerm,omitempty"`
	OAnkTerm     string `xml:"oAnkTerm,omitempty"`
	OAvgDistTerm string `xml:"oAvgDistTerm,omitempty"`
	OLossar      string `xml:"oLossar,omitempty"`
	SGTjanst     string `xml:"sGTjanst,omitempty"`
}

func (c *dhlTimeTableClient) GetTimeTable(query TimeTableQuery) (*GetTimeTableResponse, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	data := c.buildGetTimeTableRequest(query)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}
	log.Debugf("GetTimeTable Response Body: %s\n", string(*contents))

	var response SoapResponseEvelope
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

func (c *dhlTimeTableClient) buildGetTimeTableRequest(q TimeTableQuery) *SoapRequestEnvelope {

	timetable := GetTimeTableRequest{
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

	return &SoapRequestEnvelope{
		XSINameSpace:  "http://www.w3.org/2001/XMLSchema-instance",
		XSDNameSpace:  "http://www.w3.org/2001/XMLSchema",
		SOAPNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &SoapRequestBody{
			Request: &timetable,
		},
	}
}
