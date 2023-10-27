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

// NewUpdateAccountParams creates a new UpdateAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAccountParams() *UpdateAccountParams {
	return &UpdateAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccountParamsWithTimeout creates a new UpdateAccountParams object
// with the ability to set a timeout on a request.
func NewUpdateAccountParamsWithTimeout(timeout time.Duration) *UpdateAccountParams {
	return &UpdateAccountParams{
		timeout: timeout,
	}
}

// NewUpdateAccountParamsWithContext creates a new UpdateAccountParams object
// with the ability to set a context for a request.
func NewUpdateAccountParamsWithContext(ctx context.Context) *UpdateAccountParams {
	return &UpdateAccountParams{
		Context: ctx,
	}
}

// NewUpdateAccountParamsWithHTTPClient creates a new UpdateAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAccountParamsWithHTTPClient(client *http.Client) *UpdateAccountParams {
	return &UpdateAccountParams{
		HTTPClient: client,
	}
}

/*
UpdateAccountParams contains all the parameters to send to the API endpoint

	for the update account operation.

	Typically these are written to a http.Request.
*/
type UpdateAccountParams struct {

	/* Account.

	   Account details
	*/
	Account *models.Account

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccountParams) WithDefaults() *UpdateAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update account params
func (o *UpdateAccountParams) WithTimeout(timeout time.Duration) *UpdateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update account params
func (o *UpdateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update account params
func (o *UpdateAccountParams) WithContext(ctx context.Context) *UpdateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update account params
func (o *UpdateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update account params
func (o *UpdateAccountParams) WithHTTPClient(client *http.Client) *UpdateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update account params
func (o *UpdateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the update account params
func (o *UpdateAccountParams) WithAccount(account *models.Account) *UpdateAccountParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the update account params
func (o *UpdateAccountParams) SetAccount(account *models.Account) {
	o.Account = account
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Account != nil {
		if err := r.SetBodyParam(o.Account); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}