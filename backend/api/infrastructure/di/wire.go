//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	"github.com/google/wire"
	"work-management-app/infrastructure"
	"work-management-app/infrastructure/repository"
	"work-management-app/presentation/controller"
	"work-management-app/usecase"
)

// HistoryControllerの依存関係
var historySuperSet = wire.NewSet(
	infrastructure.InitDB,
	repository.NewHistoryRepository,
	usecase.NewHistoryUsecase,
	controller.NewHistoryController,
)

// ActivityController周りの依存関係
var activitySuperSet = wire.NewSet(
	infrastructure.InitDB,
	repository.NewActivityRepository,
	usecase.NewActivityUsecase,
	controller.NewActivityController,
)

// InitHistoryController HistoryControllerのインスタンスを初期化
//
//	Controllerが一番先頭の呼び出し関数のため
func InitHistoryController() (controller.HistoryController, error) {
	wire.Build(historySuperSet)
	return &controller.HistoryControllerImpl{}, nil
}

// InitActivityController InitActivityControllerのインスタンス初期化
func InitActivityController() (controller.ActivityController, error) {
	wire.Build(activitySuperSet)
	return &controller.ActivityControllerImpl{}, nil

}
