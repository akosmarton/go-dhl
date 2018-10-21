package express

import (
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
	digits := ""
	for i := 0; i < 28+rand.Intn(4); i++ {
		digits += strconv.Itoa(rand.Intn(10))
	}

	return digits
}
