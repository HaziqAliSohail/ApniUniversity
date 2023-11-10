package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type UpdateSubjectName struct {
	rt *runtime.Runtime
}

func NewUpdateSubjectName(rt *runtime.Runtime) operations.UpdateSubjectNameHandler {
	return &UpdateSubjectName{rt: rt}
}

func (h *UpdateSubjectName) Handle(params operations.UpdateSubjectNameParams) middleware.Responder {
	Subject := make(map[string]string, 1)
	Subject["name"] = params.Name.Name
	_, err := h.rt.Service().UpdateSubjectName(int(params.ID), Subject)

	if err != nil {
		log().Errorf("Failed to Update Subject Name: %s", err)
		return operations.NewUpdateSubjectNameNotFound()
	}

	return operations.NewUpdateSubjectNameOK().WithPayload(params.ID)
}
