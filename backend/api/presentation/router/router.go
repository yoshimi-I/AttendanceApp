package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yoshimi-I/AttendanceApp/infrastructure/di"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// DIコンテナ
	historyController, err := di.InitHistoryController()
	if err != nil {
	}
	activityController, err := di.InitActivityController()
	if err != nil {

	}
	// 活動の全履歴を取得
	r.Get("/study/allHistory/{userId}/{year}", historyController.GetAllHistory())

	// 指定日の勉強履歴を取得
	r.Get("/study/history/{userId}/{date}", historyController.GetHistoryByDate())

	// 活動の開始を追加
	r.Post("/study/activity", activityController.AddStartActivity())

	// 活動の終了を追加
	r.Put("/study/activity/{activityId}", activityController.AddEndActivity())

	// 指定の活動を更新
	r.Put("/study/activity/{activityId}", activityController.UpdateActivity())

	// 指定の活動を削除
	r.Delete("/study/activity/{activityId}", activityController.DeleteActivity())

	return r
}
