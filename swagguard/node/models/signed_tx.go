// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignedTx signed tx
//
// swagger:model SignedTx
type SignedTx struct {

	// At least one signature is required unless for Generalized Account Meta transactions
	// Required: true
	// Min Items: 0
	Signatures []string `json:"signatures"`

	txField GenericTx
}

// Tx gets the tx of this base type
func (m *SignedTx) Tx() GenericTx {
	return m.txField
}

// SetTx sets the tx of this base type
func (m *SignedTx) SetTx(val GenericTx) {
	m.txField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SignedTx) UnmarshalJSON(raw []byte) error {
	var data struct {
		Signatures []string `json:"signatures"`

		Tx json.RawMessage `json:"tx"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propTx, err := UnmarshalGenericTx(bytes.NewBuffer(data.Tx), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result SignedTx

	// signatures
	result.Signatures = data.Signatures

	// tx
	result.txField = propTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SignedTx) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Signatures []string `json:"signatures"`
	}{

		Signatures: m.Signatures,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Tx GenericTx `json:"tx"`
	}{

		Tx: m.txField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this signed tx
func (m *SignedTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignedTx) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("signatures", "body", m.Signatures); err != nil {
		return err
	}

	iSignaturesSize := int64(len(m.Signatures))

	if err := validate.MinItems("signatures", "body", iSignaturesSize, 0); err != nil {
		return err
	}

	return nil
}

func (m *SignedTx) validateTx(formats strfmt.Registry) error {

	if err := validate.Required("tx", "body", m.Tx()); err != nil {
		return err
	}

	if err := m.Tx().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tx")
		}
		return err
	}

	return nil
}

// ContextValidate validate this signed tx based on the context it is used
func (m *SignedTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignedTx) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Tx().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tx")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SignedTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignedTx) UnmarshalBinary(b []byte) error {
	var res SignedTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
