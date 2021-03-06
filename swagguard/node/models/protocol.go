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

// Protocol protocol
//
// swagger:model Protocol
type Protocol struct {

	// Height at which protocol becomes active
	// Required: true
	EffectiveAtHeight *uint64 `json:"effective_at_height"`

	// Protocol version (can include protocol activated by miner signalling)
	// Required: true
	Version *uint32 `json:"version"`
}

// Validate validates this protocol
func (m *Protocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveAtHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Protocol) validateEffectiveAtHeight(formats strfmt.Registry) error {

	if err := validate.Required("effective_at_height", "body", m.EffectiveAtHeight); err != nil {
		return err
	}

	return nil
}

func (m *Protocol) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this protocol based on context it is used
func (m *Protocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Protocol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Protocol) UnmarshalBinary(b []byte) error {
	var res Protocol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
