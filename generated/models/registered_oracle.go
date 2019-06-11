// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// RegisteredOracle registered oracle
// swagger:model RegisteredOracle
type RegisteredOracle struct {

	// abi version
	// Required: true
	AbiVersion *uint16 `json:"abi_version"`

	// id
	// Required: true
	ID EncodedPubkey `json:"id"`

	// query fee
	// Required: true
	QueryFee utils.BigInt `json:"query_fee"`

	// query format
	// Required: true
	QueryFormat *string `json:"query_format"`

	// response format
	// Required: true
	ResponseFormat *string `json:"response_format"`

	// ttl
	// Required: true
	TTL *uint64 `json:"ttl"`
}

// Validate validates this registered oracle
func (m *RegisteredOracle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredOracle) validateAbiVersion(formats strfmt.Registry) error {

	if err := validate.Required("abi_version", "body", m.AbiVersion); err != nil {
		return err
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

	if err := m.QueryFee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("query_fee")
		}
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
