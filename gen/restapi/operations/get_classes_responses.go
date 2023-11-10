// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// GetClassesOKCode is the HTTP code returned for type GetClassesOK
const GetClassesOKCode int = 200

/*
GetClassesOK Classes' response

swagger:response getClassesOK
*/
type GetClassesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Class `json:"body,omitempty"`
}

// NewGetClassesOK creates GetClassesOK with default headers values
func NewGetClassesOK() *GetClassesOK {

	return &GetClassesOK{}
}

// WithPayload adds the payload to the get classes o k response
func (o *GetClassesOK) WithPayload(payload []*models.Class) *GetClassesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get classes o k response
func (o *GetClassesOK) SetPayload(payload []*models.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClassesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Class, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetClassesNotFoundCode is the HTTP code returned for type GetClassesNotFound
const GetClassesNotFoundCode int = 404

/*
GetClassesNotFound Class not found

swagger:response getClassesNotFound
*/
type GetClassesNotFound struct {
}

// NewGetClassesNotFound creates GetClassesNotFound with default headers values
func NewGetClassesNotFound() *GetClassesNotFound {

	return &GetClassesNotFound{}
}

// WriteResponse to the client
func (o *GetClassesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetClassesInternalServerErrorCode is the HTTP code returned for type GetClassesInternalServerError
const GetClassesInternalServerErrorCode int = 500

/*
GetClassesInternalServerError Internal server error

swagger:response getClassesInternalServerError
*/
type GetClassesInternalServerError struct {
}

// NewGetClassesInternalServerError creates GetClassesInternalServerError with default headers values
func NewGetClassesInternalServerError() *GetClassesInternalServerError {

	return &GetClassesInternalServerError{}
}

// WriteResponse to the client
func (o *GetClassesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
