/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-05-26
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the OTPDeviceConfigurationOtpDeliveriesCooldown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTPDeviceConfigurationOtpDeliveriesCooldown{}

// OTPDeviceConfigurationOtpDeliveriesCooldown Cooldown (waiting period between OTP attempts) configuration.
type OTPDeviceConfigurationOtpDeliveriesCooldown struct {
	// Cooldown duration configuration.
	Duration int32 `json:"duration"`
	TimeUnit EnumTimeUnit `json:"timeUnit"`
}

// NewOTPDeviceConfigurationOtpDeliveriesCooldown instantiates a new OTPDeviceConfigurationOtpDeliveriesCooldown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTPDeviceConfigurationOtpDeliveriesCooldown(duration int32, timeUnit EnumTimeUnit) *OTPDeviceConfigurationOtpDeliveriesCooldown {
	this := OTPDeviceConfigurationOtpDeliveriesCooldown{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewOTPDeviceConfigurationOtpDeliveriesCooldownWithDefaults instantiates a new OTPDeviceConfigurationOtpDeliveriesCooldown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTPDeviceConfigurationOtpDeliveriesCooldownWithDefaults() *OTPDeviceConfigurationOtpDeliveriesCooldown {
	this := OTPDeviceConfigurationOtpDeliveriesCooldown{}
	return &this
}

// GetDuration returns the Duration field value
func (o *OTPDeviceConfigurationOtpDeliveriesCooldown) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtpDeliveriesCooldown) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *OTPDeviceConfigurationOtpDeliveriesCooldown) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *OTPDeviceConfigurationOtpDeliveriesCooldown) GetTimeUnit() EnumTimeUnit {
	if o == nil {
		var ret EnumTimeUnit
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtpDeliveriesCooldown) GetTimeUnitOk() (*EnumTimeUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *OTPDeviceConfigurationOtpDeliveriesCooldown) SetTimeUnit(v EnumTimeUnit) {
	o.TimeUnit = v
}

func (o OTPDeviceConfigurationOtpDeliveriesCooldown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTPDeviceConfigurationOtpDeliveriesCooldown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["timeUnit"] = o.TimeUnit
	return toSerialize, nil
}

type NullableOTPDeviceConfigurationOtpDeliveriesCooldown struct {
	value *OTPDeviceConfigurationOtpDeliveriesCooldown
	isSet bool
}

func (v NullableOTPDeviceConfigurationOtpDeliveriesCooldown) Get() *OTPDeviceConfigurationOtpDeliveriesCooldown {
	return v.value
}

func (v *NullableOTPDeviceConfigurationOtpDeliveriesCooldown) Set(val *OTPDeviceConfigurationOtpDeliveriesCooldown) {
	v.value = val
	v.isSet = true
}

func (v NullableOTPDeviceConfigurationOtpDeliveriesCooldown) IsSet() bool {
	return v.isSet
}

func (v *NullableOTPDeviceConfigurationOtpDeliveriesCooldown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTPDeviceConfigurationOtpDeliveriesCooldown(val *OTPDeviceConfigurationOtpDeliveriesCooldown) *NullableOTPDeviceConfigurationOtpDeliveriesCooldown {
	return &NullableOTPDeviceConfigurationOtpDeliveriesCooldown{value: val, isSet: true}
}

func (v NullableOTPDeviceConfigurationOtpDeliveriesCooldown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTPDeviceConfigurationOtpDeliveriesCooldown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


