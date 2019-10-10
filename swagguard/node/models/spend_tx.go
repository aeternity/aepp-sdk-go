// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v6/utils"
)

// SpendTx spend tx
// swagger:model SpendTx
type SpendTx struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// payload
	// Required: true
	Payload *string `json:"payload"`

	// recipient id
	// Required: true
	RecipientID *string `json:"recipient_id"`

	// sender id
	// Required: true
	SenderID *string `json:"sender_id"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this spend tx
func (m *SpendTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpendTx) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *SpendTx) validateRecipientID(formats strfmt.Registry) error {

	if err := validate.Required("recipient_id", "body", m.RecipientID); err != nil {
		return err
	}

	return nil
}

func (m *SpendTx) validateSenderID(formats strfmt.Registry) error {

	if err := validate.Required("sender_id", "body", m.SenderID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpendTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpendTx) UnmarshalBinary(b []byte) error {
	var res SpendTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
