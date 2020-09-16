package analyze

type InquiryAnalyzedDataRequest struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Distance            float64 `json:"distance"`
	MerchantCategory    string  `json:"merchantCategory"`
	MerchantSubCategory string  `json:"merchantSubCategory"`
	PaymentDateTime     string  `json:"paymentDatetime"`
}

type InquiryAnalyzedDataResponse struct {
	MerchantSubCategoryNumber int      `json:"merchantSubCategoryNumber"`
	AverageAmount             float64  `json:"averageAmount"`
	PurchasingPower           string   `json:"purchasingPower"`
	Age                       string   `json:"age"`
	Male                      float64  `json:"male"`
	Female                    float64  `json:"female"`
	Salary                    float64  `json:"salary"`
	TopSubCategory            []string `json:"topSubCategory"`
}
