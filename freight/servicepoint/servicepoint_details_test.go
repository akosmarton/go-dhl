package servicepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetServicePointDetailTestSuite struct {
	suite.Suite
}

func (suite *GetServicePointDetailTestSuite) SetupTest() {

}

func (suite *GetServicePointDetailTestSuite) TestGetServicePointDetailWorks() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewClient(config)

	query := ServicePointDetailQuery{}
	query.ID = "SE-648600"

	resp, err := client.GetServicePointDetail(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)

	assert.Equal(suite.T(), query.ID, resp.ServicePointDetail.Identity.ID)
}

func (suite *GetServicePointDetailTestSuite) TestGetServicePointDetailServiceAddressAddress() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewClient(config)

	query := ServicePointDetailQuery{}
	query.ID = "SE-648600"

	resp, err := client.GetServicePointDetail(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.ServicePointDetail.ServiceAddress)

	address := resp.ServicePointDetail.ServiceAddress.Address
	assert.Equal(suite.T(), "HANDLARN BERGSUNDSTRAND", address.AddresseeName)
	assert.Equal(suite.T(), "STOCKHOLM", address.City)
	assert.Equal(suite.T(), "SE", address.CountryCode)
	assert.Equal(suite.T(), "11739", address.PostCode)
	assert.Equal(suite.T(), "SLIPGATAN 11", address.Street1)
	assert.Equal(suite.T(), "", address.Street2)
}

func (suite *GetServicePointDetailTestSuite) TestGetServicePointDetailServiceAddressTelecom() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewClient(config)

	query := ServicePointDetailQuery{}
	query.ID = "SE-648600"

	resp, err := client.GetServicePointDetail(query)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.ServicePointDetail.ServiceAddress)

	telecom := resp.ServicePointDetail.ServiceAddress.Telecom
	assert.Equal(suite.T(), "", telecom.Email)
	assert.Equal(suite.T(), "", telecom.Fax)
	assert.Equal(suite.T(), "08-68433410", telecom.Phone)
}

func (suite *GetServicePointDetailTestSuite) TestGetServicePointDetailInvalidServicePointID() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewClient(config)

	query := ServicePointDetailQuery{}
	query.ID = "SE-2930192"

	resp, err := client.GetServicePointDetail(query)

	assert.Nil(suite.T(), resp)
	assert.NotNil(suite.T(), "Error occurred during webservice request", err.Error())
}

func TestGetServicePointDetailTestSuite(t *testing.T) {
	suite.Run(t, new(GetServicePointDetailTestSuite))
}
