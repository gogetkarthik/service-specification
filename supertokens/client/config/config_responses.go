// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gogetkarthik/service-specification/supertokens/models"
)

// ConfigReader is a Reader for the Config structure.
type ConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewConfigMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConfigOK creates a ConfigOK with default headers values
func NewConfigOK() *ConfigOK {
	return &ConfigOK{}
}

/*ConfigOK handles this case with default header values.

Config destails
*/
type ConfigOK struct {
	Payload *models.ConfigOutput
}

func (o *ConfigOK) Error() string {
	return fmt.Sprintf("[GET /config][%d] configOK  %+v", 200, o.Payload)
}

func (o *ConfigOK) GetPayload() *models.ConfigOutput {
	return o.Payload
}

func (o *ConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigBadRequest creates a ConfigBadRequest with default headers values
func NewConfigBadRequest() *ConfigBadRequest {
	return &ConfigBadRequest{}
}

/*ConfigBadRequest handles this case with default header values.

Invalid input
*/
type ConfigBadRequest struct {
	Payload models.Error
}

func (o *ConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /config][%d] configBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigBadRequest) GetPayload() models.Error {
	return o.Payload
}

func (o *ConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigNotFound creates a ConfigNotFound with default headers values
func NewConfigNotFound() *ConfigNotFound {
	return &ConfigNotFound{}
}

/*ConfigNotFound handles this case with default header values.

The specified resource was not found
*/
type ConfigNotFound struct {
	Payload models.Error
}

func (o *ConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /config][%d] configNotFound  %+v", 404, o.Payload)
}

func (o *ConfigNotFound) GetPayload() models.Error {
	return o.Payload
}

func (o *ConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigMethodNotAllowed creates a ConfigMethodNotAllowed with default headers values
func NewConfigMethodNotAllowed() *ConfigMethodNotAllowed {
	return &ConfigMethodNotAllowed{}
}

/*ConfigMethodNotAllowed handles this case with default header values.

Method not souported
*/
type ConfigMethodNotAllowed struct {
	Payload models.Error
}

func (o *ConfigMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /config][%d] configMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ConfigMethodNotAllowed) GetPayload() models.Error {
	return o.Payload
}

func (o *ConfigMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigInternalServerError creates a ConfigInternalServerError with default headers values
func NewConfigInternalServerError() *ConfigInternalServerError {
	return &ConfigInternalServerError{}
}

/*ConfigInternalServerError handles this case with default header values.

Internal server error
*/
type ConfigInternalServerError struct {
	Payload models.Error
}

func (o *ConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config][%d] configInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigInternalServerError) GetPayload() models.Error {
	return o.Payload
}

func (o *ConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
