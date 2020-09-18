package analyze

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type InquiryMerchantRawDataFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error)
type InquiryMerchantSummaryFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*SummaryData, error)
type InquiryMaleMerchantFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (int, error)
type InquiryFemaleMerchantFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (int, error)
type InquiryCountAgeFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) ([]CountAge, error)
type InquiryTopSubMerchantFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantDatetime string, ctx context.Context) (*[]TopSubMerchant, error)

func NewInquiryMerchantRawDataFn(db *bigquery.Client) InquiryMerchantRawDataFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error) {
		var merchants []AnalyzedData
		query := fmt.Sprintf("SELECT * FROM bootcamp1_dataviz.masterData\n"+
			"WHERE ST_MAXDISTANCE(merchant_latlog, ST_GEOGPOINT(%f, %f)) <= %f AND merchant_category='%s' AND merchant_sub_category='%s' AND time_stamp>='%s';",
			longitude, latitude, distance, merchantCategory, merchantSubCategory, merchantDatetime)
		zap.L().Debug(query)
		q := db.Query(query)
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
		zap.L().Info(fmt.Sprintf("Inquiry Raw Data - Success"))
		return &merchants, nil
	}
}

func NewInquiryMerchantSummaryFn(db *bigquery.Client) InquiryMerchantSummaryFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*SummaryData, error) {
		var summary SummaryData
		query := fmt.Sprintf("SELECT COUNT(DISTINCT merchant_id) AS merchant_sub_category_number, AVG(amount) AS average_amount, MAX(amount) AS maximum_amount, MIN(amount) AS minimum_amount, AVG(salary) AS average_salary, COUNT(gender) AS total_transaction, AVG(age) AS average_age\n"+
			"FROM bootcamp1_dataviz.masterData\n"+
			"WHERE ST_MAXDISTANCE(merchant_latlog, ST_GEOGPOINT(%f, %f)) <= %f AND merchant_category='%s' AND merchant_sub_category='%s' AND time_stamp>='%s';",
			longitude, latitude, distance, merchantCategory, merchantSubCategory, merchantDatetime)
		zap.L().Debug(query)
		q := db.Query(query)
		it, err := q.Read(ctx)
		if err != nil {
			return nil, err
		}
		err = it.Next(&summary)
		if err != nil {
			return nil, err
		}
		zap.L().Info(fmt.Sprintf("Inquiry Summary - Success"))
		return &summary, nil
	}
}

func NewInquiryMaleMerchantFn(db *bigquery.Client) InquiryMaleMerchantFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (int, error) {
		var values map[string]bigquery.Value
		query := fmt.Sprintf("SELECT COUNT(*) AS male FROM bootcamp1_dataviz.masterData\n"+
			"WHERE ST_MAXDISTANCE(merchant_latlog, ST_GEOGPOINT(%f, %f)) <= %f AND merchant_category='%s' AND merchant_sub_category='%s' AND time_stamp>='%s' AND gender='ชาย'",
			longitude, latitude, distance, merchantCategory, merchantSubCategory, merchantDatetime)
		zap.L().Debug(query)
		q := db.Query(query)
		it, err := q.Read(ctx)
		if err != nil {
			return 0, err
		}
		err = it.Next(&values)
		if err != nil {
			return 0, err
		}
		male := values["male"].(int64)
		zap.L().Info(fmt.Sprintf("Inquiry male - Success"))
		return int(male), nil
	}
}

func NewInquiryFemaleMerchantFn(db *bigquery.Client) InquiryFemaleMerchantFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (int, error) {
		var values map[string]bigquery.Value
		query := fmt.Sprintf("SELECT COUNT(*) AS female FROM bootcamp1_dataviz.masterData\n"+
			"WHERE ST_MAXDISTANCE(merchant_latlog, ST_GEOGPOINT(%f, %f)) <= %f AND merchant_category='%s' AND merchant_sub_category='%s' AND time_stamp>='%s' AND gender='หญิง'",
			longitude, latitude, distance, merchantCategory, merchantSubCategory, merchantDatetime)
		zap.L().Debug(query)
		q := db.Query(query)
		it, err := q.Read(ctx)
		if err != nil {
			return 0, err
		}
		err = it.Next(&values)
		if err != nil {
			return 0, err
		}
		female := values["female"].(int64)
		zap.L().Info(fmt.Sprintf("Inquiry male - Success"))
		return int(female), nil
	}
}

func NewInquiryCountAgeFn(db *bigquery.Client) InquiryCountAgeFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) ([]CountAge, error) {
		var values []CountAge
		query := fmt.Sprintf("SELECT age, COUNT(age) AS count_age FROM bootcamp1_dataviz.masterData\n"+
			"WHERE ST_MAXDISTANCE(merchant_latlog, ST_GEOGPOINT(%f, %f)) <= %f AND merchant_category='%s' AND merchant_sub_category='%s' AND time_stamp>='%s'\n"+
			"GROUP BY age;",
			longitude, latitude, distance, merchantCategory, merchantSubCategory, merchantDatetime)
		zap.L().Debug(query)
		q := db.Query(query)
		it, err := q.Read(ctx)
		if err != nil {
			return nil, err
		}
		for {
			var value CountAge
			err := it.Next(&value)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		zap.L().Info(fmt.Sprintf("Inquiry Age - Success"))
		return values, nil
	}
}

func NewInquiryTopSubMerchantFn(db *bigquery.Client) InquiryTopSubMerchantFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantDatetime string, ctx context.Context) (*[]TopSubMerchant, error) {
		var topSubMerchants []TopSubMerchant
		query := fmt.Sprintf("SELECT merchant_sub_category, COUNT(*) AS merchant_sub_category_number FROM bootcamp1_dataviz.masterData\n"+
			"WHERE ST_MAXDISTANCE(merchant_latlog, ST_GEOGPOINT(%f, %f)) <= %f AND merchant_category='%s' AND time_stamp>='%s'\n"+
			"GROUP BY merchant_sub_category", longitude, latitude, distance, merchantCategory, merchantDatetime)
		zap.L().Debug(query)
		q := db.Query(query)
		it, err := q.Read(ctx)
		if err != nil {
			return nil, err
		}
		for {
			var topSubMerchant TopSubMerchant
			err := it.Next(&topSubMerchant)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			topSubMerchants = append(topSubMerchants, topSubMerchant)
		}
		zap.L().Info(fmt.Sprintf("Inquiry Top Sub Category - Success"))
		return &topSubMerchants, nil
	}
}

// func NewInquiryMerchantAgeFn(db *bigquery.Client)
