// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHelloDeleteParams creates a new HelloDeleteParams object
// with the default values initialized.
func NewHelloDeleteParams() *HelloDeleteParams {

	return &HelloDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHelloDeleteParamsWithTimeout creates a new HelloDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHelloDeleteParamsWithTimeout(timeout time.Duration) *HelloDeleteParams {

	return &HelloDeleteParams{

		timeout: timeout,
	}
}

// NewHelloDeleteParamsWithContext creates a new HelloDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewHelloDeleteParamsWithContext(ctx context.Context) *HelloDeleteParams {

	return &HelloDeleteParams{

		Context: ctx,
	}
}

// NewHelloDeleteParamsWithHTTPClient creates a new HelloDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHelloDeleteParamsWithHTTPClient(client *http.Client) *HelloDeleteParams {

	return &HelloDeleteParams{
		HTTPClient: client,
	}
}

/*HelloDeleteParams contains all the parameters to send to the API endpoint
for the hello delete operation typically these are written to a http.Request
*/
type HelloDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hello delete params
func (o *HelloDeleteParams) WithTimeout(timeout time.Duration) *HelloDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hello delete params
func (o *HelloDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hello delete params
func (o *HelloDeleteParams) WithContext(ctx context.Context) *HelloDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hello delete params
func (o *HelloDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hello delete params
func (o *HelloDeleteParams) WithHTTPClient(client *http.Client) *HelloDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hello delete params
func (o *HelloDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HelloDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
