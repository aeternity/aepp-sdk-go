// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/aeternity/aepp-sdk-go/v9/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GAAttachTx g a attach tx
//
// swagger:model GAAttachTx
type GAAttachTx struct {

	// ABI version
	// Required: true
	AbiVersion *uint16 `json:"abi_version"`

	// Contract authorization function hash (hex encoded)
	// Required: true
	// Pattern: ^(0x|0X)?[a-fA-F0-9]+$
	AuthFun *string `json:"auth_fun"`

	// Contract call data
	// Required: true
	CallData *string `json:"call_data"`

	// Contract's code
	// Required: true
	Code *string `json:"code"`

	// fee
	// Required: true
	Fee *utils.BigInt `json:"fee"`

	// gas
	// Required: true
	Gas *uint64 `json:"gas"`

	// gas price
	// Required: true
	GasPrice *utils.BigInt `json:"gas_price"`

	// Owner's nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// Contract owner pub_key
	// Required: true
	OwnerID *string `json:"owner_id"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`

	// Virtual machine's version
	// Required: true
	VMVersion *uint16 `json:"vm_version"`
}

// Validate validates this g a attach tx
func (m *GAAttachTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthFun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
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

	if err := m.validateOwnerID(formats); err != nil {
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

func (m *GAAttachTx) validateAbiVersion(formats strfmt.Registry) error {

	if err := validate.Required("abi_version", "body", m.AbiVersion); err != nil {
		return err
	}

	return nil
}

func (m *GAAttachTx) validateAuthFun(formats strfmt.Registry) error {

	if err := validate.Required("auth_fun", "body", m.AuthFun); err != nil {
		return err
	}

	if err := validate.Pattern("auth_fun", "body", *m.AuthFun, `^(0x|0X)?[a-fA-F0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *GAAttachTx) validateCallData(formats strfmt.Registry) error {

	if err := validate.Required("call_data", "body", m.CallData); err != nil {
		return err
	}

	return nil
}

func (m *GAAttachTx) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *GAAttachTx) validateFee(formats strfmt.Registry) error {

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

func (m *GAAttachTx) validateGas(formats strfmt.Registry) error {

	if err := validate.Required("gas", "body", m.Gas); err != nil {
		return err
	}

	return nil
}

func (m *GAAttachTx) validateGasPrice(formats strfmt.Registry) error {

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

func (m *GAAttachTx) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("owner_id", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *GAAttachTx) validateVMVersion(formats strfmt.Registry) error {

	if err := validate.Required("vm_version", "body", m.VMVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this g a attach tx based on the context it is used
func (m *GAAttachTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *GAAttachTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GAAttachTx) contextValidateGasPrice(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GAAttachTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GAAttachTx) UnmarshalBinary(b []byte) error {
	var res GAAttachTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
