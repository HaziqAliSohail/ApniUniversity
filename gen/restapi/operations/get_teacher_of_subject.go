// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTeacherOfSubjectHandlerFunc turns a function with the right signature into a get teacher of subject handler
type GetTeacherOfSubjectHandlerFunc func(GetTeacherOfSubjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeacherOfSubjectHandlerFunc) Handle(params GetTeacherOfSubjectParams) middleware.Responder {
	return fn(params)
}

// GetTeacherOfSubjectHandler interface for that can handle valid get teacher of subject params
type GetTeacherOfSubjectHandler interface {
	Handle(GetTeacherOfSubjectParams) middleware.Responder
}

// NewGetTeacherOfSubject creates a new http.Handler for the get teacher of subject operation
func NewGetTeacherOfSubject(ctx *middleware.Context, handler GetTeacherOfSubjectHandler) *GetTeacherOfSubject {
	return &GetTeacherOfSubject{Context: ctx, Handler: handler}
}

/*
	GetTeacherOfSubject swagger:route GET /subject/{ID}/teacher getTeacherOfSubject

GetTeacherOfSubject get teacher of subject API
*/
type GetTeacherOfSubject struct {
	Context *middleware.Context
	Handler GetTeacherOfSubjectHandler
}

func (o *GetTeacherOfSubject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTeacherOfSubjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
