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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewAssignClassParams creates a new AssignClassParams object
//
// There are no default values defined in the spec.
func NewAssignClassParams() AssignClassParams {

	return AssignClassParams{}
}

// AssignClassParams contains all the bound params for the assign class operation
// typically these are obtained from a http.Request
//
// swagger:parameters assignClass
type AssignClassParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Subject ID
	  Required: true
	  In: path
	*/
	ID int64
	/*The object containing ID of the class to be assigned
	  Required: true
	  In: body
	*/
	Class AssignClassBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAssignClassParams() beforehand.
func (o *AssignClassParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("ID")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body AssignClassBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("class", "body", ""))
			} else {
				res = append(res, errors.NewParseError("class", "body", "", err))
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
				o.Class = body
			}
		}
	} else {
		res = append(res, errors.Required("class", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *AssignClassParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ID", "path", "int64", raw)
	}
	o.ID = value

	return nil
}
