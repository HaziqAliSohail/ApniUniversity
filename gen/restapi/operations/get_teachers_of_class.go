// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTeachersOfClassHandlerFunc turns a function with the right signature into a get teachers of class handler
type GetTeachersOfClassHandlerFunc func(GetTeachersOfClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeachersOfClassHandlerFunc) Handle(params GetTeachersOfClassParams) middleware.Responder {
	return fn(params)
}

// GetTeachersOfClassHandler interface for that can handle valid get teachers of class params
type GetTeachersOfClassHandler interface {
	Handle(GetTeachersOfClassParams) middleware.Responder
}

// NewGetTeachersOfClass creates a new http.Handler for the get teachers of class operation
func NewGetTeachersOfClass(ctx *middleware.Context, handler GetTeachersOfClassHandler) *GetTeachersOfClass {
	return &GetTeachersOfClass{Context: ctx, Handler: handler}
}

/*
	GetTeachersOfClass swagger:route GET /class/{ID}/teachers getTeachersOfClass

GetTeachersOfClass get teachers of class API
*/
type GetTeachersOfClass struct {
	Context *middleware.Context
	Handler GetTeachersOfClassHandler
}

func (o *GetTeachersOfClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTeachersOfClassParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
