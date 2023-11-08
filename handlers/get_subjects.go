package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetSubjects(rt *runtime.Runtime) operations.GetSubjectsHandler {
	return &getSubjects{rt: rt}
}

type getSubjects struct {
	rt *runtime.Runtime
}

// Handle the get subjects request
func (d *getSubjects) Handle(_ operations.GetSubjectsParams) middleware.Responder {
	Subjects, err := d.rt.Service().GetSubjects()
	if err != nil {

		return operations.NewGetSubjectsNotFound()
	}

	mSubjects := make([]*models.Subject, len(Subjects))
	for i, Subject := range Subjects {
		mSubjects[i] = &models.Subject{
			ID:        int64(Subject.ID),
			Name:      Subject.Name,
			ClassID:   int64(Subject.ClassID),
			CreatedAt: strfmt.DateTime(Subject.CreatedAt),
			UpdatedAt: strfmt.DateTime(Subject.UpdatedAt),
		}
	}

	return operations.NewGetSubjectsOK().WithPayload(mSubjects)

}
