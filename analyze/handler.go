package analyze

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
	"krungthai.com/khanapat/backend-web-big-query/response"
)

type handler struct {
	GetMerchantRawDataFn GetMerchantRawDataFn
	GetMerchantSummaryFn GetMerchantSummaryFn
}

func NewHandler(getMerchantRawDataFn GetMerchantRawDataFn, getMerchantSummaryFn GetMerchantSummaryFn) *handler {
	return &handler{
		GetMerchantRawDataFn: getMerchantRawDataFn,
		GetMerchantSummaryFn: getMerchantSummaryFn,
	}
}

func (h *handler) InquiryBigQueryRaw(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req InquiryAnalyzedDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		zap.L().Error("RAW ERROR", zap.Error(err))
		json.NewEncoder(w).Encode(response.NewResponseError(0, "4001", "Cannot bind request"))
		return
	}
	result, err := h.GetMerchantRawDataFn(req, ctx)
	if err != nil {
		zap.L().Error("Inquiry Raw Data", zap.Error(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)
	zap.L().Info("ok")
}

func (h *handler) InquiryBigQuerySummary(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req InquiryAnalyzedDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		zap.L().Error("RAW ERROR", zap.Error(err))
		json.NewEncoder(w).Encode(response.NewResponseError(0, "4001", "Cannot bind request"))
		return
	}
	result, err := h.GetMerchantSummaryFn(req, ctx)
	if err != nil {
		zap.L().Error("Inquiry Summary Data", zap.Error(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)
	zap.L().Info("ok")
}
