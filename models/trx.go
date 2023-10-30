package models

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	DocNo              string  `json:"docNo"`
	DocDate            string  `json:"docDate"`
	ExtDocNo           string  `json:"extDocNo"`
	IdempotencyKey     string  `json:"idempotencyKey"`
	CheckInDatetime    string  `json:"checkInDatetime"`
	CheckOutDatetime   string  `json:"checkOutDatetime"`
	CheckInTime        int64   `json:"checkInTime"`
	CheckOutTime       int64   `json:"checkOutTime"`
	DurationTime       int64   `json:"durationTime"`
	OuID               int64   `json:"ouId"`
	OuCode             string  `json:"ouCode"`
	OuName             string  `json:"ouName"`
	OuSubBranchID      int64   `json:"ouSubBranchId"`
	OuSubBranchCode    string  `json:"ouSubBranchCode"`
	OuSubBranchName    string  `json:"ouSubBranchName"`
	MerchantKey        string  `json:"merchant_key"`
	ProductID          int64   `json:"productId"`
	ProductCode        string  `json:"productCode"`
	ProductName        string  `json:"productName"`
	Price              float64 `json:"price"`
	BaseTime           int     `json:"baseTime"`
	ProgressiveTime    int     `json:"progressiveTime"`
	ProgressivePrice   float64 `json:"progressivePrice"`
	IsPct              string  `json:"isPct"`
	ProgressivePct     int     `json:"progressivePct"`
	MaxPrice           float64 `json:"maxPrice"`
	Is24H              string  `json:"is24H"`
	OvernightTime      string  `json:"overnightTime"`
	OvernightPrice     float64 `json:"overnightPrice"`
	GracePeriod        int     `json:"gracePeriod"`
	DropOffTime        int     `json:"dropOffTime"`
	ServiceFee         float64 `json:"serviceFee"`
	GrandTotal         float64 `json:"grandTotal"`
	LogTrx             string  `json:"logTrx"`
	PaymentMethod      string  `json:"paymentMethod"`
	Mdr                float64 `json:"mdr"`
	Mid                string  `json:"mid"`
	Tid                string  `json:"tid"`
	ResponseTrxCode    string  `json:"responseTrxCode"`
	Status             string  `json:"status"`
	StatusDesc         string  `json:"statusDesc"`
	VehicleNumberIn    string  `json:"vehicleNumberIn"`
	VehicleNumberOut   string  `json:"vehicleNumberOut"`
	ExtLocalDatetime   string  `json:"extLocalDatetime"`
	SettlementDatetime *string `json:"settlementDatetime"`
	DeductDatetime     *string `json:"deductDatetime"`
	PathImageIn        string  `json:"pathImageIn"`
	PathImageOut       string  `json:"pathImageOut"`
	CreatedAt          string  `json:"createdAt"`
	CreatedBy          string  `json:"createdBy"`
	UpdatedAt          string  `json:"updatedAt"`
	UpdatedBy          string  `json:"updatedBy"`
	PaymentRefDocNo    string  `json:"paymentRefDocNo"`
	RefDocNo           string  `json:"refDocNo"`
	FlgRepeat          string  `json:"flgRepeat"`
}

type TrxRequest struct {
	DocNo              string  `json:"docNo"`
	DocDate            string  `json:"docDate"`
	ExtDocNo           string  `json:"extDocNo"`
	IdempotencyKey     string  `json:"idempotencyKey"`
	CheckInDatetime    string  `json:"checkInDatetime"`
	CheckOutDatetime   string  `json:"checkOutDatetime"`
	CheckInTime        int64   `json:"checkInTime"`
	CheckOutTime       int64   `json:"checkOutTime"`
	DurationTime       int64   `json:"durationTime"`
	OuID               int64   `json:"ouId"`
	OuCode             string  `json:"ouCode"`
	OuName             string  `json:"ouName"`
	OuSubBranchID      int64   `json:"ouSubBranchId"`
	OuSubBranchCode    string  `json:"ouSubBranchCode"`
	OuSubBranchName    string  `json:"ouSubBranchName"`
	MerchantKey        string  `json:"merchant_key"`
	ProductID          int64   `json:"productId"`
	ProductCode        string  `json:"productCode"`
	ProductName        string  `json:"productName"`
	Price              float64 `json:"price"`
	BaseTime           int     `json:"baseTime"`
	ProgressiveTime    int     `json:"progressiveTime"`
	ProgressivePrice   float64 `json:"progressivePrice"`
	IsPct              string  `json:"isPct"`
	ProgressivePct     int     `json:"progressivePct"`
	MaxPrice           float64 `json:"maxPrice"`
	Is24H              string  `json:"is24H"`
	OvernightTime      string  `json:"overnightTime"`
	OvernightPrice     float64 `json:"overnightPrice"`
	GracePeriod        int     `json:"gracePeriod"`
	DropOffTime        int     `json:"dropOffTime"`
	ServiceFee         float64 `json:"serviceFee"`
	GrandTotal         float64 `json:"grandTotal"`
	LogTrx             string  `json:"logTrx"`
	PaymentMethod      string  `json:"paymentMethod"`
	Mdr                float64 `json:"mdr"`
	Mid                string  `json:"mid"`
	Tid                string  `json:"tid"`
	ResponseTrxCode    string  `json:"responseTrxCode"`
	Status             string  `json:"status"`
	StatusDesc         string  `json:"statusDesc"`
	VehicleNumberIn    string  `json:"vehicleNumberIn"`
	VehicleNumberOut   string  `json:"vehicleNumberOut"`
	ExtLocalDatetime   string  `json:"extLocalDatetime"`
	SettlementDatetime *string `json:"settlementDatetime"`
	DeductDatetime     *string `json:"deductDatetime"`
	PathImageIn        string  `json:"pathImageIn"`
	PathImageOut       string  `json:"pathImageOut"`
	CreatedAt          string  `json:"createdAt"`
	CreatedBy          string  `json:"createdBy"`
	UpdatedAt          string  `json:"updatedAt"`
	UpdatedBy          string  `json:"updatedBy"`
	PaymentRefDocNo    string  `json:"paymentRefDocNo"`
	RefDocNo           string  `json:"refDocNo"`
	FlgRepeat          string  `json:"flgRepeat"`
}

type Date struct {
	Start string `json:"startdate"`
	End   string `json:"enddate"`
}
