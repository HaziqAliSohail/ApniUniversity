package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type deleteStudent struct {
	rt *runtime.Runtime
}

func NewDeleteStudent(rt *runtime.Runtime) operations.DeleteStudentHandler {
	return &deleteStudent{rt: rt}
}

// Handle the delete entry request
func (h *deleteStudent) Handle(params operations.DeleteStudentParams) middleware.Responder {
	if _, err := h.rt.Service().DeleteStudent(int(params.ID)); err != nil {

		return operations.NewDeleteStudentNotFound()
	}

	return operations.NewDeleteStudentNoContent()
}
