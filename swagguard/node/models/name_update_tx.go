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

	utils "github.com/aeternity/aepp-sdk-go/v7/utils"
)

// NameUpdateTx name update tx
// swagger:model NameUpdateTx
type NameUpdateTx struct {

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// client ttl
	// Required: true
	ClientTTL *uint64 `json:"client_ttl"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// name id
	// Required: true
	NameID *string `json:"name_id"`

	// name ttl
	// Required: true
	NameTTL *uint64 `json:"name_ttl"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// pointers
	// Required: true
	Pointers []*NamePointer `json:"pointers"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this name update tx
func (m *NameUpdateTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameUpdateTx) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *NameUpdateTx) validateClientTTL(formats strfmt.Registry) error {

	if err := validate.Required("client_ttl", "body", m.ClientTTL); err != nil {
		return err
	}

	return nil
}

func (m *NameUpdateTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *NameUpdateTx) validateNameID(formats strfmt.Registry) error {

	if err := validate.Required("name_id", "body", m.NameID); err != nil {
		return err
	}

	return nil
}

func (m *NameUpdateTx) validateNameTTL(formats strfmt.Registry) error {

	if err := validate.Required("name_ttl", "body", m.NameTTL); err != nil {
		return err
	}

	return nil
}

func (m *NameUpdateTx) validatePointers(formats strfmt.Registry) error {

	if err := validate.Required("pointers", "body", m.Pointers); err != nil {
		return err
	}

	for i := 0; i < len(m.Pointers); i++ {
		if swag.IsZero(m.Pointers[i]) { // not required
			continue
		}

		if m.Pointers[i] != nil {
			if err := m.Pointers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pointers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameUpdateTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameUpdateTx) UnmarshalBinary(b []byte) error {
	var res NameUpdateTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
