// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClassByIDHandlerFunc turns a function with the right signature into a get class by ID handler
type GetClassByIDHandlerFunc func(GetClassByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClassByIDHandlerFunc) Handle(params GetClassByIDParams) middleware.Responder {
	return fn(params)
}

// GetClassByIDHandler interface for that can handle valid get class by ID params
type GetClassByIDHandler interface {
	Handle(GetClassByIDParams) middleware.Responder
}

// NewGetClassByID creates a new http.Handler for the get class by ID operation
func NewGetClassByID(ctx *middleware.Context, handler GetClassByIDHandler) *GetClassByID {
	return &GetClassByID{Context: ctx, Handler: handler}
}

/*
	GetClassByID swagger:route GET /class/{ID} getClassById

GetClassByID get class by ID API
*/
type GetClassByID struct {
	Context *middleware.Context
	Handler GetClassByIDHandler
}

func (o *GetClassByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClassByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}