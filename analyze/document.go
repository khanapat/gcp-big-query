package analyze

import "time"

type AnalyzedData struct {
	MerchantID           string    `bigquery:"merchant_id"`
	MerchantName         string    `bigquery:"merchant_name"`
	MerchantCategory     string    `bigquery:"merchant_category"`
	MerchantSubCategory  string    `bigquery:"merchant_sub_category"`
	MerchantLatLog       string    `bigquery:"merchant_latlog"`
	MerchantType         string    `bigquery:"merchant_type"`
	MerchantBranch       string    `bigquery:"merchant_branch"`
	IsCreditCardAccepted bool      `bigquery:"is_credit_card_accepted"`
	Amount               float64   `bigquery:"amount"`
	Province             string    `bigquery:"province"`
	TimeStamp            time.Time `bigquery:"time_stamp"`
	PaymentType          string    `bigquery:"payment_type"`
	Gender               string    `bigquery:"gender"`
	Fee                  float64   `bigquery:"fee"`
	InstallmentPlan      string    `bigquery:"installment_plan"`
	Salary               float64   `bigquery:"salary"`
	Age                  int       `bigquery:"age"`
}

type SummaryData struct {
	MerchantSubCategoryNumber int     `bigquery:"merchant_sub_category_number"`
	AverageAmount             float64 `bigquery:"average_amount"`
	MinimumAmount             float64 `bigquery:"minimum_amount"`
	MaximumAmount             float64 `bigquery:"maximum_amount"`
	AverageSalary             float64 `bigquery:"average_salary"`
	TotalTransaction          int     `bigquery:"total_transaction"`
}

type CountAge struct {
	Age      int `bigquery:"age"`
	CountAge int `bigquery:"count_age"`
}

type TopSubMerchant struct {
	MerchantSubCategory       string `bigquery:"merchant_sub_category"`
	MerchantSubCategoryNumber int    `bigquery:"merchant_sub_category_number"`
}
