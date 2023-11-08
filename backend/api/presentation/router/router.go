package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"work-management-app/infrastructure/di"
	middle "work-management-app/presentation/router/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middle.Cors().Handler)
	r.Use(middleware.Logger)

	// DIコンテナ
	// コントローラの作成
	cr, err := di.InitializeControllers()
	if err != nil {
		log.Fatal("Failed to initialize controllers: ", err)

	}

	// ユーザーの登録
	r.Get("/user", cr.UserController.CreateUser())

	// 活動の全履歴を取得
	r.Get("/study/allHistory/{userId}/{year}", cr.HistoryController.GetAllHistory())

	// 指定日の勉強履歴を取得
	r.Get("/study/history/{userId}/{date}", cr.HistoryController.GetHistoryByDate())

	// 活動の開始を追加
	r.Post("/study/activity", cr.ActivityController.AddStartActivity())

	// 活動の終了を追加
	r.Put("/study/activity/{activityId}/end", cr.ActivityController.AddEndActivity())

	// 指定の活動を更新
	r.Put("/study/activity/{activityId}/update", cr.ActivityController.UpdateActivity())

	// 指定の活動を削除
	r.Delete("/study/activity/{activityId}/delete", cr.ActivityController.DeleteActivity())

	return r
}
