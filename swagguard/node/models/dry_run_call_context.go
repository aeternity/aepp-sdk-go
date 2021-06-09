// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DryRunCallContext dry run call context
//
// swagger:model DryRunCallContext
type DryRunCallContext struct {

	// This call will have effects on the next call in this dry-run (or not)
	Stateful bool `json:"stateful,omitempty"`

	// tx hash
	TxHash string `json:"tx_hash,omitempty"`
}

// Validate validates this dry run call context
func (m *DryRunCallContext) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dry run call context based on context it is used
func (m *DryRunCallContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DryRunCallContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DryRunCallContext) UnmarshalBinary(b []byte) error {
	var res DryRunCallContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
