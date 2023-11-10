package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetClassesOfStudent(rt *runtime.Runtime) operations.GetClassesOfStudentHandler {
	return &getClassesOfStudent{rt: rt}
}

type getClassesOfStudent struct {
	rt *runtime.Runtime
}

// Handle the get Classes request
func (d *getClassesOfStudent) Handle(params operations.GetClassesOfStudentParams) middleware.Responder {
	Classes, err := d.rt.Service().GetClassesOfStudent(int(params.ID))
	if err != nil {

		return operations.NewGetClassesOfStudentNotFound()
	}

	sClasses := make([]*models.Class, len(Classes))
	for i, Class := range Classes {
		students := make([]int64, len(Class.Students))
		dBytes, _ := json.Marshal(Class.Students)
		_ = json.Unmarshal(dBytes, &students)

		sClasses[i] = &models.Class{
			ID:        int64(Class.ID),
			Name:      Class.Name,
			Students:  students,
			CreatedAt: strfmt.DateTime(Class.CreatedAt),
			UpdatedAt: strfmt.DateTime(Class.UpdatedAt),
		}
	}

	if len(sClasses) == 0 {

		return operations.NewGetClassesOfStudentNotFound()
	}

	return operations.NewGetClassesOfStudentOK().WithPayload(sClasses)
}
