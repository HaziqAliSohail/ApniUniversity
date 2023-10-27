// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"ApniUniversity/gen/models"
)

// UpdateAccountReader is a Reader for the UpdateAccount structure.
type UpdateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /account] updateAccount", response, response.Code())
	}
}

// NewUpdateAccountOK creates a UpdateAccountOK with default headers values
func NewUpdateAccountOK() *UpdateAccountOK {
	return &UpdateAccountOK{}
}

/*
UpdateAccountOK describes a response with status code 200, with default header values.

Account updated
*/
type UpdateAccountOK struct {
	Payload *models.Account
}

// IsSuccess returns true when this update account o k response has a 2xx status code
func (o *UpdateAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update account o k response has a 3xx status code
func (o *UpdateAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update account o k response has a 4xx status code
func (o *UpdateAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update account o k response has a 5xx status code
func (o *UpdateAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update account o k response a status code equal to that given
func (o *UpdateAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update account o k response
func (o *UpdateAccountOK) Code() int {
	return 200
}

func (o *UpdateAccountOK) Error() string {
	return fmt.Sprintf("[PUT /account][%d] updateAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateAccountOK) String() string {
	return fmt.Sprintf("[PUT /account][%d] updateAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateAccountOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *UpdateAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccountNotFound creates a UpdateAccountNotFound with default headers values
func NewUpdateAccountNotFound() *UpdateAccountNotFound {
	return &UpdateAccountNotFound{}
}

/*
UpdateAccountNotFound describes a response with status code 404, with default header values.

Account not found!
*/
type UpdateAccountNotFound struct {
}

// IsSuccess returns true when this update account not found response has a 2xx status code
func (o *UpdateAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update account not found response has a 3xx status code
func (o *UpdateAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update account not found response has a 4xx status code
func (o *UpdateAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update account not found response has a 5xx status code
func (o *UpdateAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update account not found response a status code equal to that given
func (o *UpdateAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update account not found response
func (o *UpdateAccountNotFound) Code() int {
	return 404
}

func (o *UpdateAccountNotFound) Error() string {
	return fmt.Sprintf("[PUT /account][%d] updateAccountNotFound ", 404)
}

func (o *UpdateAccountNotFound) String() string {
	return fmt.Sprintf("[PUT /account][%d] updateAccountNotFound ", 404)
}

func (o *UpdateAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccountInternalServerError creates a UpdateAccountInternalServerError with default headers values
func NewUpdateAccountInternalServerError() *UpdateAccountInternalServerError {
	return &UpdateAccountInternalServerError{}
}

/*
UpdateAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateAccountInternalServerError struct {
}

// IsSuccess returns true when this update account internal server error response has a 2xx status code
func (o *UpdateAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update account internal server error response has a 3xx status code
func (o *UpdateAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update account internal server error response has a 4xx status code
func (o *UpdateAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update account internal server error response has a 5xx status code
func (o *UpdateAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update account internal server error response a status code equal to that given
func (o *UpdateAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update account internal server error response
func (o *UpdateAccountInternalServerError) Code() int {
	return 500
}

func (o *UpdateAccountInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /account][%d] updateAccountInternalServerError ", 500)
}

func (o *UpdateAccountInternalServerError) String() string {
	return fmt.Sprintf("[PUT /account][%d] updateAccountInternalServerError ", 500)
}

func (o *UpdateAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
