package handlers

import (
	"encoding/json"
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
	"ApniUniversity/models"
)

type addStudent struct {
	rt *runtime.Runtime
}

func NewAddStudent(rt *runtime.Runtime) operations.AddStudentHandler {
	return &addStudent{rt: rt}
}

func (h *addStudent) Handle(params operations.AddStudentParams) middleware.Responder {
	subjects := make([]int, len(params.Student.Subjects))
	dBytes, _ := json.Marshal(params.Student.Subjects)
	_ = json.Unmarshal(dBytes, &subjects)

	id, err := h.rt.Service().AddStudent(&models.Student{
		ID:        int(params.Student.ID),
		Name:      params.Student.Name,
		GPA:       params.Student.GPA,
		Subjects:  subjects,
		CreatedAt: time.Time(params.Student.CreatedAt),
		UpdatedAt: time.Time(params.Student.UpdatedAt),
	})

	if err != nil {
		log().Errorf("Failed to Add New Student : %s", err)
		return operations.NewAddStudentConflict()
	}

	params.Student.ID = int64(id)

	return operations.NewAddStudentCreated().WithPayload(params.Student)
}
