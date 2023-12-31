// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AssignSubjectToStudentOKCode is the HTTP code returned for type AssignSubjectToStudentOK
const AssignSubjectToStudentOKCode int = 200

/*
AssignSubjectToStudentOK Subject assigned/de-assigned to Student

swagger:response assignSubjectToStudentOK
*/
type AssignSubjectToStudentOK struct {

	/*The ID of the Student!
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewAssignSubjectToStudentOK creates AssignSubjectToStudentOK with default headers values
func NewAssignSubjectToStudentOK() *AssignSubjectToStudentOK {

	return &AssignSubjectToStudentOK{}
}

// WithPayload adds the payload to the assign subject to student o k response
func (o *AssignSubjectToStudentOK) WithPayload(payload int64) *AssignSubjectToStudentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the assign subject to student o k response
func (o *AssignSubjectToStudentOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssignSubjectToStudentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AssignSubjectToStudentNotFoundCode is the HTTP code returned for type AssignSubjectToStudentNotFound
const AssignSubjectToStudentNotFoundCode int = 404

/*
AssignSubjectToStudentNotFound Student Not found!

swagger:response assignSubjectToStudentNotFound
*/
type AssignSubjectToStudentNotFound struct {
}

// NewAssignSubjectToStudentNotFound creates AssignSubjectToStudentNotFound with default headers values
func NewAssignSubjectToStudentNotFound() *AssignSubjectToStudentNotFound {

	return &AssignSubjectToStudentNotFound{}
}

// WriteResponse to the client
func (o *AssignSubjectToStudentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// AssignSubjectToStudentInternalServerErrorCode is the HTTP code returned for type AssignSubjectToStudentInternalServerError
const AssignSubjectToStudentInternalServerErrorCode int = 500

/*
AssignSubjectToStudentInternalServerError Internal Server Error

swagger:response assignSubjectToStudentInternalServerError
*/
type AssignSubjectToStudentInternalServerError struct {
}

// NewAssignSubjectToStudentInternalServerError creates AssignSubjectToStudentInternalServerError with default headers values
func NewAssignSubjectToStudentInternalServerError() *AssignSubjectToStudentInternalServerError {

	return &AssignSubjectToStudentInternalServerError{}
}

// WriteResponse to the client
func (o *AssignSubjectToStudentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
