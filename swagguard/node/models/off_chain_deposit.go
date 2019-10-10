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

	utils "github.com/aeternity/aepp-sdk-go/v6/utils"
)

// OffChainDeposit off chain deposit
// swagger:model OffChainDeposit
type OffChainDeposit struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// Depositor of tokens
	// Required: true
	From *string `json:"from"`
}

// Op gets the op of this subtype
func (m *OffChainDeposit) Op() string {
	return "OffChainDeposit"
}

// SetOp sets the op of this subtype
func (m *OffChainDeposit) SetOp(val string) {

}

// Amount gets the amount of this subtype

// From gets the from of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OffChainDeposit) UnmarshalJSON(raw []byte) error {
	var data struct {

		// amount
		// Required: true
		Amount utils.BigInt `json:"amount"`

		// Depositor of tokens
		// Required: true
		From *string `json:"from"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Op string `json:"op"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OffChainDeposit

	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.Amount = data.Amount

	result.From = data.From

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OffChainDeposit) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// amount
		// Required: true
		Amount utils.BigInt `json:"amount"`

		// Depositor of tokens
		// Required: true
		From *string `json:"from"`
	}{

		Amount: m.Amount,

		From: m.From,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Op string `json:"op"`
	}{

		Op: m.Op(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this off chain deposit
func (m *OffChainDeposit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OffChainDeposit) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *OffChainDeposit) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OffChainDeposit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffChainDeposit) UnmarshalBinary(b []byte) error {
	var res OffChainDeposit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
