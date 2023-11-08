package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetSubjectsOfClass(rt *runtime.Runtime) operations.GetSubjectsOfClassHandler {
	return &getSubjectsOfClass{rt: rt}
}

type getSubjectsOfClass struct {
	rt *runtime.Runtime
}

// Handle the get Subjects Of Class request
func (d *getSubjectsOfClass) Handle(params operations.GetSubjectsOfClassParams) middleware.Responder {
	SubjectsOfClass, err := d.rt.Service().GetSubjectsOfClass(int(params.ID))
	if err != nil {

		return operations.NewGetSubjectsOfClassNotFound()
	}

	mSubjectsOfClass := make([]*models.Subject, len(SubjectsOfClass))
	for i, Subject := range SubjectsOfClass {
		mSubjectsOfClass[i] = &models.Subject{
			ID:        int64(Subject.ID),
			Name:      Subject.Name,
			ClassID:   int64(Subject.ClassID),
			CreatedAt: strfmt.DateTime(Subject.CreatedAt),
			UpdatedAt: strfmt.DateTime(Subject.UpdatedAt),
		}
	}

	if len(mSubjectsOfClass) == 0 {

		return operations.NewGetSubjectsOfClassNotFound()
	}

	return operations.NewGetSubjectsOfClassOK().WithPayload(mSubjectsOfClass)

}
