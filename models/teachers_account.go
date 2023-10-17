package models

import (
	"github.com/fatih/structs"
	"time"
)

type TeacherAccount struct {
	ID         int       `json:"id" bson:"_id" structs:"id" db:"id"`
	TeacherPay float64   `json:"teacher_pay" bson:"teacher_pay" structs:"teacher_pay" db:"teacher_pay"`
	CreatedAt  time.Time `json:"createdAt" bson:"created_at" structs:"created_at" db:"created_at"`
}

func (a *TeacherAccount) Map() map[string]interface{} {
	return structs.Map(a)
}

// Names returns the field names of Teacher model
func (a *TeacherAccount) Names() []string {
	fields := structs.Fields(a)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
