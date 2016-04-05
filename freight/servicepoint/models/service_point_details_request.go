package models

type GetServicePointDetailRequest struct {
	XMLName         string           `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetServicePointDetailRequest"`
	ServicePointRef *ServicePointRef `xml:"http://DHL.ServicePoint.DataContracts/2008/10 ServicePointRef"`
}
