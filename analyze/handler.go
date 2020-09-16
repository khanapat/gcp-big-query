package analyze

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
	"krungthai.com/khanapat/backend-web-big-query/response"
)

type handler struct {
	GetMerchantFn        GetMerchantFn
	GetMerchantRawDataFn GetMerchantRawDataFn
}

func NewHandler(getMerchantFn GetMerchantFn, getMerchantRawDataFn GetMerchantRawDataFn) *handler {
	return &handler{
		GetMerchantFn:        getMerchantFn,
		GetMerchantRawDataFn: getMerchantRawDataFn,
	}
}

func (h *handler) InquiryRaw(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req InquiryAnalyzedDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		zap.L().Error("RAW ERROR", zap.Error(err))
		json.NewEncoder(w).Encode(response.NewResponseError(0, "4001", "Cannot bind request"))
		return
	}
	result, err := h.GetMerchantFn(req, ctx)
	if err != nil {
		zap.L().Error("RAW ERROR", zap.Error(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)
	zap.L().Info("ok")
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
		zap.L().Error("RAW ERROR", zap.Error(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)
	zap.L().Info("ok")
}
