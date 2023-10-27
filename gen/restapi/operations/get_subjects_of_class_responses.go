// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ApniUniversity/gen/models"
)

// GetSubjectsOfClassOKCode is the HTTP code returned for type GetSubjectsOfClassOK
const GetSubjectsOfClassOKCode int = 200

/*
GetSubjectsOfClassOK Class' Subjects as response

swagger:response getSubjectsOfClassOK
*/
type GetSubjectsOfClassOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Subject `json:"body,omitempty"`
}

// NewGetSubjectsOfClassOK creates GetSubjectsOfClassOK with default headers values
func NewGetSubjectsOfClassOK() *GetSubjectsOfClassOK {

	return &GetSubjectsOfClassOK{}
}

// WithPayload adds the payload to the get subjects of class o k response
func (o *GetSubjectsOfClassOK) WithPayload(payload []*models.Subject) *GetSubjectsOfClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subjects of class o k response
func (o *GetSubjectsOfClassOK) SetPayload(payload []*models.Subject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubjectsOfClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Subject, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSubjectsOfClassNotFoundCode is the HTTP code returned for type GetSubjectsOfClassNotFound
const GetSubjectsOfClassNotFoundCode int = 404

/*
GetSubjectsOfClassNotFound Class not found

swagger:response getSubjectsOfClassNotFound
*/
type GetSubjectsOfClassNotFound struct {
}

// NewGetSubjectsOfClassNotFound creates GetSubjectsOfClassNotFound with default headers values
func NewGetSubjectsOfClassNotFound() *GetSubjectsOfClassNotFound {

	return &GetSubjectsOfClassNotFound{}
}

// WriteResponse to the client
func (o *GetSubjectsOfClassNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetSubjectsOfClassInternalServerErrorCode is the HTTP code returned for type GetSubjectsOfClassInternalServerError
const GetSubjectsOfClassInternalServerErrorCode int = 500

/*
GetSubjectsOfClassInternalServerError Internal server error

swagger:response getSubjectsOfClassInternalServerError
*/
type GetSubjectsOfClassInternalServerError struct {
}

// NewGetSubjectsOfClassInternalServerError creates GetSubjectsOfClassInternalServerError with default headers values
func NewGetSubjectsOfClassInternalServerError() *GetSubjectsOfClassInternalServerError {

	return &GetSubjectsOfClassInternalServerError{}
}

// WriteResponse to the client
func (o *GetSubjectsOfClassInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
