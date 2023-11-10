package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetClassByID(rt *runtime.Runtime) operations.GetClassByIDHandler {
	return &getClassByID{rt: rt}
}

type getClassByID struct {
	rt *runtime.Runtime
}

// Handle the get class request
func (d *getClassByID) Handle(params operations.GetClassByIDParams) middleware.Responder {
	Class, err := d.rt.Service().GetClassByID(int(params.ID))
	if err != nil {

		return operations.NewGetClassByIDNotFound()
	}

	students := make([]int64, len(Class.Students))
	dBytes, _ := json.Marshal(Class.Students)
	_ = json.Unmarshal(dBytes, &students)

	return operations.NewGetClassByIDOK().WithPayload(&models.Class{
		ID:        int64(Class.ID),
		Name:      Class.Name,
		Students:  students,
		CreatedAt: strfmt.DateTime(Class.CreatedAt),
		UpdatedAt: strfmt.DateTime(Class.UpdatedAt),
	})
}
