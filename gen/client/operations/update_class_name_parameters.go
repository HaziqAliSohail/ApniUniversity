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

// NewUpdateClassNameParams creates a new UpdateClassNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateClassNameParams() *UpdateClassNameParams {
	return &UpdateClassNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClassNameParamsWithTimeout creates a new UpdateClassNameParams object
// with the ability to set a timeout on a request.
func NewUpdateClassNameParamsWithTimeout(timeout time.Duration) *UpdateClassNameParams {
	return &UpdateClassNameParams{
		timeout: timeout,
	}
}

// NewUpdateClassNameParamsWithContext creates a new UpdateClassNameParams object
// with the ability to set a context for a request.
func NewUpdateClassNameParamsWithContext(ctx context.Context) *UpdateClassNameParams {
	return &UpdateClassNameParams{
		Context: ctx,
	}
}

// NewUpdateClassNameParamsWithHTTPClient creates a new UpdateClassNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateClassNameParamsWithHTTPClient(client *http.Client) *UpdateClassNameParams {
	return &UpdateClassNameParams{
		HTTPClient: client,
	}
}

/*
UpdateClassNameParams contains all the parameters to send to the API endpoint

	for the update class name operation.

	Typically these are written to a http.Request.
*/
type UpdateClassNameParams struct {

	/* ID.

	   Class ID
	*/
	ID int64

	/* Name.

	   Updated Name to be stored
	*/
	Name UpdateClassNameBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update class name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClassNameParams) WithDefaults() *UpdateClassNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update class name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClassNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update class name params
func (o *UpdateClassNameParams) WithTimeout(timeout time.Duration) *UpdateClassNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update class name params
func (o *UpdateClassNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update class name params
func (o *UpdateClassNameParams) WithContext(ctx context.Context) *UpdateClassNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update class name params
func (o *UpdateClassNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update class name params
func (o *UpdateClassNameParams) WithHTTPClient(client *http.Client) *UpdateClassNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update class name params
func (o *UpdateClassNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update class name params
func (o *UpdateClassNameParams) WithID(id int64) *UpdateClassNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update class name params
func (o *UpdateClassNameParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the update class name params
func (o *UpdateClassNameParams) WithName(name UpdateClassNameBody) *UpdateClassNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update class name params
func (o *UpdateClassNameParams) SetName(name UpdateClassNameBody) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClassNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
