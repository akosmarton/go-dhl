package client

import "errors"

// NearestServicePointQuery used for doing GetTimeTable requests
type NearestServicePointsQuery struct {
	CountryCode      string
	Street           string
	PostCode         string
	City             string
	FeatureCodes     []string
	MaxNumberOfItems int
	BitCatCodes      []string
}

// Validate is used to validate query input
func (q *NearestServicePointsQuery) Validate() error {
	if q.CountryCode == "" {
		return errors.New("CountryCode is required")
	}

	if q.PostCode == "" {
		return errors.New("PostCode is required")
	}

	return nil
}
