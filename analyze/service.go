package analyze

import (
	"context"
	"fmt"
	"net/http"

	"krungthai.com/khanapat/backend-web-big-query/response"
)

type GetMerchantRawDataFn func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error)
type GetMerchantSummaryFn func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error)

func NewGetMerchantRawDataFn(inquiryMerchantRawDataFn InquiryMerchantRawDataFn) GetMerchantRawDataFn {
	return func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error) {
		resp, err := inquiryMerchantRawDataFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}

func NewGetMerchantSummaryFn(inquiryMerchantSummaryFn InquiryMerchantSummaryFn, inquiryMaleMerchantFn InquiryMaleMerchantFn, inquiryFemaleMerchantFn InquiryFemaleMerchantFn, inquiryTopSubMerchantFn InquiryTopSubMerchantFn) GetMerchantSummaryFn {
	return func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error) {
		summary, err := inquiryMerchantSummaryFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		maleNo, err := inquiryMaleMerchantFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		femaleNo, err := inquiryFemaleMerchantFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		topSubMerchant, err := inquiryTopSubMerchantFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		resp := InquiryAnalyzedDataResponse{
			MerchantSubCategoryNumber: summary.MerchantSubCategoryNumber,
			AverageAmount:             summary.AverageAmount,
			PurchasingPower:           fmt.Sprintf("%f-%f", summary.MinimumAmount, summary.MaximumAmount),
			Age:                       "20",
			Male:                      float64(maleNo / summary.TotalTransaction),
			Female:                    float64(femaleNo / summary.TotalTransaction),
			Salary:                    summary.AverageSalary,
			TopSubCategory:            *topSubMerchant,
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}
