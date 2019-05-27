// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Pow pow
// swagger:model Pow
type Pow []int32

// Validate validates this pow
func (m Pow) Validate(formats strfmt.Registry) error {
	var res []error

	iPowSize := int64(len(m))

	if err := validate.MinItems("", "body", iPowSize, 42); err != nil {
		return err
	}

	if err := validate.MaxItems("", "body", iPowSize, 42); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
