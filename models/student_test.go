package models

import (
	"reflect"
	"testing"
	"time"
)

func TestStudent_Map(test *testing.T) {
	type studentFields struct {
		ID        int
		Name      string
		GPA       float64
		Subjects  []int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	testCases := []struct {
		name   string
		fields studentFields
		want   map[string]interface{}
	}{
		{
			name: "Student struct mapped successfully!",
			fields: studentFields{
				ID:        2,
				Name:      "Ali",
				GPA:       3.5,
				Subjects:  []int{1, 5, 2},
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "Ali",
				"gpa":       3.5,
				"subjects":  []int{1, 5, 2},
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Student struct mapped successfully without gpa and subjects",
			fields: studentFields{
				ID:        2,
				Name:      "Ali",
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "Ali",
				"gpa":       0.0,
				"subjects":  []int(nil),
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Student struct mapped successfully without gpa,name and subjects",
			fields: studentFields{
				ID:        2,
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "",
				"gpa":       0.0,
				"subjects":  []int(nil),
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Student{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				GPA:       testCase.fields.GPA,
				Subjects:  testCase.fields.Subjects,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.Map(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Map() = %v,\n want %v", got, testCase.want)
			}
		})
	}
}

func TestStudent_GetNames(test *testing.T) {
	type studentFields struct {
		ID        int
		Name      string
		GPA       float64
		Subjects  []int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	testCases := []struct {
		name   string
		fields studentFields
		want   []string
	}{
		{
			name: "Successfully returned the list of student fields",
			fields: studentFields{
				ID:        10,
				Name:      "Asim",
				GPA:       3.5,
				Subjects:  []int{1, 5, 2},
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: []string{"id", "name", "gpa", "subjects", "createdAt", "updatedAt"},
		},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Student{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				GPA:       testCase.fields.GPA,
				Subjects:  testCase.fields.Subjects,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.GetNames(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
