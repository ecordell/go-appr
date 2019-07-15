// Code generated by go-swagger; DO NOT EDIT.

package blobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPullBlobJSONParams creates a new PullBlobJSONParams object
// with the default values initialized.
func NewPullBlobJSONParams() *PullBlobJSONParams {
	var (
		formatDefault = string("gzip")
	)
	return &PullBlobJSONParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPullBlobJSONParamsWithTimeout creates a new PullBlobJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPullBlobJSONParamsWithTimeout(timeout time.Duration) *PullBlobJSONParams {
	var (
		formatDefault = string("gzip")
	)
	return &PullBlobJSONParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewPullBlobJSONParamsWithContext creates a new PullBlobJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewPullBlobJSONParamsWithContext(ctx context.Context) *PullBlobJSONParams {
	var (
		formatDefault = string("gzip")
	)
	return &PullBlobJSONParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewPullBlobJSONParamsWithHTTPClient creates a new PullBlobJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPullBlobJSONParamsWithHTTPClient(client *http.Client) *PullBlobJSONParams {
	var (
		formatDefault = string("gzip")
	)
	return &PullBlobJSONParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*PullBlobJSONParams contains all the parameters to send to the API endpoint
for the pull blob Json operation typically these are written to a http.Request
*/
type PullBlobJSONParams struct {

	/*Digest
	  content digest

	*/
	Digest string
	/*Format
	  return format type(json or gzip)

	*/
	Format *string
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

// WithTimeout adds the timeout to the pull blob Json params
func (o *PullBlobJSONParams) WithTimeout(timeout time.Duration) *PullBlobJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pull blob Json params
func (o *PullBlobJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pull blob Json params
func (o *PullBlobJSONParams) WithContext(ctx context.Context) *PullBlobJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pull blob Json params
func (o *PullBlobJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pull blob Json params
func (o *PullBlobJSONParams) WithHTTPClient(client *http.Client) *PullBlobJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pull blob Json params
func (o *PullBlobJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigest adds the digest to the pull blob Json params
func (o *PullBlobJSONParams) WithDigest(digest string) *PullBlobJSONParams {
	o.SetDigest(digest)
	return o
}

// SetDigest adds the digest to the pull blob Json params
func (o *PullBlobJSONParams) SetDigest(digest string) {
	o.Digest = digest
}

// WithFormat adds the format to the pull blob Json params
func (o *PullBlobJSONParams) WithFormat(format *string) *PullBlobJSONParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the pull blob Json params
func (o *PullBlobJSONParams) SetFormat(format *string) {
	o.Format = format
}

// WithNamespace adds the namespace to the pull blob Json params
func (o *PullBlobJSONParams) WithNamespace(namespace string) *PullBlobJSONParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the pull blob Json params
func (o *PullBlobJSONParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the pull blob Json params
func (o *PullBlobJSONParams) WithPackage(packageVar string) *PullBlobJSONParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the pull blob Json params
func (o *PullBlobJSONParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WriteToRequest writes these params to a swagger request
func (o *PullBlobJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param digest
	if err := r.SetPathParam("digest", o.Digest); err != nil {
		return err
	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
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
