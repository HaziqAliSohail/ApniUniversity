package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
	"ApniUniversity/models"
)

type addTeacher struct {
	rt *runtime.Runtime
}

func NewAddTeacher(rt *runtime.Runtime) operations.AddTeacherHandler {
	return &addTeacher{rt: rt}
}

func (h *addTeacher) Handle(params operations.AddTeacherParams) middleware.Responder {
	id, err := h.rt.Service().AddTeacher(&models.Teacher{
		ID:        int(params.Teacher.ID),
		Name:      params.Teacher.Name,
		SubjectID: int(params.Teacher.SubjectID),
		CreatedAt: time.Time(params.Teacher.CreatedAt),
		UpdatedAt: time.Time(params.Teacher.UpdatedAt),
	})

	if err != nil {
		log().Errorf("Failed to Add New Teacher : %s", err)
		return operations.NewAddTeacherConflict()
	}

	params.Teacher.ID = int64(id)

	return operations.NewAddTeacherCreated().WithPayload(params.Teacher)
}
