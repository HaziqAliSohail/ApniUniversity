// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAccountByIDHandlerFunc turns a function with the right signature into a get account by ID handler
type GetAccountByIDHandlerFunc func(GetAccountByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountByIDHandlerFunc) Handle(params GetAccountByIDParams) middleware.Responder {
	return fn(params)
}

// GetAccountByIDHandler interface for that can handle valid get account by ID params
type GetAccountByIDHandler interface {
	Handle(GetAccountByIDParams) middleware.Responder
}

// NewGetAccountByID creates a new http.Handler for the get account by ID operation
func NewGetAccountByID(ctx *middleware.Context, handler GetAccountByIDHandler) *GetAccountByID {
	return &GetAccountByID{Context: ctx, Handler: handler}
}

/*
	GetAccountByID swagger:route GET /account/{ID} getAccountById

GetAccountByID get account by ID API
*/
type GetAccountByID struct {
	Context *middleware.Context
	Handler GetAccountByIDHandler
}

func (o *GetAccountByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
