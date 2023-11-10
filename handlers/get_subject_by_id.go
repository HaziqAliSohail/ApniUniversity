package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetSubjectByID(rt *runtime.Runtime) operations.GetSubjectByIDHandler {
	return &getSubjectByID{rt: rt}
}

type getSubjectByID struct {
	rt *runtime.Runtime
}

// Handle the get subject request
func (d *getSubjectByID) Handle(params operations.GetSubjectByIDParams) middleware.Responder {
	Subject, err := d.rt.Service().GetSubjectByID(int(params.ID))
	if err != nil {

		return operations.NewGetSubjectByIDNotFound()
	}

	return operations.NewGetSubjectByIDOK().WithPayload(&models.Subject{
		ID:        int64(Subject.ID),
		Name:      Subject.Name,
		ClassID:   int64(Subject.ClassID),
		CreatedAt: strfmt.DateTime(Subject.CreatedAt),
		UpdatedAt: strfmt.DateTime(Subject.UpdatedAt),
	})
}
