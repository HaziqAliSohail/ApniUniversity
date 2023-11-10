package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetAccounts(rt *runtime.Runtime) operations.GetAccountsHandler {
	return &getAccounts{rt: rt}
}

type getAccounts struct {
	rt *runtime.Runtime
}

// Handle the get accounts request
func (d *getAccounts) Handle(_ operations.GetAccountsParams) middleware.Responder {
	accounts, err := d.rt.Service().GetAccounts()
	if err != nil {

		return operations.NewGetAccountsNotFound()
	}
	log().Warning(len(accounts))
	mAccounts := make([]*models.Account, len(accounts))
	for i, account := range accounts {
		mAccounts[i] = &models.Account{
			ID:          int64(account.ID),
			Name:        account.Name,
			AccountType: account.AccountType,
			AccountData: account.AccountData,
			CreatedAt:   strfmt.DateTime(account.CreatedAt),
			UpdatedAt:   strfmt.DateTime(account.UpdatedAt),
		}
	}
	log().Warning(len(mAccounts))
	return operations.NewGetAccountsOK().WithPayload(mAccounts)

}
