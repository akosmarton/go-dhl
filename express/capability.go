package express

import (
	"encoding/xml"

	log "github.com/sirupsen/logrus"
)

// GetCapability takes DCTFrom, DCTTo, BkgDetailsRequest and DCTDutiable and makes DCTRequest with GetQuote.
// Returns a DCTResponse with GetCapabilityResponse
func (c *dhlExpressClient) GetCapability(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) (*DCTResponse, error) {
	data := c.buildCapabilityRequest(from, to, details, dutiable)
	contents, err := c.fetch(data)
	if err != nil {
		return nil, err
	}
	log.Debugf("GetCapability Response Body: %s\n", string(*contents))

	var dctResponse DCTResponse
	if err := xml.Unmarshal(*contents, &dctResponse); err != nil {
		return nil, err
	}

	return &dctResponse, nil
}

func (c *dhlExpressClient) buildCapabilityRequest(from *DCTFrom, to *DCTTo, details *BkgDetailsRequest, dutiable *DCTDutiable) *DCTRequest {
	sh := NewServiceHeader(c.siteID, c.password)
	q := GetCapability{
		Request: &Request{
			ServiceHeader: &sh,
		},
		From:       from,
		To:         to,
		BkgDetails: details,
		Dutiable:   dutiable,
	}
	return &DCTRequest{
		XPNameSpace:     "http://www.dhl.com",
		XP1NameSpace:    "http://www.dhl.com/datatypes",
		XP2NameSpace:    "http://www.dhl.com/DCTRequestdatatypes",
		XXSINameSpace:   "http://www.w3.org/2001/XMLSchema-instance",
		XSchemaLocation: "http://www.dhl.com DCT-req.xsd ",
		GetCapability:   &q,
	}
}
