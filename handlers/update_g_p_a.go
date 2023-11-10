package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type UpdateGPA struct {
	rt *runtime.Runtime
}

func NewUpdateGPA(rt *runtime.Runtime) operations.UpdateGPAHandler {
	return &UpdateGPA{rt: rt}
}

func (h *UpdateGPA) Handle(params operations.UpdateGPAParams) middleware.Responder {
	gpa := make(map[string]float64, 1)
	gpa["gpa"] = params.Gpa.GPA
	_, err := h.rt.Service().UpdateGPA(int(params.ID), gpa)

	if err != nil {
		log().Errorf("Failed to Update GPA of Student: %s", err)
		return operations.NewUpdateGPANotFound()
	}

	return operations.NewUpdateGPAOK().WithPayload(params.ID)
}
