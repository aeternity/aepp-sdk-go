// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CompilerError compiler error
// swagger:model CompilerError
type CompilerError struct {

	// context
	Context string `json:"context,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`

	// pos
	// Required: true
	Pos *ErrorPos `json:"pos"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this compiler error
func (m *CompilerError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompilerError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *CompilerError) validatePos(formats strfmt.Registry) error {

	if err := validate.Required("pos", "body", m.Pos); err != nil {
		return err
	}

	if m.Pos != nil {
		if err := m.Pos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pos")
			}
			return err
		}
	}

	return nil
}

func (m *CompilerError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompilerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompilerError) UnmarshalBinary(b []byte) error {
	var res CompilerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *CompilerError) String() string {
	return fmt.Sprintf("%s, %s, %s", *m.Type, *m.Message, m.Context)
}
