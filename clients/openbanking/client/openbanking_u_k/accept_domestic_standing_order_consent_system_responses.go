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

// AcceptDomesticStandingOrderConsentSystemReader is a Reader for the AcceptDomesticStandingOrderConsentSystem structure.
type AcceptDomesticStandingOrderConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptDomesticStandingOrderConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptDomesticStandingOrderConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAcceptDomesticStandingOrderConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcceptDomesticStandingOrderConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcceptDomesticStandingOrderConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAcceptDomesticStandingOrderConsentSystemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptDomesticStandingOrderConsentSystemOK creates a AcceptDomesticStandingOrderConsentSystemOK with default headers values
func NewAcceptDomesticStandingOrderConsentSystemOK() *AcceptDomesticStandingOrderConsentSystemOK {
	return &AcceptDomesticStandingOrderConsentSystemOK{}
}

/* AcceptDomesticStandingOrderConsentSystemOK describes a response with status code 200, with default header values.

Consent accepted
*/
type AcceptDomesticStandingOrderConsentSystemOK struct {
	Payload *models.ConsentAccepted
}

func (o *AcceptDomesticStandingOrderConsentSystemOK) Error() string {
	return fmt.Sprintf("[POST /open-banking/domestic-standing-order-consent/{login}/accept][%d] acceptDomesticStandingOrderConsentSystemOK  %+v", 200, o.Payload)
}
func (o *AcceptDomesticStandingOrderConsentSystemOK) GetPayload() *models.ConsentAccepted {
	return o.Payload
}

func (o *AcceptDomesticStandingOrderConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentAccepted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptDomesticStandingOrderConsentSystemUnauthorized creates a AcceptDomesticStandingOrderConsentSystemUnauthorized with default headers values
func NewAcceptDomesticStandingOrderConsentSystemUnauthorized() *AcceptDomesticStandingOrderConsentSystemUnauthorized {
	return &AcceptDomesticStandingOrderConsentSystemUnauthorized{}
}

/* AcceptDomesticStandingOrderConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type AcceptDomesticStandingOrderConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *AcceptDomesticStandingOrderConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking/domestic-standing-order-consent/{login}/accept][%d] acceptDomesticStandingOrderConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *AcceptDomesticStandingOrderConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptDomesticStandingOrderConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptDomesticStandingOrderConsentSystemForbidden creates a AcceptDomesticStandingOrderConsentSystemForbidden with default headers values
func NewAcceptDomesticStandingOrderConsentSystemForbidden() *AcceptDomesticStandingOrderConsentSystemForbidden {
	return &AcceptDomesticStandingOrderConsentSystemForbidden{}
}

/* AcceptDomesticStandingOrderConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type AcceptDomesticStandingOrderConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *AcceptDomesticStandingOrderConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking/domestic-standing-order-consent/{login}/accept][%d] acceptDomesticStandingOrderConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *AcceptDomesticStandingOrderConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptDomesticStandingOrderConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptDomesticStandingOrderConsentSystemNotFound creates a AcceptDomesticStandingOrderConsentSystemNotFound with default headers values
func NewAcceptDomesticStandingOrderConsentSystemNotFound() *AcceptDomesticStandingOrderConsentSystemNotFound {
	return &AcceptDomesticStandingOrderConsentSystemNotFound{}
}

/* AcceptDomesticStandingOrderConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type AcceptDomesticStandingOrderConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *AcceptDomesticStandingOrderConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[POST /open-banking/domestic-standing-order-consent/{login}/accept][%d] acceptDomesticStandingOrderConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *AcceptDomesticStandingOrderConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptDomesticStandingOrderConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptDomesticStandingOrderConsentSystemTooManyRequests creates a AcceptDomesticStandingOrderConsentSystemTooManyRequests with default headers values
func NewAcceptDomesticStandingOrderConsentSystemTooManyRequests() *AcceptDomesticStandingOrderConsentSystemTooManyRequests {
	return &AcceptDomesticStandingOrderConsentSystemTooManyRequests{}
}

/* AcceptDomesticStandingOrderConsentSystemTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type AcceptDomesticStandingOrderConsentSystemTooManyRequests struct {
	Payload *models.Error
}

func (o *AcceptDomesticStandingOrderConsentSystemTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking/domestic-standing-order-consent/{login}/accept][%d] acceptDomesticStandingOrderConsentSystemTooManyRequests  %+v", 429, o.Payload)
}
func (o *AcceptDomesticStandingOrderConsentSystemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptDomesticStandingOrderConsentSystemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
