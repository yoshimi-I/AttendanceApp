package service

import (
	"reflect"
	"testing"
	"time"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
)

func TestActivityServiceImpl_AddEndBreakTime(t *testing.T) {
	type fields struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	type args struct {
		userID int
		tx     repository.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityServiceImpl{
				ar: tt.fields.ar,
				hr: tt.fields.hr,
			}
			got, err := a.AddEndBreakTime(tt.args.userID, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEndBreakTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEndBreakTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityServiceImpl_AddEndWorkTime(t *testing.T) {
	type fields struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	type args struct {
		userID int
		tx     repository.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityServiceImpl{
				ar: tt.fields.ar,
				hr: tt.fields.hr,
			}
			got, err := a.AddEndWorkTime(tt.args.userID, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEndWorkTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEndWorkTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityServiceImpl_AddStarWorkTime(t *testing.T) {
	type fields struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	type args struct {
		userID int
		tx     repository.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityServiceImpl{
				ar: tt.fields.ar,
				hr: tt.fields.hr,
			}
			got, err := a.AddStarWorkTime(tt.args.userID, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStarWorkTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStarWorkTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityServiceImpl_AddStartBreakTime(t *testing.T) {
	type fields struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	type args struct {
		userID int
		tx     repository.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityServiceImpl{
				ar: tt.fields.ar,
				hr: tt.fields.hr,
			}
			got, err := a.AddStartBreakTime(tt.args.userID, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStartBreakTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStartBreakTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityServiceImpl_Delete(t *testing.T) {
	type fields struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	type args struct {
		attendance *model.Attendance
		tx         repository.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityServiceImpl{
				ar: tt.fields.ar,
				hr: tt.fields.hr,
			}
			got, err := a.Delete(tt.args.attendance, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityServiceImpl_EditTime(t *testing.T) {
	type fields struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	type args struct {
		activity *model.Attendance
		newTime  time.Time
		tx       repository.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ActivityServiceImpl{
				ar: tt.fields.ar,
				hr: tt.fields.hr,
			}
			got, err := a.EditTime(tt.args.activity, tt.args.newTime, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewActivityDomainService(t *testing.T) {
	type args struct {
		ar repository.ActivityRepository
		hr repository.HistoryRepository
	}
	tests := []struct {
		name string
		args args
		want ActivityDomainService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewActivityDomainService(tt.args.ar, tt.args.hr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActivityDomainService() = %v, want %v", got, tt.want)
			}
		})
	}
}
