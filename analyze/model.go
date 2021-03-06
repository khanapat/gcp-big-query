package analyze

type InquiryAnalyzedDataRequest struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Distance            float64 `json:"distance"`
	MerchantCategory    string  `json:"merchantCategory"`
	MerchantSubCategory string  `json:"merchantSubCategory"`
	MerchantDateTime    string  `json:"merchantDatetime"`
}

type InquiryAnalyzedDataResponse struct {
	MerchantSubCategoryNumber int              `json:"merchantSubCategoryNumber"`
	AverageAmount             float64          `json:"averageAmount"`
	PurchasingPowerMax        float64          `json:"purchasingPowerMax"`
	PurchasingPowerMin        float64          `json:"purchasingPowerMin"`
	Age                       float64          `json:"age"`
	AgeRange                  []int            `json:"ageRange"`
	Male                      float64          `json:"male"`
	Female                    float64          `json:"female"`
	Salary                    float64          `json:"salary"`
	TopSubCategory            []TopSubMerchant `json:"topSubCategory"`
}
