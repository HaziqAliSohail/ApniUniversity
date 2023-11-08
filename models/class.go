package models

import (
	"time"

	"github.com/fatih/structs"
)

type Class struct {
	ID        int       `bson:"_id" json:"id" structs:"id" db:"id"`
	Name      string    `bson:"name" json:"name" structs:"name" db:"name"`
	Students  []int     `bson:"students" json:"students" structs:"students" db:"students"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt" db:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt" db:"updatedAt" structs:"updatedAt"`
}

func (c *Class) Map() map[string]interface{} {

	return structs.Map(c)
}

func (c *Class) GetNames() []string {
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
