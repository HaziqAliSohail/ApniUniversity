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

	"ApniUniversity/gen/models"
)

// NewAddClassParams creates a new AddClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddClassParams() *AddClassParams {
	return &AddClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddClassParamsWithTimeout creates a new AddClassParams object
// with the ability to set a timeout on a request.
func NewAddClassParamsWithTimeout(timeout time.Duration) *AddClassParams {
	return &AddClassParams{
		timeout: timeout,
	}
}

// NewAddClassParamsWithContext creates a new AddClassParams object
// with the ability to set a context for a request.
func NewAddClassParamsWithContext(ctx context.Context) *AddClassParams {
	return &AddClassParams{
		Context: ctx,
	}
}

// NewAddClassParamsWithHTTPClient creates a new AddClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddClassParamsWithHTTPClient(client *http.Client) *AddClassParams {
	return &AddClassParams{
		HTTPClient: client,
	}
}

/*
AddClassParams contains all the parameters to send to the API endpoint

	for the add class operation.

	Typically these are written to a http.Request.
*/
type AddClassParams struct {

	/* Class.

	   Class to be added
	*/
	Class *models.Class

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddClassParams) WithDefaults() *AddClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add class params
func (o *AddClassParams) WithTimeout(timeout time.Duration) *AddClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add class params
func (o *AddClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add class params
func (o *AddClassParams) WithContext(ctx context.Context) *AddClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add class params
func (o *AddClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add class params
func (o *AddClassParams) WithHTTPClient(client *http.Client) *AddClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add class params
func (o *AddClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClass adds the class to the add class params
func (o *AddClassParams) WithClass(class *models.Class) *AddClassParams {
	o.SetClass(class)
	return o
}

// SetClass adds the class to the add class params
func (o *AddClassParams) SetClass(class *models.Class) {
	o.Class = class
}

// WriteToRequest writes these params to a swagger request
func (o *AddClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Class != nil {
		if err := r.SetBodyParam(o.Class); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
