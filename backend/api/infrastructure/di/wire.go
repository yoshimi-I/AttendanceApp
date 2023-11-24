//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	"github.com/google/wire"
	"work-management-app/application/usecase"
	"work-management-app/domain/service"
	"work-management-app/infrastructure/database"
	"work-management-app/infrastructure/database/repository"
	"work-management-app/presentation/controller"
)

// infrastructure
var infrastructureSet = wire.NewSet(
	database.InitDB,
)

// domainService
var domainServiceSet = wire.NewSet(
	service.NewActivityDomainService,
	service.NewHistoryDomainService,
)

// repository
var repositorySet = wire.NewSet(
	repository.NewActivityRepository,
	repository.NewHistoryRepository,
	repository.NewUserRepository,
	repository.NewTransaction,
)

// application
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
		domainServiceSet,
		repositorySet,
		usecaseSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil, nil
}
