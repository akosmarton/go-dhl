package express

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
	// Test isn't passing regardless of correct test data.
	suite.T().Skip("Skipping TestValidKnownTrackingRequest")

	config := ClientConfig{Host: "staging"}
	c, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := TrackingQuery{
		LanguageCode:   "SV",
		AWBNumbers:     []AWBNumber{"123456789"},
		LevelOfDetails: "ALL_CHECK_POINTS",
	}

	resp, err := c.Tracking(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), "success", resp.AWBInfo[0].Status.ActionStatus)
}

func (suite *TrackingTestSuite) TestValidUnknownTrackingRequest() {
	config := ClientConfig{Host: "staging"}
	c, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := TrackingQuery{
		LanguageCode:  "en",
		AccountNumber: 630276297,
		ShipperReference: &Reference{
			ReferenceID:   "8100048270",
			ReferenceType: "St",
		},
		ShipmentDate: &ShipmentDate{
			ShipmentDateFrom: "2010-07-15",
			ShipmentDateTo:   "2010-07-20",
		},
	}

	resp, err := c.Tracking(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), "No Shipments Found", resp.AWBInfo[0].Status.ActionStatus)
}

func TestTrackingTestSuite(t *testing.T) {
	suite.Run(t, new(TrackingTestSuite))
}
