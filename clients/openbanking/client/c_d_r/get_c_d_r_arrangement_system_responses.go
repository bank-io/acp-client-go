// Code generated by go-swagger; DO NOT EDIT.

package c_d_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// GetCDRArrangementSystemReader is a Reader for the GetCDRArrangementSystem structure.
type GetCDRArrangementSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCDRArrangementSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCDRArrangementSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCDRArrangementSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCDRArrangementSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCDRArrangementSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCDRArrangementSystemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCDRArrangementSystemOK creates a GetCDRArrangementSystemOK with default headers values
func NewGetCDRArrangementSystemOK() *GetCDRArrangementSystemOK {
	return &GetCDRArrangementSystemOK{}
}

/* GetCDRArrangementSystemOK describes a response with status code 200, with default header values.

GetCDRConsentResponse
*/
type GetCDRArrangementSystemOK struct {
	Payload *models.GetCDRConsentResponse
}

func (o *GetCDRArrangementSystemOK) Error() string {
	return fmt.Sprintf("[GET /cdr/cdr-arrangement/{login}][%d] getCDRArrangementSystemOK  %+v", 200, o.Payload)
}
func (o *GetCDRArrangementSystemOK) GetPayload() *models.GetCDRConsentResponse {
	return o.Payload
}

func (o *GetCDRArrangementSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCDRConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementSystemUnauthorized creates a GetCDRArrangementSystemUnauthorized with default headers values
func NewGetCDRArrangementSystemUnauthorized() *GetCDRArrangementSystemUnauthorized {
	return &GetCDRArrangementSystemUnauthorized{}
}

/* GetCDRArrangementSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetCDRArrangementSystemUnauthorized struct {
	Payload *models.Error
}

func (o *GetCDRArrangementSystemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdr/cdr-arrangement/{login}][%d] getCDRArrangementSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCDRArrangementSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementSystemForbidden creates a GetCDRArrangementSystemForbidden with default headers values
func NewGetCDRArrangementSystemForbidden() *GetCDRArrangementSystemForbidden {
	return &GetCDRArrangementSystemForbidden{}
}

/* GetCDRArrangementSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetCDRArrangementSystemForbidden struct {
	Payload *models.Error
}

func (o *GetCDRArrangementSystemForbidden) Error() string {
	return fmt.Sprintf("[GET /cdr/cdr-arrangement/{login}][%d] getCDRArrangementSystemForbidden  %+v", 403, o.Payload)
}
func (o *GetCDRArrangementSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementSystemNotFound creates a GetCDRArrangementSystemNotFound with default headers values
func NewGetCDRArrangementSystemNotFound() *GetCDRArrangementSystemNotFound {
	return &GetCDRArrangementSystemNotFound{}
}

/* GetCDRArrangementSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetCDRArrangementSystemNotFound struct {
	Payload *models.Error
}

func (o *GetCDRArrangementSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /cdr/cdr-arrangement/{login}][%d] getCDRArrangementSystemNotFound  %+v", 404, o.Payload)
}
func (o *GetCDRArrangementSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementSystemTooManyRequests creates a GetCDRArrangementSystemTooManyRequests with default headers values
func NewGetCDRArrangementSystemTooManyRequests() *GetCDRArrangementSystemTooManyRequests {
	return &GetCDRArrangementSystemTooManyRequests{}
}

/* GetCDRArrangementSystemTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetCDRArrangementSystemTooManyRequests struct {
	Payload *models.Error
}

func (o *GetCDRArrangementSystemTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cdr/cdr-arrangement/{login}][%d] getCDRArrangementSystemTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCDRArrangementSystemTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementSystemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
