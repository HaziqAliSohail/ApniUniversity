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

// NewUpdateStudentNameParams creates a new UpdateStudentNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateStudentNameParams() *UpdateStudentNameParams {
	return &UpdateStudentNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStudentNameParamsWithTimeout creates a new UpdateStudentNameParams object
// with the ability to set a timeout on a request.
func NewUpdateStudentNameParamsWithTimeout(timeout time.Duration) *UpdateStudentNameParams {
	return &UpdateStudentNameParams{
		timeout: timeout,
	}
}

// NewUpdateStudentNameParamsWithContext creates a new UpdateStudentNameParams object
// with the ability to set a context for a request.
func NewUpdateStudentNameParamsWithContext(ctx context.Context) *UpdateStudentNameParams {
	return &UpdateStudentNameParams{
		Context: ctx,
	}
}

// NewUpdateStudentNameParamsWithHTTPClient creates a new UpdateStudentNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateStudentNameParamsWithHTTPClient(client *http.Client) *UpdateStudentNameParams {
	return &UpdateStudentNameParams{
		HTTPClient: client,
	}
}

/*
UpdateStudentNameParams contains all the parameters to send to the API endpoint

	for the update student name operation.

	Typically these are written to a http.Request.
*/
type UpdateStudentNameParams struct {

	/* ID.

	   Student ID
	*/
	ID int64

	/* Name.

	   Updated Name to be stored
	*/
	Name UpdateStudentNameBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update student name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateStudentNameParams) WithDefaults() *UpdateStudentNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update student name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateStudentNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update student name params
func (o *UpdateStudentNameParams) WithTimeout(timeout time.Duration) *UpdateStudentNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update student name params
func (o *UpdateStudentNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update student name params
func (o *UpdateStudentNameParams) WithContext(ctx context.Context) *UpdateStudentNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update student name params
func (o *UpdateStudentNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update student name params
func (o *UpdateStudentNameParams) WithHTTPClient(client *http.Client) *UpdateStudentNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update student name params
func (o *UpdateStudentNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update student name params
func (o *UpdateStudentNameParams) WithID(id int64) *UpdateStudentNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update student name params
func (o *UpdateStudentNameParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the update student name params
func (o *UpdateStudentNameParams) WithName(name UpdateStudentNameBody) *UpdateStudentNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update student name params
func (o *UpdateStudentNameParams) SetName(name UpdateStudentNameBody) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStudentNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
