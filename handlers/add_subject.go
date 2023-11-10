package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
	"github.com/HaziqAliSohail/ApniUniversity/models"
)

type addSubject struct {
	rt *runtime.Runtime
}

func NewAddSubject(rt *runtime.Runtime) operations.AddSubjectHandler {
	return &addSubject{rt: rt}
}

func (h *addSubject) Handle(params operations.AddSubjectParams) middleware.Responder {

	id, err := h.rt.Service().AddSubject(&models.Subject{
		ID:        int(params.Subject.ID),
		Name:      params.Subject.Name,
		ClassID:   int(params.Subject.ClassID),
		CreatedAt: time.Time(params.Subject.CreatedAt),
		UpdatedAt: time.Time(params.Subject.UpdatedAt),
	})

	if err != nil {
		log().Errorf("Failed to Add Subject : %s", err)
		return operations.NewAddSubjectConflict()
	}

	params.Subject.ID = int64(id)

	return operations.NewAddSubjectCreated().WithPayload(params.Subject)
}
