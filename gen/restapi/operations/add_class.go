// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddClassHandlerFunc turns a function with the right signature into a add class handler
type AddClassHandlerFunc func(AddClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddClassHandlerFunc) Handle(params AddClassParams) middleware.Responder {
	return fn(params)
}

// AddClassHandler interface for that can handle valid add class params
type AddClassHandler interface {
	Handle(AddClassParams) middleware.Responder
}

// NewAddClass creates a new http.Handler for the add class operation
func NewAddClass(ctx *middleware.Context, handler AddClassHandler) *AddClass {
	return &AddClass{Context: ctx, Handler: handler}
}

/*
	AddClass swagger:route POST /class addClass

AddClass add class API
*/
type AddClass struct {
	Context *middleware.Context
	Handler AddClassHandler
}

func (o *AddClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddClassParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}