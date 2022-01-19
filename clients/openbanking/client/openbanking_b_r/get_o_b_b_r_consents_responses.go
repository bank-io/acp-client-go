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

// GetOBBRConsentsReader is a Reader for the GetOBBRConsents structure.
type GetOBBRConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOBBRConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOBBRConsentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOBBRConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOBBRConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOBBRConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOBBRConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOBBRConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOBBRConsentsOK creates a GetOBBRConsentsOK with default headers values
func NewGetOBBRConsentsOK() *GetOBBRConsentsOK {
	return &GetOBBRConsentsOK{}
}

/* GetOBBRConsentsOK describes a response with status code 200, with default header values.

OBBRConsents
*/
type GetOBBRConsentsOK struct {
	Payload *models.OBBRConsents
}

func (o *GetOBBRConsentsOK) Error() string {
	return fmt.Sprintf("[GET /servers/{aid}/open-banking-brasil/consents][%d] getOBBRConsentsOK  %+v", 200, o.Payload)
}
func (o *GetOBBRConsentsOK) GetPayload() *models.OBBRConsents {
	return o.Payload
}

func (o *GetOBBRConsentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBBRConsentsBadRequest creates a GetOBBRConsentsBadRequest with default headers values
func NewGetOBBRConsentsBadRequest() *GetOBBRConsentsBadRequest {
	return &GetOBBRConsentsBadRequest{}
}

/* GetOBBRConsentsBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type GetOBBRConsentsBadRequest struct {
	Payload *models.Error
}

func (o *GetOBBRConsentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /servers/{aid}/open-banking-brasil/consents][%d] getOBBRConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *GetOBBRConsentsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBBRConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBBRConsentsUnauthorized creates a GetOBBRConsentsUnauthorized with default headers values
func NewGetOBBRConsentsUnauthorized() *GetOBBRConsentsUnauthorized {
	return &GetOBBRConsentsUnauthorized{}
}

/* GetOBBRConsentsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetOBBRConsentsUnauthorized struct {
	Payload *models.Error
}

func (o *GetOBBRConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{aid}/open-banking-brasil/consents][%d] getOBBRConsentsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetOBBRConsentsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBBRConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBBRConsentsForbidden creates a GetOBBRConsentsForbidden with default headers values
func NewGetOBBRConsentsForbidden() *GetOBBRConsentsForbidden {
	return &GetOBBRConsentsForbidden{}
}

/* GetOBBRConsentsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetOBBRConsentsForbidden struct {
	Payload *models.Error
}

func (o *GetOBBRConsentsForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{aid}/open-banking-brasil/consents][%d] getOBBRConsentsForbidden  %+v", 403, o.Payload)
}
func (o *GetOBBRConsentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBBRConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBBRConsentsNotFound creates a GetOBBRConsentsNotFound with default headers values
func NewGetOBBRConsentsNotFound() *GetOBBRConsentsNotFound {
	return &GetOBBRConsentsNotFound{}
}

/* GetOBBRConsentsNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetOBBRConsentsNotFound struct {
	Payload *models.Error
}

func (o *GetOBBRConsentsNotFound) Error() string {
	return fmt.Sprintf("[GET /servers/{aid}/open-banking-brasil/consents][%d] getOBBRConsentsNotFound  %+v", 404, o.Payload)
}
func (o *GetOBBRConsentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBBRConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBBRConsentsTooManyRequests creates a GetOBBRConsentsTooManyRequests with default headers values
func NewGetOBBRConsentsTooManyRequests() *GetOBBRConsentsTooManyRequests {
	return &GetOBBRConsentsTooManyRequests{}
}

/* GetOBBRConsentsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetOBBRConsentsTooManyRequests struct {
	Payload *models.Error
}

func (o *GetOBBRConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{aid}/open-banking-brasil/consents][%d] getOBBRConsentsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetOBBRConsentsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBBRConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
