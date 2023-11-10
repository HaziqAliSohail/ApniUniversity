package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetStudentsOfSubject(rt *runtime.Runtime) operations.GetStudentsOfSubjectHandler {
	return &getStudentsOfSubject{rt: rt}
}

type getStudentsOfSubject struct {
	rt *runtime.Runtime
}

// Handle the get Students Of Subject request
func (d *getStudentsOfSubject) Handle(params operations.GetStudentsOfSubjectParams) middleware.Responder {
	StudentsOfSubject, err := d.rt.Service().GetStudentsOfSubject(int(params.ID))
	if err != nil {

		return operations.NewGetStudentsOfSubjectNotFound()
	}

	mStudentsOfSubject := make([]*models.Student, len(StudentsOfSubject))
	for i, Student := range StudentsOfSubject {
		subjects := make([]int64, len(Student.Subjects))
		dBytes, _ := json.Marshal(Student.Subjects)
		_ = json.Unmarshal(dBytes, &subjects)

		mStudentsOfSubject[i] = &models.Student{
			ID:        int64(Student.ID),
			Name:      Student.Name,
			GPA:       Student.GPA,
			Subjects:  subjects,
			CreatedAt: strfmt.DateTime(Student.CreatedAt),
			UpdatedAt: strfmt.DateTime(Student.UpdatedAt),
		}
	}

	if len(mStudentsOfSubject) == 0 {
		return operations.NewGetStudentsOfSubjectNotFound()
	}

	return operations.NewGetStudentsOfSubjectOK().WithPayload(mStudentsOfSubject)

}
