// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ApniUniversity/gen/models"
)

// GetTeachersOfSubjectOKCode is the HTTP code returned for type GetTeachersOfSubjectOK
const GetTeachersOfSubjectOKCode int = 200

/*
GetTeachersOfSubjectOK Subject's Teachers response

swagger:response getTeachersOfSubjectOK
*/
type GetTeachersOfSubjectOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Teacher `json:"body,omitempty"`
}

// NewGetTeachersOfSubjectOK creates GetTeachersOfSubjectOK with default headers values
func NewGetTeachersOfSubjectOK() *GetTeachersOfSubjectOK {

	return &GetTeachersOfSubjectOK{}
}

// WithPayload adds the payload to the get teachers of subject o k response
func (o *GetTeachersOfSubjectOK) WithPayload(payload []*models.Teacher) *GetTeachersOfSubjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teachers of subject o k response
func (o *GetTeachersOfSubjectOK) SetPayload(payload []*models.Teacher) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeachersOfSubjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Teacher, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTeachersOfSubjectNotFoundCode is the HTTP code returned for type GetTeachersOfSubjectNotFound
const GetTeachersOfSubjectNotFoundCode int = 404

/*
GetTeachersOfSubjectNotFound Subject not found

swagger:response getTeachersOfSubjectNotFound
*/
type GetTeachersOfSubjectNotFound struct {
}

// NewGetTeachersOfSubjectNotFound creates GetTeachersOfSubjectNotFound with default headers values
func NewGetTeachersOfSubjectNotFound() *GetTeachersOfSubjectNotFound {

	return &GetTeachersOfSubjectNotFound{}
}

// WriteResponse to the client
func (o *GetTeachersOfSubjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTeachersOfSubjectInternalServerErrorCode is the HTTP code returned for type GetTeachersOfSubjectInternalServerError
const GetTeachersOfSubjectInternalServerErrorCode int = 500

/*
GetTeachersOfSubjectInternalServerError Internal server error

swagger:response getTeachersOfSubjectInternalServerError
*/
type GetTeachersOfSubjectInternalServerError struct {
}

// NewGetTeachersOfSubjectInternalServerError creates GetTeachersOfSubjectInternalServerError with default headers values
func NewGetTeachersOfSubjectInternalServerError() *GetTeachersOfSubjectInternalServerError {

	return &GetTeachersOfSubjectInternalServerError{}
}

// WriteResponse to the client
func (o *GetTeachersOfSubjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}