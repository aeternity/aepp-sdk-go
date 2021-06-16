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

// PeerDetails peer details
//
// swagger:model PeerDetails
type PeerDetails struct {

	// Unix timestamp of when the peer was first pinged
	// Required: true
	FirstSeen *uint32 `json:"first_seen"`

	// The genesis hash the remote node reports
	// Required: true
	GenesisHash *string `json:"genesis_hash"`

	// Hostname of peer
	// Required: true
	Host *string `json:"host"`

	// Unix timestamp of when the peer was last pinged
	// Required: true
	LastSeen *uint32 `json:"last_seen"`

	// network id
	NetworkID string `json:"network_id,omitempty"`

	// node os
	NodeOs string `json:"node_os,omitempty"`

	// node revision
	NodeRevision string `json:"node_revision,omitempty"`

	// node vendor
	NodeVendor string `json:"node_vendor,omitempty"`

	// node version
	NodeVersion string `json:"node_version,omitempty"`

	// Port of peer
	// Required: true
	Port *uint32 `json:"port"`

	// The total top difficulty the node reports
	// Required: true
	TopDifficulty *uint64 `json:"top_difficulty"`

	// The top hash the remote node reports
	// Required: true
	TopHash *string `json:"top_hash"`
}

// Validate validates this peer details
func (m *PeerDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenesisHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopDifficulty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerDetails) validateFirstSeen(formats strfmt.Registry) error {

	if err := validate.Required("first_seen", "body", m.FirstSeen); err != nil {
		return err
	}

	return nil
}

func (m *PeerDetails) validateGenesisHash(formats strfmt.Registry) error {

	if err := validate.Required("genesis_hash", "body", m.GenesisHash); err != nil {
		return err
	}

	return nil
}

func (m *PeerDetails) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *PeerDetails) validateLastSeen(formats strfmt.Registry) error {

	if err := validate.Required("last_seen", "body", m.LastSeen); err != nil {
		return err
	}

	return nil
}

func (m *PeerDetails) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *PeerDetails) validateTopDifficulty(formats strfmt.Registry) error {

	if err := validate.Required("top_difficulty", "body", m.TopDifficulty); err != nil {
		return err
	}

	return nil
}

func (m *PeerDetails) validateTopHash(formats strfmt.Registry) error {

	if err := validate.Required("top_hash", "body", m.TopHash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this peer details based on context it is used
func (m *PeerDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PeerDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerDetails) UnmarshalBinary(b []byte) error {
	var res PeerDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}