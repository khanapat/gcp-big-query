package analyze

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type InquiryMerchantByLatLongFn func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error)

func NewInquiryMerchantByLatLongFn(db *sql.DB) InquiryMerchantByLatLongFn {
	return func(latitude float64, longitude float64, distance float64, merchantCategory string, merchantSubCategory string, merchantDatetime string, ctx context.Context) (*[]AnalyzedData, error) {
		var merchants []AnalyzedData
		// query := fmt.Sprintf("SELECT * FROM to_be_number_one.dbo.analyzed_data WHERE merchant_category=N'%s' AND merchant_sub_category=N'%s' AND time_stamp='%s' AND (GEOGRAPHY::Point(%f, %f, 4326)).STDistance(merchant_latlog) >= %f", merchantCategory, merchantSubCategory, merchantDatetime, latitude, longitude, distance)
		rows, err := db.QueryContext(ctx, "SELECT merchant_id, merchant_name, merchant_category, merchant_sub_category, merchant_type,\n"+
			"merchant_branch, is_credit_card_accepted, amount, province, payment_type, gender, fee, installment_plan\n"+
			"FROM to_be_number_one.dbo.analyzed_data")
		// zap.L().Debug(query)
		// rows, err := db.QueryContext(ctx, query)
		if err != nil {
			zap.L().Error(fmt.Sprintf("Inquiry Merchant - %s", err))
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var merchant AnalyzedData
			if err := rows.Scan(&merchant.MerchantID, &merchant.MerchantName, &merchant.MerchantCategory, &merchant.MerchantSubCategory,
				&merchant.MerchantType, &merchant.MerchantBranch, &merchant.IsCreditCardAccepted, &merchant.Amount, &merchant.Province,
				&merchant.PaymentType, &merchant.Gender, &merchant.Fee, &merchant.InstallmentPlan); err != nil {
				zap.L().Error(fmt.Sprintf("Scan inquiry data - %s", err))
				return nil, err
			}
			merchants = append(merchants, merchant)
		}
		zap.L().Info(fmt.Sprintf("Inquiry Data Receipt - Success"))
		return &merchants, nil
	}
}
