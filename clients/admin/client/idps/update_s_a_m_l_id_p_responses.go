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

// UpdateSAMLIDPReader is a Reader for the UpdateSAMLIDP structure.
type UpdateSAMLIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSAMLIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSAMLIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSAMLIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateSAMLIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSAMLIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSAMLIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateSAMLIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateSAMLIDPTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSAMLIDPOK creates a UpdateSAMLIDPOK with default headers values
func NewUpdateSAMLIDPOK() *UpdateSAMLIDPOK {
	return &UpdateSAMLIDPOK{}
}

/* UpdateSAMLIDPOK describes a response with status code 200, with default header values.

SAMLIDP
*/
type UpdateSAMLIDPOK struct {
	Payload *models.SAMLIDP
}

func (o *UpdateSAMLIDPOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPOK  %+v", 200, o.Payload)
}
func (o *UpdateSAMLIDPOK) GetPayload() *models.SAMLIDP {
	return o.Payload
}

func (o *UpdateSAMLIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAMLIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSAMLIDPBadRequest creates a UpdateSAMLIDPBadRequest with default headers values
func NewUpdateSAMLIDPBadRequest() *UpdateSAMLIDPBadRequest {
	return &UpdateSAMLIDPBadRequest{}
}

/* UpdateSAMLIDPBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type UpdateSAMLIDPBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSAMLIDPBadRequest) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateSAMLIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSAMLIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSAMLIDPUnauthorized creates a UpdateSAMLIDPUnauthorized with default headers values
func NewUpdateSAMLIDPUnauthorized() *UpdateSAMLIDPUnauthorized {
	return &UpdateSAMLIDPUnauthorized{}
}

/* UpdateSAMLIDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateSAMLIDPUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateSAMLIDPUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateSAMLIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSAMLIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSAMLIDPForbidden creates a UpdateSAMLIDPForbidden with default headers values
func NewUpdateSAMLIDPForbidden() *UpdateSAMLIDPForbidden {
	return &UpdateSAMLIDPForbidden{}
}

/* UpdateSAMLIDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateSAMLIDPForbidden struct {
	Payload *models.Error
}

func (o *UpdateSAMLIDPForbidden) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPForbidden  %+v", 403, o.Payload)
}
func (o *UpdateSAMLIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSAMLIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSAMLIDPNotFound creates a UpdateSAMLIDPNotFound with default headers values
func NewUpdateSAMLIDPNotFound() *UpdateSAMLIDPNotFound {
	return &UpdateSAMLIDPNotFound{}
}

/* UpdateSAMLIDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateSAMLIDPNotFound struct {
	Payload *models.Error
}

func (o *UpdateSAMLIDPNotFound) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPNotFound  %+v", 404, o.Payload)
}
func (o *UpdateSAMLIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSAMLIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSAMLIDPUnprocessableEntity creates a UpdateSAMLIDPUnprocessableEntity with default headers values
func NewUpdateSAMLIDPUnprocessableEntity() *UpdateSAMLIDPUnprocessableEntity {
	return &UpdateSAMLIDPUnprocessableEntity{}
}

/* UpdateSAMLIDPUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateSAMLIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateSAMLIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateSAMLIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSAMLIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSAMLIDPTooManyRequests creates a UpdateSAMLIDPTooManyRequests with default headers values
func NewUpdateSAMLIDPTooManyRequests() *UpdateSAMLIDPTooManyRequests {
	return &UpdateSAMLIDPTooManyRequests{}
}

/* UpdateSAMLIDPTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateSAMLIDPTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateSAMLIDPTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/saml/{iid}][%d] updateSAMLIdPTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateSAMLIDPTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSAMLIDPTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
