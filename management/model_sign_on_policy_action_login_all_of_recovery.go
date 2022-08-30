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

// SignOnPolicyActionLoginAllOfRecovery Specifies the account recovery options.
type SignOnPolicyActionLoginAllOfRecovery struct {
	// A boolean that specifies the enabled/disabled state of the account recovery action. The default is disabled when creating a new policy. When enabled, it allows the use of the forgot password flow.
	Enabled bool `json:"enabled"`
}

// NewSignOnPolicyActionLoginAllOfRecovery instantiates a new SignOnPolicyActionLoginAllOfRecovery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginAllOfRecovery(enabled bool) *SignOnPolicyActionLoginAllOfRecovery {
	this := SignOnPolicyActionLoginAllOfRecovery{}
	this.Enabled = enabled
	return &this
}

// NewSignOnPolicyActionLoginAllOfRecoveryWithDefaults instantiates a new SignOnPolicyActionLoginAllOfRecovery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginAllOfRecoveryWithDefaults() *SignOnPolicyActionLoginAllOfRecovery {
	this := SignOnPolicyActionLoginAllOfRecovery{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SignOnPolicyActionLoginAllOfRecovery) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRecovery) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SignOnPolicyActionLoginAllOfRecovery) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SignOnPolicyActionLoginAllOfRecovery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionLoginAllOfRecovery struct {
	value *SignOnPolicyActionLoginAllOfRecovery
	isSet bool
}

func (v NullableSignOnPolicyActionLoginAllOfRecovery) Get() *SignOnPolicyActionLoginAllOfRecovery {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginAllOfRecovery) Set(val *SignOnPolicyActionLoginAllOfRecovery) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginAllOfRecovery) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginAllOfRecovery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginAllOfRecovery(val *SignOnPolicyActionLoginAllOfRecovery) *NullableSignOnPolicyActionLoginAllOfRecovery {
	return &NullableSignOnPolicyActionLoginAllOfRecovery{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginAllOfRecovery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginAllOfRecovery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


