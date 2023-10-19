package models

import (
	"github.com/fatih/structs"
	"time"
)

type Teacher struct {
	ID        int       `bson:"_id,omitempty" json:"id" db:"id" structs:"id"`
	Name      string    `bson:"name" json:"name" db:"name" structs:"name"`
	SubjectID int       `bson:"subject_id" json:"subject_id" db:"subject_id" structs:"subject_id"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt" db:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt" db:"updatedAt" structs:"updatedAt"`
}

func (t *Teacher) Map() map[string]interface{} {
	return structs.Map(t)
}

func (t *Teacher) GetNames() []string {
	fieldNames := structs.Fields(t)
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
