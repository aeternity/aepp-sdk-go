// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChannelCloseMutualTxJSON channel close mutual tx JSON
//
// swagger:model ChannelCloseMutualTxJSON
type ChannelCloseMutualTxJSON struct {
	versionField *uint32

	ChannelCloseMutualTx
}

// Type gets the type of this subtype
func (m *ChannelCloseMutualTxJSON) Type() string {
	return "ChannelCloseMutualTx"
}

// SetType sets the type of this subtype
func (m *ChannelCloseMutualTxJSON) SetType(val string) {
}

// Version gets the version of this subtype
func (m *ChannelCloseMutualTxJSON) Version() *uint32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *ChannelCloseMutualTxJSON) SetVersion(val *uint32) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ChannelCloseMutualTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		ChannelCloseMutualTx
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`

		Version *uint32 `json:"version"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ChannelCloseMutualTxJSON

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.versionField = base.Version

	result.ChannelCloseMutualTx = data.ChannelCloseMutualTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ChannelCloseMutualTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		ChannelCloseMutualTx
	}{

		ChannelCloseMutualTx: m.ChannelCloseMutualTx,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`

		Version *uint32 `json:"version"`
	}{

		Type: m.Type(),

		Version: m.Version(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this channel close mutual tx JSON
func (m *ChannelCloseMutualTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with ChannelCloseMutualTx
	if err := m.ChannelCloseMutualTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelCloseMutualTxJSON) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this channel close mutual tx JSON based on the context it is used
func (m *ChannelCloseMutualTxJSON) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChannelCloseMutualTx
	if err := m.ChannelCloseMutualTx.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCloseMutualTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCloseMutualTxJSON) UnmarshalBinary(b []byte) error {
	var res ChannelCloseMutualTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
