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
)

// Channel channel
// swagger:model Channel
type Channel struct {

	// channel amount
	// Minimum: 0
	ChannelAmount *int64 `json:"channel_amount,omitempty"`

	// channel reserve
	// Minimum: 0
	ChannelReserve *int64 `json:"channel_reserve,omitempty"`

	// closes at
	ClosesAt int64 `json:"closes_at,omitempty"`

	// delegate ids
	DelegateIds []EncodedHash `json:"delegate_ids"`

	// forcing blocked until
	ForcingBlockedUntil int64 `json:"forcing_blocked_until,omitempty"`

	// id
	ID EncodedHash `json:"id,omitempty"`

	// initiator amount
	// Minimum: 0
	InitiatorAmount *int64 `json:"initiator_amount,omitempty"`

	// initiator id
	InitiatorID EncodedHash `json:"initiator_id,omitempty"`

	// lock period
	// Minimum: 0
	LockPeriod *int64 `json:"lock_period,omitempty"`

	// responder amount
	// Minimum: 0
	ResponderAmount *int64 `json:"responder_amount,omitempty"`

	// responder id
	ResponderID EncodedHash `json:"responder_id,omitempty"`

	// round
	// Minimum: 0
	Round *int64 `json:"round,omitempty"`

	// state hash
	StateHash EncodedHash `json:"state_hash,omitempty"`
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

	if err := m.validateResponderAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderID(formats); err != nil {
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

func (m *Channel) validateChannelAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelAmount) { // not required
		return nil
	}

	if err := validate.MinimumInt("channel_amount", "body", int64(*m.ChannelAmount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateChannelReserve(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelReserve) { // not required
		return nil
	}

	if err := validate.MinimumInt("channel_reserve", "body", int64(*m.ChannelReserve), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateDelegateIds(formats strfmt.Registry) error {

	if swag.IsZero(m.DelegateIds) { // not required
		return nil
	}

	for i := 0; i < len(m.DelegateIds); i++ {

		if err := m.DelegateIds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegate_ids" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Channel) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Channel) validateInitiatorAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.InitiatorAmount) { // not required
		return nil
	}

	if err := validate.MinimumInt("initiator_amount", "body", int64(*m.InitiatorAmount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateInitiatorID(formats strfmt.Registry) error {

	if swag.IsZero(m.InitiatorID) { // not required
		return nil
	}

	if err := m.InitiatorID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("initiator_id")
		}
		return err
	}

	return nil
}

func (m *Channel) validateLockPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.LockPeriod) { // not required
		return nil
	}

	if err := validate.MinimumInt("lock_period", "body", int64(*m.LockPeriod), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateResponderAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponderAmount) { // not required
		return nil
	}

	if err := validate.MinimumInt("responder_amount", "body", int64(*m.ResponderAmount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateResponderID(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponderID) { // not required
		return nil
	}

	if err := m.ResponderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder_id")
		}
		return err
	}

	return nil
}

func (m *Channel) validateRound(formats strfmt.Registry) error {

	if swag.IsZero(m.Round) { // not required
		return nil
	}

	if err := validate.MinimumInt("round", "body", int64(*m.Round), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateStateHash(formats strfmt.Registry) error {

	if swag.IsZero(m.StateHash) { // not required
		return nil
	}

	if err := m.StateHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state_hash")
		}
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
