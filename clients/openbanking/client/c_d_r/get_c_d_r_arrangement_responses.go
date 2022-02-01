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

// GetCDRArrangementReader is a Reader for the GetCDRArrangement structure.
type GetCDRArrangementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCDRArrangementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCDRArrangementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCDRArrangementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCDRArrangementUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCDRArrangementForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCDRArrangementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCDRArrangementTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCDRArrangementOK creates a GetCDRArrangementOK with default headers values
func NewGetCDRArrangementOK() *GetCDRArrangementOK {
	return &GetCDRArrangementOK{}
}

/* GetCDRArrangementOK describes a response with status code 200, with default header values.

CDRArrangement
*/
type GetCDRArrangementOK struct {
	Payload *models.CDRArrangement
}

func (o *GetCDRArrangementOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/cdr/arrangements/{arrangementID}][%d] getCDRArrangementOK  %+v", 200, o.Payload)
}
func (o *GetCDRArrangementOK) GetPayload() *models.CDRArrangement {
	return o.Payload
}

func (o *GetCDRArrangementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CDRArrangement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementBadRequest creates a GetCDRArrangementBadRequest with default headers values
func NewGetCDRArrangementBadRequest() *GetCDRArrangementBadRequest {
	return &GetCDRArrangementBadRequest{}
}

/* GetCDRArrangementBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type GetCDRArrangementBadRequest struct {
	Payload *models.Error
}

func (o *GetCDRArrangementBadRequest) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/cdr/arrangements/{arrangementID}][%d] getCDRArrangementBadRequest  %+v", 400, o.Payload)
}
func (o *GetCDRArrangementBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementUnauthorized creates a GetCDRArrangementUnauthorized with default headers values
func NewGetCDRArrangementUnauthorized() *GetCDRArrangementUnauthorized {
	return &GetCDRArrangementUnauthorized{}
}

/* GetCDRArrangementUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetCDRArrangementUnauthorized struct {
	Payload *models.Error
}

func (o *GetCDRArrangementUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/cdr/arrangements/{arrangementID}][%d] getCDRArrangementUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCDRArrangementUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementForbidden creates a GetCDRArrangementForbidden with default headers values
func NewGetCDRArrangementForbidden() *GetCDRArrangementForbidden {
	return &GetCDRArrangementForbidden{}
}

/* GetCDRArrangementForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetCDRArrangementForbidden struct {
	Payload *models.Error
}

func (o *GetCDRArrangementForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/cdr/arrangements/{arrangementID}][%d] getCDRArrangementForbidden  %+v", 403, o.Payload)
}
func (o *GetCDRArrangementForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementNotFound creates a GetCDRArrangementNotFound with default headers values
func NewGetCDRArrangementNotFound() *GetCDRArrangementNotFound {
	return &GetCDRArrangementNotFound{}
}

/* GetCDRArrangementNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetCDRArrangementNotFound struct {
	Payload *models.Error
}

func (o *GetCDRArrangementNotFound) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/cdr/arrangements/{arrangementID}][%d] getCDRArrangementNotFound  %+v", 404, o.Payload)
}
func (o *GetCDRArrangementNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDRArrangementTooManyRequests creates a GetCDRArrangementTooManyRequests with default headers values
func NewGetCDRArrangementTooManyRequests() *GetCDRArrangementTooManyRequests {
	return &GetCDRArrangementTooManyRequests{}
}

/* GetCDRArrangementTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetCDRArrangementTooManyRequests struct {
	Payload *models.Error
}

func (o *GetCDRArrangementTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/cdr/arrangements/{arrangementID}][%d] getCDRArrangementTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCDRArrangementTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCDRArrangementTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
