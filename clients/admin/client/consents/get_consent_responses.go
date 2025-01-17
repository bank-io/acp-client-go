// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// GetConsentReader is a Reader for the GetConsent structure.
type GetConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConsentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConsentOK creates a GetConsentOK with default headers values
func NewGetConsentOK() *GetConsentOK {
	return &GetConsentOK{}
}

/* GetConsentOK describes a response with status code 200, with default header values.

Consent
*/
type GetConsentOK struct {
	Payload *models.Consent
}

func (o *GetConsentOK) Error() string {
	return fmt.Sprintf("[GET /consents/{consent}][%d] getConsentOK  %+v", 200, o.Payload)
}
func (o *GetConsentOK) GetPayload() *models.Consent {
	return o.Payload
}

func (o *GetConsentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Consent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentUnauthorized creates a GetConsentUnauthorized with default headers values
func NewGetConsentUnauthorized() *GetConsentUnauthorized {
	return &GetConsentUnauthorized{}
}

/* GetConsentUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetConsentUnauthorized struct {
	Payload *models.Error
}

func (o *GetConsentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /consents/{consent}][%d] getConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *GetConsentUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentForbidden creates a GetConsentForbidden with default headers values
func NewGetConsentForbidden() *GetConsentForbidden {
	return &GetConsentForbidden{}
}

/* GetConsentForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetConsentForbidden struct {
	Payload *models.Error
}

func (o *GetConsentForbidden) Error() string {
	return fmt.Sprintf("[GET /consents/{consent}][%d] getConsentForbidden  %+v", 403, o.Payload)
}
func (o *GetConsentForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentNotFound creates a GetConsentNotFound with default headers values
func NewGetConsentNotFound() *GetConsentNotFound {
	return &GetConsentNotFound{}
}

/* GetConsentNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetConsentNotFound struct {
	Payload *models.Error
}

func (o *GetConsentNotFound) Error() string {
	return fmt.Sprintf("[GET /consents/{consent}][%d] getConsentNotFound  %+v", 404, o.Payload)
}
func (o *GetConsentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentTooManyRequests creates a GetConsentTooManyRequests with default headers values
func NewGetConsentTooManyRequests() *GetConsentTooManyRequests {
	return &GetConsentTooManyRequests{}
}

/* GetConsentTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetConsentTooManyRequests struct {
	Payload *models.Error
}

func (o *GetConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /consents/{consent}][%d] getConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetConsentTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
