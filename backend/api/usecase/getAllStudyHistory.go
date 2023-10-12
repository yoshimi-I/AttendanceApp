package usecase

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetAllStudyHistory(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")
	// ここに指定日の履歴を取得するロジックを実装
	w.Write([]byte("Study history for date: " + date))
}
