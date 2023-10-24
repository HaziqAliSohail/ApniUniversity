package models

import (
	"github.com/fatih/structs"
	"time"
)

type Student struct {
	ID        int       `bson:"_id" json:"id" structs:"id" db:"id"`
	Name      string    `bson:"name" json:"name" structs:"name" db:"name"`
	GPA       float64   `bson:"gpa" json:"GPA" structs:"gpa" db:"gpa"`
	Subjects  []int     `bson:"subjects" json:"subjects" structs:"subjects" db:"subjects"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt" structs:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt" structs:"updatedAt" db:"updatedAt"`
}

func (c *Student) Map() map[string]interface{} {

	return structs.Map(c)
}

func (c *Student) GetNames() []string {
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
