package timetable

import "errors"

// TimeTableQuery used for doing GetTimeTable requests
type TimeTableQuery struct {
	OriginCountryCode      string
	OriginPostCode         string
	OriginPlace            string
	DestinationCountryCode string
	DestinationPostCode    string
	DestinationPlace       string
	Date                   string
	EarliestSent           bool
	HolidayCheck           bool
}

// Validate is used to validate query input
func (q *TimeTableQuery) Validate() error {
	if q.OriginCountryCode == "" {
		return errors.New("OriginCountryCode is required")
	}

	if q.OriginPostCode == "" {
		return errors.New("OriginPostCode is required")
	}

	if q.DestinationCountryCode == "" {
		return errors.New("DestinationCountryCode is required")
	}

	if q.DestinationPostCode == "" {
		return errors.New("DestinationPostCode is required")
	}

	if q.Date == "" {
		return errors.New("Date is required")
	}

	return nil
}
