// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateGPAOKCode is the HTTP code returned for type UpdateGPAOK
const UpdateGPAOKCode int = 200

/*
UpdateGPAOK GPA assigned to Student

swagger:response updateGPAOK
*/
type UpdateGPAOK struct {

	/*The ID of the Student!
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewUpdateGPAOK creates UpdateGPAOK with default headers values
func NewUpdateGPAOK() *UpdateGPAOK {

	return &UpdateGPAOK{}
}

// WithPayload adds the payload to the update g p a o k response
func (o *UpdateGPAOK) WithPayload(payload int64) *UpdateGPAOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update g p a o k response
func (o *UpdateGPAOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGPAOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateGPANotFoundCode is the HTTP code returned for type UpdateGPANotFound
const UpdateGPANotFoundCode int = 404

/*
UpdateGPANotFound Student Not found!

swagger:response updateGPANotFound
*/
type UpdateGPANotFound struct {
}

// NewUpdateGPANotFound creates UpdateGPANotFound with default headers values
func NewUpdateGPANotFound() *UpdateGPANotFound {

	return &UpdateGPANotFound{}
}

// WriteResponse to the client
func (o *UpdateGPANotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateGPAInternalServerErrorCode is the HTTP code returned for type UpdateGPAInternalServerError
const UpdateGPAInternalServerErrorCode int = 500

/*
UpdateGPAInternalServerError Internal Server Error

swagger:response updateGPAInternalServerError
*/
type UpdateGPAInternalServerError struct {
}

// NewUpdateGPAInternalServerError creates UpdateGPAInternalServerError with default headers values
func NewUpdateGPAInternalServerError() *UpdateGPAInternalServerError {

	return &UpdateGPAInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateGPAInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
