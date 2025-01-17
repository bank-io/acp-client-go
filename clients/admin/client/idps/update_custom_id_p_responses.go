// Code generated by go-swagger; DO NOT EDIT.

package idps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// UpdateCustomIDPReader is a Reader for the UpdateCustomIDP structure.
type UpdateCustomIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCustomIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCustomIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCustomIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateCustomIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCustomIDPTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomIDPOK creates a UpdateCustomIDPOK with default headers values
func NewUpdateCustomIDPOK() *UpdateCustomIDPOK {
	return &UpdateCustomIDPOK{}
}

/* UpdateCustomIDPOK describes a response with status code 200, with default header values.

CustomIDP
*/
type UpdateCustomIDPOK struct {
	Payload *models.CustomIDP
}

func (o *UpdateCustomIDPOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPOK  %+v", 200, o.Payload)
}
func (o *UpdateCustomIDPOK) GetPayload() *models.CustomIDP {
	return o.Payload
}

func (o *UpdateCustomIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomIDPBadRequest creates a UpdateCustomIDPBadRequest with default headers values
func NewUpdateCustomIDPBadRequest() *UpdateCustomIDPBadRequest {
	return &UpdateCustomIDPBadRequest{}
}

/* UpdateCustomIDPBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type UpdateCustomIDPBadRequest struct {
	Payload *models.Error
}

func (o *UpdateCustomIDPBadRequest) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCustomIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomIDPUnauthorized creates a UpdateCustomIDPUnauthorized with default headers values
func NewUpdateCustomIDPUnauthorized() *UpdateCustomIDPUnauthorized {
	return &UpdateCustomIDPUnauthorized{}
}

/* UpdateCustomIDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateCustomIDPUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateCustomIDPUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateCustomIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomIDPForbidden creates a UpdateCustomIDPForbidden with default headers values
func NewUpdateCustomIDPForbidden() *UpdateCustomIDPForbidden {
	return &UpdateCustomIDPForbidden{}
}

/* UpdateCustomIDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateCustomIDPForbidden struct {
	Payload *models.Error
}

func (o *UpdateCustomIDPForbidden) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCustomIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomIDPNotFound creates a UpdateCustomIDPNotFound with default headers values
func NewUpdateCustomIDPNotFound() *UpdateCustomIDPNotFound {
	return &UpdateCustomIDPNotFound{}
}

/* UpdateCustomIDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateCustomIDPNotFound struct {
	Payload *models.Error
}

func (o *UpdateCustomIDPNotFound) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPNotFound  %+v", 404, o.Payload)
}
func (o *UpdateCustomIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomIDPUnprocessableEntity creates a UpdateCustomIDPUnprocessableEntity with default headers values
func NewUpdateCustomIDPUnprocessableEntity() *UpdateCustomIDPUnprocessableEntity {
	return &UpdateCustomIDPUnprocessableEntity{}
}

/* UpdateCustomIDPUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateCustomIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateCustomIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateCustomIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomIDPTooManyRequests creates a UpdateCustomIDPTooManyRequests with default headers values
func NewUpdateCustomIDPTooManyRequests() *UpdateCustomIDPTooManyRequests {
	return &UpdateCustomIDPTooManyRequests{}
}

/* UpdateCustomIDPTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateCustomIDPTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateCustomIDPTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/custom/{iid}][%d] updateCustomIdPTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateCustomIDPTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomIDPTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
