package usecase

import (
	"reflect"
	"testing"
	"work-management-app/application/dto/request"
	"work-management-app/application/dto/response"
	"work-management-app/domain/repository"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		ur repository.UserRepository
		ar repository.ActivityRepository
	}
	tests := []struct {
		name string
		args args
		want UserUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.ur, tt.args.ar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecaseImpl_AddUser(t *testing.T) {
	type fields struct {
		ur repository.UserRepository
		ar repository.ActivityRepository
	}
	type args struct {
		user *request.UserDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserUsecaseImpl{
				ur: tt.fields.ur,
				ar: tt.fields.ar,
			}
			got, err := u.AddUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecaseImpl_IDByUserKey(t *testing.T) {
	type fields struct {
		ur repository.UserRepository
		ar repository.ActivityRepository
	}
	type args struct {
		userKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserUsecaseImpl{
				ur: tt.fields.ur,
				ar: tt.fields.ar,
			}
			gotId, err := u.IDByUserKey(tt.args.userKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("IDByUserKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("IDByUserKey() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestUserUsecaseImpl_UserByUserKey(t *testing.T) {
	type fields struct {
		ur repository.UserRepository
		ar repository.ActivityRepository
	}
	type args struct {
		userKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserUsecaseImpl{
				ur: tt.fields.ur,
				ar: tt.fields.ar,
			}
			got, err := u.UserByUserKey(tt.args.userKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserByUserKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserByUserKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecaseImpl_UserStatusByUserKey(t *testing.T) {
	type fields struct {
		ur repository.UserRepository
		ar repository.ActivityRepository
	}
	type args struct {
		userKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserStatusDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserUsecaseImpl{
				ur: tt.fields.ur,
				ar: tt.fields.ar,
			}
			got, err := u.UserStatusByUserKey(tt.args.userKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserStatusByUserKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserStatusByUserKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
