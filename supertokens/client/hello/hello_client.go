// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hello API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hello API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	HelloDelete(params *HelloDeleteParams) (*HelloDeleteOK, error)

	HelloGet(params *HelloGetParams) (*HelloGetOK, error)

	HelloPost(params *HelloPostParams) (*HelloPostOK, error)

	HelloPut(params *HelloPutParams) (*HelloPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  HelloDelete hellos

  hello
*/
func (a *Client) HelloDelete(params *HelloDeleteParams) (*HelloDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelloDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "helloDelete",
		Method:             "DELETE",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HelloDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelloDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for helloDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HelloGet hellos

  hello
*/
func (a *Client) HelloGet(params *HelloGetParams) (*HelloGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelloGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "helloGet",
		Method:             "GET",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HelloGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelloGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for helloGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HelloPost hellos

  hello
*/
func (a *Client) HelloPost(params *HelloPostParams) (*HelloPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelloPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "helloPost",
		Method:             "POST",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HelloPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelloPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for helloPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HelloPut hellos

  hello
*/
func (a *Client) HelloPut(params *HelloPutParams) (*HelloPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelloPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "helloPut",
		Method:             "PUT",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HelloPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelloPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for helloPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
