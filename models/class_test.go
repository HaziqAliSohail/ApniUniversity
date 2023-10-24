package models

import (
	"reflect"
	"testing"
	"time"
)

func TestClass_Map(test *testing.T) {
	type classFields struct {
		ID        int
		Name      string
		Students  []int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	testCases := []struct {
		name   string
		fields classFields
		want   map[string]interface{}
	}{
		{
			name: "Class struct mapped successfully!",
			fields: classFields{
				ID:        2,
				Name:      "BSE-2",
				Students:  []int{1, 5, 2},
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "BSE-2",
				"students":  []int{1, 5, 2},
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Class struct mapped successfully without students",
			fields: classFields{
				ID:        2,
				Name:      "BPH-2",
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "BPH-2",
				"students":  []int(nil),
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Class struct mapped successfully without name and subjects",
			fields: classFields{
				ID:        2,
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":        2,
				"name":      "",
				"students":  []int(nil),
				"createdAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Class{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				Students:  testCase.fields.Students,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.Map(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Map() = %v,\n want %v", got, testCase.want)

				return
			}
		})
	}
}

func TestClass_GetNames(test *testing.T) {
	type classFields struct {
		ID        int
		Name      string
		Students  []int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	testCases := []struct {
		name   string
		fields classFields
		want   []string
	}{
		{
			name: "Successfully returned the list of class fields",
			fields: classFields{
				ID:        10,
				Name:      "BAF-2",
				Students:  []int{1, 5, 2},
				CreatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: []string{"id", "name", "students", "createdAt", "updatedAt"},
		},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Class{
				ID:        testCase.fields.ID,
				Name:      testCase.fields.Name,
				Students:  testCase.fields.Students,
				CreatedAt: testCase.fields.CreatedAt,
				UpdatedAt: testCase.fields.UpdatedAt,
			}
			if got := testItem.GetNames(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
