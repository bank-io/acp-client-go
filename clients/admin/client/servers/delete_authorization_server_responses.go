// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// DeleteAuthorizationServerReader is a Reader for the DeleteAuthorizationServer structure.
type DeleteAuthorizationServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthorizationServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAuthorizationServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAuthorizationServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAuthorizationServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAuthorizationServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAuthorizationServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAuthorizationServerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAuthorizationServerNoContent creates a DeleteAuthorizationServerNoContent with default headers values
func NewDeleteAuthorizationServerNoContent() *DeleteAuthorizationServerNoContent {
	return &DeleteAuthorizationServerNoContent{}
}

/* DeleteAuthorizationServerNoContent describes a response with status code 204, with default header values.

Authorization server has been deleted
*/
type DeleteAuthorizationServerNoContent struct {
}

func (o *DeleteAuthorizationServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}][%d] deleteAuthorizationServerNoContent ", 204)
}

func (o *DeleteAuthorizationServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAuthorizationServerBadRequest creates a DeleteAuthorizationServerBadRequest with default headers values
func NewDeleteAuthorizationServerBadRequest() *DeleteAuthorizationServerBadRequest {
	return &DeleteAuthorizationServerBadRequest{}
}

/* DeleteAuthorizationServerBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type DeleteAuthorizationServerBadRequest struct {
	Payload *models.Error
}

func (o *DeleteAuthorizationServerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}][%d] deleteAuthorizationServerBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAuthorizationServerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAuthorizationServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationServerUnauthorized creates a DeleteAuthorizationServerUnauthorized with default headers values
func NewDeleteAuthorizationServerUnauthorized() *DeleteAuthorizationServerUnauthorized {
	return &DeleteAuthorizationServerUnauthorized{}
}

/* DeleteAuthorizationServerUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type DeleteAuthorizationServerUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteAuthorizationServerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}][%d] deleteAuthorizationServerUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteAuthorizationServerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAuthorizationServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationServerForbidden creates a DeleteAuthorizationServerForbidden with default headers values
func NewDeleteAuthorizationServerForbidden() *DeleteAuthorizationServerForbidden {
	return &DeleteAuthorizationServerForbidden{}
}

/* DeleteAuthorizationServerForbidden describes a response with status code 403, with default header values.

HttpError
*/
type DeleteAuthorizationServerForbidden struct {
	Payload *models.Error
}

func (o *DeleteAuthorizationServerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}][%d] deleteAuthorizationServerForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAuthorizationServerForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAuthorizationServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationServerNotFound creates a DeleteAuthorizationServerNotFound with default headers values
func NewDeleteAuthorizationServerNotFound() *DeleteAuthorizationServerNotFound {
	return &DeleteAuthorizationServerNotFound{}
}

/* DeleteAuthorizationServerNotFound describes a response with status code 404, with default header values.

HttpError
*/
type DeleteAuthorizationServerNotFound struct {
	Payload *models.Error
}

func (o *DeleteAuthorizationServerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}][%d] deleteAuthorizationServerNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAuthorizationServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAuthorizationServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationServerTooManyRequests creates a DeleteAuthorizationServerTooManyRequests with default headers values
func NewDeleteAuthorizationServerTooManyRequests() *DeleteAuthorizationServerTooManyRequests {
	return &DeleteAuthorizationServerTooManyRequests{}
}

/* DeleteAuthorizationServerTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type DeleteAuthorizationServerTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteAuthorizationServerTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /servers/{wid}][%d] deleteAuthorizationServerTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteAuthorizationServerTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAuthorizationServerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
