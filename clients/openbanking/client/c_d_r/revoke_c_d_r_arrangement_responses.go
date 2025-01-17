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

// RevokeCDRArrangementReader is a Reader for the RevokeCDRArrangement structure.
type RevokeCDRArrangementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeCDRArrangementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRevokeCDRArrangementNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeCDRArrangementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRevokeCDRArrangementUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeCDRArrangementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRevokeCDRArrangementUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRevokeCDRArrangementTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeCDRArrangementNoContent creates a RevokeCDRArrangementNoContent with default headers values
func NewRevokeCDRArrangementNoContent() *RevokeCDRArrangementNoContent {
	return &RevokeCDRArrangementNoContent{}
}

/* RevokeCDRArrangementNoContent describes a response with status code 204, with default header values.

CDR Arrangement has been revoked
*/
type RevokeCDRArrangementNoContent struct {
}

func (o *RevokeCDRArrangementNoContent) Error() string {
	return fmt.Sprintf("[POST /arrangements/revoke][%d] revokeCDRArrangementNoContent ", 204)
}

func (o *RevokeCDRArrangementNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeCDRArrangementBadRequest creates a RevokeCDRArrangementBadRequest with default headers values
func NewRevokeCDRArrangementBadRequest() *RevokeCDRArrangementBadRequest {
	return &RevokeCDRArrangementBadRequest{}
}

/* RevokeCDRArrangementBadRequest describes a response with status code 400, with default header values.

CDR Error
*/
type RevokeCDRArrangementBadRequest struct {
	Payload *models.CDRErrorResponse
}

func (o *RevokeCDRArrangementBadRequest) Error() string {
	return fmt.Sprintf("[POST /arrangements/revoke][%d] revokeCDRArrangementBadRequest  %+v", 400, o.Payload)
}
func (o *RevokeCDRArrangementBadRequest) GetPayload() *models.CDRErrorResponse {
	return o.Payload
}

func (o *RevokeCDRArrangementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CDRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeCDRArrangementUnauthorized creates a RevokeCDRArrangementUnauthorized with default headers values
func NewRevokeCDRArrangementUnauthorized() *RevokeCDRArrangementUnauthorized {
	return &RevokeCDRArrangementUnauthorized{}
}

/* RevokeCDRArrangementUnauthorized describes a response with status code 401, with default header values.

CDR Error
*/
type RevokeCDRArrangementUnauthorized struct {
	Payload *models.CDRErrorResponse
}

func (o *RevokeCDRArrangementUnauthorized) Error() string {
	return fmt.Sprintf("[POST /arrangements/revoke][%d] revokeCDRArrangementUnauthorized  %+v", 401, o.Payload)
}
func (o *RevokeCDRArrangementUnauthorized) GetPayload() *models.CDRErrorResponse {
	return o.Payload
}

func (o *RevokeCDRArrangementUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CDRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeCDRArrangementNotFound creates a RevokeCDRArrangementNotFound with default headers values
func NewRevokeCDRArrangementNotFound() *RevokeCDRArrangementNotFound {
	return &RevokeCDRArrangementNotFound{}
}

/* RevokeCDRArrangementNotFound describes a response with status code 404, with default header values.

CDR Error
*/
type RevokeCDRArrangementNotFound struct {
	Payload *models.CDRErrorResponse
}

func (o *RevokeCDRArrangementNotFound) Error() string {
	return fmt.Sprintf("[POST /arrangements/revoke][%d] revokeCDRArrangementNotFound  %+v", 404, o.Payload)
}
func (o *RevokeCDRArrangementNotFound) GetPayload() *models.CDRErrorResponse {
	return o.Payload
}

func (o *RevokeCDRArrangementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CDRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeCDRArrangementUnprocessableEntity creates a RevokeCDRArrangementUnprocessableEntity with default headers values
func NewRevokeCDRArrangementUnprocessableEntity() *RevokeCDRArrangementUnprocessableEntity {
	return &RevokeCDRArrangementUnprocessableEntity{}
}

/* RevokeCDRArrangementUnprocessableEntity describes a response with status code 422, with default header values.

CDR Error
*/
type RevokeCDRArrangementUnprocessableEntity struct {
	Payload *models.CDRErrorResponse
}

func (o *RevokeCDRArrangementUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /arrangements/revoke][%d] revokeCDRArrangementUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *RevokeCDRArrangementUnprocessableEntity) GetPayload() *models.CDRErrorResponse {
	return o.Payload
}

func (o *RevokeCDRArrangementUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CDRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeCDRArrangementTooManyRequests creates a RevokeCDRArrangementTooManyRequests with default headers values
func NewRevokeCDRArrangementTooManyRequests() *RevokeCDRArrangementTooManyRequests {
	return &RevokeCDRArrangementTooManyRequests{}
}

/* RevokeCDRArrangementTooManyRequests describes a response with status code 429, with default header values.

CDR Error
*/
type RevokeCDRArrangementTooManyRequests struct {
	Payload *models.CDRErrorResponse
}

func (o *RevokeCDRArrangementTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /arrangements/revoke][%d] revokeCDRArrangementTooManyRequests  %+v", 429, o.Payload)
}
func (o *RevokeCDRArrangementTooManyRequests) GetPayload() *models.CDRErrorResponse {
	return o.Payload
}

func (o *RevokeCDRArrangementTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CDRErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
