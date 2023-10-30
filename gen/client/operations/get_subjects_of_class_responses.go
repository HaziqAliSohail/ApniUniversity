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

// GetSubjectsOfClassReader is a Reader for the GetSubjectsOfClass structure.
type GetSubjectsOfClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubjectsOfClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubjectsOfClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSubjectsOfClassNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSubjectsOfClassInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /class/{ID}/subjects] getSubjectsOfClass", response, response.Code())
	}
}

// NewGetSubjectsOfClassOK creates a GetSubjectsOfClassOK with default headers values
func NewGetSubjectsOfClassOK() *GetSubjectsOfClassOK {
	return &GetSubjectsOfClassOK{}
}

/*
GetSubjectsOfClassOK describes a response with status code 200, with default header values.

Class' Subjects as response
*/
type GetSubjectsOfClassOK struct {
	Payload []*models.Subject
}

// IsSuccess returns true when this get subjects of class o k response has a 2xx status code
func (o *GetSubjectsOfClassOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subjects of class o k response has a 3xx status code
func (o *GetSubjectsOfClassOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subjects of class o k response has a 4xx status code
func (o *GetSubjectsOfClassOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subjects of class o k response has a 5xx status code
func (o *GetSubjectsOfClassOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subjects of class o k response a status code equal to that given
func (o *GetSubjectsOfClassOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get subjects of class o k response
func (o *GetSubjectsOfClassOK) Code() int {
	return 200
}

func (o *GetSubjectsOfClassOK) Error() string {
	return fmt.Sprintf("[GET /class/{ID}/subjects][%d] getSubjectsOfClassOK  %+v", 200, o.Payload)
}

func (o *GetSubjectsOfClassOK) String() string {
	return fmt.Sprintf("[GET /class/{ID}/subjects][%d] getSubjectsOfClassOK  %+v", 200, o.Payload)
}

func (o *GetSubjectsOfClassOK) GetPayload() []*models.Subject {
	return o.Payload
}

func (o *GetSubjectsOfClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubjectsOfClassNotFound creates a GetSubjectsOfClassNotFound with default headers values
func NewGetSubjectsOfClassNotFound() *GetSubjectsOfClassNotFound {
	return &GetSubjectsOfClassNotFound{}
}

/*
GetSubjectsOfClassNotFound describes a response with status code 404, with default header values.

Class not found
*/
type GetSubjectsOfClassNotFound struct {
}

// IsSuccess returns true when this get subjects of class not found response has a 2xx status code
func (o *GetSubjectsOfClassNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subjects of class not found response has a 3xx status code
func (o *GetSubjectsOfClassNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subjects of class not found response has a 4xx status code
func (o *GetSubjectsOfClassNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subjects of class not found response has a 5xx status code
func (o *GetSubjectsOfClassNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get subjects of class not found response a status code equal to that given
func (o *GetSubjectsOfClassNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get subjects of class not found response
func (o *GetSubjectsOfClassNotFound) Code() int {
	return 404
}

func (o *GetSubjectsOfClassNotFound) Error() string {
	return fmt.Sprintf("[GET /class/{ID}/subjects][%d] getSubjectsOfClassNotFound ", 404)
}

func (o *GetSubjectsOfClassNotFound) String() string {
	return fmt.Sprintf("[GET /class/{ID}/subjects][%d] getSubjectsOfClassNotFound ", 404)
}

func (o *GetSubjectsOfClassNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubjectsOfClassInternalServerError creates a GetSubjectsOfClassInternalServerError with default headers values
func NewGetSubjectsOfClassInternalServerError() *GetSubjectsOfClassInternalServerError {
	return &GetSubjectsOfClassInternalServerError{}
}

/*
GetSubjectsOfClassInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSubjectsOfClassInternalServerError struct {
}

// IsSuccess returns true when this get subjects of class internal server error response has a 2xx status code
func (o *GetSubjectsOfClassInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subjects of class internal server error response has a 3xx status code
func (o *GetSubjectsOfClassInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subjects of class internal server error response has a 4xx status code
func (o *GetSubjectsOfClassInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subjects of class internal server error response has a 5xx status code
func (o *GetSubjectsOfClassInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get subjects of class internal server error response a status code equal to that given
func (o *GetSubjectsOfClassInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get subjects of class internal server error response
func (o *GetSubjectsOfClassInternalServerError) Code() int {
	return 500
}

func (o *GetSubjectsOfClassInternalServerError) Error() string {
	return fmt.Sprintf("[GET /class/{ID}/subjects][%d] getSubjectsOfClassInternalServerError ", 500)
}

func (o *GetSubjectsOfClassInternalServerError) String() string {
	return fmt.Sprintf("[GET /class/{ID}/subjects][%d] getSubjectsOfClassInternalServerError ", 500)
}

func (o *GetSubjectsOfClassInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
