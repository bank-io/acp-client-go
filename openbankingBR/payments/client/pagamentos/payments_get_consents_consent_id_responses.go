// Code generated by go-swagger; DO NOT EDIT.

package pagamentos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/openbankingBR/payments/models"
)

// PaymentsGetConsentsConsentIDReader is a Reader for the PaymentsGetConsentsConsentID structure.
type PaymentsGetConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentsGetConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentsGetConsentsConsentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentsGetConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentsGetConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentsGetConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentsGetConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPaymentsGetConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewPaymentsGetConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPaymentsGetConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentsGetConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPaymentsGetConsentsConsentIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentsGetConsentsConsentIDOK creates a PaymentsGetConsentsConsentIDOK with default headers values
func NewPaymentsGetConsentsConsentIDOK() *PaymentsGetConsentsConsentIDOK {
	return &PaymentsGetConsentsConsentIDOK{}
}

/* PaymentsGetConsentsConsentIDOK describes a response with status code 200, with default header values.

Dados do consentimento de pagamento obtidos com sucesso.
*/
type PaymentsGetConsentsConsentIDOK struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload interface{}
}

func (o *PaymentsGetConsentsConsentIDOK) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdOK  %+v", 200, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDBadRequest creates a PaymentsGetConsentsConsentIDBadRequest with default headers values
func NewPaymentsGetConsentsConsentIDBadRequest() *PaymentsGetConsentsConsentIDBadRequest {
	return &PaymentsGetConsentsConsentIDBadRequest{}
}

/* PaymentsGetConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

A requisio foi malformada, omitindo atributos obrigatrios, seja no payload ou atravs de atributos na URL.
*/
type PaymentsGetConsentsConsentIDBadRequest struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDBadRequest) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDUnauthorized creates a PaymentsGetConsentsConsentIDUnauthorized with default headers values
func NewPaymentsGetConsentsConsentIDUnauthorized() *PaymentsGetConsentsConsentIDUnauthorized {
	return &PaymentsGetConsentsConsentIDUnauthorized{}
}

/* PaymentsGetConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Cabealho de autenticao ausente/invlido ou token invlido
*/
type PaymentsGetConsentsConsentIDUnauthorized struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDUnauthorized) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDForbidden creates a PaymentsGetConsentsConsentIDForbidden with default headers values
func NewPaymentsGetConsentsConsentIDForbidden() *PaymentsGetConsentsConsentIDForbidden {
	return &PaymentsGetConsentsConsentIDForbidden{}
}

/* PaymentsGetConsentsConsentIDForbidden describes a response with status code 403, with default header values.

O token tem escopo incorreto ou uma poltica de segurana foi violada
*/
type PaymentsGetConsentsConsentIDForbidden struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDForbidden) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDNotFound creates a PaymentsGetConsentsConsentIDNotFound with default headers values
func NewPaymentsGetConsentsConsentIDNotFound() *PaymentsGetConsentsConsentIDNotFound {
	return &PaymentsGetConsentsConsentIDNotFound{}
}

/* PaymentsGetConsentsConsentIDNotFound describes a response with status code 404, with default header values.

O recurso solicitado no existe ou no foi implementado
*/
type PaymentsGetConsentsConsentIDNotFound struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdNotFound  %+v", 404, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDNotFound) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDMethodNotAllowed creates a PaymentsGetConsentsConsentIDMethodNotAllowed with default headers values
func NewPaymentsGetConsentsConsentIDMethodNotAllowed() *PaymentsGetConsentsConsentIDMethodNotAllowed {
	return &PaymentsGetConsentsConsentIDMethodNotAllowed{}
}

/* PaymentsGetConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

O consumidor tentou acessar o recurso com um mtodo no suportado
*/
type PaymentsGetConsentsConsentIDMethodNotAllowed struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDMethodNotAllowed) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDNotAcceptable creates a PaymentsGetConsentsConsentIDNotAcceptable with default headers values
func NewPaymentsGetConsentsConsentIDNotAcceptable() *PaymentsGetConsentsConsentIDNotAcceptable {
	return &PaymentsGetConsentsConsentIDNotAcceptable{}
}

/* PaymentsGetConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

A solicitao continha um cabealho Accept diferente dos tipos de mdia permitidos ou um conjunto de caracteres diferente de UTF-8
*/
type PaymentsGetConsentsConsentIDNotAcceptable struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdNotAcceptable  %+v", 406, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDNotAcceptable) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDTooManyRequests creates a PaymentsGetConsentsConsentIDTooManyRequests with default headers values
func NewPaymentsGetConsentsConsentIDTooManyRequests() *PaymentsGetConsentsConsentIDTooManyRequests {
	return &PaymentsGetConsentsConsentIDTooManyRequests{}
}

/* PaymentsGetConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

A operao foi recusada, pois muitas solicitaes foram feitas dentro de um determinado perodo ou o limite global de requisies concorrentes foi atingido
*/
type PaymentsGetConsentsConsentIDTooManyRequests struct {

	/* Cabealho que indica o tempo (em segundos) que o cliente dever aguardar para realizar uma nova tentativa de chamada. Este cabealho dever estar presente quando o cdigo HTTP de retorno for 429 Too many requests.

	 */
	RetryAfter string

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDTooManyRequests) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		o.RetryAfter = hdrRetryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDInternalServerError creates a PaymentsGetConsentsConsentIDInternalServerError with default headers values
func NewPaymentsGetConsentsConsentIDInternalServerError() *PaymentsGetConsentsConsentIDInternalServerError {
	return &PaymentsGetConsentsConsentIDInternalServerError{}
}

/* PaymentsGetConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Ocorreu um erro no gateway da API ou no microsservio
*/
type PaymentsGetConsentsConsentIDInternalServerError struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.ResponseError
}

func (o *PaymentsGetConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDInternalServerError) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsGetConsentsConsentIDDefault creates a PaymentsGetConsentsConsentIDDefault with default headers values
func NewPaymentsGetConsentsConsentIDDefault(code int) *PaymentsGetConsentsConsentIDDefault {
	return &PaymentsGetConsentsConsentIDDefault{
		_statusCode: code,
	}
}

/* PaymentsGetConsentsConsentIDDefault describes a response with status code -1, with default header values.

Erro inesperado.
*/
type PaymentsGetConsentsConsentIDDefault struct {
	_statusCode int

	Payload *models.ResponseError
}

// Code gets the status code for the payments get consents consent Id default response
func (o *PaymentsGetConsentsConsentIDDefault) Code() int {
	return o._statusCode
}

func (o *PaymentsGetConsentsConsentIDDefault) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] paymentsGetConsentsConsentId default  %+v", o._statusCode, o.Payload)
}
func (o *PaymentsGetConsentsConsentIDDefault) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *PaymentsGetConsentsConsentIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
