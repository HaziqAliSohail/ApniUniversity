package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type UpdateTeacherName struct {
	rt *runtime.Runtime
}

func NewUpdateTeacherName(rt *runtime.Runtime) operations.UpdateTeacherNameHandler {
	return &UpdateTeacherName{rt: rt}
}

func (h *UpdateTeacherName) Handle(params operations.UpdateTeacherNameParams) middleware.Responder {
	Teacher := make(map[string]string, 1)
	Teacher["name"] = params.Name.Name
	_, err := h.rt.Service().UpdateTeacherName(int(params.ID), Teacher)

	if err != nil {
		log().Errorf("Failed to Update Teacher Name: %s", err)
		return operations.NewUpdateTeacherNameNotFound()
	}

	return operations.NewUpdateTeacherNameOK().WithPayload(params.ID)
}
