package models

import (
	"github.com/fatih/structs"
	"time"
)

type Subject struct {
	ID        int       `json:"id" bson:"_id" structs:"id" db:"id"`
	Name      string    `json:"name" bson:"name" structs:"name" db:"name"`
	ClassID   int       `json:"class_id" bson:"class_id" db:"class_id" structs:"class_id"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt" db:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt" db:"updatedAt" structs:"updatedAt"`
}

func (s *Subject) Map() map[string]interface{} {
	return structs.Map(s)
}

func (s *Subject) GetNames() []string {
	fieldNames := structs.Fields(s)
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
