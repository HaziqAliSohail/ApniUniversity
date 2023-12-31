// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTeachersHandlerFunc turns a function with the right signature into a get teachers handler
type GetTeachersHandlerFunc func(GetTeachersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeachersHandlerFunc) Handle(params GetTeachersParams) middleware.Responder {
	return fn(params)
}

// GetTeachersHandler interface for that can handle valid get teachers params
type GetTeachersHandler interface {
	Handle(GetTeachersParams) middleware.Responder
}

// NewGetTeachers creates a new http.Handler for the get teachers operation
func NewGetTeachers(ctx *middleware.Context, handler GetTeachersHandler) *GetTeachers {
	return &GetTeachers{Context: ctx, Handler: handler}
}

/*
	GetTeachers swagger:route GET /teacher getTeachers

Get all Teachers
*/
type GetTeachers struct {
	Context *middleware.Context
	Handler GetTeachersHandler
}

func (o *GetTeachers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTeachersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
