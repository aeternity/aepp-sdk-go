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

// NameRevokeTx name revoke tx
// swagger:model NameRevokeTx
type NameRevokeTx struct {

	// account id
	// Required: true
	AccountID EncodedHash `json:"account_id"`

	// fee
	// Required: true
	Fee *int64 `json:"fee"`

	// name id
	// Required: true
	NameID EncodedHash `json:"name_id"`

	// nonce
	Nonce int64 `json:"nonce,omitempty"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this name revoke tx
func (m *NameRevokeTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameRevokeTx) validateAccountID(formats strfmt.Registry) error {

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_id")
		}
		return err
	}

	return nil
}

func (m *NameRevokeTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	return nil
}

func (m *NameRevokeTx) validateNameID(formats strfmt.Registry) error {

	if err := m.NameID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameRevokeTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameRevokeTx) UnmarshalBinary(b []byte) error {
	var res NameRevokeTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
