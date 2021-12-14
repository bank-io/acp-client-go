// Code generated by go-swagger; DO NOT EDIT.

package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// ListScopesReader is a Reader for the ListScopes structure.
type ListScopesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListScopesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListScopesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListScopesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListScopesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListScopesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListScopesOK creates a ListScopesOK with default headers values
func NewListScopesOK() *ListScopesOK {
	return &ListScopesOK{}
}

/* ListScopesOK describes a response with status code 200, with default header values.

Scopes with services
*/
type ListScopesOK struct {
	Payload *models.ScopesWithServices
}

func (o *ListScopesOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/scopes][%d] listScopesOK  %+v", 200, o.Payload)
}
func (o *ListScopesOK) GetPayload() *models.ScopesWithServices {
	return o.Payload
}

func (o *ListScopesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopesWithServices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScopesUnauthorized creates a ListScopesUnauthorized with default headers values
func NewListScopesUnauthorized() *ListScopesUnauthorized {
	return &ListScopesUnauthorized{}
}

/* ListScopesUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListScopesUnauthorized struct {
	Payload *models.Error
}

func (o *ListScopesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/scopes][%d] listScopesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListScopesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListScopesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScopesForbidden creates a ListScopesForbidden with default headers values
func NewListScopesForbidden() *ListScopesForbidden {
	return &ListScopesForbidden{}
}

/* ListScopesForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListScopesForbidden struct {
	Payload *models.Error
}

func (o *ListScopesForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/scopes][%d] listScopesForbidden  %+v", 403, o.Payload)
}
func (o *ListScopesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListScopesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScopesTooManyRequests creates a ListScopesTooManyRequests with default headers values
func NewListScopesTooManyRequests() *ListScopesTooManyRequests {
	return &ListScopesTooManyRequests{}
}

/* ListScopesTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type ListScopesTooManyRequests struct {
	Payload *models.Error
}

func (o *ListScopesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/scopes][%d] listScopesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListScopesTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListScopesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
