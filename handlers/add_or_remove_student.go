package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type addOrRemoveStudent struct {
	rt *runtime.Runtime
}

func NewAddOrRemoveStudent(rt *runtime.Runtime) operations.AddOrRemoveStudentHandler {
	return &addOrRemoveStudent{rt: rt}
}

func (h *addOrRemoveStudent) Handle(params operations.AddOrRemoveStudentParams) middleware.Responder {
	student := make(map[string]interface{}, 2)
	student["studentID"] = params.Student.StudentID
	student["add"] = params.Student.Add
	_, err := h.rt.Service().AddOrRemoveStudent(int(params.ID), student)

	if err != nil {
		log().Errorf("Failed to Add or Remove Student from Class : %s", err)
		return operations.NewAddOrRemoveStudentNotFound()
	}

	return operations.NewAddOrRemoveStudentOK().WithPayload(params.ID)
}
