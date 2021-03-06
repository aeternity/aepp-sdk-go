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

// UInt64 u int64
//
// swagger:model UInt64
type UInt64 uint64

// Validate validates this u int64
func (m UInt64) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinimumUint("", "body", uint64(m), 0, false); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this u int64 based on context it is used
func (m UInt64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
