package express

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoutingTestSuite struct {
	suite.Suite
}

func (suite *RoutingTestSuite) SetupTest() {

}

func (suite *RoutingTestSuite) TestRoutingRequestGlobalAM() {
	config := ClientConfig{Host: "staging"}
	c, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := RouteQuery{
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
	assert.Equal(suite.T(), "Success", resp.Note.ActionNote)
	assert.Equal(suite.T(), "AM", resp.RegionCode)
}

func (suite *RoutingTestSuite) TestRoutingRequestGlobalAP() {
	config := ClientConfig{Host: "staging"}
	c, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := RouteQuery{
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
	assert.Equal(suite.T(), "Success", resp.Note.ActionNote)
	assert.Equal(suite.T(), "AP", resp.RegionCode)
}

func (suite *RoutingTestSuite) TestRoutingRequestGlobalEU() {
	config := ClientConfig{Host: "staging"}
	c, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := RouteQuery{
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
	assert.Equal(suite.T(), "Success", resp.Note.ActionNote)
	assert.Equal(suite.T(), "EU", resp.RegionCode)
}

func (suite *RoutingTestSuite) TestInvalidRoutingRequest() {
	config := ClientConfig{Host: "staging"}
	c, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	query := RouteQuery{
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
	assert.Equal(suite.T(), "Error", resp.Response.Status.ActionStatus)
}

func TestRoutingTestSuite(t *testing.T) {
	suite.Run(t, new(RoutingTestSuite))
}
