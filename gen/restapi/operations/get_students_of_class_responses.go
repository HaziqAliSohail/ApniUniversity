// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// GetStudentsOfClassOKCode is the HTTP code returned for type GetStudentsOfClassOK
const GetStudentsOfClassOKCode int = 200

/*
GetStudentsOfClassOK Class' Students as response

swagger:response getStudentsOfClassOK
*/
type GetStudentsOfClassOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Student `json:"body,omitempty"`
}

// NewGetStudentsOfClassOK creates GetStudentsOfClassOK with default headers values
func NewGetStudentsOfClassOK() *GetStudentsOfClassOK {

	return &GetStudentsOfClassOK{}
}

// WithPayload adds the payload to the get students of class o k response
func (o *GetStudentsOfClassOK) WithPayload(payload []*models.Student) *GetStudentsOfClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get students of class o k response
func (o *GetStudentsOfClassOK) SetPayload(payload []*models.Student) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStudentsOfClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Student, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetStudentsOfClassNotFoundCode is the HTTP code returned for type GetStudentsOfClassNotFound
const GetStudentsOfClassNotFoundCode int = 404

/*
GetStudentsOfClassNotFound Class not found

swagger:response getStudentsOfClassNotFound
*/
type GetStudentsOfClassNotFound struct {
}

// NewGetStudentsOfClassNotFound creates GetStudentsOfClassNotFound with default headers values
func NewGetStudentsOfClassNotFound() *GetStudentsOfClassNotFound {

	return &GetStudentsOfClassNotFound{}
}

// WriteResponse to the client
func (o *GetStudentsOfClassNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetStudentsOfClassInternalServerErrorCode is the HTTP code returned for type GetStudentsOfClassInternalServerError
const GetStudentsOfClassInternalServerErrorCode int = 500

/*
GetStudentsOfClassInternalServerError Internal server error

swagger:response getStudentsOfClassInternalServerError
*/
type GetStudentsOfClassInternalServerError struct {
}

// NewGetStudentsOfClassInternalServerError creates GetStudentsOfClassInternalServerError with default headers values
func NewGetStudentsOfClassInternalServerError() *GetStudentsOfClassInternalServerError {

	return &GetStudentsOfClassInternalServerError{}
}

// WriteResponse to the client
func (o *GetStudentsOfClassInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
