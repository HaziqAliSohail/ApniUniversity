package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/HaziqAliSohail/ApniUniversity/data"
)

func TestAccount_Map(test *testing.T) {
	type accountFields struct {
		ID          int
		Name        string
		AccountType string
		AccountData interface{}
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	testCases := []struct {
		name   string
		fields accountFields
		want   map[string]interface{}
	}{
		{
			name: "Student Account struct mapped successfully!",
			fields: accountFields{
				ID:          3,
				Name:        "Haziq",
				AccountType: data.STUDENT,
				AccountData: &StudentAccount{StudentID: 2, TotalInstallments: 2, RemainingInstallments: 1, IsDefaulter: false},
				CreatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":          3,
				"name":        "Haziq",
				"accountType": data.STUDENT,
				"accountData": map[string]interface{}{"studentID": 2, "totalInstallments": 2, "remainingInstallments": 1, "isDefaulter": false},
				"createdAt":   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt":   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Teacher Account struct mapped successfully!",
			fields: accountFields{
				ID:          2,
				Name:        "Haziq",
				AccountType: data.TEACHER,
				AccountData: &TeacherAccount{TeacherID: 2, Salary: 40000},
				CreatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":          2,
				"name":        "Haziq",
				"accountType": data.TEACHER,
				"accountData": map[string]interface{}{"teacherID": 2, "salary": 40000},
				"createdAt":   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt":   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
		{
			name: "Student struct mapped successfully without account Data",
			fields: accountFields{
				ID:          3,
				Name:        "Ali",
				AccountType: data.STUDENT,
				AccountData: &StudentAccount{},
				CreatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":          3,
				"name":        "Ali",
				"accountType": data.STUDENT,
				"accountData": map[string]interface{}{"studentID": 0, "totalInstallments": 0, "remainingInstallments": 0, "isDefaulter": false},
				"createdAt":   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt":   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Account{
				ID:          testCase.fields.ID,
				Name:        testCase.fields.Name,
				AccountType: testCase.fields.AccountType,
				AccountData: testCase.fields.AccountData,
				CreatedAt:   testCase.fields.CreatedAt,
				UpdatedAt:   testCase.fields.UpdatedAt,
			}
			if got := testItem.Map(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Map() = %v,\n want %v", got, testCase.want)
			}
		})
	}
}

func TestAccount_GetNames(test *testing.T) {
	type accountFields struct {
		ID          int
		Name        string
		AccountType string
		AccountData interface{}
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	testCases := []struct {
		name   string
		fields accountFields
		want   []string
	}{
		{
			name: "Successfully returned the list of account fields",
			fields: accountFields{
				ID:          2,
				Name:        "Haziq",
				AccountType: data.TEACHER,
				AccountData: &TeacherAccount{TeacherID: 2, Salary: 40000},
				CreatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt:   time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: []string{"id", "name", "accountType", "accountData", "createdAt", "updatedAt"},
		},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Account{
				ID:          testCase.fields.ID,
				Name:        testCase.fields.Name,
				AccountType: testCase.fields.AccountType,
				AccountData: testCase.fields.AccountData,
				CreatedAt:   testCase.fields.CreatedAt,
				UpdatedAt:   testCase.fields.UpdatedAt,
			}
			if got := testItem.GetNames(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
