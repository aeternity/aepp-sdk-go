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

// OracleQuery oracle query
// swagger:model OracleQuery
type OracleQuery struct {

	// fee
	Fee int64 `json:"fee,omitempty"`

	// id
	ID EncodedHash `json:"id,omitempty"`

	// oracle id
	OracleID EncodedHash `json:"oracle_id,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// response
	Response string `json:"response,omitempty"`

	// response ttl
	ResponseTTL *TTL `json:"response_ttl,omitempty"`

	// sender id
	SenderID EncodedHash `json:"sender_id,omitempty"`

	// sender nonce
	// Minimum: 0
	SenderNonce *int64 `json:"sender_nonce,omitempty"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this oracle query
func (m *OracleQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderNonce(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleQuery) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *OracleQuery) validateOracleID(formats strfmt.Registry) error {

	if swag.IsZero(m.OracleID) { // not required
		return nil
	}

	if err := m.OracleID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("oracle_id")
		}
		return err
	}

	return nil
}

func (m *OracleQuery) validateResponseTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseTTL) { // not required
		return nil
	}

	if m.ResponseTTL != nil {
		if err := m.ResponseTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *OracleQuery) validateSenderID(formats strfmt.Registry) error {

	if swag.IsZero(m.SenderID) { // not required
		return nil
	}

	if err := m.SenderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sender_id")
		}
		return err
	}

	return nil
}

func (m *OracleQuery) validateSenderNonce(formats strfmt.Registry) error {

	if swag.IsZero(m.SenderNonce) { // not required
		return nil
	}

	if err := validate.MinimumInt("sender_nonce", "body", int64(*m.SenderNonce), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleQuery) UnmarshalBinary(b []byte) error {
	var res OracleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
