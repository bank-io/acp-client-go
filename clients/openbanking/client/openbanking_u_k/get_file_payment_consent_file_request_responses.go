// Code generated by go-swagger; DO NOT EDIT.

package openbanking_u_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

// GetFilePaymentConsentFileRequestReader is a Reader for the GetFilePaymentConsentFileRequest structure.
type GetFilePaymentConsentFileRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilePaymentConsentFileRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilePaymentConsentFileRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFilePaymentConsentFileRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFilePaymentConsentFileRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilePaymentConsentFileRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFilePaymentConsentFileRequestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFilePaymentConsentFileRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFilePaymentConsentFileRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFilePaymentConsentFileRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFilePaymentConsentFileRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilePaymentConsentFileRequestOK creates a GetFilePaymentConsentFileRequestOK with default headers values
func NewGetFilePaymentConsentFileRequestOK() *GetFilePaymentConsentFileRequestOK {
	return &GetFilePaymentConsentFileRequestOK{}
}

/* GetFilePaymentConsentFileRequestOK describes a response with status code 200, with default header values.

File payment consent file
*/
type GetFilePaymentConsentFileRequestOK struct {
	Payload models.FilePaymentConsentFileResponse
}

func (o *GetFilePaymentConsentFileRequestOK) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestOK  %+v", 200, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestOK) GetPayload() models.FilePaymentConsentFileResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestBadRequest creates a GetFilePaymentConsentFileRequestBadRequest with default headers values
func NewGetFilePaymentConsentFileRequestBadRequest() *GetFilePaymentConsentFileRequestBadRequest {
	return &GetFilePaymentConsentFileRequestBadRequest{}
}

/* GetFilePaymentConsentFileRequestBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestBadRequest  %+v", 400, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestUnauthorized creates a GetFilePaymentConsentFileRequestUnauthorized with default headers values
func NewGetFilePaymentConsentFileRequestUnauthorized() *GetFilePaymentConsentFileRequestUnauthorized {
	return &GetFilePaymentConsentFileRequestUnauthorized{}
}

/* GetFilePaymentConsentFileRequestUnauthorized describes a response with status code 401, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestForbidden creates a GetFilePaymentConsentFileRequestForbidden with default headers values
func NewGetFilePaymentConsentFileRequestForbidden() *GetFilePaymentConsentFileRequestForbidden {
	return &GetFilePaymentConsentFileRequestForbidden{}
}

/* GetFilePaymentConsentFileRequestForbidden describes a response with status code 403, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestForbidden  %+v", 403, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestMethodNotAllowed creates a GetFilePaymentConsentFileRequestMethodNotAllowed with default headers values
func NewGetFilePaymentConsentFileRequestMethodNotAllowed() *GetFilePaymentConsentFileRequestMethodNotAllowed {
	return &GetFilePaymentConsentFileRequestMethodNotAllowed{}
}

/* GetFilePaymentConsentFileRequestMethodNotAllowed describes a response with status code 405, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestNotAcceptable creates a GetFilePaymentConsentFileRequestNotAcceptable with default headers values
func NewGetFilePaymentConsentFileRequestNotAcceptable() *GetFilePaymentConsentFileRequestNotAcceptable {
	return &GetFilePaymentConsentFileRequestNotAcceptable{}
}

/* GetFilePaymentConsentFileRequestNotAcceptable describes a response with status code 406, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestNotAcceptable  %+v", 406, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestUnsupportedMediaType creates a GetFilePaymentConsentFileRequestUnsupportedMediaType with default headers values
func NewGetFilePaymentConsentFileRequestUnsupportedMediaType() *GetFilePaymentConsentFileRequestUnsupportedMediaType {
	return &GetFilePaymentConsentFileRequestUnsupportedMediaType{}
}

/* GetFilePaymentConsentFileRequestUnsupportedMediaType describes a response with status code 415, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestTooManyRequests creates a GetFilePaymentConsentFileRequestTooManyRequests with default headers values
func NewGetFilePaymentConsentFileRequestTooManyRequests() *GetFilePaymentConsentFileRequestTooManyRequests {
	return &GetFilePaymentConsentFileRequestTooManyRequests{}
}

/* GetFilePaymentConsentFileRequestTooManyRequests describes a response with status code 429, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentFileRequestInternalServerError creates a GetFilePaymentConsentFileRequestInternalServerError with default headers values
func NewGetFilePaymentConsentFileRequestInternalServerError() *GetFilePaymentConsentFileRequestInternalServerError {
	return &GetFilePaymentConsentFileRequestInternalServerError{}
}

/* GetFilePaymentConsentFileRequestInternalServerError describes a response with status code 500, with default header values.

Error
*/
type GetFilePaymentConsentFileRequestInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetFilePaymentConsentFileRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /open-banking/v3.1/pisp/file-payment-consents/{consentID}/file][%d] getFilePaymentConsentFileRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFilePaymentConsentFileRequestInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilePaymentConsentFileRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
