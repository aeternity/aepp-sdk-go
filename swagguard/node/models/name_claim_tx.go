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

// NameClaimTx name claim tx
//
// swagger:model NameClaimTx
type NameClaimTx struct {

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// fee
	// Required: true
	Fee *utils.BigInt `json:"fee"`

	// name
	// Required: true
	Name *string `json:"name"`

	// name fee
	NameFee utils.BigInt `json:"name_fee,omitempty"`

	// name salt
	// Required: true
	NameSalt *utils.BigInt `json:"name_salt"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this name claim tx
func (m *NameClaimTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameSalt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameClaimTx) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *NameClaimTx) validateFee(formats strfmt.Registry) error {

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

func (m *NameClaimTx) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NameClaimTx) validateNameFee(formats strfmt.Registry) error {
	if swag.IsZero(m.NameFee) { // not required
		return nil
	}

	if err := m.NameFee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name_fee")
		}
		return err
	}

	return nil
}

func (m *NameClaimTx) validateNameSalt(formats strfmt.Registry) error {

	if err := validate.Required("name_salt", "body", m.NameSalt); err != nil {
		return err
	}

	if err := validate.Required("name_salt", "body", m.NameSalt); err != nil {
		return err
	}

	if m.NameSalt != nil {
		if err := m.NameSalt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name_salt")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this name claim tx based on the context it is used
func (m *NameClaimTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNameFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNameSalt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameClaimTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NameClaimTx) contextValidateNameFee(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NameFee.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name_fee")
		}
		return err
	}

	return nil
}

func (m *NameClaimTx) contextValidateNameSalt(ctx context.Context, formats strfmt.Registry) error {

	if m.NameSalt != nil {
		if err := m.NameSalt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name_salt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameClaimTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameClaimTx) UnmarshalBinary(b []byte) error {
	var res NameClaimTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
