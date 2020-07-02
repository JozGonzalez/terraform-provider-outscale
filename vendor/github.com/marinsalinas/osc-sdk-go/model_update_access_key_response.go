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

// UpdateAccessKeyResponse struct for UpdateAccessKeyResponse
type UpdateAccessKeyResponse struct {
	AccessKey       *AccessKey       `json:"AccessKey,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *UpdateAccessKeyResponse) GetAccessKey() AccessKey {
	if o == nil || o.AccessKey == nil {
		var ret AccessKey
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessKeyResponse) GetAccessKeyOk() (AccessKey, bool) {
	if o == nil || o.AccessKey == nil {
		var ret AccessKey
		return ret, false
	}
	return *o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *UpdateAccessKeyResponse) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given AccessKey and assigns it to the AccessKey field.
func (o *UpdateAccessKeyResponse) SetAccessKey(v AccessKey) {
	o.AccessKey = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateAccessKeyResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessKeyResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateAccessKeyResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateAccessKeyResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableUpdateAccessKeyResponse struct {
	Value        UpdateAccessKeyResponse
	ExplicitNull bool
}

func (v NullableUpdateAccessKeyResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateAccessKeyResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
