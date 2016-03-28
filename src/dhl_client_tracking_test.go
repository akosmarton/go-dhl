package dhl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TrackingTestSuite struct {
	suite.Suite
}

func (suite *TrackingTestSuite) SetupTest() {

}

func (suite *TrackingTestSuite) TestValidKnownTrackingRequest() {
	config := ClientConfig{Host: "staging", Debug: true}
	client, _ := NewDHLClient("DServiceVal", "testServVal", config)

	query := TrackingQuery{
		LanguageCode:   "SV",
		AWBNumbers:     []AWBNumber{"123456789"},
		LevelOfDetails: "ALL_CHECK_POINTS",
	}

	resp, err := client.Tracking(query)

	assert.NotNil(suite.T(), resp.Response)
	assert.Nil(suite.T(), err)
}

func (suite *TrackingTestSuite) TestValidUnknownTrackingRequest() {
	config := ClientConfig{Host: "staging", Debug: true}
	client, _ := NewDHLClient("DServiceVal", "testServVal", config)

	query := TrackingQuery{
		LanguageCode:  "en",
		AccountNumber: 630276297,
		ShipmentDate: &ShipmentDate{
			ShipmentDateFrom: "2010-07-15",
			ShipmentDateTo:   "2010-07-20",
		},
	}

	resp, err := client.Tracking(query)

	assert.NotNil(suite.T(), resp.Response)
	assert.Nil(suite.T(), err)
}

func TestTrackingTestSuite(t *testing.T) {
	suite.Run(t, new(TrackingTestSuite))
}
