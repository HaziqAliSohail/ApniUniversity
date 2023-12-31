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

// NewDeleteTeacherParams creates a new DeleteTeacherParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTeacherParams() *DeleteTeacherParams {
	return &DeleteTeacherParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeacherParamsWithTimeout creates a new DeleteTeacherParams object
// with the ability to set a timeout on a request.
func NewDeleteTeacherParamsWithTimeout(timeout time.Duration) *DeleteTeacherParams {
	return &DeleteTeacherParams{
		timeout: timeout,
	}
}

// NewDeleteTeacherParamsWithContext creates a new DeleteTeacherParams object
// with the ability to set a context for a request.
func NewDeleteTeacherParamsWithContext(ctx context.Context) *DeleteTeacherParams {
	return &DeleteTeacherParams{
		Context: ctx,
	}
}

// NewDeleteTeacherParamsWithHTTPClient creates a new DeleteTeacherParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTeacherParamsWithHTTPClient(client *http.Client) *DeleteTeacherParams {
	return &DeleteTeacherParams{
		HTTPClient: client,
	}
}

/*
DeleteTeacherParams contains all the parameters to send to the API endpoint

	for the delete teacher operation.

	Typically these are written to a http.Request.
*/
type DeleteTeacherParams struct {

	/* ID.

	   ID of the Teacher
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete teacher params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeacherParams) WithDefaults() *DeleteTeacherParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete teacher params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeacherParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete teacher params
func (o *DeleteTeacherParams) WithTimeout(timeout time.Duration) *DeleteTeacherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete teacher params
func (o *DeleteTeacherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete teacher params
func (o *DeleteTeacherParams) WithContext(ctx context.Context) *DeleteTeacherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete teacher params
func (o *DeleteTeacherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete teacher params
func (o *DeleteTeacherParams) WithHTTPClient(client *http.Client) *DeleteTeacherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete teacher params
func (o *DeleteTeacherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete teacher params
func (o *DeleteTeacherParams) WithID(id int64) *DeleteTeacherParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete teacher params
func (o *DeleteTeacherParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeacherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
