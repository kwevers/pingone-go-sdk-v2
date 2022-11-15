/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode struct for SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode
type SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode struct {
	// Select this option to allow users to log in when PingOne and or PingID are not available.
	Enabled bool `json:"enabled"`
}

// NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode instantiates a new SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode(enabled bool) *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode {
	this := SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode{}
	this.Enabled = enabled
	return &this
}

// NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineModeWithDefaults instantiates a new SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineModeWithDefaults() *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode {
	this := SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode struct {
	value *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode
	isSet bool
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) Get() *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode {
	return v.value
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) Set(val *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode(val *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode {
	return &NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


