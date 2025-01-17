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

// GetDomesticScheduledPaymentConsentsConsentIDReader is a Reader for the GetDomesticScheduledPaymentConsentsConsentID structure.
type GetDomesticScheduledPaymentConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomesticScheduledPaymentConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomesticScheduledPaymentConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDomesticScheduledPaymentConsentsConsentIDOK creates a GetDomesticScheduledPaymentConsentsConsentIDOK with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDOK() *GetDomesticScheduledPaymentConsentsConsentIDOK {
	return &GetDomesticScheduledPaymentConsentsConsentIDOK{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDOK describes a response with status code 200, with default header values.

Domestic Scheduled Payment Consents Read
*/
type GetDomesticScheduledPaymentConsentsConsentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteDomesticScheduledConsentResponse5
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDOK) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdOK  %+v", 200, o.Payload)
}
func (o *GetDomesticScheduledPaymentConsentsConsentIDOK) GetPayload() *models.OBWriteDomesticScheduledConsentResponse5 {
	return o.Payload
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticScheduledPaymentConsentsConsentIDBadRequest creates a GetDomesticScheduledPaymentConsentsConsentIDBadRequest with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDBadRequest() *GetDomesticScheduledPaymentConsentsConsentIDBadRequest {
	return &GetDomesticScheduledPaymentConsentsConsentIDBadRequest{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetDomesticScheduledPaymentConsentsConsentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDomesticScheduledPaymentConsentsConsentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticScheduledPaymentConsentsConsentIDUnauthorized creates a GetDomesticScheduledPaymentConsentsConsentIDUnauthorized with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDUnauthorized() *GetDomesticScheduledPaymentConsentsConsentIDUnauthorized {
	return &GetDomesticScheduledPaymentConsentsConsentIDUnauthorized{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDomesticScheduledPaymentConsentsConsentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdUnauthorized ", 401)
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticScheduledPaymentConsentsConsentIDForbidden creates a GetDomesticScheduledPaymentConsentsConsentIDForbidden with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDForbidden() *GetDomesticScheduledPaymentConsentsConsentIDForbidden {
	return &GetDomesticScheduledPaymentConsentsConsentIDForbidden{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDomesticScheduledPaymentConsentsConsentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDomesticScheduledPaymentConsentsConsentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticScheduledPaymentConsentsConsentIDNotFound creates a GetDomesticScheduledPaymentConsentsConsentIDNotFound with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDNotFound() *GetDomesticScheduledPaymentConsentsConsentIDNotFound {
	return &GetDomesticScheduledPaymentConsentsConsentIDNotFound{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetDomesticScheduledPaymentConsentsConsentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdNotFound ", 404)
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed creates a GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed() *GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed {
	return &GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdMethodNotAllowed ", 405)
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticScheduledPaymentConsentsConsentIDNotAcceptable creates a GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDNotAcceptable() *GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable {
	return &GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdNotAcceptable ", 406)
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticScheduledPaymentConsentsConsentIDTooManyRequests creates a GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDTooManyRequests() *GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests {
	return &GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdTooManyRequests ", 429)
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticScheduledPaymentConsentsConsentIDInternalServerError creates a GetDomesticScheduledPaymentConsentsConsentIDInternalServerError with default headers values
func NewGetDomesticScheduledPaymentConsentsConsentIDInternalServerError() *GetDomesticScheduledPaymentConsentsConsentIDInternalServerError {
	return &GetDomesticScheduledPaymentConsentsConsentIDInternalServerError{}
}

/* GetDomesticScheduledPaymentConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDomesticScheduledPaymentConsentsConsentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /domestic-scheduled-payment-consents/{ConsentId}][%d] getDomesticScheduledPaymentConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDomesticScheduledPaymentConsentsConsentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticScheduledPaymentConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
