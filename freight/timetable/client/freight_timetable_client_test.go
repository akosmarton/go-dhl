package client_test

import (
	"testing"
	"time"

	timetable "github.com/shipwallet/go-dhl/freight/timetable/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetTimeTableTestSuite struct {
	suite.Suite
}

func (suite *GetTimeTableTestSuite) SetupTest() {

}

func (suite *GetTimeTableTestSuite) TestGetTimeTableSEtoSE() {
	config := timetable.ClientConfig{Host: "staging"}
	client, _ := timetable.NewTimeTableClient("TimeTableAppUser", "TTA%Pwd06", config)

	t := time.Now()
	query := timetable.TimeTableQuery{}
	query.OriginCountryCode = "SE"
	query.OriginPostCode = "14250"
	query.DestinationCountryCode = "SE"
	query.DestinationPostCode = "11733"
	query.EarliestSent = true
	query.Date = t.Format("2006-01-02")

	resp, err := client.GetTimeTable(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)

	assert.Equal(suite.T(), resp.GetTimeTableResult.SCHTimeTable.Information.PickupCountryCode, query.OriginCountryCode)
	assert.Equal(suite.T(), resp.GetTimeTableResult.SCHTimeTable.Information.PickupPostCode, query.OriginPostCode)
	assert.NotNil(suite.T(), resp.GetTimeTableResult.SCHTimeTable.Information.PickupPlace)

	assert.Equal(suite.T(), resp.GetTimeTableResult.SCHTimeTable.Information.DeliveryCountryCode, query.DestinationCountryCode)
	assert.Equal(suite.T(), resp.GetTimeTableResult.SCHTimeTable.Information.DeliveryPostCode, query.DestinationPostCode)
	assert.NotNil(suite.T(), resp.GetTimeTableResult.SCHTimeTable.Information.DeliveryPlace)

	assert.NotNil(suite.T(), len(resp.GetTimeTableResult.SCHTimeTable.TimeTable))
	assert.NotNil(suite.T(), len(resp.GetTimeTableResult.SCHTimeTable.TimeTableOrg))
}

func (suite *GetTimeTableTestSuite) TestGetTimeTableSEtoSEInvalidPostCode() {
	config := timetable.ClientConfig{Host: "staging"}
	client, _ := timetable.NewTimeTableClient("TimeTableAppUser", "TTA%Pwd06", config)

	t := time.Now()
	query := timetable.TimeTableQuery{}
	query.OriginCountryCode = "SE"
	query.OriginPostCode = "0175"
	query.DestinationCountryCode = "SE"
	query.DestinationPostCode = "117 33"
	query.EarliestSent = true
	query.Date = t.Format("2006-01-02")

	resp, err := client.GetTimeTable(query)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), resp)

	assert.Equal(suite.T(), err.Error(), "Server was unable to process request. ---> <DHLErrID>9005</DHLErrID><DHLErrMessage>Pickup post code must be 5 digits.</DHLErrMessage>")
}

func TestGetTimeTableTestSuite(t *testing.T) {
	suite.Run(t, new(GetTimeTableTestSuite))
}
