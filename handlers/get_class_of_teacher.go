package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetClassOfTeacher(rt *runtime.Runtime) operations.GetClassOfTeacherHandler {
	return &getClassOfTeacher{rt: rt}
}

type getClassOfTeacher struct {
	rt *runtime.Runtime
}

// Handle the get class request
func (d *getClassOfTeacher) Handle(params operations.GetClassOfTeacherParams) middleware.Responder {
	Class, err := d.rt.Service().GetClassOfTeacher(int(params.ID))
	if err != nil {

		return operations.NewGetClassOfTeacherNotFound()
	}

	students := make([]int64, len(Class.Students))
	dBytes, _ := json.Marshal(Class.Students)
	_ = json.Unmarshal(dBytes, &students)

	return operations.NewGetClassOfTeacherOK().WithPayload(&models.Class{
		ID:        int64(Class.ID),
		Name:      Class.Name,
		Students:  students,
		CreatedAt: strfmt.DateTime(Class.CreatedAt),
		UpdatedAt: strfmt.DateTime(Class.UpdatedAt),
	})
}
