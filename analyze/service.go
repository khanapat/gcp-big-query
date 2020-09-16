package analyze

import (
	"context"
	"net/http"

	"krungthai.com/khanapat/backend-web-big-query/response"
)

type GetMerchantFn func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error)

func NewGetMerchantFn(inquiryMerchantByLatLongFn InquiryMerchantByLatLongFn) GetMerchantFn {
	return func(req InquiryAnalyzedDataRequest, ctx context.Context) (response.Responser, error) {
		resp, err := inquiryMerchantByLatLongFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.PaymentDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}
