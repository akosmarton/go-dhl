package client_test

import (
	"testing"

	"github.com/shipwallet/go-dhl/express/client"
	"github.com/shipwallet/go-dhl/express/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TrackingTestSuite struct {
	suite.Suite
}

func (suite *TrackingTestSuite) SetupTest() {

}

func (suite *TrackingTestSuite) TestValidKnownTrackingRequest() {
	config := client.ClientConfig{Host: "staging"}
	c, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := client.TrackingQuery{
		LanguageCode:   "SV",
		AWBNumbers:     []models.AWBNumber{"123456789"},
		LevelOfDetails: "ALL_CHECK_POINTS",
	}

	resp, err := c.Tracking(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), resp.AWBInfo[0].Status.ActionStatus, "success")
}

func (suite *TrackingTestSuite) TestValidUnknownTrackingRequest() {
	config := client.ClientConfig{Host: "staging"}
	c, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := client.TrackingQuery{
		LanguageCode:  "en",
		AccountNumber: 630276297,
		ShipperReference: &models.Reference{
			ReferenceID:   "8100048270",
			ReferenceType: "St",
		},
		ShipmentDate: &models.ShipmentDate{
			ShipmentDateFrom: "2010-07-15",
			ShipmentDateTo:   "2010-07-20",
		},
	}

	resp, err := c.Tracking(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), resp.AWBInfo[0].Status.ActionStatus, "No Shipments Found")
}

func TestTrackingTestSuite(t *testing.T) {
	suite.Run(t, new(TrackingTestSuite))
}
