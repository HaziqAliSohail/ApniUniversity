// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSubjectReader is a Reader for the DeleteSubject structure.
type DeleteSubjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSubjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSubjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSubjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /subject/{ID}] deleteSubject", response, response.Code())
	}
}

// NewDeleteSubjectNoContent creates a DeleteSubjectNoContent with default headers values
func NewDeleteSubjectNoContent() *DeleteSubjectNoContent {
	return &DeleteSubjectNoContent{}
}

/*
DeleteSubjectNoContent describes a response with status code 204, with default header values.

Subject Deleted Successfully!
*/
type DeleteSubjectNoContent struct {
}

// IsSuccess returns true when this delete subject no content response has a 2xx status code
func (o *DeleteSubjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete subject no content response has a 3xx status code
func (o *DeleteSubjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject no content response has a 4xx status code
func (o *DeleteSubjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subject no content response has a 5xx status code
func (o *DeleteSubjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subject no content response a status code equal to that given
func (o *DeleteSubjectNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete subject no content response
func (o *DeleteSubjectNoContent) Code() int {
	return 204
}

func (o *DeleteSubjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subject/{ID}][%d] deleteSubjectNoContent ", 204)
}

func (o *DeleteSubjectNoContent) String() string {
	return fmt.Sprintf("[DELETE /subject/{ID}][%d] deleteSubjectNoContent ", 204)
}

func (o *DeleteSubjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubjectNotFound creates a DeleteSubjectNotFound with default headers values
func NewDeleteSubjectNotFound() *DeleteSubjectNotFound {
	return &DeleteSubjectNotFound{}
}

/*
DeleteSubjectNotFound describes a response with status code 404, with default header values.

Subject not found
*/
type DeleteSubjectNotFound struct {
}

// IsSuccess returns true when this delete subject not found response has a 2xx status code
func (o *DeleteSubjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subject not found response has a 3xx status code
func (o *DeleteSubjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject not found response has a 4xx status code
func (o *DeleteSubjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete subject not found response has a 5xx status code
func (o *DeleteSubjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subject not found response a status code equal to that given
func (o *DeleteSubjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete subject not found response
func (o *DeleteSubjectNotFound) Code() int {
	return 404
}

func (o *DeleteSubjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /subject/{ID}][%d] deleteSubjectNotFound ", 404)
}

func (o *DeleteSubjectNotFound) String() string {
	return fmt.Sprintf("[DELETE /subject/{ID}][%d] deleteSubjectNotFound ", 404)
}

func (o *DeleteSubjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubjectInternalServerError creates a DeleteSubjectInternalServerError with default headers values
func NewDeleteSubjectInternalServerError() *DeleteSubjectInternalServerError {
	return &DeleteSubjectInternalServerError{}
}

/*
DeleteSubjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteSubjectInternalServerError struct {
}

// IsSuccess returns true when this delete subject internal server error response has a 2xx status code
func (o *DeleteSubjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subject internal server error response has a 3xx status code
func (o *DeleteSubjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject internal server error response has a 4xx status code
func (o *DeleteSubjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subject internal server error response has a 5xx status code
func (o *DeleteSubjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete subject internal server error response a status code equal to that given
func (o *DeleteSubjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete subject internal server error response
func (o *DeleteSubjectInternalServerError) Code() int {
	return 500
}

func (o *DeleteSubjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /subject/{ID}][%d] deleteSubjectInternalServerError ", 500)
}

func (o *DeleteSubjectInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /subject/{ID}][%d] deleteSubjectInternalServerError ", 500)
}

func (o *DeleteSubjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
