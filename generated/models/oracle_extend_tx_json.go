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
)

// OracleExtendTxJSON oracle extend tx JSON
// swagger:model OracleExtendTxJSON
type OracleExtendTxJSON struct {
	vsnField int64

	OracleExtendTx
}

// DataSchema gets the data schema of this subtype
func (m *OracleExtendTxJSON) DataSchema() string {
	return "OracleExtendTxJSON"
}

// SetDataSchema sets the data schema of this subtype
func (m *OracleExtendTxJSON) SetDataSchema(val string) {

}

// Vsn gets the vsn of this subtype
func (m *OracleExtendTxJSON) Vsn() int64 {
	return m.vsnField
}

// SetVsn sets the vsn of this subtype
func (m *OracleExtendTxJSON) SetVsn(val int64) {
	m.vsnField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OracleExtendTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		OracleExtendTx
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DataSchema string `json:"data_schema"`

		Vsn int64 `json:"vsn,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OracleExtendTxJSON

	if base.DataSchema != result.DataSchema() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid data_schema value: %q", base.DataSchema)
	}

	result.vsnField = base.Vsn

	result.OracleExtendTx = data.OracleExtendTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OracleExtendTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		OracleExtendTx
	}{

		OracleExtendTx: m.OracleExtendTx,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DataSchema string `json:"data_schema"`

		Vsn int64 `json:"vsn,omitempty"`
	}{

		DataSchema: m.DataSchema(),

		Vsn: m.Vsn(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this oracle extend tx JSON
func (m *OracleExtendTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OracleExtendTx
	if err := m.OracleExtendTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OracleExtendTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleExtendTxJSON) UnmarshalBinary(b []byte) error {
	var res OracleExtendTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}