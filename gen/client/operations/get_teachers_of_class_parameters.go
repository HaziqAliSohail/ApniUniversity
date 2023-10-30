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

// NewGetTeachersOfClassParams creates a new GetTeachersOfClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeachersOfClassParams() *GetTeachersOfClassParams {
	return &GetTeachersOfClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeachersOfClassParamsWithTimeout creates a new GetTeachersOfClassParams object
// with the ability to set a timeout on a request.
func NewGetTeachersOfClassParamsWithTimeout(timeout time.Duration) *GetTeachersOfClassParams {
	return &GetTeachersOfClassParams{
		timeout: timeout,
	}
}

// NewGetTeachersOfClassParamsWithContext creates a new GetTeachersOfClassParams object
// with the ability to set a context for a request.
func NewGetTeachersOfClassParamsWithContext(ctx context.Context) *GetTeachersOfClassParams {
	return &GetTeachersOfClassParams{
		Context: ctx,
	}
}

// NewGetTeachersOfClassParamsWithHTTPClient creates a new GetTeachersOfClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeachersOfClassParamsWithHTTPClient(client *http.Client) *GetTeachersOfClassParams {
	return &GetTeachersOfClassParams{
		HTTPClient: client,
	}
}

/*
GetTeachersOfClassParams contains all the parameters to send to the API endpoint

	for the get teachers of class operation.

	Typically these are written to a http.Request.
*/
type GetTeachersOfClassParams struct {

	/* ID.

	   The ID of the Class
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get teachers of class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeachersOfClassParams) WithDefaults() *GetTeachersOfClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get teachers of class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeachersOfClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get teachers of class params
func (o *GetTeachersOfClassParams) WithTimeout(timeout time.Duration) *GetTeachersOfClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get teachers of class params
func (o *GetTeachersOfClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get teachers of class params
func (o *GetTeachersOfClassParams) WithContext(ctx context.Context) *GetTeachersOfClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get teachers of class params
func (o *GetTeachersOfClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get teachers of class params
func (o *GetTeachersOfClassParams) WithHTTPClient(client *http.Client) *GetTeachersOfClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get teachers of class params
func (o *GetTeachersOfClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get teachers of class params
func (o *GetTeachersOfClassParams) WithID(id int64) *GetTeachersOfClassParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get teachers of class params
func (o *GetTeachersOfClassParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeachersOfClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
