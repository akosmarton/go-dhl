package timetable

import (
	"fmt"
	"regexp"
)

var (
	dhlErrIDPattern  = regexp.MustCompile(`<DHLErrID>(.*)</DHLErrID>`)
	dhlErrMsgPattern = regexp.MustCompile(`<DHLErrMessage>(.*)</DHLErrMessage>`)
)

// DHLError is an internal DHL error representation
type DHLError struct {
	DHLErrID      string
	DHLErrMessage string
}

func (d DHLError) Error() string {
	return fmt.Sprintf("DHLError id=%s msg=%s", d.DHLErrID, d.DHLErrMessage)
}
