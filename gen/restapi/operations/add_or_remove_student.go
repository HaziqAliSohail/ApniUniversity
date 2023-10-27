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

// AddOrRemoveStudentHandlerFunc turns a function with the right signature into a add or remove student handler
type AddOrRemoveStudentHandlerFunc func(AddOrRemoveStudentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOrRemoveStudentHandlerFunc) Handle(params AddOrRemoveStudentParams) middleware.Responder {
	return fn(params)
}

// AddOrRemoveStudentHandler interface for that can handle valid add or remove student params
type AddOrRemoveStudentHandler interface {
	Handle(AddOrRemoveStudentParams) middleware.Responder
}

// NewAddOrRemoveStudent creates a new http.Handler for the add or remove student operation
func NewAddOrRemoveStudent(ctx *middleware.Context, handler AddOrRemoveStudentHandler) *AddOrRemoveStudent {
	return &AddOrRemoveStudent{Context: ctx, Handler: handler}
}

/*
	AddOrRemoveStudent swagger:route PATCH /class/{ID}/student/ addOrRemoveStudent

AddOrRemoveStudent add or remove student API
*/
type AddOrRemoveStudent struct {
	Context *middleware.Context
	Handler AddOrRemoveStudentHandler
}

func (o *AddOrRemoveStudent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddOrRemoveStudentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddOrRemoveStudentBody add or remove student body
//
// swagger:model AddOrRemoveStudentBody
type AddOrRemoveStudentBody struct {

	// add
	Add bool `json:"Add,omitempty"`

	// subject ID
	SubjectID int64 `json:"SubjectID,omitempty"`
}

// Validate validates this add or remove student body
func (o *AddOrRemoveStudentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add or remove student body based on context it is used
func (o *AddOrRemoveStudentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOrRemoveStudentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOrRemoveStudentBody) UnmarshalBinary(b []byte) error {
	var res AddOrRemoveStudentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}