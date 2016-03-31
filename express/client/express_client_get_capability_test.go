package client_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/shipwallet/go-dhl/express/client"
	"github.com/shipwallet/go-dhl/express/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetCapabilityTestSuite struct {
	suite.Suite
}

func (suite *GetCapabilityTestSuite) loadTestData(path string) (*models.GetQuote, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer xmlFile.Close()

	requestXML, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var q models.GetQuote
	xml.Unmarshal(requestXML, &q)

	return &q, nil
}

func (suite *GetCapabilityTestSuite) SetupTest() {

}

func (suite *GetCapabilityTestSuite) TestInvalidCapabilityEUToNonEUDutiableRequest() {
	config := client.ClientConfig{Host: "staging"}
	client, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &models.DCTFrom{}
	from.CountryCode = "ES"
	from.PostalCode = "28001"

	to := &models.DCTTo{}
	to.CountryCode = "HK"
	to.PostalCode = ""
	to.City = ""

	t := time.Now()
	bdr := &models.BkgDetailsRequest{}
	bdr.PaymentCountryCode = "ES"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &models.Pieces{
		Piece: []models.PieceType{
			models.PieceType{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 1.0},
		},
	}
	bdr.IsDutiable = "Y"
	bdr.NetworkTypeCode = "AL"

	du := &models.DCTDutiable{}
	du.DeclaredCurrency = "USD"
	du.DeclaredValue = 1002.00

	resp, err := client.GetCapability(from, to, bdr, du)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.GetCapabilityResponse)
	assert.NotNil(suite.T(), resp.GetCapabilityResponse.Note)
}

func (suite *GetQuoteTestSuite) TestValidCapabilityEUToNonEUDutiableRequest() {
	config := client.ClientConfig{Host: "staging"}
	client, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &models.DCTFrom{}
	from.CountryCode = "ES"
	from.PostalCode = "28001"

	to := &models.DCTTo{}
	to.CountryCode = "HK"
	to.PostalCode = "90210"
	to.City = "Hong Kong"

	t := time.Now()
	bdr := &models.BkgDetailsRequest{}
	bdr.PaymentCountryCode = "ES"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &models.Pieces{
		Piece: []models.PieceType{
			models.PieceType{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 1.0},
		},
	}
	bdr.IsDutiable = "Y"
	bdr.NetworkTypeCode = "AL"

	du := &models.DCTDutiable{}
	du.DeclaredCurrency = "USD"
	du.DeclaredValue = 1002.00

	resp, err := client.GetCapability(from, to, bdr, du)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.GetCapabilityResponse)
	assert.NotNil(suite.T(), len(resp.GetCapabilityResponse.BkgDetails.QtdShp))
	assert.Nil(suite.T(), resp.GetCapabilityResponse.Note)
}

func (suite *GetQuoteTestSuite) TestValidCapabilityEUToEUDutiableRequest() {
	config := client.ClientConfig{Host: "staging"}
	client, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &models.DCTFrom{}
	from.CountryCode = "CZ"
	from.PostalCode = "10000"

	to := &models.DCTTo{}
	to.CountryCode = "SE"
	to.PostalCode = "10054"
	to.City = "Stockholm"

	t := time.Now()
	bdr := &models.BkgDetailsRequest{}
	bdr.PaymentCountryCode = "CZ"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &models.Pieces{
		Piece: []models.PieceType{
			models.PieceType{PieceID: "1", Height: 30, Depth: 30, Width: 30, Weight: 10.0},
		},
	}
	bdr.IsDutiable = "N"
	bdr.NetworkTypeCode = "AL"

	resp, err := client.GetCapability(from, to, bdr, nil)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.GetCapabilityResponse)
	assert.NotNil(suite.T(), len(resp.GetCapabilityResponse.BkgDetails.QtdShp))
	assert.Nil(suite.T(), resp.GetCapabilityResponse.Note)
}

func (suite *GetQuoteTestSuite) TestValidCapabilityNonEUToNonEUDutiableRequest() {
	config := client.ClientConfig{Host: "staging"}
	client, _ := client.NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &models.DCTFrom{}
	from.CountryCode = "MY"
	from.PostalCode = "57000"
	from.City = "Kuala Lumpur"

	to := &models.DCTTo{}
	to.CountryCode = "AU"
	to.PostalCode = "2020"
	to.City = "Mascot"

	t := time.Now()
	bdr := &models.BkgDetailsRequest{}
	bdr.PaymentCountryCode = "MY"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &models.Pieces{
		Piece: []models.PieceType{
			models.PieceType{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 10.0},
		},
	}
	bdr.IsDutiable = "Y"
	bdr.NetworkTypeCode = "AL"

	du := &models.DCTDutiable{}
	du.DeclaredCurrency = "MYR"
	du.DeclaredValue = 1002.00

	resp, err := client.GetCapability(from, to, bdr, du)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp.GetCapabilityResponse)
	assert.NotNil(suite.T(), len(resp.GetCapabilityResponse.BkgDetails.QtdShp))
	assert.Nil(suite.T(), resp.GetCapabilityResponse.Note)
}

func TestGetCapabilityTestSuite(t *testing.T) {
	suite.Run(t, new(GetCapabilityTestSuite))
}
