package models

import (
	"ApniUniversity/data"
	"github.com/fatih/structs"
	"time"
)

type Account struct {
	ID          int         `bson:"_id" json:"id" structs:"id" db:"id"`
	Name        string      `bson:"name" json:"name" structs:"name" db:"name"`
	AccountType string      `bson:"account_type" json:"accountType" structs:"accountType" db:"accountType"`
	AccountData interface{} `bson:"account_data" json:"accountData" structs:"accountData" db:"accountData"`
	CreatedAt   time.Time   `bson:"created_at" json:"createdAt" db:"createdAt" structs:"createdAt"`
	UpdatedAt   time.Time   `bson:"updated_at" json:"updatedAt" db:"updatedAt" structs:"updatedAt"`
}

type StudentAccount struct {
	StudentID             int  `bson:"student_id" json:"studentID" structs:"studentID" db:"studentID"`
	TotalInstallments     int  `bson:"total_installments" json:"totalInstallments" structs:"totalInstallments" db:"totalInstallments"`
	RemainingInstallments int  `bson:"remaining_installments" json:"remainingInstallments" structs:"remainingInstallments" db:"remainingInstallments"`
	IsDefaulter           bool `bson:"is_defaulter" json:"isDefaulter" structs:"isDefaulter" db:"isDefaulter"`
}

type TeacherAccount struct {
	TeacherID int `bson:"teacher_id" json:"teacherID" structs:"teacherID" db:"teacherID"`
	Salary    int `bson:"salary" json:"salary" structs:"salary" db:"salary"`
}

func (c *Account) GetAccountTypes() []string {
	return []string{data.TEACHER, data.STUDENT}
}

func (c *Account) Map() map[string]interface{} {
	return structs.Map(c)
}

func (c *Account) GetNames() []string {
	fieldNames := structs.Fields(c)
	fieldStructNames := make([]string, len(fieldNames))

	for i, fieldName := range fieldNames {
		fieldString := fieldName.Name()
		fieldStructName := fieldName.Tag(structs.DefaultTagName)
		if fieldStructName == "" {
			fieldStructName = fieldString
		}

		fieldStructNames[i] = fieldStructName
	}

	return fieldStructNames
}
