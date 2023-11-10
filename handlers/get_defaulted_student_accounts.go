package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetDefaultedStudentAccounts(rt *runtime.Runtime) operations.GetDefaultedStudentAccountsHandler {
	return &getDefaultedStudentAccounts{rt: rt}
}

type getDefaultedStudentAccounts struct {
	rt *runtime.Runtime
}

// Handle the get Defaulted Student Accounts request
func (d *getDefaultedStudentAccounts) Handle(_ operations.GetDefaultedStudentAccountsParams) middleware.Responder {
	DefaultedStudentAccounts, err := d.rt.Service().GetDefaultedStudentAccounts()
	if err != nil {

		return operations.NewGetDefaultedStudentAccountsNotFound()
	}

	mDefaultedStudentAccounts := make([]*models.Account, len(DefaultedStudentAccounts))
	for i, account := range DefaultedStudentAccounts {
		mDefaultedStudentAccounts[i] = &models.Account{
			ID:          int64(account.ID),
			Name:        account.Name,
			AccountType: account.AccountType,
			AccountData: account.AccountData,
			CreatedAt:   strfmt.DateTime(account.CreatedAt),
			UpdatedAt:   strfmt.DateTime(account.UpdatedAt),
		}
	}

	return operations.NewGetDefaultedStudentAccountsOK().WithPayload(mDefaultedStudentAccounts)

}
