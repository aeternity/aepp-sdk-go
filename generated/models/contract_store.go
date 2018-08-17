// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContractStore contract store
// swagger:model ContractStore
type ContractStore struct {

	// store
	Store []*ContractStoreStore `json:"store"`
}

// Validate validates this contract store
func (m *ContractStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContractStore) validateStore(formats strfmt.Registry) error {

	if swag.IsZero(m.Store) { // not required
		return nil
	}

	for i := 0; i < len(m.Store); i++ {
		if swag.IsZero(m.Store[i]) { // not required
			continue
		}

		if m.Store[i] != nil {
			if err := m.Store[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("store" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractStore) UnmarshalBinary(b []byte) error {
	var res ContractStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}