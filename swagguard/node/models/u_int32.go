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

// UInt32 u int32
//
// swagger:model UInt32
type UInt32 uint64

// Validate validates this u int32
func (m UInt32) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinimumUint("", "body", uint64(m), 0, false); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this u int32 based on context it is used
func (m UInt32) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}