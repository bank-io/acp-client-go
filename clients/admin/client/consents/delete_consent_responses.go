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

// DeleteConsentReader is a Reader for the DeleteConsent structure.
type DeleteConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConsentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConsentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConsentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConsentNoContent creates a DeleteConsentNoContent with default headers values
func NewDeleteConsentNoContent() *DeleteConsentNoContent {
	return &DeleteConsentNoContent{}
}

/* DeleteConsentNoContent describes a response with status code 204, with default header values.

Consent has been deleted
*/
type DeleteConsentNoContent struct {
}

func (o *DeleteConsentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consent}][%d] deleteConsentNoContent ", 204)
}

func (o *DeleteConsentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConsentUnauthorized creates a DeleteConsentUnauthorized with default headers values
func NewDeleteConsentUnauthorized() *DeleteConsentUnauthorized {
	return &DeleteConsentUnauthorized{}
}

/* DeleteConsentUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type DeleteConsentUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteConsentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consent}][%d] deleteConsentUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteConsentUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConsentForbidden creates a DeleteConsentForbidden with default headers values
func NewDeleteConsentForbidden() *DeleteConsentForbidden {
	return &DeleteConsentForbidden{}
}

/* DeleteConsentForbidden describes a response with status code 403, with default header values.

HttpError
*/
type DeleteConsentForbidden struct {
	Payload *models.Error
}

func (o *DeleteConsentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consent}][%d] deleteConsentForbidden  %+v", 403, o.Payload)
}
func (o *DeleteConsentForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConsentNotFound creates a DeleteConsentNotFound with default headers values
func NewDeleteConsentNotFound() *DeleteConsentNotFound {
	return &DeleteConsentNotFound{}
}

/* DeleteConsentNotFound describes a response with status code 404, with default header values.

HttpError
*/
type DeleteConsentNotFound struct {
	Payload *models.Error
}

func (o *DeleteConsentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consent}][%d] deleteConsentNotFound  %+v", 404, o.Payload)
}
func (o *DeleteConsentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteConsentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConsentTooManyRequests creates a DeleteConsentTooManyRequests with default headers values
func NewDeleteConsentTooManyRequests() *DeleteConsentTooManyRequests {
	return &DeleteConsentTooManyRequests{}
}

/* DeleteConsentTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type DeleteConsentTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteConsentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consent}][%d] deleteConsentTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteConsentTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteConsentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
