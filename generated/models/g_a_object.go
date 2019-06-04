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

// GAObject g a object
// swagger:model GAObject
type GAObject struct {

	// caller id
	// Required: true
	CallerID EncodedPubkey `json:"caller_id"`

	// gas price
	// Required: true
	GasPrice utils.BigInt `json:"gas_price"`

	// gas used
	// Required: true
	GasUsed Uint64 `json:"gas_used"`

	// height
	// Required: true
	Height Uint64 `json:"height"`

	// inner object
	InnerObject *TxInfoObject `json:"inner_object,omitempty"`

	// The status of the call 'ok | error'.
	// Required: true
	ReturnType *string `json:"return_type"`

	// return value
	// Required: true
	ReturnValue EncodedByteArray `json:"return_value"`
}

// Validate validates this g a object
func (m *GAObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInnerObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GAObject) validateCallerID(formats strfmt.Registry) error {

	if err := m.CallerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("caller_id")
		}
		return err
	}

	return nil
}

func (m *GAObject) validateGasPrice(formats strfmt.Registry) error {

	if err := m.GasPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas_price")
		}
		return err
	}

	return nil
}

func (m *GAObject) validateGasUsed(formats strfmt.Registry) error {

	if err := m.GasUsed.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas_used")
		}
		return err
	}

	return nil
}

func (m *GAObject) validateHeight(formats strfmt.Registry) error {

	if err := m.Height.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("height")
		}
		return err
	}

	return nil
}

func (m *GAObject) validateInnerObject(formats strfmt.Registry) error {

	if swag.IsZero(m.InnerObject) { // not required
		return nil
	}

	if m.InnerObject != nil {
		if err := m.InnerObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inner_object")
			}
			return err
		}
	}

	return nil
}

func (m *GAObject) validateReturnType(formats strfmt.Registry) error {

	if err := validate.Required("return_type", "body", m.ReturnType); err != nil {
		return err
	}

	return nil
}

func (m *GAObject) validateReturnValue(formats strfmt.Registry) error {

	if err := m.ReturnValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("return_value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GAObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GAObject) UnmarshalBinary(b []byte) error {
	var res GAObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
