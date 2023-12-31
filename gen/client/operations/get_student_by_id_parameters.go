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

// NewGetStudentByIDParams creates a new GetStudentByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStudentByIDParams() *GetStudentByIDParams {
	return &GetStudentByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStudentByIDParamsWithTimeout creates a new GetStudentByIDParams object
// with the ability to set a timeout on a request.
func NewGetStudentByIDParamsWithTimeout(timeout time.Duration) *GetStudentByIDParams {
	return &GetStudentByIDParams{
		timeout: timeout,
	}
}

// NewGetStudentByIDParamsWithContext creates a new GetStudentByIDParams object
// with the ability to set a context for a request.
func NewGetStudentByIDParamsWithContext(ctx context.Context) *GetStudentByIDParams {
	return &GetStudentByIDParams{
		Context: ctx,
	}
}

// NewGetStudentByIDParamsWithHTTPClient creates a new GetStudentByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStudentByIDParamsWithHTTPClient(client *http.Client) *GetStudentByIDParams {
	return &GetStudentByIDParams{
		HTTPClient: client,
	}
}

/*
GetStudentByIDParams contains all the parameters to send to the API endpoint

	for the get student by ID operation.

	Typically these are written to a http.Request.
*/
type GetStudentByIDParams struct {

	/* ID.

	   ID of the Student
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get student by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStudentByIDParams) WithDefaults() *GetStudentByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get student by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStudentByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get student by ID params
func (o *GetStudentByIDParams) WithTimeout(timeout time.Duration) *GetStudentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get student by ID params
func (o *GetStudentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get student by ID params
func (o *GetStudentByIDParams) WithContext(ctx context.Context) *GetStudentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get student by ID params
func (o *GetStudentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get student by ID params
func (o *GetStudentByIDParams) WithHTTPClient(client *http.Client) *GetStudentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get student by ID params
func (o *GetStudentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get student by ID params
func (o *GetStudentByIDParams) WithID(id int64) *GetStudentByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get student by ID params
func (o *GetStudentByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStudentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
