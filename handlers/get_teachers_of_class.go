package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetTeachersOfClass(rt *runtime.Runtime) operations.GetTeachersOfClassHandler {
	return &getTeachersOfClass{rt: rt}
}

type getTeachersOfClass struct {
	rt *runtime.Runtime
}

// Handle the get Teachers Of Class request
func (d *getTeachersOfClass) Handle(params operations.GetTeachersOfClassParams) middleware.Responder {
	TeachersOfClass, err := d.rt.Service().GetTeachersOfClass(int(params.ID))
	if err != nil {

		return operations.NewGetTeachersOfClassNotFound()
	}

	mTeachersOfClass := make([]*models.Teacher, len(TeachersOfClass))
	for i, Teacher := range TeachersOfClass {
		mTeachersOfClass[i] = &models.Teacher{
			ID:        int64(Teacher.ID),
			Name:      Teacher.Name,
			SubjectID: int64(Teacher.SubjectID),
			CreatedAt: strfmt.DateTime(Teacher.CreatedAt),
			UpdatedAt: strfmt.DateTime(Teacher.UpdatedAt),
		}
	}

	if len(mTeachersOfClass) == 0 {
		return operations.NewGetTeachersOfClassNotFound()
	}

	return operations.NewGetTeachersOfClassOK().WithPayload(mTeachersOfClass)

}
