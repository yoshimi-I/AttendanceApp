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
	r.Post("/user", cr.UserController.CreateUser())

	// 現在の活動を取得
	r.Get("/user/status/{userKey}", cr.UserController.GetStatus())

	// 活動の全履歴を取得
	r.Get("/study/allHistory/{userKey}/{year}", cr.HistoryController.GetAllHistory())

	// 指定日の勉強履歴を取得
	r.Get("/study/history/{userKey}/{date}", cr.HistoryController.GetHistoryByDate())

	// 作業の開始を追加
	r.Post("/study/activity/work/start", cr.ActivityController.AddStartWork())

	// 休憩の開始を追加
	r.Post("/study/activity/break/start", cr.ActivityController.AddStartBreak())

	// 作業の終了を追加
	r.Post("/study/activity/work/end", cr.ActivityController.AddEndWork())

	// 休憩の終了を追加
	r.Post("/study/activity/break/end", cr.ActivityController.AddEndBreak())

	// 指定の活動を更新
	r.Put("/study/activity/update", cr.ActivityController.UpdateActivity())

	// 指定の活動を削除
	r.Delete("/study/activity/delete", cr.ActivityController.DeleteActivity())

	return r
}
