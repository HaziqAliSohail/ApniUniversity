// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// GetClassesOfStudentOKCode is the HTTP code returned for type GetClassesOfStudentOK
const GetClassesOfStudentOKCode int = 200

/*
GetClassesOfStudentOK Student's Class as response

swagger:response getClassesOfStudentOK
*/
type GetClassesOfStudentOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Class `json:"body,omitempty"`
}

// NewGetClassesOfStudentOK creates GetClassesOfStudentOK with default headers values
func NewGetClassesOfStudentOK() *GetClassesOfStudentOK {

	return &GetClassesOfStudentOK{}
}

// WithPayload adds the payload to the get classes of student o k response
func (o *GetClassesOfStudentOK) WithPayload(payload []*models.Class) *GetClassesOfStudentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get classes of student o k response
func (o *GetClassesOfStudentOK) SetPayload(payload []*models.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClassesOfStudentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetClassesOfStudentNotFoundCode is the HTTP code returned for type GetClassesOfStudentNotFound
const GetClassesOfStudentNotFoundCode int = 404

/*
GetClassesOfStudentNotFound Student not found

swagger:response getClassesOfStudentNotFound
*/
type GetClassesOfStudentNotFound struct {
}

// NewGetClassesOfStudentNotFound creates GetClassesOfStudentNotFound with default headers values
func NewGetClassesOfStudentNotFound() *GetClassesOfStudentNotFound {

	return &GetClassesOfStudentNotFound{}
}

// WriteResponse to the client
func (o *GetClassesOfStudentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetClassesOfStudentInternalServerErrorCode is the HTTP code returned for type GetClassesOfStudentInternalServerError
const GetClassesOfStudentInternalServerErrorCode int = 500

/*
GetClassesOfStudentInternalServerError Internal server error

swagger:response getClassesOfStudentInternalServerError
*/
type GetClassesOfStudentInternalServerError struct {
}

// NewGetClassesOfStudentInternalServerError creates GetClassesOfStudentInternalServerError with default headers values
func NewGetClassesOfStudentInternalServerError() *GetClassesOfStudentInternalServerError {

	return &GetClassesOfStudentInternalServerError{}
}

// WriteResponse to the client
func (o *GetClassesOfStudentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
