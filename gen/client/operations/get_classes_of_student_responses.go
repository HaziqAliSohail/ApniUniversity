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

// GetClassesOfStudentReader is a Reader for the GetClassesOfStudent structure.
type GetClassesOfStudentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClassesOfStudentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClassesOfStudentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetClassesOfStudentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClassesOfStudentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /student/{ID}/classes] getClassesOfStudent", response, response.Code())
	}
}

// NewGetClassesOfStudentOK creates a GetClassesOfStudentOK with default headers values
func NewGetClassesOfStudentOK() *GetClassesOfStudentOK {
	return &GetClassesOfStudentOK{}
}

/*
GetClassesOfStudentOK describes a response with status code 200, with default header values.

Student's Class as response
*/
type GetClassesOfStudentOK struct {
	Payload []*models.Class
}

// IsSuccess returns true when this get classes of student o k response has a 2xx status code
func (o *GetClassesOfStudentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get classes of student o k response has a 3xx status code
func (o *GetClassesOfStudentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get classes of student o k response has a 4xx status code
func (o *GetClassesOfStudentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get classes of student o k response has a 5xx status code
func (o *GetClassesOfStudentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get classes of student o k response a status code equal to that given
func (o *GetClassesOfStudentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get classes of student o k response
func (o *GetClassesOfStudentOK) Code() int {
	return 200
}

func (o *GetClassesOfStudentOK) Error() string {
	return fmt.Sprintf("[GET /student/{ID}/classes][%d] getClassesOfStudentOK  %+v", 200, o.Payload)
}

func (o *GetClassesOfStudentOK) String() string {
	return fmt.Sprintf("[GET /student/{ID}/classes][%d] getClassesOfStudentOK  %+v", 200, o.Payload)
}

func (o *GetClassesOfStudentOK) GetPayload() []*models.Class {
	return o.Payload
}

func (o *GetClassesOfStudentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClassesOfStudentNotFound creates a GetClassesOfStudentNotFound with default headers values
func NewGetClassesOfStudentNotFound() *GetClassesOfStudentNotFound {
	return &GetClassesOfStudentNotFound{}
}

/*
GetClassesOfStudentNotFound describes a response with status code 404, with default header values.

Student not found
*/
type GetClassesOfStudentNotFound struct {
}

// IsSuccess returns true when this get classes of student not found response has a 2xx status code
func (o *GetClassesOfStudentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get classes of student not found response has a 3xx status code
func (o *GetClassesOfStudentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get classes of student not found response has a 4xx status code
func (o *GetClassesOfStudentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get classes of student not found response has a 5xx status code
func (o *GetClassesOfStudentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get classes of student not found response a status code equal to that given
func (o *GetClassesOfStudentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get classes of student not found response
func (o *GetClassesOfStudentNotFound) Code() int {
	return 404
}

func (o *GetClassesOfStudentNotFound) Error() string {
	return fmt.Sprintf("[GET /student/{ID}/classes][%d] getClassesOfStudentNotFound ", 404)
}

func (o *GetClassesOfStudentNotFound) String() string {
	return fmt.Sprintf("[GET /student/{ID}/classes][%d] getClassesOfStudentNotFound ", 404)
}

func (o *GetClassesOfStudentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClassesOfStudentInternalServerError creates a GetClassesOfStudentInternalServerError with default headers values
func NewGetClassesOfStudentInternalServerError() *GetClassesOfStudentInternalServerError {
	return &GetClassesOfStudentInternalServerError{}
}

/*
GetClassesOfStudentInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetClassesOfStudentInternalServerError struct {
}

// IsSuccess returns true when this get classes of student internal server error response has a 2xx status code
func (o *GetClassesOfStudentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get classes of student internal server error response has a 3xx status code
func (o *GetClassesOfStudentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get classes of student internal server error response has a 4xx status code
func (o *GetClassesOfStudentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get classes of student internal server error response has a 5xx status code
func (o *GetClassesOfStudentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get classes of student internal server error response a status code equal to that given
func (o *GetClassesOfStudentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get classes of student internal server error response
func (o *GetClassesOfStudentInternalServerError) Code() int {
	return 500
}

func (o *GetClassesOfStudentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /student/{ID}/classes][%d] getClassesOfStudentInternalServerError ", 500)
}

func (o *GetClassesOfStudentInternalServerError) String() string {
	return fmt.Sprintf("[GET /student/{ID}/classes][%d] getClassesOfStudentInternalServerError ", 500)
}

func (o *GetClassesOfStudentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
