// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStudentsOfTeacherHandlerFunc turns a function with the right signature into a get students of teacher handler
type GetStudentsOfTeacherHandlerFunc func(GetStudentsOfTeacherParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStudentsOfTeacherHandlerFunc) Handle(params GetStudentsOfTeacherParams) middleware.Responder {
	return fn(params)
}

// GetStudentsOfTeacherHandler interface for that can handle valid get students of teacher params
type GetStudentsOfTeacherHandler interface {
	Handle(GetStudentsOfTeacherParams) middleware.Responder
}

// NewGetStudentsOfTeacher creates a new http.Handler for the get students of teacher operation
func NewGetStudentsOfTeacher(ctx *middleware.Context, handler GetStudentsOfTeacherHandler) *GetStudentsOfTeacher {
	return &GetStudentsOfTeacher{Context: ctx, Handler: handler}
}

/*
	GetStudentsOfTeacher swagger:route GET /teacher/{ID}/students getStudentsOfTeacher

GetStudentsOfTeacher get students of teacher API
*/
type GetStudentsOfTeacher struct {
	Context *middleware.Context
	Handler GetStudentsOfTeacherHandler
}

func (o *GetStudentsOfTeacher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetStudentsOfTeacherParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
