package models

import (
	"reflect"
	"testing"
)

func TestStudent_Map(t *testing.T) {

	type fields struct {
		ID      string
		Name    string
		Age     string
		GPA     float64
		ClassID string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - student struct to map",
			fields: fields{
				ID:      "12345",
				Name:    "asif",
				Age:     "22",
				GPA:     3.23,
				ClassID: "1",
			},
			want: map[string]interface{}{
				"id":      "12345",
				"name":    "asif",
				"age":     "22",
				"gpa":     3.23,
				"classID": "1",
			},
		},
		{
			name: " success - student struct to map with fewer fields",
			fields: fields{
				Name: "asif",
				Age:  "22",
			},
			want: map[string]interface{}{
				"id":      "",
				"name":    "asif",
				"age":     "22",
				"gpa":     0.0,
				"classID": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				Age:     tt.fields.Age,
				GPA:     tt.fields.GPA,
				ClassID: tt.fields.ClassID,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestStudent_Names(t *testing.T) {
	type fields struct {
		ID      string
		Name    string
		Age     string
		GPA     float64
		ClassID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:      "12345",
				Name:    "asif",
				Age:     "22",
				GPA:     3.23,
				ClassID: "1",
			},
			want: []string{"id", "name", "age", "gpa", "classID"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				Age:     tt.fields.Age,
				GPA:     tt.fields.GPA,
				ClassID: tt.fields.ClassID,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
