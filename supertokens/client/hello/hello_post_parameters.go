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

// NewHelloPostParams creates a new HelloPostParams object
// with the default values initialized.
func NewHelloPostParams() *HelloPostParams {

	return &HelloPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHelloPostParamsWithTimeout creates a new HelloPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHelloPostParamsWithTimeout(timeout time.Duration) *HelloPostParams {

	return &HelloPostParams{

		timeout: timeout,
	}
}

// NewHelloPostParamsWithContext creates a new HelloPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewHelloPostParamsWithContext(ctx context.Context) *HelloPostParams {

	return &HelloPostParams{

		Context: ctx,
	}
}

// NewHelloPostParamsWithHTTPClient creates a new HelloPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHelloPostParamsWithHTTPClient(client *http.Client) *HelloPostParams {

	return &HelloPostParams{
		HTTPClient: client,
	}
}

/*HelloPostParams contains all the parameters to send to the API endpoint
for the hello post operation typically these are written to a http.Request
*/
type HelloPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hello post params
func (o *HelloPostParams) WithTimeout(timeout time.Duration) *HelloPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hello post params
func (o *HelloPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hello post params
func (o *HelloPostParams) WithContext(ctx context.Context) *HelloPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hello post params
func (o *HelloPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hello post params
func (o *HelloPostParams) WithHTTPClient(client *http.Client) *HelloPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hello post params
func (o *HelloPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HelloPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
