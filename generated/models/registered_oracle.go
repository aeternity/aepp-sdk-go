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

// RegisteredOracle registered oracle
// swagger:model RegisteredOracle
type RegisteredOracle struct {

	// id
	// Required: true
	ID EncodedHash `json:"id"`

	// query fee
	// Required: true
	QueryFee *int64 `json:"query_fee"`

	// query format
	// Required: true
	QueryFormat *string `json:"query_format"`

	// response format
	// Required: true
	ResponseFormat *string `json:"response_format"`

	// ttl
	// Required: true
	TTL *uint64 `json:"ttl"`

	// vm version
	// Required: true
	// Minimum: 0
	VMVersion *uint64 `json:"vm_version"`
}

// Validate validates this registered oracle
func (m *RegisteredOracle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredOracle) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *RegisteredOracle) validateQueryFee(formats strfmt.Registry) error {

	if err := validate.Required("query_fee", "body", m.QueryFee); err != nil {
		return err
	}

	return nil
}

func (m *RegisteredOracle) validateQueryFormat(formats strfmt.Registry) error {

	if err := validate.Required("query_format", "body", m.QueryFormat); err != nil {
		return err
	}

	return nil
}

func (m *RegisteredOracle) validateResponseFormat(formats strfmt.Registry) error {

	if err := validate.Required("response_format", "body", m.ResponseFormat); err != nil {
		return err
	}

	return nil
}

func (m *RegisteredOracle) validateTTL(formats strfmt.Registry) error {

	if err := validate.Required("ttl", "body", m.TTL); err != nil {
		return err
	}

	return nil
}

func (m *RegisteredOracle) validateVMVersion(formats strfmt.Registry) error {

	if err := validate.Required("vm_version", "body", m.VMVersion); err != nil {
		return err
	}

	if err := validate.MinimumInt("vm_version", "body", int64(*m.VMVersion), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisteredOracle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisteredOracle) UnmarshalBinary(b []byte) error {
	var res RegisteredOracle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
