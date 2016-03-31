package models

type RouteResponse struct {
	Response             *Response
	Note                 *Note
	RegionCode           string
	GMTNegativeIndicator string
	GMTOffset            string
	ServiceArea          *ServiceArea
}

type Note struct {
	ActionNote string
	Condition  *Condition
}

type ServiceArea struct {
	ServiceAreaCode string
	Description     string
}
