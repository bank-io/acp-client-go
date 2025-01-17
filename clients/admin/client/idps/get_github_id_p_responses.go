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

// GetGithubIDPReader is a Reader for the GetGithubIDP structure.
type GetGithubIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGithubIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGithubIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGithubIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGithubIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGithubIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGithubIDPTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGithubIDPOK creates a GetGithubIDPOK with default headers values
func NewGetGithubIDPOK() *GetGithubIDPOK {
	return &GetGithubIDPOK{}
}

/* GetGithubIDPOK describes a response with status code 200, with default header values.

GithubIDP
*/
type GetGithubIDPOK struct {
	Payload *models.GithubIDP
}

func (o *GetGithubIDPOK) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/idps/github/{iid}][%d] getGithubIdPOK  %+v", 200, o.Payload)
}
func (o *GetGithubIDPOK) GetPayload() *models.GithubIDP {
	return o.Payload
}

func (o *GetGithubIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubIDPUnauthorized creates a GetGithubIDPUnauthorized with default headers values
func NewGetGithubIDPUnauthorized() *GetGithubIDPUnauthorized {
	return &GetGithubIDPUnauthorized{}
}

/* GetGithubIDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetGithubIDPUnauthorized struct {
	Payload *models.Error
}

func (o *GetGithubIDPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/idps/github/{iid}][%d] getGithubIdPUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGithubIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubIDPForbidden creates a GetGithubIDPForbidden with default headers values
func NewGetGithubIDPForbidden() *GetGithubIDPForbidden {
	return &GetGithubIDPForbidden{}
}

/* GetGithubIDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetGithubIDPForbidden struct {
	Payload *models.Error
}

func (o *GetGithubIDPForbidden) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/idps/github/{iid}][%d] getGithubIdPForbidden  %+v", 403, o.Payload)
}
func (o *GetGithubIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubIDPNotFound creates a GetGithubIDPNotFound with default headers values
func NewGetGithubIDPNotFound() *GetGithubIDPNotFound {
	return &GetGithubIDPNotFound{}
}

/* GetGithubIDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetGithubIDPNotFound struct {
	Payload *models.Error
}

func (o *GetGithubIDPNotFound) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/idps/github/{iid}][%d] getGithubIdPNotFound  %+v", 404, o.Payload)
}
func (o *GetGithubIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubIDPTooManyRequests creates a GetGithubIDPTooManyRequests with default headers values
func NewGetGithubIDPTooManyRequests() *GetGithubIDPTooManyRequests {
	return &GetGithubIDPTooManyRequests{}
}

/* GetGithubIDPTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetGithubIDPTooManyRequests struct {
	Payload *models.Error
}

func (o *GetGithubIDPTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /servers/{wid}/idps/github/{iid}][%d] getGithubIdPTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetGithubIDPTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubIDPTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
