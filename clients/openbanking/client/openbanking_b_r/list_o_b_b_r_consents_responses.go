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

// ListOBBRConsentsReader is a Reader for the ListOBBRConsents structure.
type ListOBBRConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOBBRConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOBBRConsentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListOBBRConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListOBBRConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListOBBRConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOBBRConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListOBBRConsentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListOBBRConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOBBRConsentsOK creates a ListOBBRConsentsOK with default headers values
func NewListOBBRConsentsOK() *ListOBBRConsentsOK {
	return &ListOBBRConsentsOK{}
}

/* ListOBBRConsentsOK describes a response with status code 200, with default header values.

OBBRConsents
*/
type ListOBBRConsentsOK struct {
	Payload *models.OBBRConsents
}

func (o *ListOBBRConsentsOK) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsOK  %+v", 200, o.Payload)
}
func (o *ListOBBRConsentsOK) GetPayload() *models.OBBRConsents {
	return o.Payload
}

func (o *ListOBBRConsentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OBBRConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOBBRConsentsBadRequest creates a ListOBBRConsentsBadRequest with default headers values
func NewListOBBRConsentsBadRequest() *ListOBBRConsentsBadRequest {
	return &ListOBBRConsentsBadRequest{}
}

/* ListOBBRConsentsBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type ListOBBRConsentsBadRequest struct {
	Payload *models.Error
}

func (o *ListOBBRConsentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *ListOBBRConsentsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOBBRConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOBBRConsentsUnauthorized creates a ListOBBRConsentsUnauthorized with default headers values
func NewListOBBRConsentsUnauthorized() *ListOBBRConsentsUnauthorized {
	return &ListOBBRConsentsUnauthorized{}
}

/* ListOBBRConsentsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListOBBRConsentsUnauthorized struct {
	Payload *models.Error
}

func (o *ListOBBRConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListOBBRConsentsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOBBRConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOBBRConsentsForbidden creates a ListOBBRConsentsForbidden with default headers values
func NewListOBBRConsentsForbidden() *ListOBBRConsentsForbidden {
	return &ListOBBRConsentsForbidden{}
}

/* ListOBBRConsentsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListOBBRConsentsForbidden struct {
	Payload *models.Error
}

func (o *ListOBBRConsentsForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsForbidden  %+v", 403, o.Payload)
}
func (o *ListOBBRConsentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOBBRConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOBBRConsentsNotFound creates a ListOBBRConsentsNotFound with default headers values
func NewListOBBRConsentsNotFound() *ListOBBRConsentsNotFound {
	return &ListOBBRConsentsNotFound{}
}

/* ListOBBRConsentsNotFound describes a response with status code 404, with default header values.

HttpError
*/
type ListOBBRConsentsNotFound struct {
	Payload *models.Error
}

func (o *ListOBBRConsentsNotFound) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsNotFound  %+v", 404, o.Payload)
}
func (o *ListOBBRConsentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOBBRConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOBBRConsentsUnprocessableEntity creates a ListOBBRConsentsUnprocessableEntity with default headers values
func NewListOBBRConsentsUnprocessableEntity() *ListOBBRConsentsUnprocessableEntity {
	return &ListOBBRConsentsUnprocessableEntity{}
}

/* ListOBBRConsentsUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type ListOBBRConsentsUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ListOBBRConsentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ListOBBRConsentsUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOBBRConsentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOBBRConsentsTooManyRequests creates a ListOBBRConsentsTooManyRequests with default headers values
func NewListOBBRConsentsTooManyRequests() *ListOBBRConsentsTooManyRequests {
	return &ListOBBRConsentsTooManyRequests{}
}

/* ListOBBRConsentsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListOBBRConsentsTooManyRequests struct {
	Payload *models.Error
}

func (o *ListOBBRConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /servers/{wid}/open-banking-brasil/consents][%d] listOBBRConsentsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListOBBRConsentsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOBBRConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
