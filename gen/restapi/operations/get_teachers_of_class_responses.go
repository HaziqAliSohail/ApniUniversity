// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// GetTeachersOfClassOKCode is the HTTP code returned for type GetTeachersOfClassOK
const GetTeachersOfClassOKCode int = 200

/*
GetTeachersOfClassOK Class' Teachers as response

swagger:response getTeachersOfClassOK
*/
type GetTeachersOfClassOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Teacher `json:"body,omitempty"`
}

// NewGetTeachersOfClassOK creates GetTeachersOfClassOK with default headers values
func NewGetTeachersOfClassOK() *GetTeachersOfClassOK {

	return &GetTeachersOfClassOK{}
}

// WithPayload adds the payload to the get teachers of class o k response
func (o *GetTeachersOfClassOK) WithPayload(payload []*models.Teacher) *GetTeachersOfClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teachers of class o k response
func (o *GetTeachersOfClassOK) SetPayload(payload []*models.Teacher) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeachersOfClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTeachersOfClassNotFoundCode is the HTTP code returned for type GetTeachersOfClassNotFound
const GetTeachersOfClassNotFoundCode int = 404

/*
GetTeachersOfClassNotFound Class not found

swagger:response getTeachersOfClassNotFound
*/
type GetTeachersOfClassNotFound struct {
}

// NewGetTeachersOfClassNotFound creates GetTeachersOfClassNotFound with default headers values
func NewGetTeachersOfClassNotFound() *GetTeachersOfClassNotFound {

	return &GetTeachersOfClassNotFound{}
}

// WriteResponse to the client
func (o *GetTeachersOfClassNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTeachersOfClassInternalServerErrorCode is the HTTP code returned for type GetTeachersOfClassInternalServerError
const GetTeachersOfClassInternalServerErrorCode int = 500

/*
GetTeachersOfClassInternalServerError Internal server error

swagger:response getTeachersOfClassInternalServerError
*/
type GetTeachersOfClassInternalServerError struct {
}

// NewGetTeachersOfClassInternalServerError creates GetTeachersOfClassInternalServerError with default headers values
func NewGetTeachersOfClassInternalServerError() *GetTeachersOfClassInternalServerError {

	return &GetTeachersOfClassInternalServerError{}
}

// WriteResponse to the client
func (o *GetTeachersOfClassInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
