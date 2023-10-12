package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yoshimi-I/AttendanceApp/usecase"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	// 勉強の全履歴を取得
	r.Get("/study/allHistory/", usecase.GetAllStudyHistory)

	// 指定日の勉強履歴を取得
	r.Get("/study/history/{date}", usecase.GetStudyHistoryByDate)

	// 勉強の活動を追加
	r.Post("/study/activity", usecase.PostStudyActivity)

	// 指定の勉強活動を更新
	r.Put("/study/activity/{activityId}", usecase.UpdateStudyActivity)

	// 指定の勉強活動を削除
	r.Delete("/study/activity/{activityId}", usecase.DeleteStudyActivity)

	return r
}
