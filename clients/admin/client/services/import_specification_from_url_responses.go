// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// ImportSpecificationFromURLReader is a Reader for the ImportSpecificationFromURL structure.
type ImportSpecificationFromURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportSpecificationFromURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportSpecificationFromURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportSpecificationFromURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportSpecificationFromURLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportSpecificationFromURLForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportSpecificationFromURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImportSpecificationFromURLConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewImportSpecificationFromURLUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewImportSpecificationFromURLTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportSpecificationFromURLOK creates a ImportSpecificationFromURLOK with default headers values
func NewImportSpecificationFromURLOK() *ImportSpecificationFromURLOK {
	return &ImportSpecificationFromURLOK{}
}

/* ImportSpecificationFromURLOK describes a response with status code 200, with default header values.

Import service configuration result
*/
type ImportSpecificationFromURLOK struct {
	Payload *models.ImportServiceConfigurationResult
}

func (o *ImportSpecificationFromURLOK) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlOK  %+v", 200, o.Payload)
}
func (o *ImportSpecificationFromURLOK) GetPayload() *models.ImportServiceConfigurationResult {
	return o.Payload
}

func (o *ImportSpecificationFromURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportServiceConfigurationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLBadRequest creates a ImportSpecificationFromURLBadRequest with default headers values
func NewImportSpecificationFromURLBadRequest() *ImportSpecificationFromURLBadRequest {
	return &ImportSpecificationFromURLBadRequest{}
}

/* ImportSpecificationFromURLBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type ImportSpecificationFromURLBadRequest struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlBadRequest  %+v", 400, o.Payload)
}
func (o *ImportSpecificationFromURLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLUnauthorized creates a ImportSpecificationFromURLUnauthorized with default headers values
func NewImportSpecificationFromURLUnauthorized() *ImportSpecificationFromURLUnauthorized {
	return &ImportSpecificationFromURLUnauthorized{}
}

/* ImportSpecificationFromURLUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ImportSpecificationFromURLUnauthorized struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLUnauthorized) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlUnauthorized  %+v", 401, o.Payload)
}
func (o *ImportSpecificationFromURLUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLForbidden creates a ImportSpecificationFromURLForbidden with default headers values
func NewImportSpecificationFromURLForbidden() *ImportSpecificationFromURLForbidden {
	return &ImportSpecificationFromURLForbidden{}
}

/* ImportSpecificationFromURLForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ImportSpecificationFromURLForbidden struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLForbidden) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlForbidden  %+v", 403, o.Payload)
}
func (o *ImportSpecificationFromURLForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLNotFound creates a ImportSpecificationFromURLNotFound with default headers values
func NewImportSpecificationFromURLNotFound() *ImportSpecificationFromURLNotFound {
	return &ImportSpecificationFromURLNotFound{}
}

/* ImportSpecificationFromURLNotFound describes a response with status code 404, with default header values.

HttpError
*/
type ImportSpecificationFromURLNotFound struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLNotFound) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlNotFound  %+v", 404, o.Payload)
}
func (o *ImportSpecificationFromURLNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLConflict creates a ImportSpecificationFromURLConflict with default headers values
func NewImportSpecificationFromURLConflict() *ImportSpecificationFromURLConflict {
	return &ImportSpecificationFromURLConflict{}
}

/* ImportSpecificationFromURLConflict describes a response with status code 409, with default header values.

HttpError
*/
type ImportSpecificationFromURLConflict struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLConflict) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlConflict  %+v", 409, o.Payload)
}
func (o *ImportSpecificationFromURLConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLUnprocessableEntity creates a ImportSpecificationFromURLUnprocessableEntity with default headers values
func NewImportSpecificationFromURLUnprocessableEntity() *ImportSpecificationFromURLUnprocessableEntity {
	return &ImportSpecificationFromURLUnprocessableEntity{}
}

/* ImportSpecificationFromURLUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type ImportSpecificationFromURLUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ImportSpecificationFromURLUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromURLTooManyRequests creates a ImportSpecificationFromURLTooManyRequests with default headers values
func NewImportSpecificationFromURLTooManyRequests() *ImportSpecificationFromURLTooManyRequests {
	return &ImportSpecificationFromURLTooManyRequests{}
}

/* ImportSpecificationFromURLTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ImportSpecificationFromURLTooManyRequests struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromURLTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /services/{sid}/apis/import/url][%d] importSpecificationFromUrlTooManyRequests  %+v", 429, o.Payload)
}
func (o *ImportSpecificationFromURLTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromURLTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}