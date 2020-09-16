package analyze

type AnalyzedData struct {
	MerchantID           string  `json:"merchantId"`
	MerchantName         string  `json:"merchantName"`
	MerchantCategory     string  `json:"merchantCategory"`
	MerchantSubCategory  string  `json:"merchantSubCategory"`
	MerchantType         string  `json:"merchantType"`
	MerchantBranch       string  `json:"merchantBranch"`
	IsCreditCardAccepted bool    `json:"isCreditCardAccepted"`
	Amount               float64 `json:"amount"`
	Province             string  `json:"province"`
	PaymentType          string  `json:"paymentType"`
	Gender               string  `json:"gender"`
	Fee                  float64 `json:"fee"`
	InstallmentPlan      string  `json:"installmentPlan"`
}
