// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v5/utils"
)

// ChannelWithdrawTx channel withdraw tx
// swagger:model ChannelWithdrawTx
type ChannelWithdrawTx struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// nonce
	// Required: true
	Nonce *uint64 `json:"nonce"`

	// Channel's next round
	// Required: true
	Round *uint64 `json:"round"`

	// Root hash of the channel's internal state tree after the withdraw had been applied to it
	// Required: true
	StateHash *string `json:"state_hash"`

	// to id
	// Required: true
	ToID *string `json:"to_id"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel withdraw tx
func (m *ChannelWithdrawTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelWithdrawTx) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *ChannelWithdrawTx) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelWithdrawTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *ChannelWithdrawTx) validateNonce(formats strfmt.Registry) error {

	if err := validate.Required("nonce", "body", m.Nonce); err != nil {
		return err
	}

	return nil
}

func (m *ChannelWithdrawTx) validateRound(formats strfmt.Registry) error {

	if err := validate.Required("round", "body", m.Round); err != nil {
		return err
	}

	return nil
}

func (m *ChannelWithdrawTx) validateStateHash(formats strfmt.Registry) error {

	if err := validate.Required("state_hash", "body", m.StateHash); err != nil {
		return err
	}

	return nil
}

func (m *ChannelWithdrawTx) validateToID(formats strfmt.Registry) error {

	if err := validate.Required("to_id", "body", m.ToID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelWithdrawTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelWithdrawTx) UnmarshalBinary(b []byte) error {
	var res ChannelWithdrawTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
