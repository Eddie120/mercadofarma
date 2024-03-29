// Code generated by go-swagger; DO NOT EDIT.

package shopper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mercadofarma/services/models"
)

// SignUpShopperNoContentCode is the HTTP code returned for type SignUpShopperNoContent
const SignUpShopperNoContentCode int = 204

/*
SignUpShopperNoContent No content

swagger:response signUpShopperNoContent
*/
type SignUpShopperNoContent struct {
}

// NewSignUpShopperNoContent creates SignUpShopperNoContent with default headers values
func NewSignUpShopperNoContent() *SignUpShopperNoContent {

	return &SignUpShopperNoContent{}
}

// WriteResponse to the client
func (o *SignUpShopperNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// SignUpShopperBadRequestCode is the HTTP code returned for type SignUpShopperBadRequest
const SignUpShopperBadRequestCode int = 400

/*
SignUpShopperBadRequest Invalid request

swagger:response signUpShopperBadRequest
*/
type SignUpShopperBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSignUpShopperBadRequest creates SignUpShopperBadRequest with default headers values
func NewSignUpShopperBadRequest() *SignUpShopperBadRequest {

	return &SignUpShopperBadRequest{}
}

// WithPayload adds the payload to the sign up shopper bad request response
func (o *SignUpShopperBadRequest) WithPayload(payload *models.Error) *SignUpShopperBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign up shopper bad request response
func (o *SignUpShopperBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignUpShopperBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SignUpShopperInternalServerErrorCode is the HTTP code returned for type SignUpShopperInternalServerError
const SignUpShopperInternalServerErrorCode int = 500

/*
SignUpShopperInternalServerError Internal server error

swagger:response signUpShopperInternalServerError
*/
type SignUpShopperInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSignUpShopperInternalServerError creates SignUpShopperInternalServerError with default headers values
func NewSignUpShopperInternalServerError() *SignUpShopperInternalServerError {

	return &SignUpShopperInternalServerError{}
}

// WithPayload adds the payload to the sign up shopper internal server error response
func (o *SignUpShopperInternalServerError) WithPayload(payload *models.Error) *SignUpShopperInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign up shopper internal server error response
func (o *SignUpShopperInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignUpShopperInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
