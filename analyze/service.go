package analyze

import (
	"context"
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

func NewGetMerchantSummaryFn(inquiryMerchantSummaryFn InquiryMerchantSummaryFn) GetMerchantSummaryFn {
	return func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error) {
		resp, err := inquiryMerchantSummaryFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}
