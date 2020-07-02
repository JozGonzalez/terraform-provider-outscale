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

// CreateListenerRuleResponse struct for CreateListenerRuleResponse
type CreateListenerRuleResponse struct {
	ListenerRule    *ListenerRule    `json:"ListenerRule,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetListenerRule returns the ListenerRule field value if set, zero value otherwise.
func (o *CreateListenerRuleResponse) GetListenerRule() ListenerRule {
	if o == nil || o.ListenerRule == nil {
		var ret ListenerRule
		return ret
	}
	return *o.ListenerRule
}

// GetListenerRuleOk returns a tuple with the ListenerRule field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateListenerRuleResponse) GetListenerRuleOk() (ListenerRule, bool) {
	if o == nil || o.ListenerRule == nil {
		var ret ListenerRule
		return ret, false
	}
	return *o.ListenerRule, true
}

// HasListenerRule returns a boolean if a field has been set.
func (o *CreateListenerRuleResponse) HasListenerRule() bool {
	if o != nil && o.ListenerRule != nil {
		return true
	}

	return false
}

// SetListenerRule gets a reference to the given ListenerRule and assigns it to the ListenerRule field.
func (o *CreateListenerRuleResponse) SetListenerRule(v ListenerRule) {
	o.ListenerRule = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateListenerRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateListenerRuleResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateListenerRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateListenerRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableCreateListenerRuleResponse struct {
	Value        CreateListenerRuleResponse
	ExplicitNull bool
}

func (v NullableCreateListenerRuleResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateListenerRuleResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
