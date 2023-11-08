package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type assignSubjectToStudent struct {
	rt *runtime.Runtime
}

func NewAssignSubjectToStudent(rt *runtime.Runtime) operations.AssignSubjectToStudentHandler {
	return &assignSubjectToStudent{rt: rt}
}

func (h *assignSubjectToStudent) Handle(params operations.AssignSubjectToStudentParams) middleware.Responder {
	subject := make(map[string]interface{}, 2)
	subject["subjectID"] = int(params.Subject.SubjectID)
	subject["assign"] = params.Subject.Assign
	_, err := h.rt.Service().AssignSubjectToStudent(int(params.ID), subject)

	if err != nil {
		log().Errorf("Failed to Assign Subject to Student: %s", err)
		return operations.NewAssignSubjectToStudentNotFound()
	}

	return operations.NewAssignSubjectToStudentOK().WithPayload(params.ID)
}
