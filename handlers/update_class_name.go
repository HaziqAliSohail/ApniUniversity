package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type UpdateClassName struct {
	rt *runtime.Runtime
}

func NewUpdateClassName(rt *runtime.Runtime) operations.UpdateClassNameHandler {
	return &UpdateClassName{rt: rt}
}

func (h *UpdateClassName) Handle(params operations.UpdateClassNameParams) middleware.Responder {
	class := make(map[string]string, 1)
	class["name"] = params.Name.Name
	_, err := h.rt.Service().UpdateClassName(int(params.ID), class)

	if err != nil {
		log().Errorf("Failed to Update Class Name: %s", err)
		return operations.NewUpdateClassNameNotFound()
	}

	return operations.NewUpdateClassNameOK().WithPayload(params.ID)
}
