package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetClasses(rt *runtime.Runtime) operations.GetClassesHandler {
	return &getClasses{rt: rt}
}

type getClasses struct {
	rt *runtime.Runtime
}

// Handle the get classes request
func (d *getClasses) Handle(_ operations.GetClassesParams) middleware.Responder {
	Classes, err := d.rt.Service().GetClasses()
	if err != nil {

		return operations.NewGetClassesNotFound()
	}

	mClasses := make([]*models.Class, len(Classes))
	for i, Class := range Classes {
		students := make([]int64, len(Class.Students))
		dBytes, _ := json.Marshal(Class.Students)
		_ = json.Unmarshal(dBytes, &students)

		mClasses[i] = &models.Class{
			ID:        int64(Class.ID),
			Name:      Class.Name,
			Students:  students,
			CreatedAt: strfmt.DateTime(Class.CreatedAt),
			UpdatedAt: strfmt.DateTime(Class.UpdatedAt),
		}
	}

	return operations.NewGetClassesOK().WithPayload(mClasses)

}
