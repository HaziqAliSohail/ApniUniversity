// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssignSubjectToTeacherHandlerFunc turns a function with the right signature into a assign subject to teacher handler
type AssignSubjectToTeacherHandlerFunc func(AssignSubjectToTeacherParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssignSubjectToTeacherHandlerFunc) Handle(params AssignSubjectToTeacherParams) middleware.Responder {
	return fn(params)
}

// AssignSubjectToTeacherHandler interface for that can handle valid assign subject to teacher params
type AssignSubjectToTeacherHandler interface {
	Handle(AssignSubjectToTeacherParams) middleware.Responder
}

// NewAssignSubjectToTeacher creates a new http.Handler for the assign subject to teacher operation
func NewAssignSubjectToTeacher(ctx *middleware.Context, handler AssignSubjectToTeacherHandler) *AssignSubjectToTeacher {
	return &AssignSubjectToTeacher{Context: ctx, Handler: handler}
}

/*
	AssignSubjectToTeacher swagger:route PATCH /teacher/{ID}/subject assignSubjectToTeacher

AssignSubjectToTeacher assign subject to teacher API
*/
type AssignSubjectToTeacher struct {
	Context *middleware.Context
	Handler AssignSubjectToTeacherHandler
}

func (o *AssignSubjectToTeacher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAssignSubjectToTeacherParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AssignSubjectToTeacherBody assign subject to teacher body
//
// swagger:model AssignSubjectToTeacherBody
type AssignSubjectToTeacherBody struct {

	// subject ID
	SubjectID int64 `json:"SubjectID,omitempty"`
}

// Validate validates this assign subject to teacher body
func (o *AssignSubjectToTeacherBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign subject to teacher body based on context it is used
func (o *AssignSubjectToTeacherBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignSubjectToTeacherBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignSubjectToTeacherBody) UnmarshalBinary(b []byte) error {
	var res AssignSubjectToTeacherBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}