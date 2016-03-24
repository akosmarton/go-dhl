package dhl

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DHLClientTestSuite struct {
	suite.Suite
}

func (suite *DHLClientTestSuite) SetupTest() {

}

func (suite *DHLClientTestSuite) TestItCanGetAQoute() {
	client, _ := NewDHLStagingClient("DServiceVal", "testServVal")

	from := DCTFrom{}
	from.CountryCode = "SG"
	from.PostalCode = "100000"
	from.City = "Singapore"

	to := DCTTo{}
	to.CountryCode = "GB"
	to.PostalCode = "WC2E 9LA"
	to.City = "London"

	details := BkgDetailsRequest{}
	details.PaymentCountryCode = "SG"
	t := time.Now()
	details.Date = t.Format("2006-01-02")
	details.ReadyTime = t.Format("PT15H04M")
	details.DimensionUnit = "CM"
	details.WeightUnit = "KG"

	// details.Pieces = Pieces{
	// 	Piece: []PieceType{
	// 		PieceType{PieceID: "1", Height: 1, Depth: 1, Width: 5, Weight: 5.0},
	// 	},
	// }
	//details.PaymentAccountNumber = "CASHSIN"
	//details.IsDutiable = "Y"

	dutiable := &DCTDutiable{}
	dutiable.DeclaredCurrency = "SEK"
	dutiable.DeclaredValue = 0

	fmt.Println(details)

	resp, err := client.GetQuote(from, to, details, dutiable)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(resp.GetQuoteResponse.Response.ServiceHeader.SiteID)
	fmt.Print(len(resp.GetQuoteResponse.BkgDetails.QtdShp))
	assert.NotNil(suite.T(), *resp)
}

func TestDHLClientTestSuite(t *testing.T) {
	suite.Run(t, new(DHLClientTestSuite))
}
