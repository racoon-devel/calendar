// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Password Пароль
//
// swagger:model Password
type Password string

// Validate validates this password
func (m Password) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 32); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this password based on context it is used
func (m Password) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}