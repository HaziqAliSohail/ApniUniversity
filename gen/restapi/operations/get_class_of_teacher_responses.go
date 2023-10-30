// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ApniUniversity/gen/models"
)

// GetClassOfTeacherOKCode is the HTTP code returned for type GetClassOfTeacherOK
const GetClassOfTeacherOKCode int = 200

/*
GetClassOfTeacherOK Teacher's Class response

swagger:response getClassOfTeacherOK
*/
type GetClassOfTeacherOK struct {

	/*
	  In: Body
	*/
	Payload *models.Class `json:"body,omitempty"`
}

// NewGetClassOfTeacherOK creates GetClassOfTeacherOK with default headers values
func NewGetClassOfTeacherOK() *GetClassOfTeacherOK {

	return &GetClassOfTeacherOK{}
}

// WithPayload adds the payload to the get class of teacher o k response
func (o *GetClassOfTeacherOK) WithPayload(payload *models.Class) *GetClassOfTeacherOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get class of teacher o k response
func (o *GetClassOfTeacherOK) SetPayload(payload *models.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClassOfTeacherOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetClassOfTeacherNotFoundCode is the HTTP code returned for type GetClassOfTeacherNotFound
const GetClassOfTeacherNotFoundCode int = 404

/*
GetClassOfTeacherNotFound Class not found

swagger:response getClassOfTeacherNotFound
*/
type GetClassOfTeacherNotFound struct {
}

// NewGetClassOfTeacherNotFound creates GetClassOfTeacherNotFound with default headers values
func NewGetClassOfTeacherNotFound() *GetClassOfTeacherNotFound {

	return &GetClassOfTeacherNotFound{}
}

// WriteResponse to the client
func (o *GetClassOfTeacherNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetClassOfTeacherInternalServerErrorCode is the HTTP code returned for type GetClassOfTeacherInternalServerError
const GetClassOfTeacherInternalServerErrorCode int = 500

/*
GetClassOfTeacherInternalServerError Internal server error

swagger:response getClassOfTeacherInternalServerError
*/
type GetClassOfTeacherInternalServerError struct {
}

// NewGetClassOfTeacherInternalServerError creates GetClassOfTeacherInternalServerError with default headers values
func NewGetClassOfTeacherInternalServerError() *GetClassOfTeacherInternalServerError {

	return &GetClassOfTeacherInternalServerError{}
}

// WriteResponse to the client
func (o *GetClassOfTeacherInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
