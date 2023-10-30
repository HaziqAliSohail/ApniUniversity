// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteClassNoContentCode is the HTTP code returned for type DeleteClassNoContent
const DeleteClassNoContentCode int = 204

/*
DeleteClassNoContent Class Deleted Successfully!

swagger:response deleteClassNoContent
*/
type DeleteClassNoContent struct {
}

// NewDeleteClassNoContent creates DeleteClassNoContent with default headers values
func NewDeleteClassNoContent() *DeleteClassNoContent {

	return &DeleteClassNoContent{}
}

// WriteResponse to the client
func (o *DeleteClassNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteClassNotFoundCode is the HTTP code returned for type DeleteClassNotFound
const DeleteClassNotFoundCode int = 404

/*
DeleteClassNotFound Class not found

swagger:response deleteClassNotFound
*/
type DeleteClassNotFound struct {
}

// NewDeleteClassNotFound creates DeleteClassNotFound with default headers values
func NewDeleteClassNotFound() *DeleteClassNotFound {

	return &DeleteClassNotFound{}
}

// WriteResponse to the client
func (o *DeleteClassNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteClassInternalServerErrorCode is the HTTP code returned for type DeleteClassInternalServerError
const DeleteClassInternalServerErrorCode int = 500

/*
DeleteClassInternalServerError Internal server error

swagger:response deleteClassInternalServerError
*/
type DeleteClassInternalServerError struct {
}

// NewDeleteClassInternalServerError creates DeleteClassInternalServerError with default headers values
func NewDeleteClassInternalServerError() *DeleteClassInternalServerError {

	return &DeleteClassInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteClassInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
