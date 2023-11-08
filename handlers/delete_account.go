package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "ApniUniversity"
	"ApniUniversity/gen/restapi/operations"
)

type deleteAccount struct {
	rt *runtime.Runtime
}

func NewDeleteAccount(rt *runtime.Runtime) operations.DeleteAccountHandler {
	return &deleteAccount{rt: rt}
}

// Handle the delete entry request
func (h *deleteAccount) Handle(params operations.DeleteAccountParams) middleware.Responder {
	if _, err := h.rt.Service().DeleteAccount(int(params.ID)); err != nil {

		return operations.NewDeleteAccountNotFound()
	}

	return operations.NewDeleteAccountNoContent()
}
