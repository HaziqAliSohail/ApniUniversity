// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClassesOfStudentHandlerFunc turns a function with the right signature into a get classes of student handler
type GetClassesOfStudentHandlerFunc func(GetClassesOfStudentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClassesOfStudentHandlerFunc) Handle(params GetClassesOfStudentParams) middleware.Responder {
	return fn(params)
}

// GetClassesOfStudentHandler interface for that can handle valid get classes of student params
type GetClassesOfStudentHandler interface {
	Handle(GetClassesOfStudentParams) middleware.Responder
}

// NewGetClassesOfStudent creates a new http.Handler for the get classes of student operation
func NewGetClassesOfStudent(ctx *middleware.Context, handler GetClassesOfStudentHandler) *GetClassesOfStudent {
	return &GetClassesOfStudent{Context: ctx, Handler: handler}
}

/*
	GetClassesOfStudent swagger:route GET /student/{ID}/classes getClassesOfStudent

GetClassesOfStudent get classes of student API
*/
type GetClassesOfStudent struct {
	Context *middleware.Context
	Handler GetClassesOfStudentHandler
}

func (o *GetClassesOfStudent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClassesOfStudentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
