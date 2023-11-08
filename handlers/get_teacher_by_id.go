package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetTeacherByID(rt *runtime.Runtime) operations.GetTeacherByIDHandler {
	return &getTeacherByID{rt: rt}
}

type getTeacherByID struct {
	rt *runtime.Runtime
}

// Handle the get teacher request
func (d *getTeacherByID) Handle(params operations.GetTeacherByIDParams) middleware.Responder {
	Teacher, err := d.rt.Service().GetTeacherByID(int(params.ID))
	if err != nil {

		return operations.NewGetTeacherByIDNotFound()
	}

	return operations.NewGetTeacherByIDOK().WithPayload(&models.Teacher{
		ID:        int64(Teacher.ID),
		Name:      Teacher.Name,
		SubjectID: int64(Teacher.SubjectID),
		CreatedAt: strfmt.DateTime(Teacher.CreatedAt),
		UpdatedAt: strfmt.DateTime(Teacher.UpdatedAt),
	})
}
