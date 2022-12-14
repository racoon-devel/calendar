// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetInvitesResponse get invites response
//
// swagger:model GetInvitesResponse
type GetInvitesResponse struct {

	// invites
	// Required: true
	Invites []int64 `json:"invites"`
}

// Validate validates this get invites response
func (m *GetInvitesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInvitesResponse) validateInvites(formats strfmt.Registry) error {

	if err := validate.Required("invites", "body", m.Invites); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get invites response based on context it is used
func (m *GetInvitesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetInvitesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInvitesResponse) UnmarshalBinary(b []byte) error {
	var res GetInvitesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
