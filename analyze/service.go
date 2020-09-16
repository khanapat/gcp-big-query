package analyze

import (
	"context"
	"net/http"

	"krungthai.com/khanapat/backend-web-big-query/response"
)

type GetMerchantFn func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error)
type GetMerchantRawDataFn func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error)

func NewGetMerchantFn(inquiryMerchantByLatLongFn InquiryMerchantByLatLongFn) GetMerchantFn {
	return func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error) {
		resp, err := inquiryMerchantByLatLongFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}

func NewGetMerchantRawDataFn(inquiryMerchantRawDataFn InquiryMerchantRawDataFn) GetMerchantRawDataFn {
	return func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error) {
		resp, err := inquiryMerchantRawDataFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}
