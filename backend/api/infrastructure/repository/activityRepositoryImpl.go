package repository

import (
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func (a ActivityRepositoryImpl) PostStudyActivity(activity repository.Activity) error {
	return nil
}

func (a ActivityRepositoryImpl) PutStudyActivity(id int, activity repository.Activity) error {
	return nil
}

func (a ActivityRepositoryImpl) DeleteStudyActivity(id int) error {
	return nil
}

func NewActivityRespotiroy(db *gorm.DB) repository.ActivtyRepository {
	return &ActivityRepositoryImpl{db: db}
}
