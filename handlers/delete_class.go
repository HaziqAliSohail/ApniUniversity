package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type deleteClass struct {
	rt *runtime.Runtime
}

func NewDeleteClass(rt *runtime.Runtime) operations.DeleteClassHandler {
	return &deleteClass{rt: rt}
}

// Handle the delete entry request
func (h *deleteClass) Handle(params operations.DeleteClassParams) middleware.Responder {
	if _, err := h.rt.Service().DeleteClass(int(params.ID)); err != nil {

		return operations.NewDeleteClassNotFound()
	}

	return operations.NewDeleteClassNoContent()
}
