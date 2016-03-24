package dhl

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// NewServiceHeader creates a new service header struct
func NewServiceHeader(siteID string, password string) *ServiceHeader {
	t := time.Now() //.Format(xmlDateTimeFormat)
	messRef := generateMessageReference()
	sh := &ServiceHeader{
		MessageTime:      t,
		MessageReference: messRef,
		SiteID:           siteID,
		Password:         password,
	}

	return sh
}

// ServiceHeader is generated from an XSD element
type ServiceHeader struct {
	MessageTime      time.Time `xml:"MessageTime"`
	MessageReference string    `xml:"MessageReference"`
	SiteID           string    `xml:"SiteID"`
	Password         string    `xml:"Password"`
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

// /**
//   * Randomly generates "A string, peferably number, to uniquely identify
//   * individual messages. Minimum length must be 28 and maximum length is 32".
//   */
//  function generateMessageReference () {
//    var numberOfDigits, randomDigits, digit;
//    numberOfDigits = randomInt(28, 33);
//    randomDigits = [];
//
//    for (var i = 0; i < numberOfDigits; i++) {
//      digit = randomInt(0, 10);
//      randomDigits.push(digit)
//    }
//
//    return randomDigits.join('');
//  }

// // Generates a random int in [low, high)
// function randomInt (low, high) {
//
// }
