package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

// GetV1ServicesServiceIDTypeIncomingWebhookReader is a Reader for the GetV1ServicesServiceIDTypeIncomingWebhook structure.
type GetV1ServicesServiceIDTypeIncomingWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServicesServiceIDTypeIncomingWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ServicesServiceIDTypeIncomingWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetV1ServicesServiceIDTypeIncomingWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetV1ServicesServiceIDTypeIncomingWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ServicesServiceIDTypeIncomingWebhookOK creates a GetV1ServicesServiceIDTypeIncomingWebhookOK with default headers values
func NewGetV1ServicesServiceIDTypeIncomingWebhookOK() *GetV1ServicesServiceIDTypeIncomingWebhookOK {
	return &GetV1ServicesServiceIDTypeIncomingWebhookOK{}
}

/*GetV1ServicesServiceIDTypeIncomingWebhookOK handles this case with default header values.

Service
*/
type GetV1ServicesServiceIDTypeIncomingWebhookOK struct {
	Payload *models.IncomingWebhookService
}

func (o *GetV1ServicesServiceIDTypeIncomingWebhookOK) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/?type=incoming-webhook][%d] getV1ServicesServiceIdTypeIncomingWebhookOK  %+v", 200, o.Payload)
}

func (o *GetV1ServicesServiceIDTypeIncomingWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncomingWebhookService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ServicesServiceIDTypeIncomingWebhookUnauthorized creates a GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized with default headers values
func NewGetV1ServicesServiceIDTypeIncomingWebhookUnauthorized() *GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized {
	return &GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized{}
}

/*GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized handles this case with default header values.

Unauthenticated
*/
type GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized struct {
}

func (o *GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/?type=incoming-webhook][%d] getV1ServicesServiceIdTypeIncomingWebhookUnauthorized ", 401)
}

func (o *GetV1ServicesServiceIDTypeIncomingWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetV1ServicesServiceIDTypeIncomingWebhookNotFound creates a GetV1ServicesServiceIDTypeIncomingWebhookNotFound with default headers values
func NewGetV1ServicesServiceIDTypeIncomingWebhookNotFound() *GetV1ServicesServiceIDTypeIncomingWebhookNotFound {
	return &GetV1ServicesServiceIDTypeIncomingWebhookNotFound{}
}

/*GetV1ServicesServiceIDTypeIncomingWebhookNotFound handles this case with default header values.

Not found
*/
type GetV1ServicesServiceIDTypeIncomingWebhookNotFound struct {
}

func (o *GetV1ServicesServiceIDTypeIncomingWebhookNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/?type=incoming-webhook][%d] getV1ServicesServiceIdTypeIncomingWebhookNotFound ", 404)
}

func (o *GetV1ServicesServiceIDTypeIncomingWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
