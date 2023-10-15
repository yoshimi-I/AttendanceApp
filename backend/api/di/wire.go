//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	"github.com/google/wire"
	"github.com/yoshimi-I/AttendanceApp/controller"
	"github.com/yoshimi-I/AttendanceApp/infrastructure"
	"github.com/yoshimi-I/AttendanceApp/infrastructure/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase"
)

// HistoryControllerの依存関係
var historySuperSet = wire.NewSet(
	infrastructure.InitDB,
	repository.NewHistoryRepository,
	usecase.NewHistoryUsecase,
	controller.NewHisoryController,
)

// ActivityController周りの依存関係
var activitySuperSet = wire.NewSet(
	infrastructure.InitDB,
	repository.NewActivityRespotiroy,
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
