package service

import (
	"reflect"
	"testing"
	"work-management-app/domain/model"
)

func TestHistoryServiceImpl_TotalWorkByYear(t *testing.T) {
	type args struct {
		allHistoryList []model.Attendance
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HistoryServiceImpl{}
			if got := h.TotalWorkByYear(tt.args.allHistoryList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalWorkByYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
