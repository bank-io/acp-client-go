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

// GetFilePaymentConsentSystemReader is a Reader for the GetFilePaymentConsentSystem structure.
type GetFilePaymentConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilePaymentConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilePaymentConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFilePaymentConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilePaymentConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilePaymentConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFilePaymentConsentSystemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilePaymentConsentSystemOK creates a GetFilePaymentConsentSystemOK with default headers values
func NewGetFilePaymentConsentSystemOK() *GetFilePaymentConsentSystemOK {
	return &GetFilePaymentConsentSystemOK{}
}

/* GetFilePaymentConsentSystemOK describes a response with status code 200, with default header values.

GetFilePaymentConsentResponse
*/
type GetFilePaymentConsentSystemOK struct {
	Payload *models.GetFilePaymentConsentResponse
}

func (o *GetFilePaymentConsentSystemOK) Error() string {
	return fmt.Sprintf("[GET /open-banking/file-payment-consent/{login}][%d] getFilePaymentConsentSystemOK  %+v", 200, o.Payload)
}
func (o *GetFilePaymentConsentSystemOK) GetPayload() *models.GetFilePaymentConsentResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFilePaymentConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentSystemUnauthorized creates a GetFilePaymentConsentSystemUnauthorized with default headers values
func NewGetFilePaymentConsentSystemUnauthorized() *GetFilePaymentConsentSystemUnauthorized {
	return &GetFilePaymentConsentSystemUnauthorized{}
}

/* GetFilePaymentConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetFilePaymentConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *GetFilePaymentConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /open-banking/file-payment-consent/{login}][%d] getFilePaymentConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFilePaymentConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFilePaymentConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentSystemForbidden creates a GetFilePaymentConsentSystemForbidden with default headers values
func NewGetFilePaymentConsentSystemForbidden() *GetFilePaymentConsentSystemForbidden {
	return &GetFilePaymentConsentSystemForbidden{}
}

/* GetFilePaymentConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetFilePaymentConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *GetFilePaymentConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[GET /open-banking/file-payment-consent/{login}][%d] getFilePaymentConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *GetFilePaymentConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFilePaymentConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentSystemNotFound creates a GetFilePaymentConsentSystemNotFound with default headers values
func NewGetFilePaymentConsentSystemNotFound() *GetFilePaymentConsentSystemNotFound {
	return &GetFilePaymentConsentSystemNotFound{}
}

/* GetFilePaymentConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetFilePaymentConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *GetFilePaymentConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /open-banking/file-payment-consent/{login}][%d] getFilePaymentConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *GetFilePaymentConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFilePaymentConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentSystemTooManyRequests creates a GetFilePaymentConsentSystemTooManyRequests with default headers values
func NewGetFilePaymentConsentSystemTooManyRequests() *GetFilePaymentConsentSystemTooManyRequests {
	return &GetFilePaymentConsentSystemTooManyRequests{}
}

/* GetFilePaymentConsentSystemTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetFilePaymentConsentSystemTooManyRequests struct {
	Payload *models.Error
}

func (o *GetFilePaymentConsentSystemTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /open-banking/file-payment-consent/{login}][%d] getFilePaymentConsentSystemTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFilePaymentConsentSystemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFilePaymentConsentSystemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
