// Code generated by go-swagger; DO NOT EDIT.

package session_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new session data API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for session data API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SessionData(params *SessionDataParams) (*SessionDataOK, error)

	SessionDataUpdate(params *SessionDataUpdateParams) (*SessionDataUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SessionData gets session data stored in db

  Get session data stored in db
*/
func (a *Client) SessionData(params *SessionDataParams) (*SessionDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSessionDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sessionData",
		Method:             "GET",
		PathPattern:        "/session/data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SessionDataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SessionDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sessionData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SessionDataUpdate updates session data stored in db

  Update session data stored in db
*/
func (a *Client) SessionDataUpdate(params *SessionDataUpdateParams) (*SessionDataUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSessionDataUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sessionDataUpdate",
		Method:             "PUT",
		PathPattern:        "/session/data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SessionDataUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SessionDataUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sessionDataUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}