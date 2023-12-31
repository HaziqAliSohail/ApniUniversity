// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// AddAccountReader is a Reader for the AddAccount structure.
type AddAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewAddAccountConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /account] addAccount", response, response.Code())
	}
}

// NewAddAccountCreated creates a AddAccountCreated with default headers values
func NewAddAccountCreated() *AddAccountCreated {
	return &AddAccountCreated{}
}

/*
AddAccountCreated describes a response with status code 201, with default header values.

Account added Successfully
*/
type AddAccountCreated struct {
	Payload *models.Account
}

// IsSuccess returns true when this add account created response has a 2xx status code
func (o *AddAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add account created response has a 3xx status code
func (o *AddAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add account created response has a 4xx status code
func (o *AddAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add account created response has a 5xx status code
func (o *AddAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add account created response a status code equal to that given
func (o *AddAccountCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add account created response
func (o *AddAccountCreated) Code() int {
	return 201
}

func (o *AddAccountCreated) Error() string {
	return fmt.Sprintf("[POST /account][%d] addAccountCreated  %+v", 201, o.Payload)
}

func (o *AddAccountCreated) String() string {
	return fmt.Sprintf("[POST /account][%d] addAccountCreated  %+v", 201, o.Payload)
}

func (o *AddAccountCreated) GetPayload() *models.Account {
	return o.Payload
}

func (o *AddAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAccountConflict creates a AddAccountConflict with default headers values
func NewAddAccountConflict() *AddAccountConflict {
	return &AddAccountConflict{}
}

/*
AddAccountConflict describes a response with status code 409, with default header values.

Account already exists
*/
type AddAccountConflict struct {
}

// IsSuccess returns true when this add account conflict response has a 2xx status code
func (o *AddAccountConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add account conflict response has a 3xx status code
func (o *AddAccountConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add account conflict response has a 4xx status code
func (o *AddAccountConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add account conflict response has a 5xx status code
func (o *AddAccountConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add account conflict response a status code equal to that given
func (o *AddAccountConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add account conflict response
func (o *AddAccountConflict) Code() int {
	return 409
}

func (o *AddAccountConflict) Error() string {
	return fmt.Sprintf("[POST /account][%d] addAccountConflict ", 409)
}

func (o *AddAccountConflict) String() string {
	return fmt.Sprintf("[POST /account][%d] addAccountConflict ", 409)
}

func (o *AddAccountConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAccountInternalServerError creates a AddAccountInternalServerError with default headers values
func NewAddAccountInternalServerError() *AddAccountInternalServerError {
	return &AddAccountInternalServerError{}
}

/*
AddAccountInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddAccountInternalServerError struct {
}

// IsSuccess returns true when this add account internal server error response has a 2xx status code
func (o *AddAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add account internal server error response has a 3xx status code
func (o *AddAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add account internal server error response has a 4xx status code
func (o *AddAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add account internal server error response has a 5xx status code
func (o *AddAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add account internal server error response a status code equal to that given
func (o *AddAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add account internal server error response
func (o *AddAccountInternalServerError) Code() int {
	return 500
}

func (o *AddAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /account][%d] addAccountInternalServerError ", 500)
}

func (o *AddAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /account][%d] addAccountInternalServerError ", 500)
}

func (o *AddAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
