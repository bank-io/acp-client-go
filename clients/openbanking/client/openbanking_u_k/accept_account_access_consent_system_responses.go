// Code generated by go-swagger; DO NOT EDIT.

package openbanking_u_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// AcceptAccountAccessConsentSystemReader is a Reader for the AcceptAccountAccessConsentSystem structure.
type AcceptAccountAccessConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptAccountAccessConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptAccountAccessConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAcceptAccountAccessConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcceptAccountAccessConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcceptAccountAccessConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAcceptAccountAccessConsentSystemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptAccountAccessConsentSystemOK creates a AcceptAccountAccessConsentSystemOK with default headers values
func NewAcceptAccountAccessConsentSystemOK() *AcceptAccountAccessConsentSystemOK {
	return &AcceptAccountAccessConsentSystemOK{}
}

/* AcceptAccountAccessConsentSystemOK describes a response with status code 200, with default header values.

Consent accepted
*/
type AcceptAccountAccessConsentSystemOK struct {
	Payload *models.ConsentAccepted
}

func (o *AcceptAccountAccessConsentSystemOK) Error() string {
	return fmt.Sprintf("[POST /open-banking/account-access-consent/{login}/accept][%d] acceptAccountAccessConsentSystemOK  %+v", 200, o.Payload)
}
func (o *AcceptAccountAccessConsentSystemOK) GetPayload() *models.ConsentAccepted {
	return o.Payload
}

func (o *AcceptAccountAccessConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentAccepted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptAccountAccessConsentSystemUnauthorized creates a AcceptAccountAccessConsentSystemUnauthorized with default headers values
func NewAcceptAccountAccessConsentSystemUnauthorized() *AcceptAccountAccessConsentSystemUnauthorized {
	return &AcceptAccountAccessConsentSystemUnauthorized{}
}

/* AcceptAccountAccessConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type AcceptAccountAccessConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *AcceptAccountAccessConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/account-access-consent/{login}/accept][%d] acceptAccountAccessConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *AcceptAccountAccessConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptAccountAccessConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptAccountAccessConsentSystemForbidden creates a AcceptAccountAccessConsentSystemForbidden with default headers values
func NewAcceptAccountAccessConsentSystemForbidden() *AcceptAccountAccessConsentSystemForbidden {
	return &AcceptAccountAccessConsentSystemForbidden{}
}

/* AcceptAccountAccessConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type AcceptAccountAccessConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *AcceptAccountAccessConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking/account-access-consent/{login}/accept][%d] acceptAccountAccessConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *AcceptAccountAccessConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptAccountAccessConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptAccountAccessConsentSystemNotFound creates a AcceptAccountAccessConsentSystemNotFound with default headers values
func NewAcceptAccountAccessConsentSystemNotFound() *AcceptAccountAccessConsentSystemNotFound {
	return &AcceptAccountAccessConsentSystemNotFound{}
}

/* AcceptAccountAccessConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type AcceptAccountAccessConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *AcceptAccountAccessConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[POST /open-banking/account-access-consent/{login}/accept][%d] acceptAccountAccessConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *AcceptAccountAccessConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptAccountAccessConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptAccountAccessConsentSystemTooManyRequests creates a AcceptAccountAccessConsentSystemTooManyRequests with default headers values
func NewAcceptAccountAccessConsentSystemTooManyRequests() *AcceptAccountAccessConsentSystemTooManyRequests {
	return &AcceptAccountAccessConsentSystemTooManyRequests{}
}

/* AcceptAccountAccessConsentSystemTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type AcceptAccountAccessConsentSystemTooManyRequests struct {
	Payload *models.Error
}

func (o *AcceptAccountAccessConsentSystemTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/account-access-consent/{login}/accept][%d] acceptAccountAccessConsentSystemTooManyRequests  %+v", 429, o.Payload)
}
func (o *AcceptAccountAccessConsentSystemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptAccountAccessConsentSystemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
