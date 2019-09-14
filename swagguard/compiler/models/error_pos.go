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

// ErrorPos error pos
// swagger:model ErrorPos
type ErrorPos struct {

	// col
	// Required: true
	Col *int64 `json:"col"`

	// file
	File string `json:"file,omitempty"`

	// line
	// Required: true
	Line *int64 `json:"line"`
}

// Validate validates this error pos
func (m *ErrorPos) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorPos) validateCol(formats strfmt.Registry) error {

	if err := validate.Required("col", "body", m.Col); err != nil {
		return err
	}

	return nil
}

func (m *ErrorPos) validateLine(formats strfmt.Registry) error {

	if err := validate.Required("line", "body", m.Line); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorPos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorPos) UnmarshalBinary(b []byte) error {
	var res ErrorPos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *ErrorPos) String() string {
	return fmt.Sprintf("%s:line %v, col %v", m.File, *m.Line, *m.Col)
}
