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

// ChannelCreateTx channel create tx
//
// swagger:model ChannelCreateTx
type ChannelCreateTx struct {

	// channel reserve
	// Required: true
	ChannelReserve *utils.BigInt `json:"channel_reserve"`

	// fee
	// Required: true
	Fee *utils.BigInt `json:"fee"`

	// initiator amount
	// Required: true
	InitiatorAmount *utils.BigInt `json:"initiator_amount"`

	// initiator id
	// Required: true
	InitiatorID *string `json:"initiator_id"`

	// lock period
	// Required: true
	LockPeriod *uint64 `json:"lock_period"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// responder amount
	// Required: true
	ResponderAmount *utils.BigInt `json:"responder_amount"`

	// responder id
	// Required: true
	ResponderID *string `json:"responder_id"`

	// Root hash of the channel's internal state tree
	// Required: true
	StateHash *string `json:"state_hash"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel create tx
func (m *ChannelCreateTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelReserve(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderID(formats); err != nil {
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

func (m *ChannelCreateTx) validateChannelReserve(formats strfmt.Registry) error {

	if err := validate.Required("channel_reserve", "body", m.ChannelReserve); err != nil {
		return err
	}

	if err := validate.Required("channel_reserve", "body", m.ChannelReserve); err != nil {
		return err
	}

	if m.ChannelReserve != nil {
		if err := m.ChannelReserve.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel_reserve")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelCreateTx) validateFee(formats strfmt.Registry) error {

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

func (m *ChannelCreateTx) validateInitiatorAmount(formats strfmt.Registry) error {

	if err := validate.Required("initiator_amount", "body", m.InitiatorAmount); err != nil {
		return err
	}

	if err := validate.Required("initiator_amount", "body", m.InitiatorAmount); err != nil {
		return err
	}

	if m.InitiatorAmount != nil {
		if err := m.InitiatorAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_amount")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelCreateTx) validateInitiatorID(formats strfmt.Registry) error {

	if err := validate.Required("initiator_id", "body", m.InitiatorID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateLockPeriod(formats strfmt.Registry) error {

	if err := validate.Required("lock_period", "body", m.LockPeriod); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateResponderAmount(formats strfmt.Registry) error {

	if err := validate.Required("responder_amount", "body", m.ResponderAmount); err != nil {
		return err
	}

	if err := validate.Required("responder_amount", "body", m.ResponderAmount); err != nil {
		return err
	}

	if m.ResponderAmount != nil {
		if err := m.ResponderAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responder_amount")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelCreateTx) validateResponderID(formats strfmt.Registry) error {

	if err := validate.Required("responder_id", "body", m.ResponderID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateStateHash(formats strfmt.Registry) error {

	if err := validate.Required("state_hash", "body", m.StateHash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this channel create tx based on the context it is used
func (m *ChannelCreateTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannelReserve(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiatorAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponderAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelCreateTx) contextValidateChannelReserve(ctx context.Context, formats strfmt.Registry) error {

	if m.ChannelReserve != nil {
		if err := m.ChannelReserve.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel_reserve")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelCreateTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChannelCreateTx) contextValidateInitiatorAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.InitiatorAmount != nil {
		if err := m.InitiatorAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_amount")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelCreateTx) contextValidateResponderAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.ResponderAmount != nil {
		if err := m.ResponderAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responder_amount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCreateTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCreateTx) UnmarshalBinary(b []byte) error {
	var res ChannelCreateTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
