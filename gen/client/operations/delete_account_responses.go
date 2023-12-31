// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAccountReader is a Reader for the DeleteAccount structure.
type DeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /account/{ID}] deleteAccount", response, response.Code())
	}
}

// NewDeleteAccountNoContent creates a DeleteAccountNoContent with default headers values
func NewDeleteAccountNoContent() *DeleteAccountNoContent {
	return &DeleteAccountNoContent{}
}

/*
DeleteAccountNoContent describes a response with status code 204, with default header values.

Account Deleted Successfully!
*/
type DeleteAccountNoContent struct {
}

// IsSuccess returns true when this delete account no content response has a 2xx status code
func (o *DeleteAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete account no content response has a 3xx status code
func (o *DeleteAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete account no content response has a 4xx status code
func (o *DeleteAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete account no content response has a 5xx status code
func (o *DeleteAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete account no content response a status code equal to that given
func (o *DeleteAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete account no content response
func (o *DeleteAccountNoContent) Code() int {
	return 204
}

func (o *DeleteAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /account/{ID}][%d] deleteAccountNoContent ", 204)
}

func (o *DeleteAccountNoContent) String() string {
	return fmt.Sprintf("[DELETE /account/{ID}][%d] deleteAccountNoContent ", 204)
}

func (o *DeleteAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAccountNotFound creates a DeleteAccountNotFound with default headers values
func NewDeleteAccountNotFound() *DeleteAccountNotFound {
	return &DeleteAccountNotFound{}
}

/*
DeleteAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type DeleteAccountNotFound struct {
}

// IsSuccess returns true when this delete account not found response has a 2xx status code
func (o *DeleteAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete account not found response has a 3xx status code
func (o *DeleteAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete account not found response has a 4xx status code
func (o *DeleteAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete account not found response has a 5xx status code
func (o *DeleteAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete account not found response a status code equal to that given
func (o *DeleteAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete account not found response
func (o *DeleteAccountNotFound) Code() int {
	return 404
}

func (o *DeleteAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /account/{ID}][%d] deleteAccountNotFound ", 404)
}

func (o *DeleteAccountNotFound) String() string {
	return fmt.Sprintf("[DELETE /account/{ID}][%d] deleteAccountNotFound ", 404)
}

func (o *DeleteAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAccountInternalServerError creates a DeleteAccountInternalServerError with default headers values
func NewDeleteAccountInternalServerError() *DeleteAccountInternalServerError {
	return &DeleteAccountInternalServerError{}
}

/*
DeleteAccountInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteAccountInternalServerError struct {
}

// IsSuccess returns true when this delete account internal server error response has a 2xx status code
func (o *DeleteAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete account internal server error response has a 3xx status code
func (o *DeleteAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete account internal server error response has a 4xx status code
func (o *DeleteAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete account internal server error response has a 5xx status code
func (o *DeleteAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete account internal server error response a status code equal to that given
func (o *DeleteAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete account internal server error response
func (o *DeleteAccountInternalServerError) Code() int {
	return 500
}

func (o *DeleteAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /account/{ID}][%d] deleteAccountInternalServerError ", 500)
}

func (o *DeleteAccountInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /account/{ID}][%d] deleteAccountInternalServerError ", 500)
}

func (o *DeleteAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
