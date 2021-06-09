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

// NameHash name hash
//
// swagger:model NameHash
type NameHash struct {

	// name id
	// Required: true
	NameID *string `json:"name_id"`
}

// Validate validates this name hash
func (m *NameHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNameID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameHash) validateNameID(formats strfmt.Registry) error {

	if err := validate.Required("name_id", "body", m.NameID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this name hash based on context it is used
func (m *NameHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NameHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameHash) UnmarshalBinary(b []byte) error {
	var res NameHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
