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

	"github.com/cloudentity/acp-client-go/clients/openbankingUK/payments/models"
)

// CreateFilePaymentConsentsReader is a Reader for the CreateFilePaymentConsents structure.
type CreateFilePaymentConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFilePaymentConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFilePaymentConsentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFilePaymentConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFilePaymentConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFilePaymentConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFilePaymentConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateFilePaymentConsentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateFilePaymentConsentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateFilePaymentConsentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateFilePaymentConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateFilePaymentConsentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFilePaymentConsentsCreated creates a CreateFilePaymentConsentsCreated with default headers values
func NewCreateFilePaymentConsentsCreated() *CreateFilePaymentConsentsCreated {
	return &CreateFilePaymentConsentsCreated{}
}

/* CreateFilePaymentConsentsCreated describes a response with status code 201, with default header values.

File Payment Consents Created
*/
type CreateFilePaymentConsentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteFileConsentResponse4
}

func (o *CreateFilePaymentConsentsCreated) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsCreated  %+v", 201, o.Payload)
}
func (o *CreateFilePaymentConsentsCreated) GetPayload() *models.OBWriteFileConsentResponse4 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsBadRequest creates a CreateFilePaymentConsentsBadRequest with default headers values
func NewCreateFilePaymentConsentsBadRequest() *CreateFilePaymentConsentsBadRequest {
	return &CreateFilePaymentConsentsBadRequest{}
}

/* CreateFilePaymentConsentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFilePaymentConsentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentConsentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFilePaymentConsentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsUnauthorized creates a CreateFilePaymentConsentsUnauthorized with default headers values
func NewCreateFilePaymentConsentsUnauthorized() *CreateFilePaymentConsentsUnauthorized {
	return &CreateFilePaymentConsentsUnauthorized{}
}

/* CreateFilePaymentConsentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateFilePaymentConsentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsUnauthorized ", 401)
}

func (o *CreateFilePaymentConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsForbidden creates a CreateFilePaymentConsentsForbidden with default headers values
func NewCreateFilePaymentConsentsForbidden() *CreateFilePaymentConsentsForbidden {
	return &CreateFilePaymentConsentsForbidden{}
}

/* CreateFilePaymentConsentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateFilePaymentConsentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentConsentsForbidden) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateFilePaymentConsentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsNotFound creates a CreateFilePaymentConsentsNotFound with default headers values
func NewCreateFilePaymentConsentsNotFound() *CreateFilePaymentConsentsNotFound {
	return &CreateFilePaymentConsentsNotFound{}
}

/* CreateFilePaymentConsentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateFilePaymentConsentsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsNotFound) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsNotFound ", 404)
}

func (o *CreateFilePaymentConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsMethodNotAllowed creates a CreateFilePaymentConsentsMethodNotAllowed with default headers values
func NewCreateFilePaymentConsentsMethodNotAllowed() *CreateFilePaymentConsentsMethodNotAllowed {
	return &CreateFilePaymentConsentsMethodNotAllowed{}
}

/* CreateFilePaymentConsentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateFilePaymentConsentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsMethodNotAllowed ", 405)
}

func (o *CreateFilePaymentConsentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsNotAcceptable creates a CreateFilePaymentConsentsNotAcceptable with default headers values
func NewCreateFilePaymentConsentsNotAcceptable() *CreateFilePaymentConsentsNotAcceptable {
	return &CreateFilePaymentConsentsNotAcceptable{}
}

/* CreateFilePaymentConsentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateFilePaymentConsentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsNotAcceptable ", 406)
}

func (o *CreateFilePaymentConsentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsUnsupportedMediaType creates a CreateFilePaymentConsentsUnsupportedMediaType with default headers values
func NewCreateFilePaymentConsentsUnsupportedMediaType() *CreateFilePaymentConsentsUnsupportedMediaType {
	return &CreateFilePaymentConsentsUnsupportedMediaType{}
}

/* CreateFilePaymentConsentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateFilePaymentConsentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsUnsupportedMediaType ", 415)
}

func (o *CreateFilePaymentConsentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsTooManyRequests creates a CreateFilePaymentConsentsTooManyRequests with default headers values
func NewCreateFilePaymentConsentsTooManyRequests() *CreateFilePaymentConsentsTooManyRequests {
	return &CreateFilePaymentConsentsTooManyRequests{}
}

/* CreateFilePaymentConsentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateFilePaymentConsentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsTooManyRequests ", 429)
}

func (o *CreateFilePaymentConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsInternalServerError creates a CreateFilePaymentConsentsInternalServerError with default headers values
func NewCreateFilePaymentConsentsInternalServerError() *CreateFilePaymentConsentsInternalServerError {
	return &CreateFilePaymentConsentsInternalServerError{}
}

/* CreateFilePaymentConsentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateFilePaymentConsentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentConsentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents][%d] createFilePaymentConsentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateFilePaymentConsentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
