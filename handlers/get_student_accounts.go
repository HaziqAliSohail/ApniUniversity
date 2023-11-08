package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "ApniUniversity"
	"ApniUniversity/gen/models"
	"ApniUniversity/gen/restapi/operations"
)

func NewGetStudentAccounts(rt *runtime.Runtime) operations.GetStudentAccountsHandler {
	return &getStudentAccounts{rt: rt}
}

type getStudentAccounts struct {
	rt *runtime.Runtime
}

// Handle the get Student Accounts request
func (d *getStudentAccounts) Handle(_ operations.GetStudentAccountsParams) middleware.Responder {
	StudentAccounts, err := d.rt.Service().GetStudentAccounts()
	if err != nil {

		return operations.NewGetStudentAccountsNotFound()
	}

	mStudentAccounts := make([]*models.Account, len(StudentAccounts))
	for i, account := range StudentAccounts {
		mStudentAccounts[i] = &models.Account{
			ID:          int64(account.ID),
			Name:        account.Name,
			AccountType: account.AccountType,
			AccountData: account.AccountData,
			CreatedAt:   strfmt.DateTime(account.CreatedAt),
			UpdatedAt:   strfmt.DateTime(account.UpdatedAt),
		}
	}

	return operations.NewGetStudentAccountsOK().WithPayload(mStudentAccounts)

}
