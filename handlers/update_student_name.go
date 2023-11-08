package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type UpdateStudentName struct {
	rt *runtime.Runtime
}

func NewUpdateStudentName(rt *runtime.Runtime) operations.UpdateStudentNameHandler {
	return &UpdateStudentName{rt: rt}
}

func (h *UpdateStudentName) Handle(params operations.UpdateStudentNameParams) middleware.Responder {
	Student := make(map[string]string, 1)
	Student["name"] = params.Name.Name
	_, err := h.rt.Service().UpdateStudentName(int(params.ID), Student)

	if err != nil {
		log().Errorf("Failed to Update Student Name: %s", err)
		return operations.NewUpdateStudentNameNotFound()
	}

	return operations.NewUpdateStudentNameOK().WithPayload(params.ID)
}
