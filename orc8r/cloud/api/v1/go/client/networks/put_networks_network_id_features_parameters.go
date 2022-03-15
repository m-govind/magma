// Code generated by go-swagger; DO NOT EDIT.

package networks

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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutNetworksNetworkIDFeaturesParams creates a new PutNetworksNetworkIDFeaturesParams object
// with the default values initialized.
func NewPutNetworksNetworkIDFeaturesParams() *PutNetworksNetworkIDFeaturesParams {
	var ()
	return &PutNetworksNetworkIDFeaturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDFeaturesParamsWithTimeout creates a new PutNetworksNetworkIDFeaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDFeaturesParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDFeaturesParams {
	var ()
	return &PutNetworksNetworkIDFeaturesParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDFeaturesParamsWithContext creates a new PutNetworksNetworkIDFeaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDFeaturesParamsWithContext(ctx context.Context) *PutNetworksNetworkIDFeaturesParams {
	var ()
	return &PutNetworksNetworkIDFeaturesParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDFeaturesParamsWithHTTPClient creates a new PutNetworksNetworkIDFeaturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDFeaturesParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDFeaturesParams {
	var ()
	return &PutNetworksNetworkIDFeaturesParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDFeaturesParams contains all the parameters to send to the API endpoint
for the put networks network ID features operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDFeaturesParams struct {

	/*NetworkFeatures
	  New network features for the network

	*/
	NetworkFeatures *models.NetworkFeatures
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) WithContext(ctx context.Context) *PutNetworksNetworkIDFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkFeatures adds the networkFeatures to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) WithNetworkFeatures(networkFeatures *models.NetworkFeatures) *PutNetworksNetworkIDFeaturesParams {
	o.SetNetworkFeatures(networkFeatures)
	return o
}

// SetNetworkFeatures adds the networkFeatures to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) SetNetworkFeatures(networkFeatures *models.NetworkFeatures) {
	o.NetworkFeatures = networkFeatures
}

// WithNetworkID adds the networkID to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) WithNetworkID(networkID string) *PutNetworksNetworkIDFeaturesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID features params
func (o *PutNetworksNetworkIDFeaturesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NetworkFeatures != nil {
		if err := r.SetBodyParam(o.NetworkFeatures); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}