package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTeacher_Map(test *testing.T) {
	type teacherFields struct {
		ID        int
		Name      string
		SubjectID int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	testCases := []struct {
		name   string
		fields teacherFields
		want   map[string]interface{}
	}{
		{
			name: "Teacher struct mapped successfully!",
			fields: teacherFields{
				ID:        21,
				Name:      "Haziq",
				SubjectID: 2,
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":         21,
				"name":       "Haziq",
				"subject_id": 2,
				"createdAt":  time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt":  time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Teacher struct mapped successfully without time input",
			fields: teacherFields{
				ID:        21,
				Name:      "Haziq",
				SubjectID: 2,
			},
			want: map[string]interface{}{
				"id":         21,
				"name":       "Haziq",
				"subject_id": 2,
				"createdAt":  time.Time{},
				"updatedAt":  time.Time{},
			},
		},
		{
			name: "Teacher struct mapped successfully without time input and name",
			fields: teacherFields{
				ID:        21,
				SubjectID: 2,
			},
			want: map[string]interface{}{
				"id":         21,
				"name":       "",
				"subject_id": 2,
				"createdAt":  time.Time{},
				"updatedAt":  time.Time{},
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Teacher{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				SubjectID: testCase.fields.SubjectID,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.Map(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Map() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestTeacher_GetNames(test *testing.T) {
	type teacherFields struct {
		ID        int
		Name      string
		SubjectID int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	testCases := []struct {
		name   string
		fields teacherFields
		want   []string
	}{
		{
			name: "Successfully returned the list of teacher fields",
			fields: teacherFields{
				ID:        21,
				Name:      "Haziq",
				SubjectID: 2,
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: []string{"id", "name", "subject_id", "createdAt", "updatedAt"},
		},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Teacher{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				SubjectID: testCase.fields.SubjectID,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.GetNames(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
