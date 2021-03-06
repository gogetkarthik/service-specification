// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SessionDeleteInput Delete session(s)
//
// swagger:model SessionDeleteInput
type SessionDeleteInput struct {

	// session handles
	SessionHandles []string `json:"sessionHandles"`

	// user ID
	UserID UserID `json:"userID,omitempty"`
}

// Validate validates this session delete input
func (m *SessionDeleteInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionDeleteInput) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := m.UserID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userID")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SessionDeleteInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionDeleteInput) UnmarshalBinary(b []byte) error {
	var res SessionDeleteInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
