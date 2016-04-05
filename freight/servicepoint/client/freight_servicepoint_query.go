package client

import "errors"

// NearestServicePointsQuery used for doing GetNearestServicePoints requests
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

// ServicePointDetailQuery used for doing GetServicePointDetail requests
type ServicePointDetailQuery struct {
	ID          string
	DisplayName string
}

// Validate is used to validate query input
func (q *ServicePointDetailQuery) Validate() error {
	if q.ID == "" {
		return errors.New("ID is required")
	}

	return nil
}
