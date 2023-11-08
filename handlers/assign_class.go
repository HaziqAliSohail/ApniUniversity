package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type assignClass struct {
	rt *runtime.Runtime
}

func NewAssignClass(rt *runtime.Runtime) operations.AssignClassHandler {
	return &assignClass{rt: rt}
}

func (h *assignClass) Handle(params operations.AssignClassParams) middleware.Responder {
	class := make(map[string]int, 1)
	class["classID"] = int(params.Class.ClassID)
	_, err := h.rt.Service().AssignClass(int(params.ID), class)

	if err != nil {
		log().Errorf("Failed to Assign Class to Subject: %s", err)
		return operations.NewAssignClassNotFound()
	}

	return operations.NewAssignClassOK().WithPayload(params.ID)
}
