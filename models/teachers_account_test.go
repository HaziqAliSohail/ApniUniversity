package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTeacherAccount_Map(t *testing.T) {
	type fields struct {
		ID         int
		TeacherPay float64
		CreatedAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - account struct to map",
			fields: fields{
				ID:         12,
				TeacherPay: 2000.0,
				CreatedAt:  time.Time{},
			},
			want: map[string]interface{}{
				"id":          12,
				"teacher_pay": 2000.0,
				"created_at":  time.Time{},
			},
		},
		{
			name: " success - student struct to map with fewer fields",
			fields: fields{
				TeacherPay: 30000.0,
				CreatedAt:  time.Time{},
			},
			want: map[string]interface{}{
				"id":          0,
				"teacher_pay": 30000.0,
				"created_at":  time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TeacherAccount{
				ID:         tt.fields.ID,
				TeacherPay: tt.fields.TeacherPay,
				CreatedAt:  tt.fields.CreatedAt,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeacherAccount_Names(t *testing.T) {
	type fields struct {
		ID         int
		TeacherPay float64
		CreatedAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:         1,
				TeacherPay: 30000.0,
				CreatedAt:  time.Time{},
			},
			want: []string{"id", "teacher_pay", "created_at"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TeacherAccount{
				ID:         tt.fields.ID,
				TeacherPay: tt.fields.TeacherPay,
				CreatedAt:  tt.fields.CreatedAt,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
