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

// PeerConnections peer connections
//
// swagger:model PeerConnections
type PeerConnections struct {

	// Number of inbound peer connections
	// Required: true
	Inbound *uint32 `json:"inbound"`

	// Number of outbound peer connections
	// Required: true
	Outbound *uint32 `json:"outbound"`
}

// Validate validates this peer connections
func (m *PeerConnections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInbound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerConnections) validateInbound(formats strfmt.Registry) error {

	if err := validate.Required("inbound", "body", m.Inbound); err != nil {
		return err
	}

	return nil
}

func (m *PeerConnections) validateOutbound(formats strfmt.Registry) error {

	if err := validate.Required("outbound", "body", m.Outbound); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this peer connections based on context it is used
func (m *PeerConnections) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PeerConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerConnections) UnmarshalBinary(b []byte) error {
	var res PeerConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
