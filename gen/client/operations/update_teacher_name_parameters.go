// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUpdateTeacherNameParams creates a new UpdateTeacherNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTeacherNameParams() *UpdateTeacherNameParams {
	return &UpdateTeacherNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTeacherNameParamsWithTimeout creates a new UpdateTeacherNameParams object
// with the ability to set a timeout on a request.
func NewUpdateTeacherNameParamsWithTimeout(timeout time.Duration) *UpdateTeacherNameParams {
	return &UpdateTeacherNameParams{
		timeout: timeout,
	}
}

// NewUpdateTeacherNameParamsWithContext creates a new UpdateTeacherNameParams object
// with the ability to set a context for a request.
func NewUpdateTeacherNameParamsWithContext(ctx context.Context) *UpdateTeacherNameParams {
	return &UpdateTeacherNameParams{
		Context: ctx,
	}
}

// NewUpdateTeacherNameParamsWithHTTPClient creates a new UpdateTeacherNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTeacherNameParamsWithHTTPClient(client *http.Client) *UpdateTeacherNameParams {
	return &UpdateTeacherNameParams{
		HTTPClient: client,
	}
}

/*
UpdateTeacherNameParams contains all the parameters to send to the API endpoint

	for the update teacher name operation.

	Typically these are written to a http.Request.
*/
type UpdateTeacherNameParams struct {

	/* ID.

	   Teacher ID
	*/
	ID int64

	/* Name.

	   Updated Name to be stored
	*/
	Name UpdateTeacherNameBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update teacher name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTeacherNameParams) WithDefaults() *UpdateTeacherNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update teacher name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTeacherNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update teacher name params
func (o *UpdateTeacherNameParams) WithTimeout(timeout time.Duration) *UpdateTeacherNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update teacher name params
func (o *UpdateTeacherNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update teacher name params
func (o *UpdateTeacherNameParams) WithContext(ctx context.Context) *UpdateTeacherNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update teacher name params
func (o *UpdateTeacherNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update teacher name params
func (o *UpdateTeacherNameParams) WithHTTPClient(client *http.Client) *UpdateTeacherNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update teacher name params
func (o *UpdateTeacherNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update teacher name params
func (o *UpdateTeacherNameParams) WithID(id int64) *UpdateTeacherNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update teacher name params
func (o *UpdateTeacherNameParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the update teacher name params
func (o *UpdateTeacherNameParams) WithName(name UpdateTeacherNameBody) *UpdateTeacherNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update teacher name params
func (o *UpdateTeacherNameParams) SetName(name UpdateTeacherNameBody) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTeacherNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
