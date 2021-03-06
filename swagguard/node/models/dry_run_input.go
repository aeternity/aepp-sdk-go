// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DryRunInput dry run input
//
// swagger:model DryRunInput
type DryRunInput struct {

	// Accounts
	Accounts []*DryRunAccount `json:"accounts"`

	// top
	Top string `json:"top,omitempty"`

	// Collect and return on-chain tx events that would result from the call
	TxEvents *bool `json:"tx_events,omitempty"`

	// Txs
	// Required: true
	Txs []*DryRunInputItem `json:"txs"`
}

// Validate validates this dry run input
func (m *DryRunInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DryRunInput) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Accounts); i++ {
		if swag.IsZero(m.Accounts[i]) { // not required
			continue
		}

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DryRunInput) validateTxs(formats strfmt.Registry) error {

	if err := validate.Required("txs", "body", m.Txs); err != nil {
		return err
	}

	for i := 0; i < len(m.Txs); i++ {
		if swag.IsZero(m.Txs[i]) { // not required
			continue
		}

		if m.Txs[i] != nil {
			if err := m.Txs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dry run input based on the context it is used
func (m *DryRunInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTxs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DryRunInput) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Accounts); i++ {

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DryRunInput) contextValidateTxs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Txs); i++ {

		if m.Txs[i] != nil {
			if err := m.Txs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DryRunInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DryRunInput) UnmarshalBinary(b []byte) error {
	var res DryRunInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
