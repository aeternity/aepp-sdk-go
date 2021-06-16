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

// ChannelSetDelegatesTx channel set delegates tx
//
// swagger:model ChannelSetDelegatesTx
type ChannelSetDelegatesTx struct {

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// fee
	// Required: true
	Fee *utils.BigInt `json:"fee"`

	// from id
	// Required: true
	FromID *string `json:"from_id"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// payload
	// Required: true
	Payload *string `json:"payload"`

	// round
	// Required: true
	Round *uint64 `json:"round"`

	// state hash
	// Required: true
	StateHash *string `json:"state_hash"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel set delegates tx
func (m *ChannelSetDelegatesTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelSetDelegatesTx) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSetDelegatesTx) validateFee(formats strfmt.Registry) error {

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

func (m *ChannelSetDelegatesTx) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", m.FromID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSetDelegatesTx) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSetDelegatesTx) validateRound(formats strfmt.Registry) error {

	if err := validate.Required("round", "body", m.Round); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSetDelegatesTx) validateStateHash(formats strfmt.Registry) error {

	if err := validate.Required("state_hash", "body", m.StateHash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this channel set delegates tx based on the context it is used
func (m *ChannelSetDelegatesTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelSetDelegatesTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ChannelSetDelegatesTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelSetDelegatesTx) UnmarshalBinary(b []byte) error {
	var res ChannelSetDelegatesTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}