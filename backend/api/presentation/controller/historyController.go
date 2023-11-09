package controller

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"work-management-app/usecase"

	"github.com/go-chi/chi/v5"
)

type HistoryController interface {
	GetAllHistory() http.HandlerFunc
	GetHistoryByDate() http.HandlerFunc
}

type HistoryControllerImpl struct {
	// ここで使うusecaseを全て実装
	hu usecase.HistoryUsecase
}

func NewHistoryController(hu usecase.HistoryUsecase) HistoryController {
	return &HistoryControllerImpl{hu: hu}
}

func (h *HistoryControllerImpl) GetAllHistory() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// パラメータ取得
		userKey := chi.URLParam(request, "userKey")

		year, err := strconv.Atoi(chi.URLParam(request, "year"))
		if err != nil {
			http.Error(writer, "Invalid year format", http.StatusBadRequest)
			return
		}

		// Usecaseから勉強の全履歴を取得
		activities, err := h.hu.AllHistory(userKey, year)
		if err != nil {
			http.Error(writer, "Failed to retrieve study history", http.StatusInternalServerError)
			return
		}

		// 取得した履歴をJSONとしてレスポンスに書き込む
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(activities); err != nil {
			http.Error(writer, "Failed to encode study history to JSON", http.StatusInternalServerError)
			return
		}

	}
}

func (h *HistoryControllerImpl) GetHistoryByDate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userKey := chi.URLParam(request, "userKey")

		date := chi.URLParam(request, "date")

		// dateのフォーマットを確認
		r := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
		if !r.MatchString(date) {
			http.Error(writer, "invalid date format", http.StatusInternalServerError)
			return
		}

		// Usecaseを使用して指定された日付の勉強履歴を取得
		activity, err := h.hu.HistoryByDate(userKey, date)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		// レスポンスヘッダーにContent-Typeを設定
		writer.Header().Set("Content-Type", "application/json")

		// 勉強履歴をJSON形式でエンコードしてレスポンスボディに書き込む
		if err := json.NewEncoder(writer).Encode(activity); err != nil {
			http.Error(writer, "Failed to encode study activity to JSON", http.StatusInternalServerError)
			return
		}
	}
}
