// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/service-specification/models"
)

// HelloPutReader is a Reader for the HelloPut structure.
type HelloPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelloPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelloPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHelloPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHelloPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewHelloPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHelloPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHelloPutOK creates a HelloPutOK with default headers values
func NewHelloPutOK() *HelloPutOK {
	return &HelloPutOK{}
}

/*HelloPutOK handles this case with default header values.

Hello
*/
type HelloPutOK struct {
	Payload models.HelloOutput
}

func (o *HelloPutOK) Error() string {
	return fmt.Sprintf("[PUT /hello][%d] helloPutOK  %+v", 200, o.Payload)
}

func (o *HelloPutOK) GetPayload() models.HelloOutput {
	return o.Payload
}

func (o *HelloPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelloPutBadRequest creates a HelloPutBadRequest with default headers values
func NewHelloPutBadRequest() *HelloPutBadRequest {
	return &HelloPutBadRequest{}
}

/*HelloPutBadRequest handles this case with default header values.

Invalid input
*/
type HelloPutBadRequest struct {
	Payload models.Error
}

func (o *HelloPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /hello][%d] helloPutBadRequest  %+v", 400, o.Payload)
}

func (o *HelloPutBadRequest) GetPayload() models.Error {
	return o.Payload
}

func (o *HelloPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelloPutNotFound creates a HelloPutNotFound with default headers values
func NewHelloPutNotFound() *HelloPutNotFound {
	return &HelloPutNotFound{}
}

/*HelloPutNotFound handles this case with default header values.

The specified resource was not found
*/
type HelloPutNotFound struct {
	Payload models.Error
}

func (o *HelloPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /hello][%d] helloPutNotFound  %+v", 404, o.Payload)
}

func (o *HelloPutNotFound) GetPayload() models.Error {
	return o.Payload
}

func (o *HelloPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelloPutMethodNotAllowed creates a HelloPutMethodNotAllowed with default headers values
func NewHelloPutMethodNotAllowed() *HelloPutMethodNotAllowed {
	return &HelloPutMethodNotAllowed{}
}

/*HelloPutMethodNotAllowed handles this case with default header values.

Method not souported
*/
type HelloPutMethodNotAllowed struct {
	Payload models.Error
}

func (o *HelloPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /hello][%d] helloPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *HelloPutMethodNotAllowed) GetPayload() models.Error {
	return o.Payload
}

func (o *HelloPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelloPutInternalServerError creates a HelloPutInternalServerError with default headers values
func NewHelloPutInternalServerError() *HelloPutInternalServerError {
	return &HelloPutInternalServerError{}
}

/*HelloPutInternalServerError handles this case with default header values.

Internal server error
*/
type HelloPutInternalServerError struct {
	Payload models.Error
}

func (o *HelloPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /hello][%d] helloPutInternalServerError  %+v", 500, o.Payload)
}

func (o *HelloPutInternalServerError) GetPayload() models.Error {
	return o.Payload
}

func (o *HelloPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
