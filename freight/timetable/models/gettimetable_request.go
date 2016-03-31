package models

type GetTimeTableRequest struct {
	XMLName             string `xml:"http://dhltimetable.dhl.com/ GetTimeTable"`
	User                string `xml:"strUser"`
	Password            string `xml:"strPwd"`
	Type                int    `xml:"intType"` // 1 = styckesgods
	CountryCode         string `xml:"strCountryCode"`
	PickupPostCode      string `xml:"strPickupPostCode"`
	PickupPlace         string `xml:"strPickupPlace"`
	DeliveryCountryCode string `xml:"strDeliveryCountryCode"` //SE or NO
	DeliveryPostCode    string `xml:"strDeliveryPostCode"`
	DeliveryPlace       string `xml:"strDeliveryPlace"`
	EarliestSent        bool   `xml:"blnEarliestSent"`
	Date                string `xml:"dtDate"` // Datetime string YYYY-MM-DD
	HolidayCheck        bool   `xml:"blnHolidayCheck"`
}

type SoapRequestEnvelope struct {
	XMLName       string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	XSINameSpace  string `xml:"xmlns:xsi,attr"`
	XSDNameSpace  string `xml:"xmlns:xsd,attr"`
	SOAPNameSpace string `xml:"xmlns:soap,attr"`
	Body          *SoapRequestBody
}

type SoapRequestBody struct {
	XMLName string `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Request *GetTimeTableRequest
}
