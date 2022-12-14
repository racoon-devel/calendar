// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateUserError create user error
//
// swagger:model CreateUserError
type CreateUserError struct {

	// Код ошибки
	// 409 - пользователь уже существует
	// 500 - ошибка на стороне сервера
	//
	// Required: true
	// Enum: [409 500]
	Code interface{} `json:"code"`

	// Сообщение
	Message *string `json:"message,omitempty"`
}

// Validate validates this create user error
func (m *CreateUserError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createUserErrorTypeCodePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[409,500]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createUserErrorTypeCodePropEnum = append(createUserErrorTypeCodePropEnum, v)
	}
}

// prop value enum
func (m *CreateUserError) validateCodeEnum(path, location string, value interface{}) error {
	if err := validate.EnumCase(path, location, value, createUserErrorTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateUserError) validateCode(formats strfmt.Registry) error {

	if m.Code == nil {
		return errors.Required("code", "body", nil)
	}

	return nil
}

// ContextValidate validates this create user error based on context it is used
func (m *CreateUserError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUserError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUserError) UnmarshalBinary(b []byte) error {
	var res CreateUserError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
