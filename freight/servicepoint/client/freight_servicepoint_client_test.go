package client_test

import (
	"fmt"
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

func (suite *GetNearestServicePointsTestSuite) TestGetTimeTableSEtoSE() {
	config := servicepoint.ClientConfig{Host: "staging", Debug: true}
	client, _ := servicepoint.NewServicePointClient(config)

	query := servicepoint.NearestServicePointsQuery{}
	query.CountryCode = "SE"
	query.Street = "Mälarvarvsbacken 8"
	query.PostCode = "11733"
	query.MaxNumberOfItems = 2

	resp, err := client.GetNearestServicePoints(query)

	fmt.Printf("%+v\n", resp)

	assert.Nil(suite.T(), err)
	//assert.NotNil(suite.T(), resp)
}

func TestGetNearestServicePointsTestSuite(t *testing.T) {
	suite.Run(t, new(GetNearestServicePointsTestSuite))
}
