package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type assignSubjectToTeacher struct {
	rt *runtime.Runtime
}

func NewAssignSubjectToTeacher(rt *runtime.Runtime) operations.AssignSubjectToTeacherHandler {
	return &assignSubjectToTeacher{rt: rt}
}

func (h *assignSubjectToTeacher) Handle(params operations.AssignSubjectToTeacherParams) middleware.Responder {
	subject := make(map[string]int, 1)
	subject["subjectID"] = int(params.Subject.SubjectID)

	_, err := h.rt.Service().AssignSubjectToTeacher(int(params.ID), subject)

	if err != nil {
		log().Errorf("Failed to Assign Subject to Teacher: %s", err)
		return operations.NewAssignSubjectToTeacherNotFound()
	}

	return operations.NewAssignSubjectToTeacherOK().WithPayload(params.ID)
}
