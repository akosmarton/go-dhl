package client_test

import (
	"testing"

	"github.com/shipwallet/go-dhl/express/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoutingTestSuite struct {
	suite.Suite
}

func (suite *RoutingTestSuite) SetupTest() {

}

func (suite *RoutingTestSuite) TestRoutingRequestGlobalAM() {
	config := client.ClientConfig{Host: "staging"}
	c, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := client.RouteQuery{
		RegionCode:        "AM",
		RequestType:       "O",
		Address1:          "Suit 333",
		Address2:          "333 Twin",
		PostalCode:        "94089",
		City:              "North Dakhota",
		Division:          "California",
		CountryCode:       "US",
		CountryName:       "United States of America",
		OriginCountryCode: "US",
	}

	resp, err := c.Routing(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), resp.Note.ActionNote, "Success")
	assert.Equal(suite.T(), resp.RegionCode, "AM")
}

func (suite *RoutingTestSuite) TestRoutingRequestGlobalAP() {
	config := client.ClientConfig{Host: "staging"}
	c, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := client.RouteQuery{
		RegionCode:        "AP",
		RequestType:       "D",
		Address1:          "13, Jalan SS23/15",
		Address2:          "Taman SEA",
		Address3:          "SEA building",
		PostalCode:        "47400",
		City:              "Petaling Jaya",
		Division:          "PJ",
		CountryCode:       "MY",
		CountryName:       "Malaysia",
		OriginCountryCode: "MY",
	}

	resp, err := c.Routing(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), resp.Note.ActionNote, "Success")
	assert.Equal(suite.T(), resp.RegionCode, "AP")
}

func (suite *RoutingTestSuite) TestRoutingRequestGlobalEU() {
	config := client.ClientConfig{Host: "staging"}
	c, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := client.RouteQuery{
		RegionCode:        "EU",
		RequestType:       "O",
		Address1:          "Oracle Parkway,",
		Address2:          "Thames Valley Park (TVP)",
		Address3:          "Berkshire",
		PostalCode:        "RG6 1RA",
		City:              "Reading",
		Division:          "RE",
		CountryCode:       "GB",
		CountryName:       "United Kingdom",
		OriginCountryCode: "GB",
	}

	resp, err := c.Routing(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), resp.Note.ActionNote, "Success")
	assert.Equal(suite.T(), resp.RegionCode, "EU")
}

func (suite *RoutingTestSuite) TestInvalidRoutingRequest() {
	config := client.ClientConfig{Host: "staging"}
	c, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := client.RouteQuery{
		RegionCode:        "AP",
		RequestType:       "F",
		Address1:          "13, Jalan SS23/15",
		Address2:          "Taman SEA",
		Address3:          "SEA building",
		PostalCode:        "47400",
		City:              "Petaling Jaya",
		Division:          "PJ",
		CountryCode:       "MY",
		CountryName:       "Malaysia",
		OriginCountryCode: "MY",
	}

	resp, err := c.Routing(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.Response)
	assert.Equal(suite.T(), resp.Response.Status.ActionStatus, "Error")
}

func TestRoutingTestSuite(t *testing.T) {
	suite.Run(t, new(RoutingTestSuite))
}
