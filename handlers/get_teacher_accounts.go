package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetTeacherAccounts(rt *runtime.Runtime) operations.GetTeacherAccountsHandler {
	return &getTeacherAccounts{rt: rt}
}

type getTeacherAccounts struct {
	rt *runtime.Runtime
}

// Handle the get Teacher Accounts request
func (d *getTeacherAccounts) Handle(_ operations.GetTeacherAccountsParams) middleware.Responder {
	TeacherAccounts, err := d.rt.Service().GetTeacherAccounts()
	if err != nil {

		return operations.NewGetTeacherAccountsNotFound()
	}

	mTeacherAccounts := make([]*models.Account, len(TeacherAccounts))
	for i, account := range TeacherAccounts {
		mTeacherAccounts[i] = &models.Account{
			ID:          int64(account.ID),
			Name:        account.Name,
			AccountType: account.AccountType,
			AccountData: account.AccountData,
			CreatedAt:   strfmt.DateTime(account.CreatedAt),
			UpdatedAt:   strfmt.DateTime(account.UpdatedAt),
		}
	}

	return operations.NewGetTeacherAccountsOK().WithPayload(mTeacherAccounts)

}
