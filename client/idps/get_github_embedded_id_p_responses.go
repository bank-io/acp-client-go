// Code generated by go-swagger; DO NOT EDIT.

package idps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// GetGithubEmbeddedIDPReader is a Reader for the GetGithubEmbeddedIDP structure.
type GetGithubEmbeddedIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGithubEmbeddedIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGithubEmbeddedIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGithubEmbeddedIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGithubEmbeddedIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGithubEmbeddedIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGithubEmbeddedIDPOK creates a GetGithubEmbeddedIDPOK with default headers values
func NewGetGithubEmbeddedIDPOK() *GetGithubEmbeddedIDPOK {
	return &GetGithubEmbeddedIDPOK{}
}

/* GetGithubEmbeddedIDPOK describes a response with status code 200, with default header values.

GithubEmbeddedIDP
*/
type GetGithubEmbeddedIDPOK struct {
	Payload *models.GithubEmbeddedIDP
}

func (o *GetGithubEmbeddedIDPOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/github_embedded/{iid}][%d] getGithubEmbeddedIdPOK  %+v", 200, o.Payload)
}
func (o *GetGithubEmbeddedIDPOK) GetPayload() *models.GithubEmbeddedIDP {
	return o.Payload
}

func (o *GetGithubEmbeddedIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubEmbeddedIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubEmbeddedIDPUnauthorized creates a GetGithubEmbeddedIDPUnauthorized with default headers values
func NewGetGithubEmbeddedIDPUnauthorized() *GetGithubEmbeddedIDPUnauthorized {
	return &GetGithubEmbeddedIDPUnauthorized{}
}

/* GetGithubEmbeddedIDPUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetGithubEmbeddedIDPUnauthorized struct {
	Payload *models.Error
}

func (o *GetGithubEmbeddedIDPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/github_embedded/{iid}][%d] getGithubEmbeddedIdPUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGithubEmbeddedIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubEmbeddedIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubEmbeddedIDPForbidden creates a GetGithubEmbeddedIDPForbidden with default headers values
func NewGetGithubEmbeddedIDPForbidden() *GetGithubEmbeddedIDPForbidden {
	return &GetGithubEmbeddedIDPForbidden{}
}

/* GetGithubEmbeddedIDPForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetGithubEmbeddedIDPForbidden struct {
	Payload *models.Error
}

func (o *GetGithubEmbeddedIDPForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/github_embedded/{iid}][%d] getGithubEmbeddedIdPForbidden  %+v", 403, o.Payload)
}
func (o *GetGithubEmbeddedIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubEmbeddedIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGithubEmbeddedIDPNotFound creates a GetGithubEmbeddedIDPNotFound with default headers values
func NewGetGithubEmbeddedIDPNotFound() *GetGithubEmbeddedIDPNotFound {
	return &GetGithubEmbeddedIDPNotFound{}
}

/* GetGithubEmbeddedIDPNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetGithubEmbeddedIDPNotFound struct {
	Payload *models.Error
}

func (o *GetGithubEmbeddedIDPNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/github_embedded/{iid}][%d] getGithubEmbeddedIdPNotFound  %+v", 404, o.Payload)
}
func (o *GetGithubEmbeddedIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGithubEmbeddedIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}