package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

// PostV1ProjectsReader is a Reader for the PostV1Projects structure.
type PostV1ProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1ProjectsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostV1ProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostV1ProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostV1ProjectsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1ProjectsCreated creates a PostV1ProjectsCreated with default headers values
func NewPostV1ProjectsCreated() *PostV1ProjectsCreated {
	return &PostV1ProjectsCreated{}
}

/*PostV1ProjectsCreated handles this case with default header values.

Created project
*/
type PostV1ProjectsCreated struct {
	Payload *models.Project
}

func (o *PostV1ProjectsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/projects/][%d] postV1ProjectsCreated  %+v", 201, o.Payload)
}

func (o *PostV1ProjectsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ProjectsUnauthorized creates a PostV1ProjectsUnauthorized with default headers values
func NewPostV1ProjectsUnauthorized() *PostV1ProjectsUnauthorized {
	return &PostV1ProjectsUnauthorized{}
}

/*PostV1ProjectsUnauthorized handles this case with default header values.

Unauthenticated
*/
type PostV1ProjectsUnauthorized struct {
}

func (o *PostV1ProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/projects/][%d] postV1ProjectsUnauthorized ", 401)
}

func (o *PostV1ProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostV1ProjectsForbidden creates a PostV1ProjectsForbidden with default headers values
func NewPostV1ProjectsForbidden() *PostV1ProjectsForbidden {
	return &PostV1ProjectsForbidden{}
}

/*PostV1ProjectsForbidden handles this case with default header values.

Forbidden
*/
type PostV1ProjectsForbidden struct {
}

func (o *PostV1ProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/projects/][%d] postV1ProjectsForbidden ", 403)
}

func (o *PostV1ProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostV1ProjectsUnprocessableEntity creates a PostV1ProjectsUnprocessableEntity with default headers values
func NewPostV1ProjectsUnprocessableEntity() *PostV1ProjectsUnprocessableEntity {
	return &PostV1ProjectsUnprocessableEntity{}
}

/*PostV1ProjectsUnprocessableEntity handles this case with default header values.

Validation error
*/
type PostV1ProjectsUnprocessableEntity struct {
}

func (o *PostV1ProjectsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/projects/][%d] postV1ProjectsUnprocessableEntity ", 422)
}

func (o *PostV1ProjectsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
