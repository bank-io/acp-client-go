// Code generated by go-swagger; DO NOT EDIT.

package openbanking_u_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// RejectFilePaymentConsentSystemReader is a Reader for the RejectFilePaymentConsentSystem structure.
type RejectFilePaymentConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectFilePaymentConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRejectFilePaymentConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRejectFilePaymentConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRejectFilePaymentConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRejectFilePaymentConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRejectFilePaymentConsentSystemOK creates a RejectFilePaymentConsentSystemOK with default headers values
func NewRejectFilePaymentConsentSystemOK() *RejectFilePaymentConsentSystemOK {
	return &RejectFilePaymentConsentSystemOK{}
}

/* RejectFilePaymentConsentSystemOK describes a response with status code 200, with default header values.

ConsentRejected
*/
type RejectFilePaymentConsentSystemOK struct {
	Payload *models.ConsentRejected
}

func (o *RejectFilePaymentConsentSystemOK) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/file-payment-consent/{login}/reject][%d] rejectFilePaymentConsentSystemOK  %+v", 200, o.Payload)
}
func (o *RejectFilePaymentConsentSystemOK) GetPayload() *models.ConsentRejected {
	return o.Payload
}

func (o *RejectFilePaymentConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentRejected)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectFilePaymentConsentSystemUnauthorized creates a RejectFilePaymentConsentSystemUnauthorized with default headers values
func NewRejectFilePaymentConsentSystemUnauthorized() *RejectFilePaymentConsentSystemUnauthorized {
	return &RejectFilePaymentConsentSystemUnauthorized{}
}

/* RejectFilePaymentConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type RejectFilePaymentConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *RejectFilePaymentConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/file-payment-consent/{login}/reject][%d] rejectFilePaymentConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *RejectFilePaymentConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectFilePaymentConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectFilePaymentConsentSystemForbidden creates a RejectFilePaymentConsentSystemForbidden with default headers values
func NewRejectFilePaymentConsentSystemForbidden() *RejectFilePaymentConsentSystemForbidden {
	return &RejectFilePaymentConsentSystemForbidden{}
}

/* RejectFilePaymentConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type RejectFilePaymentConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *RejectFilePaymentConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/file-payment-consent/{login}/reject][%d] rejectFilePaymentConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *RejectFilePaymentConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectFilePaymentConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectFilePaymentConsentSystemNotFound creates a RejectFilePaymentConsentSystemNotFound with default headers values
func NewRejectFilePaymentConsentSystemNotFound() *RejectFilePaymentConsentSystemNotFound {
	return &RejectFilePaymentConsentSystemNotFound{}
}

/* RejectFilePaymentConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type RejectFilePaymentConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *RejectFilePaymentConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/file-payment-consent/{login}/reject][%d] rejectFilePaymentConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *RejectFilePaymentConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RejectFilePaymentConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}