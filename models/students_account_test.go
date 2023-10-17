package models

import (
	"reflect"
	"testing"
)

func TestStudentAccount_Map(t *testing.T) {
	type fields struct {
		ID                    int
		StudentFee            float64
		DefaultStudent        bool
		StudentOnInstallment  bool
		TotalInstallments     int
		RemainingInstallments int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - convert Student Account struct to map",
			fields: fields{
				ID:                    1,
				StudentFee:            2000.0,
				DefaultStudent:        false,
				StudentOnInstallment:  false,
				TotalInstallments:     0,
				RemainingInstallments: 0,
			},
			want: map[string]interface{}{
				"id":                     1,
				"student_fee":            2000.0,
				"default_student":        false,
				"student_on_installment": false,
				"total_installments":     0,
				"remaining_installments": 0,
			},
		},
		{
			name: "success - convert student account struct to map2",
			fields: fields{
				ID:                    1,
				StudentFee:            1000.0,
				DefaultStudent:        false,
				StudentOnInstallment:  true,
				TotalInstallments:     6,
				RemainingInstallments: 2,
			},
			want: map[string]interface{}{
				"id":                     1,
				"student_fee":            1000.0,
				"default_student":        false,
				"student_on_installment": true,
				"total_installments":     6,
				"remaining_installments": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentAccount{
				ID:                    tt.fields.ID,
				StudentFee:            tt.fields.StudentFee,
				DefaultStudent:        tt.fields.DefaultStudent,
				StudentOnInstallment:  tt.fields.StudentOnInstallment,
				TotalInstallments:     tt.fields.TotalInstallments,
				RemainingInstallments: tt.fields.RemainingInstallments,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentAccount_Names(t *testing.T) {
	type fields struct {
		ID                    int
		StudentFee            float64
		DefaultStudent        bool
		StudentOnInstallment  bool
		TotalInstallments     int
		RemainingInstallments int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student account struct fields",
			fields: fields{
				ID:                    125,
				StudentFee:            1500.0,
				DefaultStudent:        false,
				StudentOnInstallment:  false,
				TotalInstallments:     0,
				RemainingInstallments: 0,
			},
			want: []string{"id", "student_fee", "default_student", "student_on_installment", "total_installments", "remaining_installments"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentAccount{
				ID:                    tt.fields.ID,
				StudentFee:            tt.fields.StudentFee,
				DefaultStudent:        tt.fields.DefaultStudent,
				StudentOnInstallment:  tt.fields.StudentOnInstallment,
				TotalInstallments:     tt.fields.TotalInstallments,
				RemainingInstallments: tt.fields.RemainingInstallments,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
