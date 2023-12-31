// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateTeacherNameOKCode is the HTTP code returned for type UpdateTeacherNameOK
const UpdateTeacherNameOKCode int = 200

/*
UpdateTeacherNameOK Teacher updated

swagger:response updateTeacherNameOK
*/
type UpdateTeacherNameOK struct {

	/*The ID of the Teacher!
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewUpdateTeacherNameOK creates UpdateTeacherNameOK with default headers values
func NewUpdateTeacherNameOK() *UpdateTeacherNameOK {

	return &UpdateTeacherNameOK{}
}

// WithPayload adds the payload to the update teacher name o k response
func (o *UpdateTeacherNameOK) WithPayload(payload int64) *UpdateTeacherNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update teacher name o k response
func (o *UpdateTeacherNameOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeacherNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateTeacherNameNotFoundCode is the HTTP code returned for type UpdateTeacherNameNotFound
const UpdateTeacherNameNotFoundCode int = 404

/*
UpdateTeacherNameNotFound Teacher not found!

swagger:response updateTeacherNameNotFound
*/
type UpdateTeacherNameNotFound struct {
}

// NewUpdateTeacherNameNotFound creates UpdateTeacherNameNotFound with default headers values
func NewUpdateTeacherNameNotFound() *UpdateTeacherNameNotFound {

	return &UpdateTeacherNameNotFound{}
}

// WriteResponse to the client
func (o *UpdateTeacherNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateTeacherNameInternalServerErrorCode is the HTTP code returned for type UpdateTeacherNameInternalServerError
const UpdateTeacherNameInternalServerErrorCode int = 500

/*
UpdateTeacherNameInternalServerError Internal Server Error

swagger:response updateTeacherNameInternalServerError
*/
type UpdateTeacherNameInternalServerError struct {
}

// NewUpdateTeacherNameInternalServerError creates UpdateTeacherNameInternalServerError with default headers values
func NewUpdateTeacherNameInternalServerError() *UpdateTeacherNameInternalServerError {

	return &UpdateTeacherNameInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateTeacherNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
