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

// GetV1ServicesServiceIDTypeWebsocketReader is a Reader for the GetV1ServicesServiceIDTypeWebsocket structure.
type GetV1ServicesServiceIDTypeWebsocketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServicesServiceIDTypeWebsocketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ServicesServiceIDTypeWebsocketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetV1ServicesServiceIDTypeWebsocketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetV1ServicesServiceIDTypeWebsocketNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ServicesServiceIDTypeWebsocketOK creates a GetV1ServicesServiceIDTypeWebsocketOK with default headers values
func NewGetV1ServicesServiceIDTypeWebsocketOK() *GetV1ServicesServiceIDTypeWebsocketOK {
	return &GetV1ServicesServiceIDTypeWebsocketOK{}
}

/*GetV1ServicesServiceIDTypeWebsocketOK handles this case with default header values.

Service
*/
type GetV1ServicesServiceIDTypeWebsocketOK struct {
	Payload *models.WebSocketService
}

func (o *GetV1ServicesServiceIDTypeWebsocketOK) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/?type=websocket][%d] getV1ServicesServiceIdTypeWebsocketOK  %+v", 200, o.Payload)
}

func (o *GetV1ServicesServiceIDTypeWebsocketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebSocketService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ServicesServiceIDTypeWebsocketUnauthorized creates a GetV1ServicesServiceIDTypeWebsocketUnauthorized with default headers values
func NewGetV1ServicesServiceIDTypeWebsocketUnauthorized() *GetV1ServicesServiceIDTypeWebsocketUnauthorized {
	return &GetV1ServicesServiceIDTypeWebsocketUnauthorized{}
}

/*GetV1ServicesServiceIDTypeWebsocketUnauthorized handles this case with default header values.

Unauthenticated
*/
type GetV1ServicesServiceIDTypeWebsocketUnauthorized struct {
}

func (o *GetV1ServicesServiceIDTypeWebsocketUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/?type=websocket][%d] getV1ServicesServiceIdTypeWebsocketUnauthorized ", 401)
}

func (o *GetV1ServicesServiceIDTypeWebsocketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetV1ServicesServiceIDTypeWebsocketNotFound creates a GetV1ServicesServiceIDTypeWebsocketNotFound with default headers values
func NewGetV1ServicesServiceIDTypeWebsocketNotFound() *GetV1ServicesServiceIDTypeWebsocketNotFound {
	return &GetV1ServicesServiceIDTypeWebsocketNotFound{}
}

/*GetV1ServicesServiceIDTypeWebsocketNotFound handles this case with default header values.

Not found
*/
type GetV1ServicesServiceIDTypeWebsocketNotFound struct {
}

func (o *GetV1ServicesServiceIDTypeWebsocketNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/?type=websocket][%d] getV1ServicesServiceIdTypeWebsocketNotFound ", 404)
}

func (o *GetV1ServicesServiceIDTypeWebsocketNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
