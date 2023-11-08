package handlers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/data"
	"ApniUniversity/gen/restapi/operations"
	"ApniUniversity/models"
)

type UpdateAccount struct {
	rt *runtime.Runtime
}

func NewUpdateAccount(rt *runtime.Runtime) operations.UpdateAccountHandler {
	return &UpdateAccount{rt: rt}
}

func (h *UpdateAccount) Handle(params operations.UpdateAccountParams) middleware.Responder {
	params.Account.AccountType = strings.ToLower(params.Account.AccountType)
	if params.Account.AccountType == data.TEACHER {
		var accountData models.TeacherAccount
		dBytes, _ := json.Marshal(params.Account.AccountData)
		_ = json.Unmarshal(dBytes, &accountData)
		params.Account.AccountData = accountData
	} else if params.Account.AccountType == data.STUDENT {
		var accountData models.StudentAccount
		dBytes, _ := json.Marshal(params.Account.AccountData)
		_ = json.Unmarshal(dBytes, &accountData)
		params.Account.AccountData = accountData
	}

	_, err := h.rt.Service().UpdateAccount(&models.Account{
		ID:          int(params.Account.ID),
		Name:        params.Account.Name,
		AccountType: params.Account.AccountType,
		AccountData: params.Account.AccountData,
		CreatedAt:   time.Time(params.Account.CreatedAt),
		UpdatedAt:   time.Time(params.Account.UpdatedAt),
	})

	if err != nil {
		log().Errorf("Failed to Update Account: %s", err)
		return operations.NewUpdateAccountNotFound()
	}

	return operations.NewUpdateAccountOK().WithPayload(params.Account)
}
