package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type deleteSubject struct {
	rt *runtime.Runtime
}

func NewDeleteSubject(rt *runtime.Runtime) operations.DeleteSubjectHandler {
	return &deleteSubject{rt: rt}
}

// Handle the delete entry request
func (h *deleteSubject) Handle(params operations.DeleteSubjectParams) middleware.Responder {
	if _, err := h.rt.Service().DeleteSubject(int(params.ID)); err != nil {

		return operations.NewDeleteSubjectNotFound()
	}

	return operations.NewDeleteSubjectNoContent()
}
