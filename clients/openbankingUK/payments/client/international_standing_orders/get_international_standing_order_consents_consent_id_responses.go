// Code generated by go-swagger; DO NOT EDIT.

package international_standing_orders

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

// GetInternationalStandingOrderConsentsConsentIDReader is a Reader for the GetInternationalStandingOrderConsentsConsentID structure.
type GetInternationalStandingOrderConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternationalStandingOrderConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternationalStandingOrderConsentsConsentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInternationalStandingOrderConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInternationalStandingOrderConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInternationalStandingOrderConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInternationalStandingOrderConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInternationalStandingOrderConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetInternationalStandingOrderConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInternationalStandingOrderConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternationalStandingOrderConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInternationalStandingOrderConsentsConsentIDOK creates a GetInternationalStandingOrderConsentsConsentIDOK with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDOK() *GetInternationalStandingOrderConsentsConsentIDOK {
	return &GetInternationalStandingOrderConsentsConsentIDOK{}
}

/* GetInternationalStandingOrderConsentsConsentIDOK describes a response with status code 200, with default header values.

International Standing Order Consents Read
*/
type GetInternationalStandingOrderConsentsConsentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteInternationalStandingOrderConsentResponse7
}

func (o *GetInternationalStandingOrderConsentsConsentIDOK) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdOK  %+v", 200, o.Payload)
}
func (o *GetInternationalStandingOrderConsentsConsentIDOK) GetPayload() *models.OBWriteInternationalStandingOrderConsentResponse7 {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentsConsentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteInternationalStandingOrderConsentResponse7)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternationalStandingOrderConsentsConsentIDBadRequest creates a GetInternationalStandingOrderConsentsConsentIDBadRequest with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDBadRequest() *GetInternationalStandingOrderConsentsConsentIDBadRequest {
	return &GetInternationalStandingOrderConsentsConsentIDBadRequest{}
}

/* GetInternationalStandingOrderConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetInternationalStandingOrderConsentsConsentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalStandingOrderConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetInternationalStandingOrderConsentsConsentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrderConsentsConsentIDUnauthorized creates a GetInternationalStandingOrderConsentsConsentIDUnauthorized with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDUnauthorized() *GetInternationalStandingOrderConsentsConsentIDUnauthorized {
	return &GetInternationalStandingOrderConsentsConsentIDUnauthorized{}
}

/* GetInternationalStandingOrderConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInternationalStandingOrderConsentsConsentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrderConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdUnauthorized ", 401)
}

func (o *GetInternationalStandingOrderConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrderConsentsConsentIDForbidden creates a GetInternationalStandingOrderConsentsConsentIDForbidden with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDForbidden() *GetInternationalStandingOrderConsentsConsentIDForbidden {
	return &GetInternationalStandingOrderConsentsConsentIDForbidden{}
}

/* GetInternationalStandingOrderConsentsConsentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInternationalStandingOrderConsentsConsentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalStandingOrderConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetInternationalStandingOrderConsentsConsentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrderConsentsConsentIDNotFound creates a GetInternationalStandingOrderConsentsConsentIDNotFound with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDNotFound() *GetInternationalStandingOrderConsentsConsentIDNotFound {
	return &GetInternationalStandingOrderConsentsConsentIDNotFound{}
}

/* GetInternationalStandingOrderConsentsConsentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetInternationalStandingOrderConsentsConsentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrderConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdNotFound ", 404)
}

func (o *GetInternationalStandingOrderConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrderConsentsConsentIDMethodNotAllowed creates a GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDMethodNotAllowed() *GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed {
	return &GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed{}
}

/* GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdMethodNotAllowed ", 405)
}

func (o *GetInternationalStandingOrderConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrderConsentsConsentIDNotAcceptable creates a GetInternationalStandingOrderConsentsConsentIDNotAcceptable with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDNotAcceptable() *GetInternationalStandingOrderConsentsConsentIDNotAcceptable {
	return &GetInternationalStandingOrderConsentsConsentIDNotAcceptable{}
}

/* GetInternationalStandingOrderConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetInternationalStandingOrderConsentsConsentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrderConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdNotAcceptable ", 406)
}

func (o *GetInternationalStandingOrderConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrderConsentsConsentIDTooManyRequests creates a GetInternationalStandingOrderConsentsConsentIDTooManyRequests with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDTooManyRequests() *GetInternationalStandingOrderConsentsConsentIDTooManyRequests {
	return &GetInternationalStandingOrderConsentsConsentIDTooManyRequests{}
}

/* GetInternationalStandingOrderConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetInternationalStandingOrderConsentsConsentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrderConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdTooManyRequests ", 429)
}

func (o *GetInternationalStandingOrderConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrderConsentsConsentIDInternalServerError creates a GetInternationalStandingOrderConsentsConsentIDInternalServerError with default headers values
func NewGetInternationalStandingOrderConsentsConsentIDInternalServerError() *GetInternationalStandingOrderConsentsConsentIDInternalServerError {
	return &GetInternationalStandingOrderConsentsConsentIDInternalServerError{}
}

/* GetInternationalStandingOrderConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInternationalStandingOrderConsentsConsentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalStandingOrderConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /international-standing-order-consents/{ConsentId}][%d] getInternationalStandingOrderConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetInternationalStandingOrderConsentsConsentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalStandingOrderConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
