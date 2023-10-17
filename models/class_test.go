package models

import (
	"reflect"
	"testing"
)

func TestClass_Map(t *testing.T) {

	type fields struct {
		ID   string
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - class struct to map",
			fields: fields{
				ID:   "12345",
				Name: "ali",
			},
			want: map[string]interface{}{
				"id":   "12345",
				"name": "ali",
			},
		},
		{
			name: " success - class struct to map with fewer fields",
			fields: fields{
				Name: "ali",
			},
			want: map[string]interface{}{
				"id":   "",
				"name": "ali",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Class{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestClass_Names(t *testing.T) {
	type fields struct {
		ID   string
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of class struct fields",
			fields: fields{
				ID:   "12345",
				Name: "ali",
			},
			want: []string{"id", "name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Class{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
