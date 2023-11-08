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

// infrastructure
var infrastructureSet = wire.NewSet(
	infrastructure.InitDB,
)

// repository
var repositorySet = wire.NewSet(
	repository.NewActivityRepository,
	repository.NewHistoryRepository,
	repository.NewUserRepository,
)

// usecase
var usecaseSet = wire.NewSet(
	usecase.NewActivityUsecase,
	usecase.NewHistoryUsecase,
	usecase.NewUserUsecase,
)

// controller
var controllerSet = wire.NewSet(
	controller.NewActivityController,
	controller.NewHistoryController,
	controller.NewUserController,
)

type ControllersSet struct {
	UserController     controller.UserController
	HistoryController  controller.HistoryController
	ActivityController controller.ActivityController
}

func InitializeControllers() (*ControllersSet, error) {
	wire.Build(
		infrastructureSet,
		repositorySet,
		usecaseSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil, nil
}
