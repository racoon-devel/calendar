// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/racoon-devel/calendar/internal/server/models"
)

// LoginUserOKCode is the HTTP code returned for type LoginUserOK
const LoginUserOKCode int = 200

/*
LoginUserOK Вход выполнен успешно

swagger:response loginUserOK
*/
type LoginUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {

	return &LoginUserOK{}
}

// WithPayload adds the payload to the login user o k response
func (o *LoginUserOK) WithPayload(payload *models.LoginResponse) *LoginUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user o k response
func (o *LoginUserOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserForbiddenCode is the HTTP code returned for type LoginUserForbidden
const LoginUserForbiddenCode int = 403

/*
LoginUserForbidden Доступ запрещен

swagger:response loginUserForbidden
*/
type LoginUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.LoginError `json:"body,omitempty"`
}

// NewLoginUserForbidden creates LoginUserForbidden with default headers values
func NewLoginUserForbidden() *LoginUserForbidden {

	return &LoginUserForbidden{}
}

// WithPayload adds the payload to the login user forbidden response
func (o *LoginUserForbidden) WithPayload(payload *models.LoginError) *LoginUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user forbidden response
func (o *LoginUserForbidden) SetPayload(payload *models.LoginError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserInternalServerErrorCode is the HTTP code returned for type LoginUserInternalServerError
const LoginUserInternalServerErrorCode int = 500

/*
LoginUserInternalServerError Ошибка на стороне сервера

swagger:response loginUserInternalServerError
*/
type LoginUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LoginError `json:"body,omitempty"`
}

// NewLoginUserInternalServerError creates LoginUserInternalServerError with default headers values
func NewLoginUserInternalServerError() *LoginUserInternalServerError {

	return &LoginUserInternalServerError{}
}

// WithPayload adds the payload to the login user internal server error response
func (o *LoginUserInternalServerError) WithPayload(payload *models.LoginError) *LoginUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user internal server error response
func (o *LoginUserInternalServerError) SetPayload(payload *models.LoginError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
