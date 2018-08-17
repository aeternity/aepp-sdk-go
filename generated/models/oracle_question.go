// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OracleQuestion oracle question
// swagger:model OracleQuestion
type OracleQuestion struct {

	// expires at
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// query fee
	QueryFee int64 `json:"query_fee,omitempty"`

	// query id
	QueryID string `json:"query_id,omitempty"`
}

// Validate validates this oracle question
func (m *OracleQuestion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleQuestion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleQuestion) UnmarshalBinary(b []byte) error {
	var res OracleQuestion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}