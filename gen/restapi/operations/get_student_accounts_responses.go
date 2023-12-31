// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/HaziqAliSohail/ApniUniversity/gen/models"
)

// GetStudentAccountsOKCode is the HTTP code returned for type GetStudentAccountsOK
const GetStudentAccountsOKCode int = 200

/*
GetStudentAccountsOK Student Accounts' response

swagger:response getStudentAccountsOK
*/
type GetStudentAccountsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Account `json:"body,omitempty"`
}

// NewGetStudentAccountsOK creates GetStudentAccountsOK with default headers values
func NewGetStudentAccountsOK() *GetStudentAccountsOK {

	return &GetStudentAccountsOK{}
}

// WithPayload adds the payload to the get student accounts o k response
func (o *GetStudentAccountsOK) WithPayload(payload []*models.Account) *GetStudentAccountsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get student accounts o k response
func (o *GetStudentAccountsOK) SetPayload(payload []*models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStudentAccountsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Account, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetStudentAccountsNotFoundCode is the HTTP code returned for type GetStudentAccountsNotFound
const GetStudentAccountsNotFoundCode int = 404

/*
GetStudentAccountsNotFound Student Account not found

swagger:response getStudentAccountsNotFound
*/
type GetStudentAccountsNotFound struct {
}

// NewGetStudentAccountsNotFound creates GetStudentAccountsNotFound with default headers values
func NewGetStudentAccountsNotFound() *GetStudentAccountsNotFound {

	return &GetStudentAccountsNotFound{}
}

// WriteResponse to the client
func (o *GetStudentAccountsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetStudentAccountsInternalServerErrorCode is the HTTP code returned for type GetStudentAccountsInternalServerError
const GetStudentAccountsInternalServerErrorCode int = 500

/*
GetStudentAccountsInternalServerError Internal server error

swagger:response getStudentAccountsInternalServerError
*/
type GetStudentAccountsInternalServerError struct {
}

// NewGetStudentAccountsInternalServerError creates GetStudentAccountsInternalServerError with default headers values
func NewGetStudentAccountsInternalServerError() *GetStudentAccountsInternalServerError {

	return &GetStudentAccountsInternalServerError{}
}

// WriteResponse to the client
func (o *GetStudentAccountsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
