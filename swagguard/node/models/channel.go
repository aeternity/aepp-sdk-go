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

// Channel channel
// swagger:model Channel
type Channel struct {

	// channel amount
	// Required: true
	ChannelAmount utils.BigInt `json:"channel_amount"`

	// channel reserve
	// Required: true
	ChannelReserve utils.BigInt `json:"channel_reserve"`

	// delegate ids
	// Required: true
	DelegateIds []string `json:"delegate_ids"`

	// id
	// Required: true
	ID *string `json:"id"`

	// initiator amount
	// Required: true
	InitiatorAmount utils.BigInt `json:"initiator_amount"`

	// initiator id
	// Required: true
	InitiatorID *string `json:"initiator_id"`

	// lock period
	// Required: true
	LockPeriod *uint64 `json:"lock_period"`

	// locked until
	// Required: true
	LockedUntil *uint64 `json:"locked_until"`

	// responder amount
	// Required: true
	ResponderAmount utils.BigInt `json:"responder_amount"`

	// responder id
	// Required: true
	ResponderID *string `json:"responder_id"`

	// round
	// Required: true
	Round *uint64 `json:"round"`

	// solo round
	// Required: true
	SoloRound *uint64 `json:"solo_round"`

	// state hash
	// Required: true
	StateHash *string `json:"state_hash"`
}

// Validate validates this channel
func (m *Channel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelReserve(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegateIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateLockedUntil(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoloRound(formats); err != nil {
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

func (m *Channel) validateChannelAmount(formats strfmt.Registry) error {

	if err := m.ChannelAmount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("channel_amount")
		}
		return err
	}

	return nil
}

func (m *Channel) validateChannelReserve(formats strfmt.Registry) error {

	if err := m.ChannelReserve.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("channel_reserve")
		}
		return err
	}

	return nil
}

func (m *Channel) validateDelegateIds(formats strfmt.Registry) error {

	if err := validate.Required("delegate_ids", "body", m.DelegateIds); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateInitiatorAmount(formats strfmt.Registry) error {

	if err := m.InitiatorAmount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("initiator_amount")
		}
		return err
	}

	return nil
}

func (m *Channel) validateInitiatorID(formats strfmt.Registry) error {

	if err := validate.Required("initiator_id", "body", m.InitiatorID); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateLockPeriod(formats strfmt.Registry) error {

	if err := validate.Required("lock_period", "body", m.LockPeriod); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateLockedUntil(formats strfmt.Registry) error {

	if err := validate.Required("locked_until", "body", m.LockedUntil); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateResponderAmount(formats strfmt.Registry) error {

	if err := m.ResponderAmount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder_amount")
		}
		return err
	}

	return nil
}

func (m *Channel) validateResponderID(formats strfmt.Registry) error {

	if err := validate.Required("responder_id", "body", m.ResponderID); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateRound(formats strfmt.Registry) error {

	if err := validate.Required("round", "body", m.Round); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateSoloRound(formats strfmt.Registry) error {

	if err := validate.Required("solo_round", "body", m.SoloRound); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateStateHash(formats strfmt.Registry) error {

	if err := validate.Required("state_hash", "body", m.StateHash); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Channel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Channel) UnmarshalBinary(b []byte) error {
	var res Channel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
