package models

import "github.com/fatih/structs"

type StudentAccount struct {
	ID                    int     `json:"id" bson:"id" structs:"id" db:"id"`
	StudentFee            float64 `json:"studentFee" bson:"student_fee" structs:"student_fee" db:"student_fee"`
	DefaultStudent        bool    `json:"defaultStudent" bson:"default_student" structs:"default_student" db:"default_student"`
	StudentOnInstallment  bool    `json:"studentsOnInstallment" bson:"student_on_installment" structs:"student_on_installment" db:"student_on_installment"`
	TotalInstallments     int     `json:"totalInstallments" bson:"total_installments" structs:"total_installments" db:"total_installments"`
	RemainingInstallments int     `json:"remainingInstallments" bson:"remaining_installments" structs:"remaining_installments" db:"remaining_installments"`
}

func (t *StudentAccount) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the field names of Teacher model
func (t *StudentAccount) Names() []string {
	fields := structs.Fields(t)
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
