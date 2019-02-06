// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OracleExtendTx oracle extend tx
// swagger:model OracleExtendTx
type OracleExtendTx struct {

	// fee
	// Required: true
	Fee *uint64 `json:"fee"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// oracle id
	// Required: true
	OracleID EncodedHash `json:"oracle_id"`

	// oracle ttl
	// Required: true
	OracleTTL *RelativeTTL `json:"oracle_ttl"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this oracle extend tx
func (m *OracleExtendTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleExtendTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	return nil
}

func (m *OracleExtendTx) validateOracleID(formats strfmt.Registry) error {

	if err := m.OracleID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("oracle_id")
		}
		return err
	}

	return nil
}

func (m *OracleExtendTx) validateOracleTTL(formats strfmt.Registry) error {

	if err := validate.Required("oracle_ttl", "body", m.OracleTTL); err != nil {
		return err
	}

	if m.OracleTTL != nil {
		if err := m.OracleTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_ttl")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleExtendTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleExtendTx) UnmarshalBinary(b []byte) error {
	var res OracleExtendTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
