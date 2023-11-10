package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetSubjectsOfStudent(rt *runtime.Runtime) operations.GetSubjectsOfStudentHandler {
	return &getSubjectsOfStudent{rt: rt}
}

type getSubjectsOfStudent struct {
	rt *runtime.Runtime
}

// Handle the get Subjects Of Student request
func (d *getSubjectsOfStudent) Handle(params operations.GetSubjectsOfStudentParams) middleware.Responder {
	SubjectsOfStudent, err := d.rt.Service().GetSubjectsOfStudent(int(params.ID))
	if err != nil {

		return operations.NewGetSubjectsOfStudentNotFound()
	}

	mSubjectsOfStudent := make([]*models.Subject, len(SubjectsOfStudent))
	for i, Subject := range SubjectsOfStudent {
		mSubjectsOfStudent[i] = &models.Subject{
			ID:        int64(Subject.ID),
			Name:      Subject.Name,
			ClassID:   int64(Subject.ClassID),
			CreatedAt: strfmt.DateTime(Subject.CreatedAt),
			UpdatedAt: strfmt.DateTime(Subject.UpdatedAt),
		}
	}

	if len(mSubjectsOfStudent) == 0 {

		return operations.NewGetSubjectsOfStudentNotFound()
	}

	return operations.NewGetSubjectsOfStudentOK().WithPayload(mSubjectsOfStudent)

}
