// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateContractUnsignedTx create contract unsigned tx
//
// swagger:model CreateContractUnsignedTx
type CreateContractUnsignedTx struct {
	UnsignedTx

	// Address of the contract to be created
	// Required: true
	ContractID *string `json:"contract_id"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CreateContractUnsignedTx) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UnsignedTx
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UnsignedTx = aO0

	// AO1
	var dataAO1 struct {
		ContractID *string `json:"contract_id"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ContractID = dataAO1.ContractID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CreateContractUnsignedTx) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.UnsignedTx)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ContractID *string `json:"contract_id"`
	}

	dataAO1.ContractID = m.ContractID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create contract unsigned tx
func (m *CreateContractUnsignedTx) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UnsignedTx
	if err := m.UnsignedTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateContractUnsignedTx) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create contract unsigned tx based on the context it is used
func (m *CreateContractUnsignedTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UnsignedTx
	if err := m.UnsignedTx.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateContractUnsignedTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateContractUnsignedTx) UnmarshalBinary(b []byte) error {
	var res CreateContractUnsignedTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
