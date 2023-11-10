package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type deleteTeacher struct {
	rt *runtime.Runtime
}

func NewDeleteTeacher(rt *runtime.Runtime) operations.DeleteTeacherHandler {
	return &deleteTeacher{rt: rt}
}

// Handle the delete entry request
func (h *deleteTeacher) Handle(params operations.DeleteTeacherParams) middleware.Responder {
	if _, err := h.rt.Service().DeleteTeacher(int(params.ID)); err != nil {

		return operations.NewDeleteTeacherNotFound()
	}

	return operations.NewDeleteTeacherNoContent()
}
