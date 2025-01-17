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

// RejectOBBRCustomerPaymentConsentSystemReader is a Reader for the RejectOBBRCustomerPaymentConsentSystem structure.
type RejectOBBRCustomerPaymentConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectOBBRCustomerPaymentConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRejectOBBRCustomerPaymentConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRejectOBBRCustomerPaymentConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRejectOBBRCustomerPaymentConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRejectOBBRCustomerPaymentConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRejectOBBRCustomerPaymentConsentSystemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRejectOBBRCustomerPaymentConsentSystemOK creates a RejectOBBRCustomerPaymentConsentSystemOK with default headers values
func NewRejectOBBRCustomerPaymentConsentSystemOK() *RejectOBBRCustomerPaymentConsentSystemOK {
	return &RejectOBBRCustomerPaymentConsentSystemOK{}
}

/* RejectOBBRCustomerPaymentConsentSystemOK describes a response with status code 200, with default header values.

Consent rejected
*/
type RejectOBBRCustomerPaymentConsentSystemOK struct {
	Payload *models.ConsentRejected
}

func (o *RejectOBBRCustomerPaymentConsentSystemOK) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/payment/{login}/reject][%d] rejectOBBRCustomerPaymentConsentSystemOK  %+v", 200, o.Payload)
}
func (o *RejectOBBRCustomerPaymentConsentSystemOK) GetPayload() *models.ConsentRejected {
	return o.Payload
}

func (o *RejectOBBRCustomerPaymentConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentRejected)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectOBBRCustomerPaymentConsentSystemUnauthorized creates a RejectOBBRCustomerPaymentConsentSystemUnauthorized with default headers values
func NewRejectOBBRCustomerPaymentConsentSystemUnauthorized() *RejectOBBRCustomerPaymentConsentSystemUnauthorized {
	return &RejectOBBRCustomerPaymentConsentSystemUnauthorized{}
}

/* RejectOBBRCustomerPaymentConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type RejectOBBRCustomerPaymentConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *RejectOBBRCustomerPaymentConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/payment/{login}/reject][%d] rejectOBBRCustomerPaymentConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *RejectOBBRCustomerPaymentConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectOBBRCustomerPaymentConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectOBBRCustomerPaymentConsentSystemForbidden creates a RejectOBBRCustomerPaymentConsentSystemForbidden with default headers values
func NewRejectOBBRCustomerPaymentConsentSystemForbidden() *RejectOBBRCustomerPaymentConsentSystemForbidden {
	return &RejectOBBRCustomerPaymentConsentSystemForbidden{}
}

/* RejectOBBRCustomerPaymentConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type RejectOBBRCustomerPaymentConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *RejectOBBRCustomerPaymentConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/payment/{login}/reject][%d] rejectOBBRCustomerPaymentConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *RejectOBBRCustomerPaymentConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectOBBRCustomerPaymentConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectOBBRCustomerPaymentConsentSystemNotFound creates a RejectOBBRCustomerPaymentConsentSystemNotFound with default headers values
func NewRejectOBBRCustomerPaymentConsentSystemNotFound() *RejectOBBRCustomerPaymentConsentSystemNotFound {
	return &RejectOBBRCustomerPaymentConsentSystemNotFound{}
}

/* RejectOBBRCustomerPaymentConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type RejectOBBRCustomerPaymentConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *RejectOBBRCustomerPaymentConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/payment/{login}/reject][%d] rejectOBBRCustomerPaymentConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *RejectOBBRCustomerPaymentConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectOBBRCustomerPaymentConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectOBBRCustomerPaymentConsentSystemTooManyRequests creates a RejectOBBRCustomerPaymentConsentSystemTooManyRequests with default headers values
func NewRejectOBBRCustomerPaymentConsentSystemTooManyRequests() *RejectOBBRCustomerPaymentConsentSystemTooManyRequests {
	return &RejectOBBRCustomerPaymentConsentSystemTooManyRequests{}
}

/* RejectOBBRCustomerPaymentConsentSystemTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type RejectOBBRCustomerPaymentConsentSystemTooManyRequests struct {
	Payload *models.Error
}

func (o *RejectOBBRCustomerPaymentConsentSystemTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /open-banking-brasil/payment/{login}/reject][%d] rejectOBBRCustomerPaymentConsentSystemTooManyRequests  %+v", 429, o.Payload)
}
func (o *RejectOBBRCustomerPaymentConsentSystemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectOBBRCustomerPaymentConsentSystemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
