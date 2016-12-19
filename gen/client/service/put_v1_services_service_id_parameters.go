package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutV1ServicesServiceIDParams creates a new PutV1ServicesServiceIDParams object
// with the default values initialized.
func NewPutV1ServicesServiceIDParams() *PutV1ServicesServiceIDParams {
	var ()
	return &PutV1ServicesServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1ServicesServiceIDParamsWithTimeout creates a new PutV1ServicesServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutV1ServicesServiceIDParamsWithTimeout(timeout time.Duration) *PutV1ServicesServiceIDParams {
	var ()
	return &PutV1ServicesServiceIDParams{

		timeout: timeout,
	}
}

// NewPutV1ServicesServiceIDParamsWithContext creates a new PutV1ServicesServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutV1ServicesServiceIDParamsWithContext(ctx context.Context) *PutV1ServicesServiceIDParams {
	var ()
	return &PutV1ServicesServiceIDParams{

		Context: ctx,
	}
}

/*PutV1ServicesServiceIDParams contains all the parameters to send to the API endpoint
for the put v1 services service ID operation typically these are written to a http.Request
*/
type PutV1ServicesServiceIDParams struct {

	/*Service*/
	Service interface{}
	/*ServiceID
	  ID of service

	*/
	ServiceID int64

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) WithTimeout(timeout time.Duration) *PutV1ServicesServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) WithContext(ctx context.Context) *PutV1ServicesServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithService adds the service to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) WithService(service interface{}) *PutV1ServicesServiceIDParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) SetService(service interface{}) {
	o.Service = service
}

// WithServiceID adds the serviceID to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) WithServiceID(serviceID int64) *PutV1ServicesServiceIDParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the put v1 services service ID params
func (o *PutV1ServicesServiceIDParams) SetServiceID(serviceID int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1ServicesServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Service); err != nil {
		return err
	}

	// path param serviceId
	if err := r.SetPathParam("serviceId", swag.FormatInt64(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
