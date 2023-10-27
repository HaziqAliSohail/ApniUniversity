// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTeachersOfSubjectHandlerFunc turns a function with the right signature into a get teachers of subject handler
type GetTeachersOfSubjectHandlerFunc func(GetTeachersOfSubjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeachersOfSubjectHandlerFunc) Handle(params GetTeachersOfSubjectParams) middleware.Responder {
	return fn(params)
}

// GetTeachersOfSubjectHandler interface for that can handle valid get teachers of subject params
type GetTeachersOfSubjectHandler interface {
	Handle(GetTeachersOfSubjectParams) middleware.Responder
}

// NewGetTeachersOfSubject creates a new http.Handler for the get teachers of subject operation
func NewGetTeachersOfSubject(ctx *middleware.Context, handler GetTeachersOfSubjectHandler) *GetTeachersOfSubject {
	return &GetTeachersOfSubject{Context: ctx, Handler: handler}
}

/*
	GetTeachersOfSubject swagger:route GET /subject/{ID}/teacher getTeachersOfSubject

GetTeachersOfSubject get teachers of subject API
*/
type GetTeachersOfSubject struct {
	Context *middleware.Context
	Handler GetTeachersOfSubjectHandler
}

func (o *GetTeachersOfSubject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTeachersOfSubjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}