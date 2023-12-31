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

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// NewAddSubjectParams creates a new AddSubjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSubjectParams() *AddSubjectParams {
	return &AddSubjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSubjectParamsWithTimeout creates a new AddSubjectParams object
// with the ability to set a timeout on a request.
func NewAddSubjectParamsWithTimeout(timeout time.Duration) *AddSubjectParams {
	return &AddSubjectParams{
		timeout: timeout,
	}
}

// NewAddSubjectParamsWithContext creates a new AddSubjectParams object
// with the ability to set a context for a request.
func NewAddSubjectParamsWithContext(ctx context.Context) *AddSubjectParams {
	return &AddSubjectParams{
		Context: ctx,
	}
}

// NewAddSubjectParamsWithHTTPClient creates a new AddSubjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddSubjectParamsWithHTTPClient(client *http.Client) *AddSubjectParams {
	return &AddSubjectParams{
		HTTPClient: client,
	}
}

/*
AddSubjectParams contains all the parameters to send to the API endpoint

	for the add subject operation.

	Typically these are written to a http.Request.
*/
type AddSubjectParams struct {

	/* Subject.

	   Subject to be added
	*/
	Subject *models.Subject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSubjectParams) WithDefaults() *AddSubjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSubjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add subject params
func (o *AddSubjectParams) WithTimeout(timeout time.Duration) *AddSubjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add subject params
func (o *AddSubjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add subject params
func (o *AddSubjectParams) WithContext(ctx context.Context) *AddSubjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add subject params
func (o *AddSubjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add subject params
func (o *AddSubjectParams) WithHTTPClient(client *http.Client) *AddSubjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add subject params
func (o *AddSubjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubject adds the subject to the add subject params
func (o *AddSubjectParams) WithSubject(subject *models.Subject) *AddSubjectParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the add subject params
func (o *AddSubjectParams) SetSubject(subject *models.Subject) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *AddSubjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Subject != nil {
		if err := r.SetBodyParam(o.Subject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
