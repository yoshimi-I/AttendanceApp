package usecase

import (
	"reflect"
	"testing"
	"work-management-app/application/dto/request"
	"work-management-app/application/dto/response"
	"work-management-app/domain/repository"
	"work-management-app/domain/service"
)

func TestActivityUsecaseImpl_AddEndBreak(t *testing.T) {
	type fields struct {
		db repository.Transaction
		ar repository.ActivityRepository
		ur repository.UserRepository
		as service.ActivityDomainService
	}
	type args struct {
		breakInfo *request.ActivityRequestDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ActivityResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityUsecaseImpl{
				db: tt.fields.db,
				ar: tt.fields.ar,
				ur: tt.fields.ur,
				as: tt.fields.as,
			}
			got, err := a.AddEndBreak(tt.args.breakInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEndBreak() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEndBreak() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityUsecaseImpl_AddEndWork(t *testing.T) {
	type fields struct {
		db repository.Transaction
		ar repository.ActivityRepository
		ur repository.UserRepository
		as service.ActivityDomainService
	}
	type args struct {
		work *request.ActivityRequestDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ActivityResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityUsecaseImpl{
				db: tt.fields.db,
				ar: tt.fields.ar,
				ur: tt.fields.ur,
				as: tt.fields.as,
			}
			got, err := a.AddEndWork(tt.args.work)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEndWork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEndWork() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityUsecaseImpl_AddStarWork(t *testing.T) {
	type fields struct {
		db repository.Transaction
		ar repository.ActivityRepository
		ur repository.UserRepository
		as service.ActivityDomainService
	}
	type args struct {
		work *request.ActivityRequestDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ActivityResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityUsecaseImpl{
				db: tt.fields.db,
				ar: tt.fields.ar,
				ur: tt.fields.ur,
				as: tt.fields.as,
			}
			got, err := a.AddStarWork(tt.args.work)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStarWork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStarWork() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityUsecaseImpl_AddStartBreak(t *testing.T) {
	type fields struct {
		db repository.Transaction
		ar repository.ActivityRepository
		ur repository.UserRepository
		as service.ActivityDomainService
	}
	type args struct {
		breakInfo *request.ActivityRequestDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ActivityResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityUsecaseImpl{
				db: tt.fields.db,
				ar: tt.fields.ar,
				ur: tt.fields.ur,
				as: tt.fields.as,
			}
			got, err := a.AddStartBreak(tt.args.breakInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStartBreak() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStartBreak() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityUsecaseImpl_DeleteByActivityID(t *testing.T) {
	type fields struct {
		db repository.Transaction
		ar repository.ActivityRepository
		ur repository.UserRepository
		as service.ActivityDomainService
	}
	type args struct {
		activity *request.ActivityDeleteRequestDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityUsecaseImpl{
				db: tt.fields.db,
				ar: tt.fields.ar,
				ur: tt.fields.ur,
				as: tt.fields.as,
			}
			if err := a.DeleteByActivityID(tt.args.activity); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByActivityID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivityUsecaseImpl_Update(t *testing.T) {
	type fields struct {
		db repository.Transaction
		ar repository.ActivityRepository
		ur repository.UserRepository
		as service.ActivityDomainService
	}
	type args struct {
		activity *request.ActivityEditRequestDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ActivityResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityUsecaseImpl{
				db: tt.fields.db,
				ar: tt.fields.ar,
				ur: tt.fields.ur,
				as: tt.fields.as,
			}
			got, err := a.Update(tt.args.activity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
