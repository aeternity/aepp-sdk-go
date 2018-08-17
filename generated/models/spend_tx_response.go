// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SpendTxResponse spend tx response
// swagger:model SpendTxResponse
type SpendTxResponse struct {

	// tx hash
	TxHash EncodedHash `json:"tx_hash,omitempty"`
}

// Validate validates this spend tx response
func (m *SpendTxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTxHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpendTxResponse) validateTxHash(formats strfmt.Registry) error {

	if swag.IsZero(m.TxHash) { // not required
		return nil
	}

	if err := m.TxHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tx_hash")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpendTxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpendTxResponse) UnmarshalBinary(b []byte) error {
	var res SpendTxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}