package express

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetQuoteTestSuite struct {
	suite.Suite
}

func (suite *GetQuoteTestSuite) loadTestData(path string) (*GetQuote, error) {
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

	var q GetQuote
	xml.Unmarshal(requestXML, &q)

	return &q, nil
}

func (suite *GetQuoteTestSuite) SetupTest() {

}

func (suite *GetQuoteTestSuite) TestValidQuoteAPPriceBreakdownRASRequest() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &DCTFrom{}
	from.CountryCode = "ID"
	from.PostalCode = "31251"

	to := &DCTTo{}
	to.CountryCode = "JP"
	to.PostalCode = "9811513"

	t := time.Now()
	bdr := &BkgDetailsRequest{}
	bdr.PaymentCountryCode = "ID"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &Pieces{
		Piece: []PieceType{
			{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 1.0},
		},
	}
	bdr.IsDutiable = "N"
	bdr.NetworkTypeCode = "AL"
	bdr.InsuredValue = 400.0
	bdr.InsuredCurrency = "IDR"

	du := &DCTDutiable{}
	du.DeclaredCurrency = "EUR"
	du.DeclaredValue = 9.0

	resp, err := client.GetQuote(from, to, bdr, du)

	assert.NotNil(suite.T(), resp.GetQuoteResponse)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), resp.GetQuoteResponse.Note)
}

func (suite *GetQuoteTestSuite) TestValidQuoteEUPriceBreakdownRASRequest() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &DCTFrom{}
	from.CountryCode = "BE"
	from.PostalCode = "1020"

	to := &DCTTo{}
	to.CountryCode = "US"
	to.PostalCode = "86001"

	t := time.Now()
	bdr := &BkgDetailsRequest{}
	bdr.PaymentCountryCode = "BE"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &Pieces{
		Piece: []PieceType{
			{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 1.0},
		},
	}
	bdr.IsDutiable = "Y"
	bdr.NetworkTypeCode = "AL"
	bdr.QtdShp = &QtdShpRequest{
		LocalProductCode: "S",
		QtdShpExChrg: &QtdShpExChrgRequest{
			SpecialServiceType:      "I",
			LocalSpecialServiceType: "II",
		},
	}
	bdr.InsuredValue = 400.0
	bdr.InsuredCurrency = "EUR"

	du := &DCTDutiable{}
	du.DeclaredCurrency = "EUR"
	du.DeclaredValue = 9.0

	resp, err := client.GetQuote(from, to, bdr, du)

	assert.NotNil(suite.T(), resp.GetQuoteResponse)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), resp.GetQuoteResponse.Note)
}

func (suite *GetQuoteTestSuite) TestValidQuoteNonEUNonEUWithAcctProdServiceRequest() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &DCTFrom{}
	from.CountryCode = "SG"
	from.PostalCode = "520110"

	to := &DCTTo{}
	to.CountryCode = "AU"
	to.PostalCode = "2007"

	t := time.Now()
	bdr := &BkgDetailsRequest{}
	bdr.PaymentCountryCode = "SG"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &Pieces{
		Piece: []PieceType{
			{PieceID: "1", Height: 1, Depth: 1, Width: 1, Weight: 5.0},
		},
	}
	bdr.PaymentAccountNumber = "CASHSIN"
	bdr.IsDutiable = "N"
	bdr.NetworkTypeCode = "AL"
	bdr.QtdShp = &QtdShpRequest{
		GlobalProductCode: "D",
		LocalProductCode:  "D",
		QtdShpExChrg: &QtdShpExChrgRequest{
			SpecialServiceType: "AA",
		},
	}

	du := &DCTDutiable{}
	du.DeclaredCurrency = "EUR"
	du.DeclaredValue = 1.0

	resp, err := client.GetQuote(from, to, bdr, du)

	assert.NotNil(suite.T(), resp.GetQuoteResponse)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), resp.GetQuoteResponse.Note)
}

func (suite *GetQuoteTestSuite) TestValidQuoteEUToNonEUWithAcctProdInsuranceRequest() {
	config := ClientConfig{Host: "staging"}
	client, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &DCTFrom{}
	from.CountryCode = "BE"
	from.PostalCode = "1020"

	to := &DCTTo{}
	to.CountryCode = "AU"
	to.PostalCode = "2020"

	t := time.Now()
	bdr := &BkgDetailsRequest{}
	bdr.PaymentCountryCode = "BE"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &Pieces{
		Piece: []PieceType{
			{PieceID: "1", Height: 20, Depth: 30, Width: 10, Weight: 20.0},
		},
	}
	bdr.PaymentAccountNumber = "272317228"
	bdr.IsDutiable = "Y"
	bdr.NetworkTypeCode = "AL"
	bdr.QtdShp = &QtdShpRequest{
		GlobalProductCode: "P",
		LocalProductCode:  "S",
	}
	bdr.InsuredValue = 800.0
	bdr.InsuredCurrency = "EUR"

	du := &DCTDutiable{}
	du.DeclaredCurrency = "EUR"
	du.DeclaredValue = 1002.0

	resp, err := client.GetQuote(from, to, bdr, du)

	assert.NotNil(suite.T(), resp.GetQuoteResponse)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), resp.GetQuoteResponse.Note)
}

func TestGetQuoteTestSuite(t *testing.T) {
	suite.Run(t, new(GetQuoteTestSuite))
}
