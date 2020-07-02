/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// ReadNetAccessPointsResponse struct for ReadNetAccessPointsResponse
type ReadNetAccessPointsResponse struct {
	// One or more Net access points.
	NetAccessPoints *[]NetAccessPoint `json:"NetAccessPoints,omitempty"`
	ResponseContext *ResponseContext  `json:"ResponseContext,omitempty"`
}

// GetNetAccessPoints returns the NetAccessPoints field value if set, zero value otherwise.
func (o *ReadNetAccessPointsResponse) GetNetAccessPoints() []NetAccessPoint {
	if o == nil || o.NetAccessPoints == nil {
		var ret []NetAccessPoint
		return ret
	}
	return *o.NetAccessPoints
}

// GetNetAccessPointsOk returns a tuple with the NetAccessPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetAccessPointsResponse) GetNetAccessPointsOk() ([]NetAccessPoint, bool) {
	if o == nil || o.NetAccessPoints == nil {
		var ret []NetAccessPoint
		return ret, false
	}
	return *o.NetAccessPoints, true
}

// HasNetAccessPoints returns a boolean if a field has been set.
func (o *ReadNetAccessPointsResponse) HasNetAccessPoints() bool {
	if o != nil && o.NetAccessPoints != nil {
		return true
	}

	return false
}

// SetNetAccessPoints gets a reference to the given []NetAccessPoint and assigns it to the NetAccessPoints field.
func (o *ReadNetAccessPointsResponse) SetNetAccessPoints(v []NetAccessPoint) {
	o.NetAccessPoints = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNetAccessPointsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetAccessPointsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNetAccessPointsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNetAccessPointsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadNetAccessPointsResponse struct {
	Value        ReadNetAccessPointsResponse
	ExplicitNull bool
}

func (v NullableReadNetAccessPointsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadNetAccessPointsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
