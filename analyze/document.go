package analyze

import "time"

type AnalyzedData struct {
	MerchantID           string    `json:"merchantId"`
	MerchantName         string    `json:"merchantName"`
	MerchantCategory     string    `json:"merchantCategory"`
	MerchantSubCategory  string    `json:"merchantSubCategory"`
	MerchantLatitude     string    `json:"merchantLatitude"`
	MerchantLongitude    string    `json:"merchantLongitude"`
	MerchantType         string    `json:"merchantType"`
	MerchantBranch       string    `json:"merchantBranch"`
	IsCreditCardAccepted bool      `json:"isCreditCardAccepted"`
	Amount               float64   `json:"amount"`
	Province             string    `json:"province"`
	TimeStamp            time.Time `json:"timeStamp"`
	PaymentType          string    `json:"paymentType"`
	Gender               string    `json:"gender"`
	Fee                  float64   `json:"fee"`
	InstallmentPlan      string    `json:"installmentPlan"`
	Salary               float64   `json:"salary"`
	Age                  int       `json:"age"`
}

type SummaryData struct {
	MerchantSubCategoryNumber int     `json:"merchantSubCategoryNumber"`
	AverageAmount             float64 `json:"averageAmount"`
	MinimumAmount             float64 `json:"minimumAmount"`
	MaximumAmount             float64 `json:"maximumAmount"`
	AverageSalary             float64 `json:"averageSalary"`
}
