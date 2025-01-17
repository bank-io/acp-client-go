// Code generated by go-swagger; DO NOT EDIT.

package domestic_scheduled_payments

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

// CreateDomesticScheduledPaymentConsentsReader is a Reader for the CreateDomesticScheduledPaymentConsents structure.
type CreateDomesticScheduledPaymentConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDomesticScheduledPaymentConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDomesticScheduledPaymentConsentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDomesticScheduledPaymentConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDomesticScheduledPaymentConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDomesticScheduledPaymentConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDomesticScheduledPaymentConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateDomesticScheduledPaymentConsentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateDomesticScheduledPaymentConsentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDomesticScheduledPaymentConsentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDomesticScheduledPaymentConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDomesticScheduledPaymentConsentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDomesticScheduledPaymentConsentsCreated creates a CreateDomesticScheduledPaymentConsentsCreated with default headers values
func NewCreateDomesticScheduledPaymentConsentsCreated() *CreateDomesticScheduledPaymentConsentsCreated {
	return &CreateDomesticScheduledPaymentConsentsCreated{}
}

/* CreateDomesticScheduledPaymentConsentsCreated describes a response with status code 201, with default header values.

Domestic Scheduled Payment Consents Created
*/
type CreateDomesticScheduledPaymentConsentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteDomesticScheduledConsentResponse5
}

func (o *CreateDomesticScheduledPaymentConsentsCreated) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsCreated  %+v", 201, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentsCreated) GetPayload() *models.OBWriteDomesticScheduledConsentResponse5 {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteDomesticScheduledConsentResponse5)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentsBadRequest creates a CreateDomesticScheduledPaymentConsentsBadRequest with default headers values
func NewCreateDomesticScheduledPaymentConsentsBadRequest() *CreateDomesticScheduledPaymentConsentsBadRequest {
	return &CreateDomesticScheduledPaymentConsentsBadRequest{}
}

/* CreateDomesticScheduledPaymentConsentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDomesticScheduledPaymentConsentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticScheduledPaymentConsentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticScheduledPaymentConsentsUnauthorized creates a CreateDomesticScheduledPaymentConsentsUnauthorized with default headers values
func NewCreateDomesticScheduledPaymentConsentsUnauthorized() *CreateDomesticScheduledPaymentConsentsUnauthorized {
	return &CreateDomesticScheduledPaymentConsentsUnauthorized{}
}

/* CreateDomesticScheduledPaymentConsentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDomesticScheduledPaymentConsentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticScheduledPaymentConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsUnauthorized ", 401)
}

func (o *CreateDomesticScheduledPaymentConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentsForbidden creates a CreateDomesticScheduledPaymentConsentsForbidden with default headers values
func NewCreateDomesticScheduledPaymentConsentsForbidden() *CreateDomesticScheduledPaymentConsentsForbidden {
	return &CreateDomesticScheduledPaymentConsentsForbidden{}
}

/* CreateDomesticScheduledPaymentConsentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateDomesticScheduledPaymentConsentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticScheduledPaymentConsentsForbidden) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticScheduledPaymentConsentsNotFound creates a CreateDomesticScheduledPaymentConsentsNotFound with default headers values
func NewCreateDomesticScheduledPaymentConsentsNotFound() *CreateDomesticScheduledPaymentConsentsNotFound {
	return &CreateDomesticScheduledPaymentConsentsNotFound{}
}

/* CreateDomesticScheduledPaymentConsentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateDomesticScheduledPaymentConsentsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticScheduledPaymentConsentsNotFound) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsNotFound ", 404)
}

func (o *CreateDomesticScheduledPaymentConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentsMethodNotAllowed creates a CreateDomesticScheduledPaymentConsentsMethodNotAllowed with default headers values
func NewCreateDomesticScheduledPaymentConsentsMethodNotAllowed() *CreateDomesticScheduledPaymentConsentsMethodNotAllowed {
	return &CreateDomesticScheduledPaymentConsentsMethodNotAllowed{}
}

/* CreateDomesticScheduledPaymentConsentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateDomesticScheduledPaymentConsentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticScheduledPaymentConsentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsMethodNotAllowed ", 405)
}

func (o *CreateDomesticScheduledPaymentConsentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentsNotAcceptable creates a CreateDomesticScheduledPaymentConsentsNotAcceptable with default headers values
func NewCreateDomesticScheduledPaymentConsentsNotAcceptable() *CreateDomesticScheduledPaymentConsentsNotAcceptable {
	return &CreateDomesticScheduledPaymentConsentsNotAcceptable{}
}

/* CreateDomesticScheduledPaymentConsentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateDomesticScheduledPaymentConsentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticScheduledPaymentConsentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsNotAcceptable ", 406)
}

func (o *CreateDomesticScheduledPaymentConsentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentsUnsupportedMediaType creates a CreateDomesticScheduledPaymentConsentsUnsupportedMediaType with default headers values
func NewCreateDomesticScheduledPaymentConsentsUnsupportedMediaType() *CreateDomesticScheduledPaymentConsentsUnsupportedMediaType {
	return &CreateDomesticScheduledPaymentConsentsUnsupportedMediaType{}
}

/* CreateDomesticScheduledPaymentConsentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateDomesticScheduledPaymentConsentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticScheduledPaymentConsentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsUnsupportedMediaType ", 415)
}

func (o *CreateDomesticScheduledPaymentConsentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticScheduledPaymentConsentsTooManyRequests creates a CreateDomesticScheduledPaymentConsentsTooManyRequests with default headers values
func NewCreateDomesticScheduledPaymentConsentsTooManyRequests() *CreateDomesticScheduledPaymentConsentsTooManyRequests {
	return &CreateDomesticScheduledPaymentConsentsTooManyRequests{}
}

/* CreateDomesticScheduledPaymentConsentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateDomesticScheduledPaymentConsentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticScheduledPaymentConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsTooManyRequests ", 429)
}

func (o *CreateDomesticScheduledPaymentConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticScheduledPaymentConsentsInternalServerError creates a CreateDomesticScheduledPaymentConsentsInternalServerError with default headers values
func NewCreateDomesticScheduledPaymentConsentsInternalServerError() *CreateDomesticScheduledPaymentConsentsInternalServerError {
	return &CreateDomesticScheduledPaymentConsentsInternalServerError{}
}

/* CreateDomesticScheduledPaymentConsentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateDomesticScheduledPaymentConsentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticScheduledPaymentConsentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /domestic-scheduled-payment-consents][%d] createDomesticScheduledPaymentConsentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDomesticScheduledPaymentConsentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticScheduledPaymentConsentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
