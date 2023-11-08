package handlers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/data"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
	"github.com/HaziqAliSohail/ApniUniversity/models"
)

type addAccount struct {
	rt *runtime.Runtime
}

func NewAddAccount(rt *runtime.Runtime) operations.AddAccountHandler {
	return &addAccount{rt: rt}
}

func (h *addAccount) Handle(params operations.AddAccountParams) middleware.Responder {
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

	id, err := h.rt.Service().AddAccount(&models.Account{
		ID:          int(params.Account.ID),
		Name:        params.Account.Name,
		AccountType: params.Account.AccountType,
		AccountData: params.Account.AccountData,
		CreatedAt:   time.Time(params.Account.CreatedAt),
		UpdatedAt:   time.Time(params.Account.UpdatedAt),
	})

	if err != nil {
		log().Errorf("Failed to Add Account : %s", err)
		return operations.NewAddAccountConflict()
	}

	params.Account.ID = int64(id)

	return operations.NewAddAccountCreated().WithPayload(params.Account)
}
