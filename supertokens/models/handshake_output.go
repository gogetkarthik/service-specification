// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HandshakeOutput Handshake details
//
// swagger:model HandshakeOutput
type HandshakeOutput struct {

	// access token blacklisting enabled
	AccessTokenBlacklistingEnabled bool `json:"accessTokenBlacklistingEnabled,omitempty"`

	// access token path
	AccessTokenPath string `json:"accessTokenPath,omitempty"`

	// cookie domain
	CookieDomain string `json:"cookieDomain,omitempty"`

	// cookie same site
	CookieSameSite string `json:"cookieSameSite,omitempty"`

	// cookie secure
	CookieSecure bool `json:"cookieSecure,omitempty"`

	// enable anti csrf
	EnableAntiCsrf bool `json:"enableAntiCsrf,omitempty"`

	// id refresh token path
	IDRefreshTokenPath string `json:"idRefreshTokenPath,omitempty"`

	// jwt signing public key
	JwtSigningPublicKey JWTSigningPublicKey `json:"jwtSigningPublicKey,omitempty"`

	// jwt signing public key expiry time
	JwtSigningPublicKeyExpiryTime int64 `json:"jwtSigningPublicKeyExpiryTime,omitempty"`

	// refresh token path
	RefreshTokenPath string `json:"refreshTokenPath,omitempty"`

	// status
	Status Status `json:"status,omitempty"`
}

// Validate validates this handshake output
func (m *HandshakeOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJwtSigningPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandshakeOutput) validateJwtSigningPublicKey(formats strfmt.Registry) error {

	if swag.IsZero(m.JwtSigningPublicKey) { // not required
		return nil
	}

	if err := m.JwtSigningPublicKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("jwtSigningPublicKey")
		}
		return err
	}

	return nil
}

func (m *HandshakeOutput) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HandshakeOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandshakeOutput) UnmarshalBinary(b []byte) error {
	var res HandshakeOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
