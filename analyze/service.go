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

func NewGetMerchantSummaryFn(inquiryMerchantSummaryFn InquiryMerchantSummaryFn, inquiryMaleMerchantFn InquiryMaleMerchantFn, inquiryFemaleMerchantFn InquiryFemaleMerchantFn, inquiryCountAgeFn InquiryCountAgeFn, inquiryTopSubMerchantFn InquiryTopSubMerchantFn) GetMerchantSummaryFn {
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
		countAge, err := inquiryCountAgeFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantSubCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		topSubMerchant, err := inquiryTopSubMerchantFn(req.Latitude, req.Longitude, req.Distance, req.MerchantCategory, req.MerchantDateTime, ctx)
		if err != nil {
			return response.NewResponseError(http.StatusInternalServerError, "501", "Merchant not found."), err
		}
		var oneFive int
		var twoEight int
		var fourOne int
		var fiveFour int
		var sixSeven int
		for _, data := range countAge {
			age := data.Age
			switch {
			case age >= 67 && age <= 80:
				sixSeven += data.CountAge
			case age >= 54:
				fiveFour += data.CountAge
			case age >= 41:
				fourOne += data.CountAge
			case age >= 28:
				twoEight += data.CountAge
			case age >= 15:
				oneFive += data.CountAge
			}
		}
		var agingRange []int
		agingRange = append(agingRange, oneFive, twoEight, fourOne, fiveFour, sixSeven)
		resp := InquiryAnalyzedDataResponse{
			MerchantSubCategoryNumber: summary.MerchantSubCategoryNumber,
			AverageAmount:             summary.AverageAmount / float64(summary.MerchantSubCategoryNumber),
			PurchasingPowerMax:        summary.MaximumAmount,
			PurchasingPowerMin:        summary.MinimumAmount,
			Age:                       summary.AverageAge,
			AgeRange:                  agingRange,
			Male:                      (float64(maleNo) / float64(summary.TotalTransaction)) * 100,
			Female:                    (float64(femaleNo) / float64(summary.TotalTransaction)) * 100,
			Salary:                    summary.AverageSalary,
			TopSubCategory:            *topSubMerchant,
		}
		return response.NewResponse(http.StatusOK, "200", "success", resp), nil
	}
}
