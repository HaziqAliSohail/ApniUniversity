// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateStudentNameReader is a Reader for the UpdateStudentName structure.
type UpdateStudentNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStudentNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStudentNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateStudentNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateStudentNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /student/{ID}/name] updateStudentName", response, response.Code())
	}
}

// NewUpdateStudentNameOK creates a UpdateStudentNameOK with default headers values
func NewUpdateStudentNameOK() *UpdateStudentNameOK {
	return &UpdateStudentNameOK{}
}

/*
UpdateStudentNameOK describes a response with status code 200, with default header values.

Student Name updated Successfully!
*/
type UpdateStudentNameOK struct {
	Payload int64
}

// IsSuccess returns true when this update student name o k response has a 2xx status code
func (o *UpdateStudentNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update student name o k response has a 3xx status code
func (o *UpdateStudentNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update student name o k response has a 4xx status code
func (o *UpdateStudentNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update student name o k response has a 5xx status code
func (o *UpdateStudentNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update student name o k response a status code equal to that given
func (o *UpdateStudentNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update student name o k response
func (o *UpdateStudentNameOK) Code() int {
	return 200
}

func (o *UpdateStudentNameOK) Error() string {
	return fmt.Sprintf("[PATCH /student/{ID}/name][%d] updateStudentNameOK  %+v", 200, o.Payload)
}

func (o *UpdateStudentNameOK) String() string {
	return fmt.Sprintf("[PATCH /student/{ID}/name][%d] updateStudentNameOK  %+v", 200, o.Payload)
}

func (o *UpdateStudentNameOK) GetPayload() int64 {
	return o.Payload
}

func (o *UpdateStudentNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStudentNameNotFound creates a UpdateStudentNameNotFound with default headers values
func NewUpdateStudentNameNotFound() *UpdateStudentNameNotFound {
	return &UpdateStudentNameNotFound{}
}

/*
UpdateStudentNameNotFound describes a response with status code 404, with default header values.

Student not found!
*/
type UpdateStudentNameNotFound struct {
}

// IsSuccess returns true when this update student name not found response has a 2xx status code
func (o *UpdateStudentNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update student name not found response has a 3xx status code
func (o *UpdateStudentNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update student name not found response has a 4xx status code
func (o *UpdateStudentNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update student name not found response has a 5xx status code
func (o *UpdateStudentNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update student name not found response a status code equal to that given
func (o *UpdateStudentNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update student name not found response
func (o *UpdateStudentNameNotFound) Code() int {
	return 404
}

func (o *UpdateStudentNameNotFound) Error() string {
	return fmt.Sprintf("[PATCH /student/{ID}/name][%d] updateStudentNameNotFound ", 404)
}

func (o *UpdateStudentNameNotFound) String() string {
	return fmt.Sprintf("[PATCH /student/{ID}/name][%d] updateStudentNameNotFound ", 404)
}

func (o *UpdateStudentNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStudentNameInternalServerError creates a UpdateStudentNameInternalServerError with default headers values
func NewUpdateStudentNameInternalServerError() *UpdateStudentNameInternalServerError {
	return &UpdateStudentNameInternalServerError{}
}

/*
UpdateStudentNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateStudentNameInternalServerError struct {
}

// IsSuccess returns true when this update student name internal server error response has a 2xx status code
func (o *UpdateStudentNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update student name internal server error response has a 3xx status code
func (o *UpdateStudentNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update student name internal server error response has a 4xx status code
func (o *UpdateStudentNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update student name internal server error response has a 5xx status code
func (o *UpdateStudentNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update student name internal server error response a status code equal to that given
func (o *UpdateStudentNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update student name internal server error response
func (o *UpdateStudentNameInternalServerError) Code() int {
	return 500
}

func (o *UpdateStudentNameInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /student/{ID}/name][%d] updateStudentNameInternalServerError ", 500)
}

func (o *UpdateStudentNameInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /student/{ID}/name][%d] updateStudentNameInternalServerError ", 500)
}

func (o *UpdateStudentNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
UpdateStudentNameBody update student name body
swagger:model UpdateStudentNameBody
*/
type UpdateStudentNameBody struct {

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this update student name body
func (o *UpdateStudentNameBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update student name body based on context it is used
func (o *UpdateStudentNameBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateStudentNameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateStudentNameBody) UnmarshalBinary(b []byte) error {
	var res UpdateStudentNameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}