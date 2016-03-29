package dhl

// DCTResponse response object
type DCTResponse struct {
	GetQuoteResponse      *GetQuoteResponse      `xml:"GetQuoteResponse"`
	GetCapabilityResponse *GetCapabilityResponse `xml:"GetCapabilityResponse"`
}

// GetQuoteResponse response object
type GetQuoteResponse struct {
	Response   *Response           `xml:"Response"`
	BkgDetails *BkgDetailsResponse `xml:"BkgDetails"`
	Srvs       *Srvs               `xml:"Srvs"`
	Note       *NoteType           `xml:"Note"`
}

// GetCapabilityResponse response object
type GetCapabilityResponse struct {
	Response   *Response           `xml:"Response"`
	BkgDetails *BkgDetailsResponse `xml:"BkgDetails"`
	Srvs       *Srvs               `xml:"Srvs"`
	Note       *NoteType           `xml:"Note"`
}

// Response response object
type Response struct {
	ServiceHeader *ServiceHeader `xml:"ServiceHeader,omitempty"`
	Status        *Status        `xml:"Status,omitempty"`
}

// BkgDetailsResponse response object
type BkgDetailsResponse struct {
	OriginServiceArea      *OrgnSvcAreaType `xml:"OriginServiceArea"`
	DestinationServiceArea *DestSvcAreaType `xml:"DestinationServiceArea"`
	QtdShp                 []QtdShpResponse `xml:"QtdShp"`
	CalcNextDayInd         string           `xml:"CalcNextDayInd"`
}

// OrgnSvcAreaType response object
type OrgnSvcAreaType struct {
	FacilityCode    string `xml:"FacilityCode"`
	ServiceAreaCode string `xml:"ServiceAreaCode"`
}

// DestSvcAreaType response object
type DestSvcAreaType struct {
	FacilityCode    string `xml:"FacilityCode"`
	ServiceAreaCode string `xml:"ServiceAreaCode"`
}

// QtdShpResponse response object
type QtdShpResponse struct {
	GlobalProductCode         string                  `xml:"GlobalProductCode,omitempty"`
	LocalProductCode          string                  `xml:"LocalProductCode,omitempty"`
	ProductShortName          string                  `xml:"ProductShortName,omitempty"`
	LocalProductName          string                  `xml:"LocalProductName,omitempty"`
	NetworkTypeCode           string                  `xml:"NetworkTypeCode,omitempty"`
	POfferedCustAgreement     string                  `xml:"POfferedCustAgreement,omitempty"`
	TransInd                  string                  `xml:"TransInd,omitempty"`
	PickupDate                string                  `xml:"PickupDate,omitempty"`
	PickupCutoffTime          string                  `xml:"PickupCutoffTime,omitempty"`
	BookingTime               string                  `xml:"BookingTime,omitempty"`
	CurrencyCode              string                  `xml:"CurrencyCode,omitempty"`
	ExchangeRate              float32                 `xml:"ExchangeRate"`
	WeightCharge              float32                 `xml:"WeightCharge"`
	WeightChargeTax           float32                 `xml:"WeightChargeTax"`
	WeightChargeTaxRate       float32                 `xml:"weightChargeTaxRate"`
	TotalTransitDays          int                     `xml:"TotalTransitDays"`
	PickupPostalLocAddDays    int                     `xml:"PickupPostalLocAddDays"`
	DeliveryPostalLocAddDays  int                     `xml:"DeliveryPostalLocAddDays"`
	PickupNonDHLCourierCode   string                  `xml:"PickupNonDHLCourierCode,omitempty"`
	DeliveryNonDHLCourierCode string                  `xml:"DeliveryNonDHLCourierCode,omitempty"`
	DeliveryCheckpointReturn  string                  `xml:"DeliveryCheckpointReturn,omitempty"`
	DeliveryDate              string                  `xml:"DeliveryDate,omitempty"`
	DeliveryTime              string                  `xml:"DeliveryTime,omitempty"`
	DeliveryTimeGMTOffset     string                  `xml:"DeliveryTimeGMTOffset,omitempty"`
	DimensionalWeight         float32                 `xml:"DimensionalWeight"`
	WeightUnit                string                  `xml:"WeightUnit,omitempty"`
	PickupDayOfWeekNum        string                  `xml:"PickupDayOfWeekNum,omitempty"`
	DestinationDayOfWeekNum   string                  `xml:"DestinationDayOfWeekNum,omitempty"`
	QtdShpExChrg              *QtdShpExChrgResponse   `xml:"QtdShpExChrg"`
	PricingDate               string                  `xml:"PricingDate,omitempty"`
	ShippingCharge            float32                 `xml:"ShippingCharge"`
	TotalTaxAmount            float32                 `xml:"TotalTaxAmount"`
	QtdSInAdCur               *QtdSInAdCurType        `xml:"QtdSInAdCur"`
	WeightChargeTaxDet        *WeightChargeTaxDetType `xml:"WeightChargeTaxDet"`
}

// QtdShpExChrgResponse response object
type QtdShpExChrgResponse struct {
	SpecialServiceType      string                   `xml:"SpecialServiceType,omitempty"`
	LocalSpecialServiceType string                   `xml:"LocalSpecialServiceType,omitempty"`
	GlobalServiceName       string                   `xml:"GlobalServiceName,omitempty"`
	LocalServiceTypeName    string                   `xml:"LocalServiceTypeName,omitempty"`
	SOfferedCustAgreement   string                   `xml:"SOfferedCustAgreement,omitempty"`
	ChargeCodeType          string                   `xml:"ChargeCodeType,omitempty"`
	InsPrmRateInPercentage  float32                  `xml:"InsPrmRateInPercentage"`
	CurrencyCode            string                   `xml:"CurrencyCode,omitempty"`
	ChargeValue             float32                  `xml:"ChargeValue"`
	ChargeTaxAmount         float32                  `xml:"ChargeTaxAmount"`
	ChargeTaxRate           float32                  `xml:"ChargeTaxRate"`
	ChargeTaxAmountDet      *ChargeTaxAmountDetType  `xml:"ChargeTaxAmountDet"`
	QtdSExtrChrgInAdCur     *QtdSExtrChrgInAdCurType `xml:"QtdSExtrChrgInAdCur"`
}

// QtdSInAdCurType response object
type QtdSInAdCurType struct {
	CustomsValue         float32                 `xml:"CustomsValue"`
	ExchangeRate         float32                 `xml:"ExchangeRate"`
	CurrencyCode         string                  `xml:"CurrencyCode,omitempty"`
	CurrencyRoleTypeCode string                  `xml:"CurrencyRoleTypeCode,omitempty"`
	WeightCharge         float32                 `xml:"WeightCharge"`
	TotalAmount          float32                 `xml:"TotalAmount"`
	TotalTaxAmount       float32                 `xml:"TotalTaxAmount"`
	WeightChargeTax      float32                 `xml:"WeightChargeTax"`
	WeightChargeTaxDet   *WeightChargeTaxDetType `xml:"WeightChargeTaxDet"`
}

// WeightChargeTaxDetType response object
type WeightChargeTaxDetType struct {
	TaxTypeRate     float32 `xml:"TaxTypeRate"`
	TaxTypeCode     string  `xml:"TaxTypeCode,omitempty"`
	WeightChargeTax float32 `xml:"WeightChargeTax"`
	BaseAmt         float32 `xml:"BaseAmt"`
}

// ChargeTaxAmountDetType response object
type ChargeTaxAmountDetType struct {
	TaxTypeRate     float32 `xml:"TaxTypeRate"`
	TaxTypeCode     string  `xml:"TaxTypeCode,omitempty"`
	WeightChargeTax float32 `xml:"WeightChargeTax"`
	BaseAmt         float32 `xml:"BaseAmt"`
}

// QtdSExtrChrgInAdCurType response object
type QtdSExtrChrgInAdCurType struct {
	ChargeValue          float32                 `xml:"ChargeValue"`
	ChargeExchangeRate   float32                 `xml:"ChargeExchangeRate"`
	ChargeTaxAmount      float32                 `xml:"ChargeTaxAmount"`
	CurrencyCode         string                  `xml:"CurrencyCode,omitempty"`
	CurrencyRoleTypeCode string                  `xml:"CurrencyRoleTypeCode,omitempty"`
	ChargeTaxAmountDet   *ChargeTaxAmountDetType `xml:"ChargeTaxAmountDet"`
}

// Srvs response object
type Srvs struct {
	Srv []SrvType `xml:"Srv"`
}

// SrvType response object
type SrvType struct {
	GlobalProductCode string       `xml:"GlobalProductCode,omitempty"`
	MrkSrv            []MrkSrvType `xml:"MrkSrv"`
	SBTP              *SBTPType    `xml:"SBTP"`
}

// MrkSrvType response object
type MrkSrvType struct {
	LocalProductCode string `xml:"LocalProductCode,omitempty"`
	LocalServiceType string `xml:"LocalServiceType,omitempty"`

	ProductShortName  string `xml:"ProductShortName,omitempty"`
	GlobalServiceName string `xml:"GlobalServiceName,omitempty"`

	LocalProductName     string `xml:"LocalProductName,omitempty"`
	LocalServiceTypeName string `xml:"LocalServiceTypeName,omitempty"`

	ProductDesc string `xml:"ProductDesc,omitempty"`
	ServiceDesc string `xml:"ServiceDesc,omitempty"`

	NetworkTypeCode string `xml:"NetworkTypeCode,omitempty"`

	POfferedCustAgreement string `xml:"POfferedCustAgreement,omitempty"`
	SOfferedCustAgreement string `xml:"SOfferedCustAgreement,omitempty"`

	TransInd       string `xml:"TransInd,omitempty"`
	ChargeCodeType string `xml:"ChargeCodeType,omitempty"`
	MrkSrvInd      string `xml:"MrkSrvInd"`
}

// SBTPType response object
type SBTPType struct {
	Prod *ProdType `xml:"Prod"`
}

// ProdType response object
type ProdType struct {
	VldSrvComb []VldSrvComb `xml:"VldSrvComb"`
}

// VldSrvComb response object
type VldSrvComb struct {
	SpecialServiceType string     `xml:"SpecialServiceType,omitempty"`
	LocalServiceType   []string   `xml:"LocalServiceType,omitempty"`
	CombRSrv           []CombRSrv `xml:"CombRSrv"`
}

// CombRSrv response object
type CombRSrv struct {
	RestrictedSpecialServiceType string   `xml:"RestrictedSpecialServiceType,omitempty"`
	RestrictedLocalServiceType   []string `xml:"RestrictedLocalServiceType,omitempty"`
}

// NoteType response object
type NoteType struct {
	Condition *Condition `xml:"Condition"`
}

// Condition response object
type Condition struct {
	ConditionCode string `xml:"ConditionCode"`
	ConditionData string `xml:"ConditionData"`
}
