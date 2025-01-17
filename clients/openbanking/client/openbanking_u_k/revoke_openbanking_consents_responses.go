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

// RevokeOpenbankingConsentsReader is a Reader for the RevokeOpenbankingConsents structure.
type RevokeOpenbankingConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeOpenbankingConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeOpenbankingConsentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeOpenbankingConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRevokeOpenbankingConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeOpenbankingConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeOpenbankingConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRevokeOpenbankingConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeOpenbankingConsentsOK creates a RevokeOpenbankingConsentsOK with default headers values
func NewRevokeOpenbankingConsentsOK() *RevokeOpenbankingConsentsOK {
	return &RevokeOpenbankingConsentsOK{}
}

/* RevokeOpenbankingConsentsOK describes a response with status code 200, with default header values.

ConsentsRemovedResponse
*/
type RevokeOpenbankingConsentsOK struct {
	Payload *models.ConsentsRemovedResponse
}

func (o *RevokeOpenbankingConsentsOK) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}/open-banking/consents][%d] revokeOpenbankingConsentsOK  %+v", 200, o.Payload)
}
func (o *RevokeOpenbankingConsentsOK) GetPayload() *models.ConsentsRemovedResponse {
	return o.Payload
}

func (o *RevokeOpenbankingConsentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentsRemovedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentsBadRequest creates a RevokeOpenbankingConsentsBadRequest with default headers values
func NewRevokeOpenbankingConsentsBadRequest() *RevokeOpenbankingConsentsBadRequest {
	return &RevokeOpenbankingConsentsBadRequest{}
}

/* RevokeOpenbankingConsentsBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type RevokeOpenbankingConsentsBadRequest struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}/open-banking/consents][%d] revokeOpenbankingConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *RevokeOpenbankingConsentsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentsUnauthorized creates a RevokeOpenbankingConsentsUnauthorized with default headers values
func NewRevokeOpenbankingConsentsUnauthorized() *RevokeOpenbankingConsentsUnauthorized {
	return &RevokeOpenbankingConsentsUnauthorized{}
}

/* RevokeOpenbankingConsentsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type RevokeOpenbankingConsentsUnauthorized struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}/open-banking/consents][%d] revokeOpenbankingConsentsUnauthorized  %+v", 401, o.Payload)
}
func (o *RevokeOpenbankingConsentsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentsForbidden creates a RevokeOpenbankingConsentsForbidden with default headers values
func NewRevokeOpenbankingConsentsForbidden() *RevokeOpenbankingConsentsForbidden {
	return &RevokeOpenbankingConsentsForbidden{}
}

/* RevokeOpenbankingConsentsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type RevokeOpenbankingConsentsForbidden struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}/open-banking/consents][%d] revokeOpenbankingConsentsForbidden  %+v", 403, o.Payload)
}
func (o *RevokeOpenbankingConsentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentsNotFound creates a RevokeOpenbankingConsentsNotFound with default headers values
func NewRevokeOpenbankingConsentsNotFound() *RevokeOpenbankingConsentsNotFound {
	return &RevokeOpenbankingConsentsNotFound{}
}

/* RevokeOpenbankingConsentsNotFound describes a response with status code 404, with default header values.

HttpError
*/
type RevokeOpenbankingConsentsNotFound struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}/open-banking/consents][%d] revokeOpenbankingConsentsNotFound  %+v", 404, o.Payload)
}
func (o *RevokeOpenbankingConsentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentsTooManyRequests creates a RevokeOpenbankingConsentsTooManyRequests with default headers values
func NewRevokeOpenbankingConsentsTooManyRequests() *RevokeOpenbankingConsentsTooManyRequests {
	return &RevokeOpenbankingConsentsTooManyRequests{}
}

/* RevokeOpenbankingConsentsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type RevokeOpenbankingConsentsTooManyRequests struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}/open-banking/consents][%d] revokeOpenbankingConsentsTooManyRequests  %+v", 429, o.Payload)
}
func (o *RevokeOpenbankingConsentsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
