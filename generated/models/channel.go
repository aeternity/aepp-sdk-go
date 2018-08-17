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

	// channel reserve
	// Minimum: 0
	ChannelReserve *int64 `json:"channel_reserve,omitempty"`

	// closes at
	ClosesAt int64 `json:"closes_at,omitempty"`

	// delegates
	Delegates []EncodedHash `json:"delegates"`

	// id
	ID EncodedHash `json:"id,omitempty"`

	// initiator
	Initiator EncodedHash `json:"initiator,omitempty"`

	// initiator amount
	// Minimum: 0
	InitiatorAmount *int64 `json:"initiator_amount,omitempty"`

	// lock period
	// Minimum: 0
	LockPeriod *int64 `json:"lock_period,omitempty"`

	// responder
	Responder EncodedHash `json:"responder,omitempty"`

	// responder amount
	// Minimum: 0
	ResponderAmount *int64 `json:"responder_amount,omitempty"`

	// round
	// Minimum: 0
	Round *int64 `json:"round,omitempty"`

	// state hash
	StateHash EncodedHash `json:"state_hash,omitempty"`

	// total amount
	// Minimum: 0
	TotalAmount *int64 `json:"total_amount,omitempty"`
}

// Validate validates this channel
func (m *Channel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelReserve(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
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

func (m *Channel) validateDelegates(formats strfmt.Registry) error {

	if swag.IsZero(m.Delegates) { // not required
		return nil
	}

	for i := 0; i < len(m.Delegates); i++ {

		if err := m.Delegates[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegates" + "." + strconv.Itoa(i))
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

func (m *Channel) validateInitiator(formats strfmt.Registry) error {

	if swag.IsZero(m.Initiator) { // not required
		return nil
	}

	if err := m.Initiator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("initiator")
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

func (m *Channel) validateLockPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.LockPeriod) { // not required
		return nil
	}

	if err := validate.MinimumInt("lock_period", "body", int64(*m.LockPeriod), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Channel) validateResponder(formats strfmt.Registry) error {

	if swag.IsZero(m.Responder) { // not required
		return nil
	}

	if err := m.Responder.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder")
		}
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

func (m *Channel) validateTotalAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalAmount) { // not required
		return nil
	}

	if err := validate.MinimumInt("total_amount", "body", int64(*m.TotalAmount), 0, false); err != nil {
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