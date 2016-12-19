package module

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1ModulesModuleIDParams creates a new GetV1ModulesModuleIDParams object
// with the default values initialized.
func NewGetV1ModulesModuleIDParams() *GetV1ModulesModuleIDParams {
	var ()
	return &GetV1ModulesModuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ModulesModuleIDParamsWithTimeout creates a new GetV1ModulesModuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1ModulesModuleIDParamsWithTimeout(timeout time.Duration) *GetV1ModulesModuleIDParams {
	var ()
	return &GetV1ModulesModuleIDParams{

		timeout: timeout,
	}
}

// NewGetV1ModulesModuleIDParamsWithContext creates a new GetV1ModulesModuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1ModulesModuleIDParamsWithContext(ctx context.Context) *GetV1ModulesModuleIDParams {
	var ()
	return &GetV1ModulesModuleIDParams{

		Context: ctx,
	}
}

/*GetV1ModulesModuleIDParams contains all the parameters to send to the API endpoint
for the get v1 modules module ID operation typically these are written to a http.Request
*/
type GetV1ModulesModuleIDParams struct {

	/*ModuleID
	  ID of module

	*/
	ModuleID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get v1 modules module ID params
func (o *GetV1ModulesModuleIDParams) WithTimeout(timeout time.Duration) *GetV1ModulesModuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 modules module ID params
func (o *GetV1ModulesModuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 modules module ID params
func (o *GetV1ModulesModuleIDParams) WithContext(ctx context.Context) *GetV1ModulesModuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 modules module ID params
func (o *GetV1ModulesModuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithModuleID adds the moduleID to the get v1 modules module ID params
func (o *GetV1ModulesModuleIDParams) WithModuleID(moduleID string) *GetV1ModulesModuleIDParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the get v1 modules module ID params
func (o *GetV1ModulesModuleIDParams) SetModuleID(moduleID string) {
	o.ModuleID = moduleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ModulesModuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param moduleId
	if err := r.SetPathParam("moduleId", o.ModuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
