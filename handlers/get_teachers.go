package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetTeachers(rt *runtime.Runtime) operations.GetTeachersHandler {
	return &getTeachers{rt: rt}
}

type getTeachers struct {
	rt *runtime.Runtime
}

// Handle the get teachers request
func (d *getTeachers) Handle(_ operations.GetTeachersParams) middleware.Responder {
	Teachers, err := d.rt.Service().GetTeachers()
	if err != nil {

		return operations.NewGetTeachersNotFound()
	}

	mTeachers := make([]*models.Teacher, len(Teachers))
	for i, Teacher := range Teachers {
		mTeachers[i] = &models.Teacher{
			ID:        int64(Teacher.ID),
			Name:      Teacher.Name,
			SubjectID: int64(Teacher.SubjectID),
			CreatedAt: strfmt.DateTime(Teacher.CreatedAt),
			UpdatedAt: strfmt.DateTime(Teacher.UpdatedAt),
		}
	}

	return operations.NewGetTeachersOK().WithPayload(mTeachers)

}
