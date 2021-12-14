// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/admin/models"
)

// GetStatsReader is a Reader for the GetStats structure.
type GetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStatsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatsOK creates a GetStatsOK with default headers values
func NewGetStatsOK() *GetStatsOK {
	return &GetStatsOK{}
}

/* GetStatsOK describes a response with status code 200, with default header values.

Stats
*/
type GetStatsOK struct {
	Payload *models.Stats
}

func (o *GetStatsOK) Error() string {
	return fmt.Sprintf("[GET /stats/{wid}/overview][%d] getStatsOK  %+v", 200, o.Payload)
}
func (o *GetStatsOK) GetPayload() *models.Stats {
	return o.Payload
}

func (o *GetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsUnauthorized creates a GetStatsUnauthorized with default headers values
func NewGetStatsUnauthorized() *GetStatsUnauthorized {
	return &GetStatsUnauthorized{}
}

/* GetStatsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetStatsUnauthorized struct {
	Payload *models.Error
}

func (o *GetStatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /stats/{wid}/overview][%d] getStatsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetStatsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsForbidden creates a GetStatsForbidden with default headers values
func NewGetStatsForbidden() *GetStatsForbidden {
	return &GetStatsForbidden{}
}

/* GetStatsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetStatsForbidden struct {
	Payload *models.Error
}

func (o *GetStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /stats/{wid}/overview][%d] getStatsForbidden  %+v", 403, o.Payload)
}
func (o *GetStatsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsTooManyRequests creates a GetStatsTooManyRequests with default headers values
func NewGetStatsTooManyRequests() *GetStatsTooManyRequests {
	return &GetStatsTooManyRequests{}
}

/* GetStatsTooManyRequests describes a response with status code 429, with default header values.

HttpError
*/
type GetStatsTooManyRequests struct {
	Payload *models.Error
}

func (o *GetStatsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /stats/{wid}/overview][%d] getStatsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetStatsTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
