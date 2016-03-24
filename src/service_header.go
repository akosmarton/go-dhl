package dhl

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// NewServiceHeader creates a new service header struct
func NewServiceHeader(siteID string, password string) ServiceHeader {
	t := time.Now().Format(xmlDateTimeFormat)
	messRef := generateMessageReference()
	sh := ServiceHeader{
		MessageTime:      t,
		MessageReference: messRef,
		SiteID:           siteID,
		Password:         password,
	}

	return sh
}

// ServiceHeader is generated from an XSD element
type ServiceHeader struct {
	MessageTime      string `xml:"MessageTime"`
	MessageReference string `xml:"MessageReference"`
	SiteID           string `xml:"SiteID"`
	Password         string `xml:"Password"`
}

func generateMessageReference() string {
	noDigits := randomInt(28, 33)
	digits := ""

	for i := 0; i < noDigits; i++ {
		digit := randomInt(0, 10)
		digits += strconv.Itoa(digit)
	}

	return digits
}

func randomInt(low int, high int) int {
	d := rand.Float64()*float64((high-low)) + float64(low)
	return int(math.Floor(d))
}
