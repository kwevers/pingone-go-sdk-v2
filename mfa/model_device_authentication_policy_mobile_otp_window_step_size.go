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

// DeviceAuthenticationPolicyMobileOtpWindowStepSize struct for DeviceAuthenticationPolicyMobileOtpWindowStepSize
type DeviceAuthenticationPolicyMobileOtpWindowStepSize struct {
	Duration int32 `json:"duration"`
	TimeUnit EnumTimeUnit `json:"timeUnit"`
}

// NewDeviceAuthenticationPolicyMobileOtpWindowStepSize instantiates a new DeviceAuthenticationPolicyMobileOtpWindowStepSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileOtpWindowStepSize(duration int32, timeUnit EnumTimeUnit) *DeviceAuthenticationPolicyMobileOtpWindowStepSize {
	this := DeviceAuthenticationPolicyMobileOtpWindowStepSize{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewDeviceAuthenticationPolicyMobileOtpWindowStepSizeWithDefaults instantiates a new DeviceAuthenticationPolicyMobileOtpWindowStepSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileOtpWindowStepSizeWithDefaults() *DeviceAuthenticationPolicyMobileOtpWindowStepSize {
	this := DeviceAuthenticationPolicyMobileOtpWindowStepSize{}
	return &this
}

// GetDuration returns the Duration field value
func (o *DeviceAuthenticationPolicyMobileOtpWindowStepSize) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileOtpWindowStepSize) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeviceAuthenticationPolicyMobileOtpWindowStepSize) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *DeviceAuthenticationPolicyMobileOtpWindowStepSize) GetTimeUnit() EnumTimeUnit {
	if o == nil {
		var ret EnumTimeUnit
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileOtpWindowStepSize) GetTimeUnitOk() (*EnumTimeUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *DeviceAuthenticationPolicyMobileOtpWindowStepSize) SetTimeUnit(v EnumTimeUnit) {
	o.TimeUnit = v
}

func (o DeviceAuthenticationPolicyMobileOtpWindowStepSize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize struct {
	value *DeviceAuthenticationPolicyMobileOtpWindowStepSize
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize) Get() *DeviceAuthenticationPolicyMobileOtpWindowStepSize {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize) Set(val *DeviceAuthenticationPolicyMobileOtpWindowStepSize) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileOtpWindowStepSize(val *DeviceAuthenticationPolicyMobileOtpWindowStepSize) *NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize {
	return &NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindowStepSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

