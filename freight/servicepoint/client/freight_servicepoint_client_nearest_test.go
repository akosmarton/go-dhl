package client_test

import (
	"testing"

	servicepoint "github.com/shipwallet/go-dhl/freight/servicepoint/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetNearestServicePointsTestSuite struct {
	suite.Suite
}

func (suite *GetNearestServicePointsTestSuite) SetupTest() {

}

func (suite *GetNearestServicePointsTestSuite) TestGetNearestServicePointsWorks() {
	config := servicepoint.ClientConfig{Host: "staging"}
	client, _ := servicepoint.NewServicePointClient(config)

	query := servicepoint.NearestServicePointsQuery{}
	query.CountryCode = "SE"
	query.Street = "Mälarvarvsbacken 8"
	query.PostCode = "11733"
	query.MaxNumberOfItems = 1

	resp, err := client.GetNearestServicePoints(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)

	assert.Equal(suite.T(), len(resp.ServicePoints.NearbyServicePoint), query.MaxNumberOfItems)
}

func (suite *GetNearestServicePointsTestSuite) TestGetNearestServicePointsHaveNearbyServicePoint() {
	config := servicepoint.ClientConfig{Host: "staging"}
	client, _ := servicepoint.NewServicePointClient(config)

	query := servicepoint.NearestServicePointsQuery{}
	query.CountryCode = "SE"
	query.Street = "Mälarvarvsbacken 8"
	query.PostCode = "11733"
	query.MaxNumberOfItems = 1

	resp, err := client.GetNearestServicePoints(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.ServicePoints.NearbyServicePoint)

	point := resp.ServicePoints.NearbyServicePoint[0]
	//Identity
	assert.Equal(suite.T(), "HANDLARN BERGSUNDSTRAND", point.Identity.DisplayName)
	assert.Equal(suite.T(), "SE-648600", point.Identity.ID)

	// Address
	assert.NotNil(suite.T(), point.Distance)
	assert.Equal(suite.T(), "SLIPGATAN 11", point.StreetName)
	assert.Equal(suite.T(), "11739", point.PostCode)
	assert.Equal(suite.T(), "STOCKHOLM", point.City)
	assert.NotNil(suite.T(), point.RouteDistance)

	// FeatureCodes
	assert.Equal(suite.T(), 2, len(point.FeatureCodes.FeatureCode))
}

func (suite *GetNearestServicePointsTestSuite) TestGetNearestServicePointsErrorResponseReturnedAsErrors() {
	config := servicepoint.ClientConfig{Host: "staging"}
	client, _ := servicepoint.NewServicePointClient(config)

	query := servicepoint.NearestServicePointsQuery{}
	query.CountryCode = "SE"
	query.Street = "Mälarvarvsbacken 8"
	query.PostCode = "1113f"
	query.MaxNumberOfItems = 1

	_, err := client.GetNearestServicePoints(query)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "Error occurred during webservice request", err.Error())
}

func TestGetNearestServicePointsTestSuite(t *testing.T) {
	suite.Run(t, new(GetNearestServicePointsTestSuite))
}
