package usecase

import (
	"reflect"
	"testing"
	"work-management-app/application/dto/response"
	"work-management-app/domain/repository"
	"work-management-app/domain/service"
)

func TestHistoryUsecaseImpl_AllHistory(t *testing.T) {
	type fields struct {
		hr repository.HistoryRepository
		ur repository.UserRepository
		hs service.HistoryDomainService
	}
	type args struct {
		userKey string
		year    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []response.ActivityTimeResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HistoryUsecaseImpl{
				hr: tt.fields.hr,
				ur: tt.fields.ur,
				hs: tt.fields.hs,
			}
			got, err := h.AllHistory(tt.args.userKey, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoryUsecaseImpl_HistoryByDate(t *testing.T) {
	type fields struct {
		hr repository.HistoryRepository
		ur repository.UserRepository
		hs service.HistoryDomainService
	}
	type args struct {
		userKey string
		date    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.HistoryByDateDto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HistoryUsecaseImpl{
				hr: tt.fields.hr,
				ur: tt.fields.ur,
				hs: tt.fields.hs,
			}
			got, err := h.HistoryByDate(tt.args.userKey, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoryByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HistoryByDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHistoryUsecase(t *testing.T) {
	type args struct {
		hr repository.HistoryRepository
		ur repository.UserRepository
		hs service.HistoryDomainService
	}
	tests := []struct {
		name string
		args args
		want HistoryUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHistoryUsecase(tt.args.hr, tt.args.ur, tt.args.hs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHistoryUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
