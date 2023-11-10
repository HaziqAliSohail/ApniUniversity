package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

func NewGetAccountByID(rt *runtime.Runtime) operations.GetAccountByIDHandler {
	return &getAccountByID{rt: rt}
}

type getAccountByID struct {
	rt *runtime.Runtime
}

// Handle the get account request
func (d *getAccountByID) Handle(params operations.GetAccountByIDParams) middleware.Responder {
	account, err := d.rt.Service().GetAccountByID(int(params.ID))
	if err != nil {

		return operations.NewGetAccountByIDNotFound()
	}

	return operations.NewGetAccountByIDOK().WithPayload(&models.Account{
		ID:          int64(account.ID),
		Name:        account.Name,
		AccountType: account.AccountType,
		AccountData: account.AccountData,
		CreatedAt:   strfmt.DateTime(account.CreatedAt),
		UpdatedAt:   strfmt.DateTime(account.UpdatedAt),
	})
}
