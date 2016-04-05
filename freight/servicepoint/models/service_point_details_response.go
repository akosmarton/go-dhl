package models

type GetServicePointDetailResponse struct {
	XMLName            string `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetServicePointDetailResponse"`
	ServicePointDetail *ServicePointDetail
}

type ServicePointDetail struct {
	XMLName        string           `xml:"ServicePointDetail"`
	Identity       *ServicePointRef `xml:"http://DHL.ServicePoint.DataContracts/2008/10 Identity,omitempty"`
	ServiceAddress *ServiceAddress  `xml:"http://DHL.ServicePoint.DataContracts/2008/10 ServiceAddress,omitempty"`
	FeatureCodes   *FeatureCodes    `xml:"http://DHL.ServicePoint.DataContracts/2008/10 FeatureCodes,omitempty"`
}

type ServiceAddress struct {
	Address           *PhysicalAddress `xml:"Address,omitempty"`
	ContactPersonName string           `xml:"ContactPersonName,omitempty"`
	Telecom           *TelecomInfo     `xml:"Telecom,omitempty"`
}

type PhysicalAddress struct {
	AddresseeName string `xml:"AddresseeName,omitempty"`
	City          string `xml:"City,omitempty"`
	CountryCode   string `xml:"CountryCode,omitempty"`
	PostCode      string `xml:"PostCode,omitempty"`
	Street1       string `xml:"Street1,omitempty"`
	Street2       string `xml:"Street2,omitempty"`
}

type TelecomInfo struct {
	Email string `xml:"Email,omitempty"`
	Fax   string `xml:"Fax,omitempty"`
	Phone string `xml:"Phone,omitempty"`
}
