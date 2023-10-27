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

// NewDeleteSubjectParams creates a new DeleteSubjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSubjectParams() *DeleteSubjectParams {
	return &DeleteSubjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubjectParamsWithTimeout creates a new DeleteSubjectParams object
// with the ability to set a timeout on a request.
func NewDeleteSubjectParamsWithTimeout(timeout time.Duration) *DeleteSubjectParams {
	return &DeleteSubjectParams{
		timeout: timeout,
	}
}

// NewDeleteSubjectParamsWithContext creates a new DeleteSubjectParams object
// with the ability to set a context for a request.
func NewDeleteSubjectParamsWithContext(ctx context.Context) *DeleteSubjectParams {
	return &DeleteSubjectParams{
		Context: ctx,
	}
}

// NewDeleteSubjectParamsWithHTTPClient creates a new DeleteSubjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSubjectParamsWithHTTPClient(client *http.Client) *DeleteSubjectParams {
	return &DeleteSubjectParams{
		HTTPClient: client,
	}
}

/*
DeleteSubjectParams contains all the parameters to send to the API endpoint

	for the delete subject operation.

	Typically these are written to a http.Request.
*/
type DeleteSubjectParams struct {

	/* ID.

	   ID of the Subject
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubjectParams) WithDefaults() *DeleteSubjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete subject params
func (o *DeleteSubjectParams) WithTimeout(timeout time.Duration) *DeleteSubjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subject params
func (o *DeleteSubjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subject params
func (o *DeleteSubjectParams) WithContext(ctx context.Context) *DeleteSubjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subject params
func (o *DeleteSubjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subject params
func (o *DeleteSubjectParams) WithHTTPClient(client *http.Client) *DeleteSubjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subject params
func (o *DeleteSubjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete subject params
func (o *DeleteSubjectParams) WithID(id int64) *DeleteSubjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete subject params
func (o *DeleteSubjectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
