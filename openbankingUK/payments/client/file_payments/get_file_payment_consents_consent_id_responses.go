// Code generated by go-swagger; DO NOT EDIT.

package file_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/acp-client-go/openbankingUK/payments/models"
)

// GetFilePaymentConsentsConsentIDReader is a Reader for the GetFilePaymentConsentsConsentID structure.
type GetFilePaymentConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilePaymentConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilePaymentConsentsConsentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFilePaymentConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFilePaymentConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilePaymentConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilePaymentConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFilePaymentConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFilePaymentConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFilePaymentConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFilePaymentConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilePaymentConsentsConsentIDOK creates a GetFilePaymentConsentsConsentIDOK with default headers values
func NewGetFilePaymentConsentsConsentIDOK() *GetFilePaymentConsentsConsentIDOK {
	return &GetFilePaymentConsentsConsentIDOK{}
}

/* GetFilePaymentConsentsConsentIDOK describes a response with status code 200, with default header values.

File Payment Consents Read
*/
type GetFilePaymentConsentsConsentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteFileConsentResponse4
}

func (o *GetFilePaymentConsentsConsentIDOK) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdOK  %+v", 200, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDOK) GetPayload() *models.OBWriteFileConsentResponse4 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBWriteFileConsentResponse4)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDBadRequest creates a GetFilePaymentConsentsConsentIDBadRequest with default headers values
func NewGetFilePaymentConsentsConsentIDBadRequest() *GetFilePaymentConsentsConsentIDBadRequest {
	return &GetFilePaymentConsentsConsentIDBadRequest{}
}

/* GetFilePaymentConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFilePaymentConsentsConsentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDUnauthorized creates a GetFilePaymentConsentsConsentIDUnauthorized with default headers values
func NewGetFilePaymentConsentsConsentIDUnauthorized() *GetFilePaymentConsentsConsentIDUnauthorized {
	return &GetFilePaymentConsentsConsentIDUnauthorized{}
}

/* GetFilePaymentConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFilePaymentConsentsConsentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdUnauthorized ", 401)
}

func (o *GetFilePaymentConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDForbidden creates a GetFilePaymentConsentsConsentIDForbidden with default headers values
func NewGetFilePaymentConsentsConsentIDForbidden() *GetFilePaymentConsentsConsentIDForbidden {
	return &GetFilePaymentConsentsConsentIDForbidden{}
}

/* GetFilePaymentConsentsConsentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFilePaymentConsentsConsentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDNotFound creates a GetFilePaymentConsentsConsentIDNotFound with default headers values
func NewGetFilePaymentConsentsConsentIDNotFound() *GetFilePaymentConsentsConsentIDNotFound {
	return &GetFilePaymentConsentsConsentIDNotFound{}
}

/* GetFilePaymentConsentsConsentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetFilePaymentConsentsConsentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdNotFound ", 404)
}

func (o *GetFilePaymentConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDMethodNotAllowed creates a GetFilePaymentConsentsConsentIDMethodNotAllowed with default headers values
func NewGetFilePaymentConsentsConsentIDMethodNotAllowed() *GetFilePaymentConsentsConsentIDMethodNotAllowed {
	return &GetFilePaymentConsentsConsentIDMethodNotAllowed{}
}

/* GetFilePaymentConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetFilePaymentConsentsConsentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdMethodNotAllowed ", 405)
}

func (o *GetFilePaymentConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDNotAcceptable creates a GetFilePaymentConsentsConsentIDNotAcceptable with default headers values
func NewGetFilePaymentConsentsConsentIDNotAcceptable() *GetFilePaymentConsentsConsentIDNotAcceptable {
	return &GetFilePaymentConsentsConsentIDNotAcceptable{}
}

/* GetFilePaymentConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetFilePaymentConsentsConsentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdNotAcceptable ", 406)
}

func (o *GetFilePaymentConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDTooManyRequests creates a GetFilePaymentConsentsConsentIDTooManyRequests with default headers values
func NewGetFilePaymentConsentsConsentIDTooManyRequests() *GetFilePaymentConsentsConsentIDTooManyRequests {
	return &GetFilePaymentConsentsConsentIDTooManyRequests{}
}

/* GetFilePaymentConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFilePaymentConsentsConsentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdTooManyRequests ", 429)
}

func (o *GetFilePaymentConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		valretryAfter, err := swag.ConvertInt64(hdrRetryAfter)
		if err != nil {
			return errors.InvalidType("Retry-After", "header", "int64", hdrRetryAfter)
		}
		o.RetryAfter = valretryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentConsentsConsentIDInternalServerError creates a GetFilePaymentConsentsConsentIDInternalServerError with default headers values
func NewGetFilePaymentConsentsConsentIDInternalServerError() *GetFilePaymentConsentsConsentIDInternalServerError {
	return &GetFilePaymentConsentsConsentIDInternalServerError{}
}

/* GetFilePaymentConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFilePaymentConsentsConsentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /file-payment-consents/{ConsentId}][%d] getFilePaymentConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFilePaymentConsentsConsentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
