package models

var (
	xmlDateFormat     = "2006-01-02"
	xmlDateTimeFormat = "2006-01-02T15:04:05"
	xmlTimeForm       = "PT15H04M"
)

// DCTRequest request object
type DCTRequest struct {
	XMLName         string         `xml:"p:DCTRequest"`
	XPNameSpace     string         `xml:"xmlns:p,attr"`
	XP1NameSpace    string         `xml:"xmlns:p1,attr"`
	XP2NameSpace    string         `xml:"xmlns:p2,attr"`
	XXSINameSpace   string         `xml:"xmlns:xsi,attr"`
	XSchemaLocation string         `xml:"xsi:schemaLocation,attr"`
	GetQuote        *GetQuote      `xml:"GetQuote,omitempty"`
	GetCapability   *GetCapability `xml:"GetCapability,omitempty"`
}

// GetQuote request object
type GetQuote struct {
	Request    *Request           `xml:"Request,omitempty"`
	From       *DCTFrom           `xml:"From,omitempty"`
	BkgDetails *BkgDetailsRequest `xml:"BkgDetails,omitempty"`
	To         *DCTTo             `xml:"To,omitempty"`
	Dutiable   *DCTDutiable       `xml:"Dutiable,omitempty"`
}

// GetCapability request object
type GetCapability struct {
	Request    *Request           `xml:"Request,omitempty"`
	From       *DCTFrom           `xml:"From,omitempty"`
	BkgDetails *BkgDetailsRequest `xml:"BkgDetails,omitempty"`
	To         *DCTTo             `xml:"To,omitempty"`
	Dutiable   *DCTDutiable       `xml:"Dutiable,omitempty"`
}

// Request request object
type Request struct {
	ServiceHeader *ServiceHeader `xml:"ServiceHeader,omitempty"`
}

// DCTFrom request object
type DCTFrom struct {
	CountryCode string `xml:"CountryCode,omitempty"`
	PostalCode  string `xml:"Postalcode,omitempty"`
	City        string `xml:"City,omitempty"`
	Suburb      string `xml:"Suburb,omitempty"`
	VatNo       string `xml:"VatNo,omitempty"`
}

// DCTTo request object
type DCTTo struct {
	CountryCode string `xml:"CountryCode,omitempty"`
	PostalCode  string `xml:"Postalcode,omitempty"`
	City        string `xml:"City,omitempty"`
	Suburb      string `xml:"Suburb,omitempty"`
	VatNo       string `xml:"VatNo,omitempty"`
}

// BkgDetailsRequest request object
type BkgDetailsRequest struct {
	PaymentCountryCode   string         `xml:"PaymentCountryCode,omitempty"`
	Date                 string         `xml:"Date,omitempty"`               // Date format YYYY-MM-DD (xmlDateFormat)
	ReadyTime            string         `xml:"ReadyTime,omitempty"`          // Time in hours and minutes (HH:MM)
	ReadyTimeGMTOffset   string         `xml:"ReadyTimeGMTOffset,omitempty"` // Time in hours and minutes (+-HH:MM)
	DimensionUnit        string         `xml:"DimensionUnit,omitempty"`      // Dimension Unit I (inches);Centimeters (CM)
	WeightUnit           string         `xml:"WeightUnit,omitempty"`         //Kilogram (KG),Pounds (LB)
	NumberOfPieces       int            `xml:"NumberOfPieces,omitempty"`
	ShipmentWeight       float32        `xml:"ShipmentWeight,omitempty"`
	Volume               float32        `xml:"Volume,omitempty"`
	MaxPieceWeight       float32        `xml:"MaxPieceWeight,omitempty"`
	MaxPieceHeight       float32        `xml:"MaxPieceHeight,omitempty"`
	MaxPieceDepth        float32        `xml:"MaxPieceDepth,omitempty"`
	MaxPieceWidth        float32        `xml:"MaxPieceWidth,omitempty"`
	Pieces               *Pieces        `xml:"Pieces,omitempty"`
	PaymentAccountNumber string         `xml:"PaymentAccountNumber,omitempty"`
	IsDutiable           string         `xml:"IsDutiable,omitempty"`      //Y - Dutiable/Non-Doc, N - Non-dutiable/Doc
	NetworkTypeCode      string         `xml:"NetworkTypeCode,omitempty"` // DD - TD - AL
	QtdShp               *QtdShpRequest `xml:"QtdShp,omitempty"`
	CODAmount            float32        `xml:"CODAmount,omitempty"`
	CODCurrencyCode      string         `xml:"CODCurrencyCode,omitempty"`
	CODAccountNumber     string         `xml:"CODAccountNumber,omitempty"`
	InsuredValue         float32        `xml:"InsuredValue,omitempty"`
	InsuredCurrency      string         `xml:"InsuredCurrency,omitempty"`
}

// Pieces request object
type Pieces struct {
	Piece []PieceType `xml:"Piece"`
}

// PieceType request object
type PieceType struct {
	PieceID         string  `xml:"PieceID,omitempty"`
	PackageTypeCode string  `xml:"PackageTypeCode,omitempty"` // Default BOX - FLY, COY, NCY, PAL, DBL, BOX
	Height          float32 `xml:"Height,omitempty"`          // required if width and depth are specified
	Depth           float32 `xml:"Depth,omitempty"`           // required if width and height are specified
	Width           float32 `xml:"Width,omitempty"`           // required if height and depth are specified
	Weight          float32 `xml:"Weight,omitempty"`
}

// QtdShpRequest request object
type QtdShpRequest struct {
	GlobalProductCode string               `xml:"GlobalProductCode,omitempty"`
	LocalProductCode  string               `xml:"LocalProductCode,omitempty"`
	QtdShpExChrg      *QtdShpExChrgRequest `xml:"QtdShpExChrg,omitempty"`
}

// QtdShpExChrgRequest request object
type QtdShpExChrgRequest struct {
	SpecialServiceType      string `xml:"SpecialServiceType,omitempty"`
	LocalSpecialServiceType string `xml:"LocalSpecialServiceType,omitempty"`
}

// DCTDutiable request object
type DCTDutiable struct {
	DeclaredCurrency string  `xml:"DeclaredCurrency"`
	DeclaredValue    float32 `xml:"DeclaredValue"`
}
