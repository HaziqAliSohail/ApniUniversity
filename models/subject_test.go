package models

import (
	"reflect"
	"testing"
	"time"
)

func TestSubject_Map(test *testing.T) {
	type subjectFields struct {
		ID        int
		Name      string
		ClassID   int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	testCases := []struct {
		name   string
		fields subjectFields
		want   map[string]interface{}
	}{
		{
			name: "Subject struct mapped successfully!",
			fields: subjectFields{
				ID:        2,
				Name:      "OOP",
				ClassID:   1,
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "OOP",
				"class_id":  1,
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Subject struct mapped successfully without time input",
			fields: subjectFields{
				ID:      21,
				Name:    "OOP",
				ClassID: 2,
			},
			want: map[string]interface{}{
				"id":        21,
				"name":      "OOP",
				"class_id":  2,
				"createdAt": time.Time{},
				"updatedAt": time.Time{},
			},
		},
		{
			name: "Subject struct mapped successfully without time input and subject name",
			fields: subjectFields{
				ID:      12,
				ClassID: 2,
			},
			want: map[string]interface{}{
				"id":        12,
				"name":      "",
				"class_id":  2,
				"createdAt": time.Time{},
				"updatedAt": time.Time{},
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Subject{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				ClassID:   testCase.fields.ClassID,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.Map(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Map() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestSubject_GetNames(test *testing.T) {
	type subjectFields struct {
		ID        int
		Name      string
		ClassID   int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	testCases := []struct {
		name   string
		fields subjectFields
		want   []string
	}{
		{
			name: "Successfully returned the list of subject fields",
			fields: subjectFields{
				ID:        10,
				Name:      "OOP",
				ClassID:   1,
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: []string{"id", "name", "class_id", "createdAt", "updatedAt"},
		},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Subject{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				ClassID:   testCase.fields.ClassID,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.GetNames(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
