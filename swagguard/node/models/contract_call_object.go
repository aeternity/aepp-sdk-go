// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v6/utils"
)

// ContractCallObject contract call object
// swagger:model ContractCallObject
type ContractCallObject struct {

	// caller id
	// Required: true
	CallerID *string `json:"caller_id"`

	// caller nonce
	// Required: true
	CallerNonce *uint64 `json:"caller_nonce"`

	// contract id
	// Required: true
	ContractID *string `json:"contract_id"`

	// gas price
	// Required: true
	GasPrice utils.BigInt `json:"gas_price"`

	// gas used
	// Required: true
	GasUsed *uint64 `json:"gas_used"`

	// height
	// Required: true
	Height *uint64 `json:"height"`

	// log
	// Required: true
	Log []*Event `json:"log"`

	// The status of the call 'ok | error | revert'.
	// Required: true
	ReturnType *string `json:"return_type"`

	// return value
	// Required: true
	ReturnValue *string `json:"return_value"`
}

// Validate validates this contract call object
func (m *ContractCallObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallerNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractID(formats); err != nil {
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

	if err := m.validateLog(formats); err != nil {
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

func (m *ContractCallObject) validateCallerID(formats strfmt.Registry) error {

	if err := validate.Required("caller_id", "body", m.CallerID); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallObject) validateCallerNonce(formats strfmt.Registry) error {

	if err := validate.Required("caller_nonce", "body", m.CallerNonce); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallObject) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallObject) validateGasPrice(formats strfmt.Registry) error {

	if err := m.GasPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas_price")
		}
		return err
	}

	return nil
}

func (m *ContractCallObject) validateGasUsed(formats strfmt.Registry) error {

	if err := validate.Required("gas_used", "body", m.GasUsed); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallObject) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallObject) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	for i := 0; i < len(m.Log); i++ {
		if swag.IsZero(m.Log[i]) { // not required
			continue
		}

		if m.Log[i] != nil {
			if err := m.Log[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContractCallObject) validateReturnType(formats strfmt.Registry) error {

	if err := validate.Required("return_type", "body", m.ReturnType); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallObject) validateReturnValue(formats strfmt.Registry) error {

	if err := validate.Required("return_value", "body", m.ReturnValue); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractCallObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractCallObject) UnmarshalBinary(b []byte) error {
	var res ContractCallObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
