package models

type GetTimeTableResponse struct {
	XMLName            string              `xml:"http://dhltimetable.dhl.com/ GetTimeTableResponse"`
	GetTimeTableResult *GetTimeTableResult `xml:"GetTimeTableResult,omitempty"`
}

type GetTimeTableResult struct {
	SCHTimeTable *SCHTimeTable `xml:"http://dhltimetable.dhl.com/schTimeTable.xsd schTimeTable,omitempty"`
}

type SCHTimeTable struct {
	XMLName      string         `xml:"http://dhltimetable.dhl.com/schTimeTable.xsd schTimeTable"`
	Information  *Information   `xml:"Information,omitempty"`
	TimeTable    []TimeTable    `xml:"TimeTable,omitempty"`
	TimeTableOrg []TimeTableOrg `xml:"TimeTableOrg,omitempty"`
}

type Information struct {
	PickupCountryCode   string `xml:"PickupCountryCode,omitempty"`
	PickupPostCode      string `xml:"PickupPostCode,omitempty"`
	PickupPlace         string `xml:"PickupPlace,omitempty"`
	DeliveryCountryCode string `xml:"DeliveryCountryCode,omitempty"`
	DeliveryPostCode    string `xml:"DeliveryPostCode,omitempty"`
	DeliveryPlace       string `xml:"DeliveryPlace,omitempty"`
}

type TimeTable struct {
	Product            string `xml:"Product,omitempty"`
	PickupDateString   string `xml:"PickupDateString,omitempty"`
	DeliveryDateString string `xml:"DeliveryDateString,omitempty"`
	Options            string `xml:"Options,omitempty"`
	PickupDate         string `xml:"PickupDate,omitempty"`   // Datetime string
	DeliveryDate       string `xml:"DeliveryDate,omitempty"` // Datetime string
}

type TimeTableOrg struct {
	SAvgTerminal string `xml:"sAvgTerminal,omitempty"`
	OPickup      string `xml:"oPickup,omitempty"`
	OAvgTerm     string `xml:"oAvgTerm,omitempty"`
	OAnkTerm     string `xml:"oAnkTerm,omitempty"`
	OAvgDistTerm string `xml:"oAvgDistTerm,omitempty"`
	OLossar      string `xml:"oLossar,omitempty"`
	SGTjanst     string `xml:"sGTjanst,omitempty"`
}

type SoapResponseEvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	XSINameSpace  string `xml:"xmlns:xsi,attr"`
	XSDNameSpace  string `xml:"xmlns:xsd,attr"`
	SOAPNameSpace string `xml:"xmlns:soap,attr"`
	Body          *SoapResponseBody
}

type SoapResponseBody struct {
	XMLName  string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Response *GetTimeTableResponse
	Fault    *SoapFault
}

type SoapFault struct {
	XMLName     string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}
