// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateChannelParams creates a new CreateChannelParams object
// with the default values initialized.
func NewCreateChannelParams() *CreateChannelParams {
	var ()
	return &CreateChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChannelParamsWithTimeout creates a new CreateChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateChannelParamsWithTimeout(timeout time.Duration) *CreateChannelParams {
	var ()
	return &CreateChannelParams{

		timeout: timeout,
	}
}

// NewCreateChannelParamsWithContext creates a new CreateChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateChannelParamsWithContext(ctx context.Context) *CreateChannelParams {
	var ()
	return &CreateChannelParams{

		Context: ctx,
	}
}

// NewCreateChannelParamsWithHTTPClient creates a new CreateChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateChannelParamsWithHTTPClient(client *http.Client) *CreateChannelParams {
	var ()
	return &CreateChannelParams{
		HTTPClient: client,
	}
}

/*CreateChannelParams contains all the parameters to send to the API endpoint
for the create channel operation typically these are written to a http.Request
*/
type CreateChannelParams struct {

	/*Name
	  Channel name

	*/
	Name string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create channel params
func (o *CreateChannelParams) WithTimeout(timeout time.Duration) *CreateChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create channel params
func (o *CreateChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create channel params
func (o *CreateChannelParams) WithContext(ctx context.Context) *CreateChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create channel params
func (o *CreateChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create channel params
func (o *CreateChannelParams) WithHTTPClient(client *http.Client) *CreateChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create channel params
func (o *CreateChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the create channel params
func (o *CreateChannelParams) WithName(name string) *CreateChannelParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create channel params
func (o *CreateChannelParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the create channel params
func (o *CreateChannelParams) WithNamespace(namespace string) *CreateChannelParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create channel params
func (o *CreateChannelParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the create channel params
func (o *CreateChannelParams) WithPackage(packageVar string) *CreateChannelParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the create channel params
func (o *CreateChannelParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
