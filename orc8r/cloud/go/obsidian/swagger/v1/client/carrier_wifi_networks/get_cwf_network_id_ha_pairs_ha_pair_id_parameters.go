// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewGetCwfNetworkIDHaPairsHaPairIDParams creates a new GetCwfNetworkIDHaPairsHaPairIDParams object
// with the default values initialized.
func NewGetCwfNetworkIDHaPairsHaPairIDParams() *GetCwfNetworkIDHaPairsHaPairIDParams {
	var ()
	return &GetCwfNetworkIDHaPairsHaPairIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDParamsWithTimeout creates a new GetCwfNetworkIDHaPairsHaPairIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCwfNetworkIDHaPairsHaPairIDParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDHaPairsHaPairIDParams {
	var ()
	return &GetCwfNetworkIDHaPairsHaPairIDParams{

		timeout: timeout,
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDParamsWithContext creates a new GetCwfNetworkIDHaPairsHaPairIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCwfNetworkIDHaPairsHaPairIDParamsWithContext(ctx context.Context) *GetCwfNetworkIDHaPairsHaPairIDParams {
	var ()
	return &GetCwfNetworkIDHaPairsHaPairIDParams{

		Context: ctx,
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDParamsWithHTTPClient creates a new GetCwfNetworkIDHaPairsHaPairIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCwfNetworkIDHaPairsHaPairIDParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDHaPairsHaPairIDParams {
	var ()
	return &GetCwfNetworkIDHaPairsHaPairIDParams{
		HTTPClient: client,
	}
}

/*GetCwfNetworkIDHaPairsHaPairIDParams contains all the parameters to send to the API endpoint
for the get cwf network ID ha pairs ha pair ID operation typically these are written to a http.Request
*/
type GetCwfNetworkIDHaPairsHaPairIDParams struct {

	/*HaPairID
	  HA Gateway Pair ID

	*/
	HaPairID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDHaPairsHaPairIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) WithContext(ctx context.Context) *GetCwfNetworkIDHaPairsHaPairIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDHaPairsHaPairIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHaPairID adds the haPairID to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) WithHaPairID(haPairID string) *GetCwfNetworkIDHaPairsHaPairIDParams {
	o.SetHaPairID(haPairID)
	return o
}

// SetHaPairID adds the haPairId to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) SetHaPairID(haPairID string) {
	o.HaPairID = haPairID
}

// WithNetworkID adds the networkID to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) WithNetworkID(networkID string) *GetCwfNetworkIDHaPairsHaPairIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID ha pairs ha pair ID params
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDHaPairsHaPairIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ha_pair_id
	if err := r.SetPathParam("ha_pair_id", o.HaPairID); err != nil {
		return err
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