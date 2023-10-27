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

// AssignSubjectToStudentReader is a Reader for the AssignSubjectToStudent structure.
type AssignSubjectToStudentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignSubjectToStudentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignSubjectToStudentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAssignSubjectToStudentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignSubjectToStudentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /student/{ID}/subject] assignSubjectToStudent", response, response.Code())
	}
}

// NewAssignSubjectToStudentOK creates a AssignSubjectToStudentOK with default headers values
func NewAssignSubjectToStudentOK() *AssignSubjectToStudentOK {
	return &AssignSubjectToStudentOK{}
}

/*
AssignSubjectToStudentOK describes a response with status code 200, with default header values.

Subject assigned to Student
*/
type AssignSubjectToStudentOK struct {
	Payload int64
}

// IsSuccess returns true when this assign subject to student o k response has a 2xx status code
func (o *AssignSubjectToStudentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign subject to student o k response has a 3xx status code
func (o *AssignSubjectToStudentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign subject to student o k response has a 4xx status code
func (o *AssignSubjectToStudentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign subject to student o k response has a 5xx status code
func (o *AssignSubjectToStudentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign subject to student o k response a status code equal to that given
func (o *AssignSubjectToStudentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign subject to student o k response
func (o *AssignSubjectToStudentOK) Code() int {
	return 200
}

func (o *AssignSubjectToStudentOK) Error() string {
	return fmt.Sprintf("[PATCH /student/{ID}/subject][%d] assignSubjectToStudentOK  %+v", 200, o.Payload)
}

func (o *AssignSubjectToStudentOK) String() string {
	return fmt.Sprintf("[PATCH /student/{ID}/subject][%d] assignSubjectToStudentOK  %+v", 200, o.Payload)
}

func (o *AssignSubjectToStudentOK) GetPayload() int64 {
	return o.Payload
}

func (o *AssignSubjectToStudentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignSubjectToStudentNotFound creates a AssignSubjectToStudentNotFound with default headers values
func NewAssignSubjectToStudentNotFound() *AssignSubjectToStudentNotFound {
	return &AssignSubjectToStudentNotFound{}
}

/*
AssignSubjectToStudentNotFound describes a response with status code 404, with default header values.

Student Not found!
*/
type AssignSubjectToStudentNotFound struct {
}

// IsSuccess returns true when this assign subject to student not found response has a 2xx status code
func (o *AssignSubjectToStudentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign subject to student not found response has a 3xx status code
func (o *AssignSubjectToStudentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign subject to student not found response has a 4xx status code
func (o *AssignSubjectToStudentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign subject to student not found response has a 5xx status code
func (o *AssignSubjectToStudentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this assign subject to student not found response a status code equal to that given
func (o *AssignSubjectToStudentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the assign subject to student not found response
func (o *AssignSubjectToStudentNotFound) Code() int {
	return 404
}

func (o *AssignSubjectToStudentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /student/{ID}/subject][%d] assignSubjectToStudentNotFound ", 404)
}

func (o *AssignSubjectToStudentNotFound) String() string {
	return fmt.Sprintf("[PATCH /student/{ID}/subject][%d] assignSubjectToStudentNotFound ", 404)
}

func (o *AssignSubjectToStudentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSubjectToStudentInternalServerError creates a AssignSubjectToStudentInternalServerError with default headers values
func NewAssignSubjectToStudentInternalServerError() *AssignSubjectToStudentInternalServerError {
	return &AssignSubjectToStudentInternalServerError{}
}

/*
AssignSubjectToStudentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AssignSubjectToStudentInternalServerError struct {
}

// IsSuccess returns true when this assign subject to student internal server error response has a 2xx status code
func (o *AssignSubjectToStudentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign subject to student internal server error response has a 3xx status code
func (o *AssignSubjectToStudentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign subject to student internal server error response has a 4xx status code
func (o *AssignSubjectToStudentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign subject to student internal server error response has a 5xx status code
func (o *AssignSubjectToStudentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assign subject to student internal server error response a status code equal to that given
func (o *AssignSubjectToStudentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the assign subject to student internal server error response
func (o *AssignSubjectToStudentInternalServerError) Code() int {
	return 500
}

func (o *AssignSubjectToStudentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /student/{ID}/subject][%d] assignSubjectToStudentInternalServerError ", 500)
}

func (o *AssignSubjectToStudentInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /student/{ID}/subject][%d] assignSubjectToStudentInternalServerError ", 500)
}

func (o *AssignSubjectToStudentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
AssignSubjectToStudentBody assign subject to student body
swagger:model AssignSubjectToStudentBody
*/
type AssignSubjectToStudentBody struct {

	// assign
	Assign bool `json:"Assign,omitempty"`

	// subject ID
	SubjectID int64 `json:"SubjectID,omitempty"`
}

// Validate validates this assign subject to student body
func (o *AssignSubjectToStudentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign subject to student body based on context it is used
func (o *AssignSubjectToStudentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignSubjectToStudentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignSubjectToStudentBody) UnmarshalBinary(b []byte) error {
	var res AssignSubjectToStudentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}