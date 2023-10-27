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

// NewGetSubjectsOfStudentParams creates a new GetSubjectsOfStudentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSubjectsOfStudentParams() *GetSubjectsOfStudentParams {
	return &GetSubjectsOfStudentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubjectsOfStudentParamsWithTimeout creates a new GetSubjectsOfStudentParams object
// with the ability to set a timeout on a request.
func NewGetSubjectsOfStudentParamsWithTimeout(timeout time.Duration) *GetSubjectsOfStudentParams {
	return &GetSubjectsOfStudentParams{
		timeout: timeout,
	}
}

// NewGetSubjectsOfStudentParamsWithContext creates a new GetSubjectsOfStudentParams object
// with the ability to set a context for a request.
func NewGetSubjectsOfStudentParamsWithContext(ctx context.Context) *GetSubjectsOfStudentParams {
	return &GetSubjectsOfStudentParams{
		Context: ctx,
	}
}

// NewGetSubjectsOfStudentParamsWithHTTPClient creates a new GetSubjectsOfStudentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSubjectsOfStudentParamsWithHTTPClient(client *http.Client) *GetSubjectsOfStudentParams {
	return &GetSubjectsOfStudentParams{
		HTTPClient: client,
	}
}

/*
GetSubjectsOfStudentParams contains all the parameters to send to the API endpoint

	for the get subjects of student operation.

	Typically these are written to a http.Request.
*/
type GetSubjectsOfStudentParams struct {

	/* ID.

	   The ID of the Student
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get subjects of student params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSubjectsOfStudentParams) WithDefaults() *GetSubjectsOfStudentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get subjects of student params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSubjectsOfStudentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get subjects of student params
func (o *GetSubjectsOfStudentParams) WithTimeout(timeout time.Duration) *GetSubjectsOfStudentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subjects of student params
func (o *GetSubjectsOfStudentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subjects of student params
func (o *GetSubjectsOfStudentParams) WithContext(ctx context.Context) *GetSubjectsOfStudentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subjects of student params
func (o *GetSubjectsOfStudentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subjects of student params
func (o *GetSubjectsOfStudentParams) WithHTTPClient(client *http.Client) *GetSubjectsOfStudentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subjects of student params
func (o *GetSubjectsOfStudentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get subjects of student params
func (o *GetSubjectsOfStudentParams) WithID(id int64) *GetSubjectsOfStudentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get subjects of student params
func (o *GetSubjectsOfStudentParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubjectsOfStudentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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