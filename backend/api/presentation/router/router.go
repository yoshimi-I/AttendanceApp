package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yoshimi-I/AttendanceApp/di"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// DIコンテナ
	historyController, err := di.InitHistoryController()
	activityController, err := di.InitActivityController()
	if err != nil {
		panic(err)
	}
	// 勉強の全履歴を取得
	r.Get("/study/allHistory/", historyController.GetAllHistory())

	// 指定日の勉強履歴を取得
	r.Get("/study/history/{date}", historyController.GetHistoryByDate())

	// 勉強の活動を追加
	r.Post("/study/activity", activityController.AddActivity())

	// 指定の勉強活動を更新
	r.Put("/study/activity/{activityId}", activityController.UpdateActivity())

	// 指定の勉強活動を削除
	r.Delete("/study/activity/{activityId}", activityController.DeleteActivity())

	return r
}
