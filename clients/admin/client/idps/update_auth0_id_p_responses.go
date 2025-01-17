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

// UpdateAuth0IDPReader is a Reader for the UpdateAuth0IDP structure.
type UpdateAuth0IDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAuth0IDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAuth0IDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAuth0IDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAuth0IDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAuth0IDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAuth0IDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateAuth0IDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateAuth0IDPTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAuth0IDPOK creates a UpdateAuth0IDPOK with default headers values
func NewUpdateAuth0IDPOK() *UpdateAuth0IDPOK {
	return &UpdateAuth0IDPOK{}
}

/* UpdateAuth0IDPOK describes a response with status code 200, with default header values.

Auth0IDP
*/
type UpdateAuth0IDPOK struct {
	Payload *models.Auth0IDP
}

func (o *UpdateAuth0IDPOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPOK  %+v", 200, o.Payload)
}
func (o *UpdateAuth0IDPOK) GetPayload() *models.Auth0IDP {
	return o.Payload
}

func (o *UpdateAuth0IDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Auth0IDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAuth0IDPBadRequest creates a UpdateAuth0IDPBadRequest with default headers values
func NewUpdateAuth0IDPBadRequest() *UpdateAuth0IDPBadRequest {
	return &UpdateAuth0IDPBadRequest{}
}

/* UpdateAuth0IDPBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type UpdateAuth0IDPBadRequest struct {
	Payload *models.Error
}

func (o *UpdateAuth0IDPBadRequest) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateAuth0IDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAuth0IDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAuth0IDPUnauthorized creates a UpdateAuth0IDPUnauthorized with default headers values
func NewUpdateAuth0IDPUnauthorized() *UpdateAuth0IDPUnauthorized {
	return &UpdateAuth0IDPUnauthorized{}
}

/* UpdateAuth0IDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type UpdateAuth0IDPUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateAuth0IDPUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateAuth0IDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAuth0IDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAuth0IDPForbidden creates a UpdateAuth0IDPForbidden with default headers values
func NewUpdateAuth0IDPForbidden() *UpdateAuth0IDPForbidden {
	return &UpdateAuth0IDPForbidden{}
}

/* UpdateAuth0IDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type UpdateAuth0IDPForbidden struct {
	Payload *models.Error
}

func (o *UpdateAuth0IDPForbidden) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPForbidden  %+v", 403, o.Payload)
}
func (o *UpdateAuth0IDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAuth0IDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAuth0IDPNotFound creates a UpdateAuth0IDPNotFound with default headers values
func NewUpdateAuth0IDPNotFound() *UpdateAuth0IDPNotFound {
	return &UpdateAuth0IDPNotFound{}
}

/* UpdateAuth0IDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type UpdateAuth0IDPNotFound struct {
	Payload *models.Error
}

func (o *UpdateAuth0IDPNotFound) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPNotFound  %+v", 404, o.Payload)
}
func (o *UpdateAuth0IDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAuth0IDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAuth0IDPUnprocessableEntity creates a UpdateAuth0IDPUnprocessableEntity with default headers values
func NewUpdateAuth0IDPUnprocessableEntity() *UpdateAuth0IDPUnprocessableEntity {
	return &UpdateAuth0IDPUnprocessableEntity{}
}

/* UpdateAuth0IDPUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type UpdateAuth0IDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateAuth0IDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateAuth0IDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAuth0IDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAuth0IDPTooManyRequests creates a UpdateAuth0IDPTooManyRequests with default headers values
func NewUpdateAuth0IDPTooManyRequests() *UpdateAuth0IDPTooManyRequests {
	return &UpdateAuth0IDPTooManyRequests{}
}

/* UpdateAuth0IDPTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type UpdateAuth0IDPTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateAuth0IDPTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /servers/{wid}/idps/auth0/{iid}][%d] updateAuth0IdPTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateAuth0IDPTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAuth0IDPTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
