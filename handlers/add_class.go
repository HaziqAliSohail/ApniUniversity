package handlers

import (
	"encoding/json"
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
	"ApniUniversity/models"
)

type addClass struct {
	rt *runtime.Runtime
}

func NewAddClass(rt *runtime.Runtime) operations.AddClassHandler {
	return &addClass{rt: rt}
}

func (h *addClass) Handle(params operations.AddClassParams) middleware.Responder {
	students := make([]int, len(params.Class.Students))
	dBytes, _ := json.Marshal(params.Class.Students)
	_ = json.Unmarshal(dBytes, &students)

	id, err := h.rt.Service().AddClass(&models.Class{
		ID:        int(params.Class.ID),
		Name:      params.Class.Name,
		Students:  students,
		CreatedAt: time.Time(params.Class.CreatedAt),
		UpdatedAt: time.Time(params.Class.UpdatedAt),
	})

	if err != nil {
		log().Errorf("Failed to Add Class : %s", err)
		return operations.NewAddClassConflict()
	}

	params.Class.ID = int64(id)

	return operations.NewAddClassCreated().WithPayload(params.Class)
}
