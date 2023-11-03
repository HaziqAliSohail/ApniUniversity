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

// GetClassesReader is a Reader for the GetClasses structure.
type GetClassesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClassesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClassesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetClassesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClassesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /class] getClasses", response, response.Code())
	}
}

// NewGetClassesOK creates a GetClassesOK with default headers values
func NewGetClassesOK() *GetClassesOK {
	return &GetClassesOK{}
}

/*
GetClassesOK describes a response with status code 200, with default header values.

Classes' response
*/
type GetClassesOK struct {
	Payload []*models.Class
}

// IsSuccess returns true when this get classes o k response has a 2xx status code
func (o *GetClassesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get classes o k response has a 3xx status code
func (o *GetClassesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get classes o k response has a 4xx status code
func (o *GetClassesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get classes o k response has a 5xx status code
func (o *GetClassesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get classes o k response a status code equal to that given
func (o *GetClassesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get classes o k response
func (o *GetClassesOK) Code() int {
	return 200
}

func (o *GetClassesOK) Error() string {
	return fmt.Sprintf("[GET /class][%d] getClassesOK  %+v", 200, o.Payload)
}

func (o *GetClassesOK) String() string {
	return fmt.Sprintf("[GET /class][%d] getClassesOK  %+v", 200, o.Payload)
}

func (o *GetClassesOK) GetPayload() []*models.Class {
	return o.Payload
}

func (o *GetClassesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClassesNotFound creates a GetClassesNotFound with default headers values
func NewGetClassesNotFound() *GetClassesNotFound {
	return &GetClassesNotFound{}
}

/*
GetClassesNotFound describes a response with status code 404, with default header values.

Class not found
*/
type GetClassesNotFound struct {
}

// IsSuccess returns true when this get classes not found response has a 2xx status code
func (o *GetClassesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get classes not found response has a 3xx status code
func (o *GetClassesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get classes not found response has a 4xx status code
func (o *GetClassesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get classes not found response has a 5xx status code
func (o *GetClassesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get classes not found response a status code equal to that given
func (o *GetClassesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get classes not found response
func (o *GetClassesNotFound) Code() int {
	return 404
}

func (o *GetClassesNotFound) Error() string {
	return fmt.Sprintf("[GET /class][%d] getClassesNotFound ", 404)
}

func (o *GetClassesNotFound) String() string {
	return fmt.Sprintf("[GET /class][%d] getClassesNotFound ", 404)
}

func (o *GetClassesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClassesInternalServerError creates a GetClassesInternalServerError with default headers values
func NewGetClassesInternalServerError() *GetClassesInternalServerError {
	return &GetClassesInternalServerError{}
}

/*
GetClassesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetClassesInternalServerError struct {
}

// IsSuccess returns true when this get classes internal server error response has a 2xx status code
func (o *GetClassesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get classes internal server error response has a 3xx status code
func (o *GetClassesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get classes internal server error response has a 4xx status code
func (o *GetClassesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get classes internal server error response has a 5xx status code
func (o *GetClassesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get classes internal server error response a status code equal to that given
func (o *GetClassesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get classes internal server error response
func (o *GetClassesInternalServerError) Code() int {
	return 500
}

func (o *GetClassesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /class][%d] getClassesInternalServerError ", 500)
}

func (o *GetClassesInternalServerError) String() string {
	return fmt.Sprintf("[GET /class][%d] getClassesInternalServerError ", 500)
}

func (o *GetClassesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}