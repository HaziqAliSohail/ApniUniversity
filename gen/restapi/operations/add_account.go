// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddAccountHandlerFunc turns a function with the right signature into a add account handler
type AddAccountHandlerFunc func(AddAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddAccountHandlerFunc) Handle(params AddAccountParams) middleware.Responder {
	return fn(params)
}

// AddAccountHandler interface for that can handle valid add account params
type AddAccountHandler interface {
	Handle(AddAccountParams) middleware.Responder
}

// NewAddAccount creates a new http.Handler for the add account operation
func NewAddAccount(ctx *middleware.Context, handler AddAccountHandler) *AddAccount {
	return &AddAccount{Context: ctx, Handler: handler}
}

/*
	AddAccount swagger:route POST /account addAccount

AddAccount add account API
*/
type AddAccount struct {
	Context *middleware.Context
	Handler AddAccountHandler
}

func (o *AddAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddAccountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
