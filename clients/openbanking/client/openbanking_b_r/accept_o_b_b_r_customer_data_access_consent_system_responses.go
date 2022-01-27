// Code generated by go-swagger; DO NOT EDIT.

package openbanking_b_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// AcceptOBBRCustomerDataAccessConsentSystemReader is a Reader for the AcceptOBBRCustomerDataAccessConsentSystem structure.
type AcceptOBBRCustomerDataAccessConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptOBBRCustomerDataAccessConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptOBBRCustomerDataAccessConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAcceptOBBRCustomerDataAccessConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcceptOBBRCustomerDataAccessConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcceptOBBRCustomerDataAccessConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAcceptOBBRCustomerDataAccessConsentSystemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptOBBRCustomerDataAccessConsentSystemOK creates a AcceptOBBRCustomerDataAccessConsentSystemOK with default headers values
func NewAcceptOBBRCustomerDataAccessConsentSystemOK() *AcceptOBBRCustomerDataAccessConsentSystemOK {
	return &AcceptOBBRCustomerDataAccessConsentSystemOK{}
}

/* AcceptOBBRCustomerDataAccessConsentSystemOK describes a response with status code 200, with default header values.

Consent accepted
*/
type AcceptOBBRCustomerDataAccessConsentSystemOK struct {
	Payload *models.ConsentAccepted
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemOK) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/consent/{login}/accept][%d] acceptOBBRCustomerDataAccessConsentSystemOK  %+v", 200, o.Payload)
}
func (o *AcceptOBBRCustomerDataAccessConsentSystemOK) GetPayload() *models.ConsentAccepted {
	return o.Payload
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentAccepted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptOBBRCustomerDataAccessConsentSystemUnauthorized creates a AcceptOBBRCustomerDataAccessConsentSystemUnauthorized with default headers values
func NewAcceptOBBRCustomerDataAccessConsentSystemUnauthorized() *AcceptOBBRCustomerDataAccessConsentSystemUnauthorized {
	return &AcceptOBBRCustomerDataAccessConsentSystemUnauthorized{}
}

/* AcceptOBBRCustomerDataAccessConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type AcceptOBBRCustomerDataAccessConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/consent/{login}/accept][%d] acceptOBBRCustomerDataAccessConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *AcceptOBBRCustomerDataAccessConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptOBBRCustomerDataAccessConsentSystemForbidden creates a AcceptOBBRCustomerDataAccessConsentSystemForbidden with default headers values
func NewAcceptOBBRCustomerDataAccessConsentSystemForbidden() *AcceptOBBRCustomerDataAccessConsentSystemForbidden {
	return &AcceptOBBRCustomerDataAccessConsentSystemForbidden{}
}

/* AcceptOBBRCustomerDataAccessConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type AcceptOBBRCustomerDataAccessConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/consent/{login}/accept][%d] acceptOBBRCustomerDataAccessConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *AcceptOBBRCustomerDataAccessConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptOBBRCustomerDataAccessConsentSystemNotFound creates a AcceptOBBRCustomerDataAccessConsentSystemNotFound with default headers values
func NewAcceptOBBRCustomerDataAccessConsentSystemNotFound() *AcceptOBBRCustomerDataAccessConsentSystemNotFound {
	return &AcceptOBBRCustomerDataAccessConsentSystemNotFound{}
}

/* AcceptOBBRCustomerDataAccessConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type AcceptOBBRCustomerDataAccessConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/consent/{login}/accept][%d] acceptOBBRCustomerDataAccessConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *AcceptOBBRCustomerDataAccessConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptOBBRCustomerDataAccessConsentSystemTooManyRequests creates a AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests with default headers values
func NewAcceptOBBRCustomerDataAccessConsentSystemTooManyRequests() *AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests {
	return &AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests{}
}

/* AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests struct {
	Payload *models.Error
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/consent/{login}/accept][%d] acceptOBBRCustomerDataAccessConsentSystemTooManyRequests  %+v", 429, o.Payload)
}
func (o *AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptOBBRCustomerDataAccessConsentSystemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}