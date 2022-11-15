/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment struct for DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment
type DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment struct {
	// Set to true if you want the application to allow Auto Enrollment. Auto Enrollment means that the user can authenticate for the first time from an unpaired device, and the successful authentication will result in the pairing of the device for MFA.
	Enabled bool `json:"enabled"`
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment(enabled bool) *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment{}
	this.Enabled = enabled
	return &this
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollmentWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollmentWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) SetEnabled(v bool) {
	o.Enabled = v
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment struct {
	value *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) Get() *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) Set(val *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment(val *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) *NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment {
	return &NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


