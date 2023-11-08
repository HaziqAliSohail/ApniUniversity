package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetTeacherOfSubject(rt *runtime.Runtime) operations.GetTeacherOfSubjectHandler {
	return &getTeacherOfSubject{rt: rt}
}

type getTeacherOfSubject struct {
	rt *runtime.Runtime
}

// Handle the get teacher of Subject request
func (d *getTeacherOfSubject) Handle(params operations.GetTeacherOfSubjectParams) middleware.Responder {
	Teacher, err := d.rt.Service().GetTeacherOfSubject(int(params.ID))
	if err != nil {

		return operations.NewGetTeacherOfSubjectNotFound()
	}

	return operations.NewGetTeacherOfSubjectOK().WithPayload(&models.Teacher{
		ID:        int64(Teacher.ID),
		Name:      Teacher.Name,
		SubjectID: int64(Teacher.SubjectID),
		CreatedAt: strfmt.DateTime(Teacher.CreatedAt),
		UpdatedAt: strfmt.DateTime(Teacher.UpdatedAt),
	})
}
