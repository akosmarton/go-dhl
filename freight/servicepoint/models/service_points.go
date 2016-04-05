package models

type ServicePoints struct {
	XMLName            string               `xml:"ServicePoints"`
	NearbyServicePoint []NearbyServicePoint `xml:"http://DHL.ServicePoint.DataContracts/2008/10 NearbyServicePoint"`
}
