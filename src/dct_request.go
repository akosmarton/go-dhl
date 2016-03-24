package dhl

var (
	xmlDateFormat     = "2006-01-02"
	xmlDateTimeFormat = "2006-01-02T15:04:05"
)

// DCTRequest request object
type DCTRequest struct {
	GetQuote      GetQuote      `xml:"GetQuote"`
	GetCapability GetCapability `xml:"GetCapability"`
}

// GetQuote request object
type GetQuote struct {
	Request    Request        `xml:"Request"`
	From       DCTFrom        `xml:"From"`
	BkgDetails BkgDetailsType `xml:"BkgDetails"`
	To         DCTTo          `xml:"To"`
	Dutiable   DCTDutiable    `xml:"Dutiable"`
}

// GetCapability request object
type GetCapability struct {
	Request    Request        `xml:"Request"`
	From       DCTFrom        `xml:"From"`
	BkgDetails BkgDetailsType `xml:"BkgDetails"`
	To         DCTTo          `xml:"To"`
	Dutiable   DCTDutiable    `xml:"Dutiable"`
}

// Request request object
type Request struct {
	ServiceHeader ServiceHeader `xml:"ServiceHeader"`
}

// DCTFrom request object
type DCTFrom struct {
	CountryCode string `xml:"CountryCode"`
	PostalCode  string `xml:"PostalCode"`
	City        string `xml:"City"`
	Suburb      string `xml:"Suburb"`
	VatNo       string `xml:"VatNo"`
}

// DCTTo request object
type DCTTo struct {
	CountryCode string `xml:"CountryCode"`
	PostalCode  string `xml:"PostalCode"`
	City        string `xml:"City"`
	Suburb      string `xml:"Suburb"`
	VatNo       string `xml:"VatNo"`
}

// BkgDetailsType request object
type BkgDetailsType struct {
	PaymentCountryCode   string      `xml:"PaymentCountryCode"`
	Date                 string      `xml:"Date"`               // Date format YYYY-MM-DD (xmlDateFormat)
	ReadyTime            string      `xml:"ReadyTime"`          // Time in hours and minutes (HH:MM)
	ReadyTimeGMTOffset   string      `xml:"ReadyTimeGMTOffset"` // Time in hours and minutes (+-HH:MM)
	DimensionUnit        string      `xml:"DimensionUnit"`      // Dimension Unit I (inches);Centimeters (CM)
	WeightUnit           string      `xml:"WeightUnit"`         //Kilogram (KG),Pounds (LB)
	NumberOfPieces       int         `xml:"NumberOfPieces"`
	ShipmentWeight       float32     `xml:"ShipmentWeight"`
	Volume               float32     `xml:"Volume"`
	MaxPieceWeight       float32     `xml:"MaxPieceWeight"`
	MaxPieceHeight       float32     `xml:"MaxPieceHeight"`
	MaxPieceDepth        float32     `xml:"MaxPieceDepth"`
	MaxPieceWidth        float32     `xml:"MaxPieceWidth"`
	Pieces               []PieceType `xml:"Pieces"`
	PaymentAccountNumber string      `xml:"PaymentAccountNumber"`
	IsDutiable           string      `xml:"IsDutiable"`      //Y - Dutiable/Non-Doc, N - Non-dutiable/Doc
	NetworkTypeCode      string      `xml:"NetworkTypeCode"` // DD - TD - AL
	QtdShp               QtdShpType  `xml:"QtdShp"`
	CODAmount            float32     `xml:"CODAmount"`
	CODCurrencyCode      string      `xml:"CODCurrencyCode"`
	CODAccountNumber     string      `xml:"CODAccountNumber"`
	InsuredValue         float32     `xml:"InsuredValue"`
	InsuredCurrency      string      `xml:"CODAccountNumber"`
}

// PieceType request object
type PieceType struct {
	PieceID         string  `xml:"PieceID"`
	PackageTypeCode string  `xml:"PackageTypeCode"` // Default BOX - FLY, COY, NCY, PAL, DBL, BOX
	Height          float32 `xml:"Height"`          // required if width and depth are specified
	Depth           float32 `xml:"Depth"`           // required if width and height are specified
	Width           float32 `xml:"Width"`           // required if height and depth are specified
	Weight          float32 `xml:"Weight"`
}

// QtdShpType request object
type QtdShpType struct {
	GlobalProductCode string           `xml:"GlobalProductCode"`
	LocalProductCode  string           `xml:"LocalProductCode"`
	QtdShpExChrg      QtdShpExChrgType `xml:"QtdShpExChrg"`
}

// QtdShpExChrgType request object
type QtdShpExChrgType struct {
	SpecialServiceType      string `xml:"SpecialServiceType"`
	LocalSpecialServiceType string `xml:"LocalSpecialServiceType"`
}

// DCTDutiable request object
type DCTDutiable struct {
	DeclaredCurrency string  `xml:"DeclaredCurrency"`
	DeclaredValue    float32 `xml:"DeclaredValue"`
}
