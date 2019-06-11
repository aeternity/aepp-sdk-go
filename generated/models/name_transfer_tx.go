// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// NameTransferTx name transfer tx
// swagger:model NameTransferTx
type NameTransferTx struct {

	// account id
	// Required: true
	AccountID EncodedPubkey `json:"account_id"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// name id
	// Required: true
	NameID EncodedValue `json:"name_id"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// recipient id
	// Required: true
	RecipientID EncodedPubkey `json:"recipient_id"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this name transfer tx
func (m *NameTransferTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameTransferTx) validateAccountID(formats strfmt.Registry) error {

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_id")
		}
		return err
	}

	return nil
}

func (m *NameTransferTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *NameTransferTx) validateNameID(formats strfmt.Registry) error {

	if err := m.NameID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name_id")
		}
		return err
	}

	return nil
}

func (m *NameTransferTx) validateRecipientID(formats strfmt.Registry) error {

	if err := m.RecipientID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("recipient_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameTransferTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameTransferTx) UnmarshalBinary(b []byte) error {
	var res NameTransferTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
