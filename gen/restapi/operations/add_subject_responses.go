// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// AddSubjectCreatedCode is the HTTP code returned for type AddSubjectCreated
const AddSubjectCreatedCode int = 201

/*
AddSubjectCreated Subject added Successfully

swagger:response addSubjectCreated
*/
type AddSubjectCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Subject `json:"body,omitempty"`
}

// NewAddSubjectCreated creates AddSubjectCreated with default headers values
func NewAddSubjectCreated() *AddSubjectCreated {

	return &AddSubjectCreated{}
}

// WithPayload adds the payload to the add subject created response
func (o *AddSubjectCreated) WithPayload(payload *models.Subject) *AddSubjectCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subject created response
func (o *AddSubjectCreated) SetPayload(payload *models.Subject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubjectCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSubjectConflictCode is the HTTP code returned for type AddSubjectConflict
const AddSubjectConflictCode int = 409

/*
AddSubjectConflict Subject already exists

swagger:response addSubjectConflict
*/
type AddSubjectConflict struct {
}

// NewAddSubjectConflict creates AddSubjectConflict with default headers values
func NewAddSubjectConflict() *AddSubjectConflict {

	return &AddSubjectConflict{}
}

// WriteResponse to the client
func (o *AddSubjectConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// AddSubjectInternalServerErrorCode is the HTTP code returned for type AddSubjectInternalServerError
const AddSubjectInternalServerErrorCode int = 500

/*
AddSubjectInternalServerError Internal server error

swagger:response addSubjectInternalServerError
*/
type AddSubjectInternalServerError struct {
}

// NewAddSubjectInternalServerError creates AddSubjectInternalServerError with default headers values
func NewAddSubjectInternalServerError() *AddSubjectInternalServerError {

	return &AddSubjectInternalServerError{}
}

// WriteResponse to the client
func (o *AddSubjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
