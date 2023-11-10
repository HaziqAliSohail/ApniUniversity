// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// AddClassCreatedCode is the HTTP code returned for type AddClassCreated
const AddClassCreatedCode int = 201

/*
AddClassCreated Class added Successfully

swagger:response addClassCreated
*/
type AddClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Class `json:"body,omitempty"`
}

// NewAddClassCreated creates AddClassCreated with default headers values
func NewAddClassCreated() *AddClassCreated {

	return &AddClassCreated{}
}

// WithPayload adds the payload to the add class created response
func (o *AddClassCreated) WithPayload(payload *models.Class) *AddClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add class created response
func (o *AddClassCreated) SetPayload(payload *models.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddClassConflictCode is the HTTP code returned for type AddClassConflict
const AddClassConflictCode int = 409

/*
AddClassConflict Class already exists

swagger:response addClassConflict
*/
type AddClassConflict struct {
}

// NewAddClassConflict creates AddClassConflict with default headers values
func NewAddClassConflict() *AddClassConflict {

	return &AddClassConflict{}
}

// WriteResponse to the client
func (o *AddClassConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// AddClassInternalServerErrorCode is the HTTP code returned for type AddClassInternalServerError
const AddClassInternalServerErrorCode int = 500

/*
AddClassInternalServerError Internal server error

swagger:response addClassInternalServerError
*/
type AddClassInternalServerError struct {
}

// NewAddClassInternalServerError creates AddClassInternalServerError with default headers values
func NewAddClassInternalServerError() *AddClassInternalServerError {

	return &AddClassInternalServerError{}
}

// WriteResponse to the client
func (o *AddClassInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
