//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	"github.com/google/wire"
	"github.com/yoshimi-I/AttendanceApp/infrastructure"
	"github.com/yoshimi-I/AttendanceApp/infrastructure/repository"
	controller2 "github.com/yoshimi-I/AttendanceApp/presentation/controller"
	"github.com/yoshimi-I/AttendanceApp/usecase"
)

// HistoryControllerの依存関係
var historySuperSet = wire.NewSet(
	infrastructure.InitDB,
	repository.NewHistoryRepository,
	usecase.NewHistoryUsecase,
	controller2.NewHisoryController,
)

// ActivityController周りの依存関係
var activitySuperSet = wire.NewSet(
	infrastructure.InitDB,
	repository.NewActivityRespotiroy,
	usecase.NewActivityUsecase,
	controller2.NewActivityController,
)

// InitHistoryController HistoryControllerのインスタンスを初期化
//
//	Controllerが一番先頭の呼び出し関数のため
func InitHistoryController() (controller2.HistoryController, error) {
	wire.Build(historySuperSet)
	return &controller2.HistoryControllerImpl{}, nil
}

// InitActivityController InitActivityControllerのインスタンス初期化
func InitActivityController() (controller2.ActivityController, error) {
	wire.Build(activitySuperSet)
	return &controller2.ActivityControllerImpl{}, nil

}
