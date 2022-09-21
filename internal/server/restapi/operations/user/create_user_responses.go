// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/racoon-devel/calendar/internal/server/models"
)

// CreateUserCreatedCode is the HTTP code returned for type CreateUserCreated
const CreateUserCreatedCode int = 201

/*
CreateUserCreated OK

swagger:response createUserCreated
*/
type CreateUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreateUserResponse `json:"body,omitempty"`
}

// NewCreateUserCreated creates CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {

	return &CreateUserCreated{}
}

// WithPayload adds the payload to the create user created response
func (o *CreateUserCreated) WithPayload(payload *models.CreateUserResponse) *CreateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user created response
func (o *CreateUserCreated) SetPayload(payload *models.CreateUserResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserConflictCode is the HTTP code returned for type CreateUserConflict
const CreateUserConflictCode int = 409

/*
CreateUserConflict Не удалось выполнить запрос (пользователь уже существует)

swagger:response createUserConflict
*/
type CreateUserConflict struct {

	/*
	  In: Body
	*/
	Payload *models.CreateUserError `json:"body,omitempty"`
}

// NewCreateUserConflict creates CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {

	return &CreateUserConflict{}
}

// WithPayload adds the payload to the create user conflict response
func (o *CreateUserConflict) WithPayload(payload *models.CreateUserError) *CreateUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user conflict response
func (o *CreateUserConflict) SetPayload(payload *models.CreateUserError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserInternalServerErrorCode is the HTTP code returned for type CreateUserInternalServerError
const CreateUserInternalServerErrorCode int = 500

/*
CreateUserInternalServerError Ошибка на стороне сервера

swagger:response createUserInternalServerError
*/
type CreateUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.CreateUserError `json:"body,omitempty"`
}

// NewCreateUserInternalServerError creates CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {

	return &CreateUserInternalServerError{}
}

// WithPayload adds the payload to the create user internal server error response
func (o *CreateUserInternalServerError) WithPayload(payload *models.CreateUserError) *CreateUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user internal server error response
func (o *CreateUserInternalServerError) SetPayload(payload *models.CreateUserError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
