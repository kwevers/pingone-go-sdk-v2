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

// SignOnPolicyActionAgreementAllOf struct for SignOnPolicyActionAgreementAllOf
type SignOnPolicyActionAgreementAllOf struct {
	Agreement SignOnPolicyActionAgreementAllOfAgreement `json:"agreement"`
	// When enabled, the `Do Not Accept` button will terminate the Flow and display an error message to the user.
	DisableDeclineOption *bool `json:"disableDeclineOption,omitempty"`
}

// NewSignOnPolicyActionAgreementAllOf instantiates a new SignOnPolicyActionAgreementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionAgreementAllOf(agreement SignOnPolicyActionAgreementAllOfAgreement) *SignOnPolicyActionAgreementAllOf {
	this := SignOnPolicyActionAgreementAllOf{}
	this.Agreement = agreement
	return &this
}

// NewSignOnPolicyActionAgreementAllOfWithDefaults instantiates a new SignOnPolicyActionAgreementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionAgreementAllOfWithDefaults() *SignOnPolicyActionAgreementAllOf {
	this := SignOnPolicyActionAgreementAllOf{}
	return &this
}

// GetAgreement returns the Agreement field value
func (o *SignOnPolicyActionAgreementAllOf) GetAgreement() SignOnPolicyActionAgreementAllOfAgreement {
	if o == nil {
		var ret SignOnPolicyActionAgreementAllOfAgreement
		return ret
	}

	return o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreementAllOf) GetAgreementOk() (*SignOnPolicyActionAgreementAllOfAgreement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agreement, true
}

// SetAgreement sets field value
func (o *SignOnPolicyActionAgreementAllOf) SetAgreement(v SignOnPolicyActionAgreementAllOfAgreement) {
	o.Agreement = v
}

// GetDisableDeclineOption returns the DisableDeclineOption field value if set, zero value otherwise.
func (o *SignOnPolicyActionAgreementAllOf) GetDisableDeclineOption() bool {
	if o == nil || o.DisableDeclineOption == nil {
		var ret bool
		return ret
	}
	return *o.DisableDeclineOption
}

// GetDisableDeclineOptionOk returns a tuple with the DisableDeclineOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreementAllOf) GetDisableDeclineOptionOk() (*bool, bool) {
	if o == nil || o.DisableDeclineOption == nil {
		return nil, false
	}
	return o.DisableDeclineOption, true
}

// HasDisableDeclineOption returns a boolean if a field has been set.
func (o *SignOnPolicyActionAgreementAllOf) HasDisableDeclineOption() bool {
	if o != nil && o.DisableDeclineOption != nil {
		return true
	}

	return false
}

// SetDisableDeclineOption gets a reference to the given bool and assigns it to the DisableDeclineOption field.
func (o *SignOnPolicyActionAgreementAllOf) SetDisableDeclineOption(v bool) {
	o.DisableDeclineOption = &v
}

func (o SignOnPolicyActionAgreementAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["agreement"] = o.Agreement
	}
	if o.DisableDeclineOption != nil {
		toSerialize["disableDeclineOption"] = o.DisableDeclineOption
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionAgreementAllOf struct {
	value *SignOnPolicyActionAgreementAllOf
	isSet bool
}

func (v NullableSignOnPolicyActionAgreementAllOf) Get() *SignOnPolicyActionAgreementAllOf {
	return v.value
}

func (v *NullableSignOnPolicyActionAgreementAllOf) Set(val *SignOnPolicyActionAgreementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionAgreementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionAgreementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionAgreementAllOf(val *SignOnPolicyActionAgreementAllOf) *NullableSignOnPolicyActionAgreementAllOf {
	return &NullableSignOnPolicyActionAgreementAllOf{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionAgreementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionAgreementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


