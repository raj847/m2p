package service

import (
	"m2p/models"
	"m2p/repository"
	"m2p/service/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type TrxService struct {
	trxRepository *repository.TrxRepository
}

func NewTrxService(
	trxRepository *repository.TrxRepository,
) *TrxService {
	return &TrxService{
		trxRepository: trxRepository,
	}
}

func (t *TrxService) Create(c echo.Context) error {

	var tReq models.Date
	// var DataPsql *models.Trx
	err := c.Bind(&tReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	data, _, err := t.trxRepository.GetData(tReq.Start, c.Request().Context(), tReq.End)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to get data from mongo",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	for _, v := range data {
		DataPsql := &models.Trx{
			DocNo:              v.DocNo,
			DocDate:            v.DocDate,
			ExtDocNo:           v.ExtDocNo,
			IdempotencyKey:     v.IdempotencyKey,
			CheckInDatetime:    v.CheckInDatetime,
			CheckOutDatetime:   v.CheckOutDatetime,
			CheckInTime:        v.CheckInTime,
			CheckOutTime:       v.CheckOutTime,
			DurationTime:       v.DurationTime,
			OuID:               v.OuID,
			OuCode:             v.OuCode,
			OuName:             v.OuName,
			OuSubBranchID:      v.OuSubBranchID,
			OuSubBranchCode:    v.OuSubBranchCode,
			OuSubBranchName:    v.OuSubBranchName,
			MerchantKey:        v.MerchantKey,
			ProductID:          v.ProductID,
			ProductCode:        v.ProductCode,
			ProductName:        v.ProductName,
			Price:              v.Price,
			BaseTime:           v.BaseTime,
			ProgressiveTime:    v.ProgressiveTime,
			ProgressivePrice:   v.ProgressivePrice,
			IsPct:              v.IsPct,
			ProgressivePct:     v.ProgressivePct,
			MaxPrice:           v.MaxPrice,
			Is24H:              v.Is24H,
			OvernightTime:      v.OvernightTime,
			OvernightPrice:     v.OvernightPrice,
			GracePeriod:        v.GracePeriod,
			DropOffTime:        v.DropOffTime,
			ServiceFee:         v.ServiceFee,
			GrandTotal:         v.GrandTotal,
			LogTrx:             v.LogTrx,
			PaymentMethod:      v.PaymentMethod,
			Mdr:                v.Mdr,
			Mid:                v.Mid,
			Tid:                v.Tid,
			ResponseTrxCode:    v.ResponseTrxCode,
			Status:             v.Status,
			StatusDesc:         v.StatusDesc,
			VehicleNumberIn:    v.VehicleNumberIn,
			VehicleNumberOut:   v.VehicleNumberOut,
			ExtLocalDatetime:   v.ExtLocalDatetime,
			SettlementDatetime: v.SettlementDatetime,
			DeductDatetime:     v.DeductDatetime,
			PathImageIn:        v.PathImageIn,
			PathImageOut:       v.PathImageOut,
			CreatedAt:          v.CreatedAt,
			CreatedBy:          v.CreatedBy,
			UpdatedAt:          v.UpdatedAt,
			UpdatedBy:          v.UpdatedBy,
			PaymentRefDocNo:    v.PaymentRefDocNo,
			RefDocNo:           v.RefDocNo,
			FlgRepeat:          v.FlgRepeat,
		}
		err = t.trxRepository.CreateTrx(c.Request().Context(), DataPsql)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to add data to postgres",
						Code:    "BAD_REQUEST",
					},
				},
			})
		}
	}

	return c.JSON(http.StatusOK, &data)
}
