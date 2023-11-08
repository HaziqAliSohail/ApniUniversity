// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// NewAddStudentParams creates a new AddStudentParams object
//
// There are no default values defined in the spec.
func NewAddStudentParams() AddStudentParams {

	return AddStudentParams{}
}

// AddStudentParams contains all the bound params for the add student operation
// typically these are obtained from a http.Request
//
// swagger:parameters addStudent
type AddStudentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Student to be added
	  Required: true
	  In: body
	*/
	Student *models.Student
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddStudentParams() beforehand.
func (o *AddStudentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Student
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("student", "body", ""))
			} else {
				res = append(res, errors.NewParseError("student", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Student = &body
			}
		}
	} else {
		res = append(res, errors.Required("student", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
