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

// ChannelSettleTx channel settle tx
//
// swagger:model ChannelSettleTx
type ChannelSettleTx struct {

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// fee
	// Required: true
	Fee *utils.BigInt `json:"fee"`

	// from id
	// Required: true
	FromID *string `json:"from_id"`

	// initiator amount final
	// Required: true
	InitiatorAmountFinal *utils.BigInt `json:"initiator_amount_final"`

	// nonce
	// Required: true
	Nonce *uint64 `json:"nonce"`

	// responder amount final
	// Required: true
	ResponderAmountFinal *utils.BigInt `json:"responder_amount_final"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel settle tx
func (m *ChannelSettleTx) Validate(formats strfmt.Registry) error {
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

func (m *ChannelSettleTx) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSettleTx) validateFee(formats strfmt.Registry) error {

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

func (m *ChannelSettleTx) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", m.FromID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSettleTx) validateInitiatorAmountFinal(formats strfmt.Registry) error {

	if err := validate.Required("initiator_amount_final", "body", m.InitiatorAmountFinal); err != nil {
		return err
	}

	if err := validate.Required("initiator_amount_final", "body", m.InitiatorAmountFinal); err != nil {
		return err
	}

	if m.InitiatorAmountFinal != nil {
		if err := m.InitiatorAmountFinal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_amount_final")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelSettleTx) validateNonce(formats strfmt.Registry) error {

	if err := validate.Required("nonce", "body", m.Nonce); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSettleTx) validateResponderAmountFinal(formats strfmt.Registry) error {

	if err := validate.Required("responder_amount_final", "body", m.ResponderAmountFinal); err != nil {
		return err
	}

	if err := validate.Required("responder_amount_final", "body", m.ResponderAmountFinal); err != nil {
		return err
	}

	if m.ResponderAmountFinal != nil {
		if err := m.ResponderAmountFinal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responder_amount_final")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this channel settle tx based on the context it is used
func (m *ChannelSettleTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiatorAmountFinal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponderAmountFinal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelSettleTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChannelSettleTx) contextValidateInitiatorAmountFinal(ctx context.Context, formats strfmt.Registry) error {

	if m.InitiatorAmountFinal != nil {
		if err := m.InitiatorAmountFinal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_amount_final")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelSettleTx) contextValidateResponderAmountFinal(ctx context.Context, formats strfmt.Registry) error {

	if m.ResponderAmountFinal != nil {
		if err := m.ResponderAmountFinal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responder_amount_final")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelSettleTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelSettleTx) UnmarshalBinary(b []byte) error {
	var res ChannelSettleTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
