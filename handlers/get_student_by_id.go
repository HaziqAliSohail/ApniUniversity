package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetStudentByID(rt *runtime.Runtime) operations.GetStudentByIDHandler {
	return &getStudentByID{rt: rt}
}

type getStudentByID struct {
	rt *runtime.Runtime
}

// Handle the get student request
func (d *getStudentByID) Handle(params operations.GetStudentByIDParams) middleware.Responder {
	Student, err := d.rt.Service().GetStudentByID(int(params.ID))
	if err != nil {

		return operations.NewGetStudentByIDNotFound()
	}

	subjects := make([]int64, len(Student.Subjects))
	dBytes, _ := json.Marshal(Student.Subjects)
	_ = json.Unmarshal(dBytes, &subjects)

	return operations.NewGetStudentByIDOK().WithPayload(&models.Student{
		ID:        int64(Student.ID),
		Name:      Student.Name,
		GPA:       Student.GPA,
		Subjects:  subjects,
		CreatedAt: strfmt.DateTime(Student.CreatedAt),
		UpdatedAt: strfmt.DateTime(Student.UpdatedAt),
	})
}
