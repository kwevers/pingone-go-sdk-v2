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

// DeviceAuthenticationPolicyOfflineDeviceOtpFailure struct for DeviceAuthenticationPolicyOfflineDeviceOtpFailure
type DeviceAuthenticationPolicyOfflineDeviceOtpFailure struct {
	// The maximum number of times that the OTP entry can fail for a user, before they are blocked.
	Count int32 `json:"count"`
	CoolDown DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown `json:"coolDown"`
}

// NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(count int32, coolDown DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) *DeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	this := DeviceAuthenticationPolicyOfflineDeviceOtpFailure{}
	this.Count = count
	this.CoolDown = coolDown
	return &this
}

// NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	this := DeviceAuthenticationPolicyOfflineDeviceOtpFailure{}
	return &this
}

// GetCount returns the Count field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) SetCount(v int32) {
	o.Count = v
}

// GetCoolDown returns the CoolDown field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCoolDown() DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown
		return ret
	}

	return o.CoolDown
}

// GetCoolDownOk returns a tuple with the CoolDown field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCoolDownOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CoolDown, true
}

// SetCoolDown sets field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) SetCoolDown(v DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) {
	o.CoolDown = v
}

func (o DeviceAuthenticationPolicyOfflineDeviceOtpFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["coolDown"] = o.CoolDown
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure struct {
	value *DeviceAuthenticationPolicyOfflineDeviceOtpFailure
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure) Get() *DeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure) Set(val *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure(val *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	return &NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


