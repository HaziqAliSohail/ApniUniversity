// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAccountsHandlerFunc turns a function with the right signature into a get accounts handler
type GetAccountsHandlerFunc func(GetAccountsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountsHandlerFunc) Handle(params GetAccountsParams) middleware.Responder {
	return fn(params)
}

// GetAccountsHandler interface for that can handle valid get accounts params
type GetAccountsHandler interface {
	Handle(GetAccountsParams) middleware.Responder
}

// NewGetAccounts creates a new http.Handler for the get accounts operation
func NewGetAccounts(ctx *middleware.Context, handler GetAccountsHandler) *GetAccounts {
	return &GetAccounts{Context: ctx, Handler: handler}
}

/*
	GetAccounts swagger:route GET /account getAccounts

Get all Accounts
*/
type GetAccounts struct {
	Context *middleware.Context
	Handler GetAccountsHandler
}

func (o *GetAccounts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
