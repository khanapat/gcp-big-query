package analyze

import (
	"context"
	"database/sql"
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type InquiryMerchantByLatLongFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error)

// type InquiryMerchantSummaryFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*SummaryData, error)

type InquiryMerchantRawDataFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error)

func NewInquiryMerchantByLatLongFn(db *sql.DB) InquiryMerchantByLatLongFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error) {
		var merchants []AnalyzedData
		rows, err := db.QueryContext(ctx, "GetLocation @merchantCate=?, @merchantSubCate=?, @merchantTime=?, @lat=?, @long=?, @distance=?", merchantCategory, merchantSubCategory, merchantDatetime, latitude, longitude, distance)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var merchant AnalyzedData
			if err := rows.Scan(&merchant.MerchantID, &merchant.MerchantName, &merchant.MerchantCategory, &merchant.MerchantSubCategory,
				&merchant.MerchantLatitude, &merchant.MerchantLongitude, &merchant.MerchantType, &merchant.MerchantBranch,
				&merchant.IsCreditCardAccepted, &merchant.Amount, &merchant.Province, &merchant.TimeStamp, &merchant.PaymentType,
				&merchant.Gender, &merchant.Fee, &merchant.InstallmentPlan, &merchant.Salary, &merchant.Age); err != nil {
				return nil, err
			}
			merchants = append(merchants, merchant)
		}
		zap.L().Info(fmt.Sprintf("Inquiry Data Receipt - Success"))
		return &merchants, nil
	}
}

// func NewInquiryMerchantSummaryFn(db *sql.DB) InquiryMerchantSummaryFn {
// 	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*SummaryData, error) {
// 		var summaryData SummaryData
// 		row := db.QueryRowContext(ctx, "SELECT COUNT(DISTINCT merchant_id), AVG(amount), MAX(amount), MIN(amount), AVG(salary)\n"+
// 			"FROM to_be_number_one.dbo.analyzed_data")

// 	}
// }

func NewInquiryMerchantRawDataFn(db *bigquery.Client) InquiryMerchantRawDataFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error) {
		var merchants []AnalyzedData
		q := db.Query("ok")
		it, err := q.Read(ctx)
		if err != nil {
			return nil, err
		}
		for {
			var merchant AnalyzedData
			err := it.Next(&merchant)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			merchants = append(merchants, merchant)
		}
		zap.L().Info(fmt.Sprintf("Inquiry Data Receipt - Success"))
		return &merchants, nil
	}
}
