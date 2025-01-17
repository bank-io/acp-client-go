// Code generated by go-swagger; DO NOT EDIT.

package vanity_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// DeleteVanityDomainReader is a Reader for the DeleteVanityDomain structure.
type DeleteVanityDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVanityDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteVanityDomainNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteVanityDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVanityDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVanityDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteVanityDomainTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVanityDomainNoContent creates a DeleteVanityDomainNoContent with default headers values
func NewDeleteVanityDomainNoContent() *DeleteVanityDomainNoContent {
	return &DeleteVanityDomainNoContent{}
}

/* DeleteVanityDomainNoContent describes a response with status code 204, with default header values.

Vanity domain has been deleted
*/
type DeleteVanityDomainNoContent struct {
}

func (o *DeleteVanityDomainNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vanity-domains][%d] deleteVanityDomainNoContent ", 204)
}

func (o *DeleteVanityDomainNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVanityDomainUnauthorized creates a DeleteVanityDomainUnauthorized with default headers values
func NewDeleteVanityDomainUnauthorized() *DeleteVanityDomainUnauthorized {
	return &DeleteVanityDomainUnauthorized{}
}

/* DeleteVanityDomainUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type DeleteVanityDomainUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteVanityDomainUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /vanity-domains][%d] deleteVanityDomainUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteVanityDomainUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVanityDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVanityDomainForbidden creates a DeleteVanityDomainForbidden with default headers values
func NewDeleteVanityDomainForbidden() *DeleteVanityDomainForbidden {
	return &DeleteVanityDomainForbidden{}
}

/* DeleteVanityDomainForbidden describes a response with status code 403, with default header values.

HttpError
*/
type DeleteVanityDomainForbidden struct {
	Payload *models.Error
}

func (o *DeleteVanityDomainForbidden) Error() string {
	return fmt.Sprintf("[DELETE /vanity-domains][%d] deleteVanityDomainForbidden  %+v", 403, o.Payload)
}
func (o *DeleteVanityDomainForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVanityDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVanityDomainNotFound creates a DeleteVanityDomainNotFound with default headers values
func NewDeleteVanityDomainNotFound() *DeleteVanityDomainNotFound {
	return &DeleteVanityDomainNotFound{}
}

/* DeleteVanityDomainNotFound describes a response with status code 404, with default header values.

HttpError
*/
type DeleteVanityDomainNotFound struct {
	Payload *models.Error
}

func (o *DeleteVanityDomainNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vanity-domains][%d] deleteVanityDomainNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVanityDomainNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVanityDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVanityDomainTooManyRequests creates a DeleteVanityDomainTooManyRequests with default headers values
func NewDeleteVanityDomainTooManyRequests() *DeleteVanityDomainTooManyRequests {
	return &DeleteVanityDomainTooManyRequests{}
}

/* DeleteVanityDomainTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type DeleteVanityDomainTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteVanityDomainTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /vanity-domains][%d] deleteVanityDomainTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteVanityDomainTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVanityDomainTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
