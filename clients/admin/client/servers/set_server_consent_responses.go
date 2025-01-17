// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// SetServerConsentReader is a Reader for the SetServerConsent structure.
type SetServerConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetServerConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetServerConsentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetServerConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetServerConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetServerConsentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSetServerConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetServerConsentOK creates a SetServerConsentOK with default headers values
func NewSetServerConsentOK() *SetServerConsentOK {
	return &SetServerConsentOK{}
}

/* SetServerConsentOK describes a response with status code 200, with default header values.

ServerConsent
*/
type SetServerConsentOK struct {
	Payload *models.ServerConsent
}

func (o *SetServerConsentOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/server-consent][%d] setServerConsentOK  %+v", 200, o.Payload)
}
func (o *SetServerConsentOK) GetPayload() *models.ServerConsent {
	return o.Payload
}

func (o *SetServerConsentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerConsent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetServerConsentUnauthorized creates a SetServerConsentUnauthorized with default headers values
func NewSetServerConsentUnauthorized() *SetServerConsentUnauthorized {
	return &SetServerConsentUnauthorized{}
}

/* SetServerConsentUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type SetServerConsentUnauthorized struct {
	Payload *models.Error
}

func (o *SetServerConsentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/server-consent][%d] setServerConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *SetServerConsentUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetServerConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetServerConsentForbidden creates a SetServerConsentForbidden with default headers values
func NewSetServerConsentForbidden() *SetServerConsentForbidden {
	return &SetServerConsentForbidden{}
}

/* SetServerConsentForbidden describes a response with status code 403, with default header values.

HttpError
*/
type SetServerConsentForbidden struct {
	Payload *models.Error
}

func (o *SetServerConsentForbidden) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/server-consent][%d] setServerConsentForbidden  %+v", 403, o.Payload)
}
func (o *SetServerConsentForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetServerConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetServerConsentNotFound creates a SetServerConsentNotFound with default headers values
func NewSetServerConsentNotFound() *SetServerConsentNotFound {
	return &SetServerConsentNotFound{}
}

/* SetServerConsentNotFound describes a response with status code 404, with default header values.

HttpError
*/
type SetServerConsentNotFound struct {
	Payload *models.Error
}

func (o *SetServerConsentNotFound) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/server-consent][%d] setServerConsentNotFound  %+v", 404, o.Payload)
}
func (o *SetServerConsentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetServerConsentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetServerConsentTooManyRequests creates a SetServerConsentTooManyRequests with default headers values
func NewSetServerConsentTooManyRequests() *SetServerConsentTooManyRequests {
	return &SetServerConsentTooManyRequests{}
}

/* SetServerConsentTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type SetServerConsentTooManyRequests struct {
	Payload *models.Error
}

func (o *SetServerConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/server-consent][%d] setServerConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *SetServerConsentTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetServerConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
