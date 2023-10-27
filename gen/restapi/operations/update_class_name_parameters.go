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

// NewUpdateClassNameParams creates a new UpdateClassNameParams object
//
// There are no default values defined in the spec.
func NewUpdateClassNameParams() UpdateClassNameParams {

	return UpdateClassNameParams{}
}

// UpdateClassNameParams contains all the bound params for the update class name operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateClassName
type UpdateClassNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Class ID
	  Required: true
	  In: path
	*/
	ID int64
	/*Updated Name to be stored
	  Required: true
	  In: body
	*/
	Name UpdateClassNameBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateClassNameParams() beforehand.
func (o *UpdateClassNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("ID")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body UpdateClassNameBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("name", "body", ""))
			} else {
				res = append(res, errors.NewParseError("name", "body", "", err))
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
				o.Name = body
			}
		}
	} else {
		res = append(res, errors.Required("name", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdateClassNameParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
