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
)

// NewGetDefaultedStudentAccountsParams creates a new GetDefaultedStudentAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDefaultedStudentAccountsParams() *GetDefaultedStudentAccountsParams {
	return &GetDefaultedStudentAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefaultedStudentAccountsParamsWithTimeout creates a new GetDefaultedStudentAccountsParams object
// with the ability to set a timeout on a request.
func NewGetDefaultedStudentAccountsParamsWithTimeout(timeout time.Duration) *GetDefaultedStudentAccountsParams {
	return &GetDefaultedStudentAccountsParams{
		timeout: timeout,
	}
}

// NewGetDefaultedStudentAccountsParamsWithContext creates a new GetDefaultedStudentAccountsParams object
// with the ability to set a context for a request.
func NewGetDefaultedStudentAccountsParamsWithContext(ctx context.Context) *GetDefaultedStudentAccountsParams {
	return &GetDefaultedStudentAccountsParams{
		Context: ctx,
	}
}

// NewGetDefaultedStudentAccountsParamsWithHTTPClient creates a new GetDefaultedStudentAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDefaultedStudentAccountsParamsWithHTTPClient(client *http.Client) *GetDefaultedStudentAccountsParams {
	return &GetDefaultedStudentAccountsParams{
		HTTPClient: client,
	}
}

/*
GetDefaultedStudentAccountsParams contains all the parameters to send to the API endpoint

	for the get defaulted student accounts operation.

	Typically these are written to a http.Request.
*/
type GetDefaultedStudentAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get defaulted student accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDefaultedStudentAccountsParams) WithDefaults() *GetDefaultedStudentAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get defaulted student accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDefaultedStudentAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get defaulted student accounts params
func (o *GetDefaultedStudentAccountsParams) WithTimeout(timeout time.Duration) *GetDefaultedStudentAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get defaulted student accounts params
func (o *GetDefaultedStudentAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get defaulted student accounts params
func (o *GetDefaultedStudentAccountsParams) WithContext(ctx context.Context) *GetDefaultedStudentAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get defaulted student accounts params
func (o *GetDefaultedStudentAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get defaulted student accounts params
func (o *GetDefaultedStudentAccountsParams) WithHTTPClient(client *http.Client) *GetDefaultedStudentAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get defaulted student accounts params
func (o *GetDefaultedStudentAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefaultedStudentAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
