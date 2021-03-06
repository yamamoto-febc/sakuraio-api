package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// WebSocketService web socket service
// swagger:model WebSocketService
type WebSocketService struct {

	// j w t
	JWT string `json:"JWT,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Project ID
	Project int64 `json:"project,omitempty"`

	// "websocket"
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this web socket service
func (m *WebSocketService) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
