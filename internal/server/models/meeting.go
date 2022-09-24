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

// Meeting meeting
//
// swagger:model Meeting
type Meeting struct {

	// Описание встречи
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// Длительность встречи (минуты)
	// Required: true
	// Maximum: 1440
	// Minimum: 5
	Duration *int64 `json:"duration"`

	// Идентификатор встречи
	ID int64 `json:"id,omitempty"`

	// Уведомить пользователя о встрече перед ней (единицы измерения - минуты!)
	// Minimum: 1
	Notify int64 `json:"notify,omitempty"`

	// Приватность деталей встречи
	Private bool `json:"private,omitempty"`

	// Повторение задачи (формат RRULE, RFC5545)
	Rrule string `json:"rrule,omitempty"`

	// Временной интервал начала встречи (формат - RFC3339)
	// Example: 1996-12-19T16:39:57Z
	// Required: true
	StartTime *string `json:"startTime"`

	// Заголовок встречи
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Title *string `json:"title"`

	// users
	Users []int64 `json:"users"`
}

// Validate validates this meeting
func (m *Meeting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Meeting) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 1000); err != nil {
		return err
	}

	return nil
}

func (m *Meeting) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	if err := validate.MinimumInt("duration", "body", *m.Duration, 5, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("duration", "body", *m.Duration, 1440, false); err != nil {
		return err
	}

	return nil
}

func (m *Meeting) validateNotify(formats strfmt.Registry) error {
	if swag.IsZero(m.Notify) { // not required
		return nil
	}

	if err := validate.MinimumInt("notify", "body", m.Notify, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Meeting) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *Meeting) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 64); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this meeting based on context it is used
func (m *Meeting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Meeting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Meeting) UnmarshalBinary(b []byte) error {
	var res Meeting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
