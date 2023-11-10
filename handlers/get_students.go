package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetStudents(rt *runtime.Runtime) operations.GetStudentsHandler {
	return &getStudents{rt: rt}
}

type getStudents struct {
	rt *runtime.Runtime
}

// Handle the get students request
func (d *getStudents) Handle(_ operations.GetStudentsParams) middleware.Responder {
	Students, err := d.rt.Service().GetStudents()
	if err != nil {

		return operations.NewGetStudentsNotFound()
	}

	mStudents := make([]*models.Student, len(Students))
	for i, Student := range Students {
		subjects := make([]int64, len(Student.Subjects))
		dBytes, _ := json.Marshal(Student.Subjects)
		_ = json.Unmarshal(dBytes, &subjects)

		mStudents[i] = &models.Student{
			ID:        int64(Student.ID),
			Name:      Student.Name,
			GPA:       Student.GPA,
			Subjects:  subjects,
			CreatedAt: strfmt.DateTime(Student.CreatedAt),
			UpdatedAt: strfmt.DateTime(Student.UpdatedAt),
		}
	}

	return operations.NewGetStudentsOK().WithPayload(mStudents)

}
