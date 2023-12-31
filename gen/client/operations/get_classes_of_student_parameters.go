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

// NewGetClassesOfStudentParams creates a new GetClassesOfStudentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClassesOfStudentParams() *GetClassesOfStudentParams {
	return &GetClassesOfStudentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClassesOfStudentParamsWithTimeout creates a new GetClassesOfStudentParams object
// with the ability to set a timeout on a request.
func NewGetClassesOfStudentParamsWithTimeout(timeout time.Duration) *GetClassesOfStudentParams {
	return &GetClassesOfStudentParams{
		timeout: timeout,
	}
}

// NewGetClassesOfStudentParamsWithContext creates a new GetClassesOfStudentParams object
// with the ability to set a context for a request.
func NewGetClassesOfStudentParamsWithContext(ctx context.Context) *GetClassesOfStudentParams {
	return &GetClassesOfStudentParams{
		Context: ctx,
	}
}

// NewGetClassesOfStudentParamsWithHTTPClient creates a new GetClassesOfStudentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClassesOfStudentParamsWithHTTPClient(client *http.Client) *GetClassesOfStudentParams {
	return &GetClassesOfStudentParams{
		HTTPClient: client,
	}
}

/*
GetClassesOfStudentParams contains all the parameters to send to the API endpoint

	for the get classes of student operation.

	Typically these are written to a http.Request.
*/
type GetClassesOfStudentParams struct {

	/* ID.

	   The ID of the Student
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get classes of student params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClassesOfStudentParams) WithDefaults() *GetClassesOfStudentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get classes of student params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClassesOfStudentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get classes of student params
func (o *GetClassesOfStudentParams) WithTimeout(timeout time.Duration) *GetClassesOfStudentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get classes of student params
func (o *GetClassesOfStudentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get classes of student params
func (o *GetClassesOfStudentParams) WithContext(ctx context.Context) *GetClassesOfStudentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get classes of student params
func (o *GetClassesOfStudentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get classes of student params
func (o *GetClassesOfStudentParams) WithHTTPClient(client *http.Client) *GetClassesOfStudentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get classes of student params
func (o *GetClassesOfStudentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get classes of student params
func (o *GetClassesOfStudentParams) WithID(id int64) *GetClassesOfStudentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get classes of student params
func (o *GetClassesOfStudentParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetClassesOfStudentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
