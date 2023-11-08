package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetStudentsOfTeacher(rt *runtime.Runtime) operations.GetStudentsOfTeacherHandler {
	return &getStudentsOfTeacher{rt: rt}
}

type getStudentsOfTeacher struct {
	rt *runtime.Runtime
}

// Handle the get Students Of Teacher request
func (d *getStudentsOfTeacher) Handle(params operations.GetStudentsOfTeacherParams) middleware.Responder {
	StudentsOfTeacher, err := d.rt.Service().GetStudentsOfTeacher(int(params.ID))
	if err != nil {

		return operations.NewGetStudentsOfTeacherNotFound()
	}

	mStudentsOfTeacher := make([]*models.Student, len(StudentsOfTeacher))
	for i, Student := range StudentsOfTeacher {
		subjects := make([]int64, len(Student.Subjects))
		dBytes, _ := json.Marshal(Student.Subjects)
		_ = json.Unmarshal(dBytes, &subjects)

		mStudentsOfTeacher[i] = &models.Student{
			ID:        int64(Student.ID),
			Name:      Student.Name,
			GPA:       Student.GPA,
			Subjects:  subjects,
			CreatedAt: strfmt.DateTime(Student.CreatedAt),
			UpdatedAt: strfmt.DateTime(Student.UpdatedAt),
		}
	}

	if len(mStudentsOfTeacher) == 0 {
		return operations.NewGetStudentsOfTeacherNotFound()
	}

	return operations.NewGetStudentsOfTeacherOK().WithPayload(mStudentsOfTeacher)

}
