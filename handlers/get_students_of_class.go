package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetStudentsOfClass(rt *runtime.Runtime) operations.GetStudentsOfClassHandler {
	return &getStudentsOfClass{rt: rt}
}

type getStudentsOfClass struct {
	rt *runtime.Runtime
}

// Handle the get students of class request
func (d *getStudentsOfClass) Handle(params operations.GetStudentsOfClassParams) middleware.Responder {
	StudentsOfClass, err := d.rt.Service().GetStudentsOfClass(int(params.ID))
	if err != nil {

		return operations.NewGetStudentsOfClassNotFound()
	}

	mStudentsOfClass := make([]*models.Student, len(StudentsOfClass))
	for i, Student := range StudentsOfClass {
		subjects := make([]int64, len(Student.Subjects))
		dBytes, _ := json.Marshal(Student.Subjects)
		_ = json.Unmarshal(dBytes, &subjects)

		mStudentsOfClass[i] = &models.Student{
			ID:        int64(Student.ID),
			Name:      Student.Name,
			GPA:       Student.GPA,
			Subjects:  subjects,
			CreatedAt: strfmt.DateTime(Student.CreatedAt),
			UpdatedAt: strfmt.DateTime(Student.UpdatedAt),
		}
	}

	return operations.NewGetStudentsOfClassOK().WithPayload(mStudentsOfClass)

}
