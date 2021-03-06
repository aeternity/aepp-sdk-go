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

// GAAttachTxJSON g a attach tx JSON
//
// swagger:model GAAttachTxJSON
type GAAttachTxJSON struct {
	versionField *uint32

	GAAttachTx
}

// Type gets the type of this subtype
func (m *GAAttachTxJSON) Type() string {
	return "GAAttachTx"
}

// SetType sets the type of this subtype
func (m *GAAttachTxJSON) SetType(val string) {
}

// Version gets the version of this subtype
func (m *GAAttachTxJSON) Version() *uint32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *GAAttachTxJSON) SetVersion(val *uint32) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GAAttachTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		GAAttachTx
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

	var result GAAttachTxJSON

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.versionField = base.Version

	result.GAAttachTx = data.GAAttachTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GAAttachTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		GAAttachTx
	}{

		GAAttachTx: m.GAAttachTx,
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

// Validate validates this g a attach tx JSON
func (m *GAAttachTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with GAAttachTx
	if err := m.GAAttachTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GAAttachTxJSON) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this g a attach tx JSON based on the context it is used
func (m *GAAttachTxJSON) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GAAttachTx
	if err := m.GAAttachTx.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GAAttachTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GAAttachTxJSON) UnmarshalBinary(b []byte) error {
	var res GAAttachTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
