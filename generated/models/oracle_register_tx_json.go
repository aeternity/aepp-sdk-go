// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OracleRegisterTxJSON oracle register tx JSON
// swagger:model OracleRegisterTxJSON
type OracleRegisterTxJSON struct {
	versionField *int64

	OracleRegisterTx
}

// Type gets the type of this subtype
func (m *OracleRegisterTxJSON) Type() string {
	return "OracleRegisterTx"
}

// SetType sets the type of this subtype
func (m *OracleRegisterTxJSON) SetType(val string) {

}

// Version gets the version of this subtype
func (m *OracleRegisterTxJSON) Version() *int64 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *OracleRegisterTxJSON) SetVersion(val *int64) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OracleRegisterTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		OracleRegisterTx
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

		Version *int64 `json:"version"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OracleRegisterTxJSON

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.versionField = base.Version

	result.OracleRegisterTx = data.OracleRegisterTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OracleRegisterTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		OracleRegisterTx
	}{

		OracleRegisterTx: m.OracleRegisterTx,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`

		Version *int64 `json:"version"`
	}{

		Type: m.Type(),

		Version: m.Version(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this oracle register tx JSON
func (m *OracleRegisterTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with OracleRegisterTx
	if err := m.OracleRegisterTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleRegisterTxJSON) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleRegisterTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleRegisterTxJSON) UnmarshalBinary(b []byte) error {
	var res OracleRegisterTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
