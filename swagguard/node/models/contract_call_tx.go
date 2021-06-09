// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/aeternity/aepp-sdk-go/v8/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContractCallTx contract call tx
//
// swagger:model ContractCallTx
type ContractCallTx struct {

	// ABI version
	// Required: true
	AbiVersion *uint16 `json:"abi_version"`

	// amount
	// Required: true
	Amount *utils.BigInt `json:"amount"`

	// Contract call data
	// Required: true
	CallData *string `json:"call_data"`

	// Contract caller pub_key
	// Required: true
	CallerID *string `json:"caller_id"`

	// Contract's pub_key
	// Required: true
	ContractID *string `json:"contract_id"`

	// fee
	// Required: true
	Fee *utils.BigInt `json:"fee"`

	// gas
	// Required: true
	Gas *uint64 `json:"gas"`

	// gas price
	// Required: true
	GasPrice *utils.BigInt `json:"gas_price"`

	// Caller's nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this contract call tx
func (m *ContractCallTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContractCallTx) validateAbiVersion(formats strfmt.Registry) error {

	if err := validate.Required("abi_version", "body", m.AbiVersion); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallTx) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *ContractCallTx) validateCallData(formats strfmt.Registry) error {

	if err := validate.Required("call_data", "body", m.CallData); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallTx) validateCallerID(formats strfmt.Registry) error {

	if err := validate.Required("caller_id", "body", m.CallerID); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallTx) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *ContractCallTx) validateGas(formats strfmt.Registry) error {

	if err := validate.Required("gas", "body", m.Gas); err != nil {
		return err
	}

	return nil
}

func (m *ContractCallTx) validateGasPrice(formats strfmt.Registry) error {

	if err := validate.Required("gas_price", "body", m.GasPrice); err != nil {
		return err
	}

	if err := validate.Required("gas_price", "body", m.GasPrice); err != nil {
		return err
	}

	if m.GasPrice != nil {
		if err := m.GasPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gas_price")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this contract call tx based on the context it is used
func (m *ContractCallTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGasPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContractCallTx) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {
		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *ContractCallTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if m.Fee != nil {
		if err := m.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *ContractCallTx) contextValidateGasPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.GasPrice != nil {
		if err := m.GasPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gas_price")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractCallTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractCallTx) UnmarshalBinary(b []byte) error {
	var res ContractCallTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
