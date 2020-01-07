// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v8/utils"
)

// ChannelCloseMutualTx channel close mutual tx
// swagger:model ChannelCloseMutualTx
type ChannelCloseMutualTx struct {

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// fee
	// Required: true
	Fee *uint64 `json:"fee"`

	// from id
	// Required: true
	FromID *string `json:"from_id"`

	// initiator amount final
	// Required: true
	InitiatorAmountFinal utils.BigInt `json:"initiator_amount_final"`

	// nonce
	// Required: true
	Nonce *uint64 `json:"nonce"`

	// responder amount final
	// Required: true
	ResponderAmountFinal utils.BigInt `json:"responder_amount_final"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel close mutual tx
func (m *ChannelCloseMutualTx) Validate(formats strfmt.Registry) error {
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

	if err := m.validateInitiatorAmountFinal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderAmountFinal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelCloseMutualTx) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", m.FromID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateInitiatorAmountFinal(formats strfmt.Registry) error {

	if err := m.InitiatorAmountFinal.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("initiator_amount_final")
		}
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateNonce(formats strfmt.Registry) error {

	if err := validate.Required("nonce", "body", m.Nonce); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateResponderAmountFinal(formats strfmt.Registry) error {

	if err := m.ResponderAmountFinal.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder_amount_final")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCloseMutualTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCloseMutualTx) UnmarshalBinary(b []byte) error {
	var res ChannelCloseMutualTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
