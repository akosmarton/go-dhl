package models

type GetNearestServicePointsResponse struct {
	XMLName       string `xml:"http://DHL.ServicePoint.ServiceContracts/2008/10 GetNearestServicePointsResponse"`
	ServicePoints *ServicePoints
}
