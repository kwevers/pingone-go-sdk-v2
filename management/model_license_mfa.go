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

// LicenseMfa struct for LicenseMfa
type LicenseMfa struct {
	// A read-only boolean that specifies whether push notifications are allowed. For TRIAL (unpaid) licenses, the default value is true. For other license package types, adoption of the feature determines the default value.
	AllowPushNotification *bool `json:"allowPushNotification,omitempty"`
	// A read-only boolean that specifies whether the license supports sending notifications outside of the environment's whitelist.
	AllowMfaNotificationsOutsideWhitelist *bool `json:"allowMfaNotificationsOutsideWhitelist,omitempty"`
	// A read-only boolean that specifies whether FIDO2 devices are allowed. For TRIAL (unpaid) licenses, the default value is true. For other license package types, adoption of the feature determines the default value.
	AllowFido2Devices *bool `json:"allowFido2Devices,omitempty"`
}

// NewLicenseMfa instantiates a new LicenseMfa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseMfa() *LicenseMfa {
	this := LicenseMfa{}
	return &this
}

// NewLicenseMfaWithDefaults instantiates a new LicenseMfa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseMfaWithDefaults() *LicenseMfa {
	this := LicenseMfa{}
	return &this
}

// GetAllowPushNotification returns the AllowPushNotification field value if set, zero value otherwise.
func (o *LicenseMfa) GetAllowPushNotification() bool {
	if o == nil || o.AllowPushNotification == nil {
		var ret bool
		return ret
	}
	return *o.AllowPushNotification
}

// GetAllowPushNotificationOk returns a tuple with the AllowPushNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseMfa) GetAllowPushNotificationOk() (*bool, bool) {
	if o == nil || o.AllowPushNotification == nil {
		return nil, false
	}
	return o.AllowPushNotification, true
}

// HasAllowPushNotification returns a boolean if a field has been set.
func (o *LicenseMfa) HasAllowPushNotification() bool {
	if o != nil && o.AllowPushNotification != nil {
		return true
	}

	return false
}

// SetAllowPushNotification gets a reference to the given bool and assigns it to the AllowPushNotification field.
func (o *LicenseMfa) SetAllowPushNotification(v bool) {
	o.AllowPushNotification = &v
}

// GetAllowMfaNotificationsOutsideWhitelist returns the AllowMfaNotificationsOutsideWhitelist field value if set, zero value otherwise.
func (o *LicenseMfa) GetAllowMfaNotificationsOutsideWhitelist() bool {
	if o == nil || o.AllowMfaNotificationsOutsideWhitelist == nil {
		var ret bool
		return ret
	}
	return *o.AllowMfaNotificationsOutsideWhitelist
}

// GetAllowMfaNotificationsOutsideWhitelistOk returns a tuple with the AllowMfaNotificationsOutsideWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseMfa) GetAllowMfaNotificationsOutsideWhitelistOk() (*bool, bool) {
	if o == nil || o.AllowMfaNotificationsOutsideWhitelist == nil {
		return nil, false
	}
	return o.AllowMfaNotificationsOutsideWhitelist, true
}

// HasAllowMfaNotificationsOutsideWhitelist returns a boolean if a field has been set.
func (o *LicenseMfa) HasAllowMfaNotificationsOutsideWhitelist() bool {
	if o != nil && o.AllowMfaNotificationsOutsideWhitelist != nil {
		return true
	}

	return false
}

// SetAllowMfaNotificationsOutsideWhitelist gets a reference to the given bool and assigns it to the AllowMfaNotificationsOutsideWhitelist field.
func (o *LicenseMfa) SetAllowMfaNotificationsOutsideWhitelist(v bool) {
	o.AllowMfaNotificationsOutsideWhitelist = &v
}

// GetAllowFido2Devices returns the AllowFido2Devices field value if set, zero value otherwise.
func (o *LicenseMfa) GetAllowFido2Devices() bool {
	if o == nil || o.AllowFido2Devices == nil {
		var ret bool
		return ret
	}
	return *o.AllowFido2Devices
}

// GetAllowFido2DevicesOk returns a tuple with the AllowFido2Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseMfa) GetAllowFido2DevicesOk() (*bool, bool) {
	if o == nil || o.AllowFido2Devices == nil {
		return nil, false
	}
	return o.AllowFido2Devices, true
}

// HasAllowFido2Devices returns a boolean if a field has been set.
func (o *LicenseMfa) HasAllowFido2Devices() bool {
	if o != nil && o.AllowFido2Devices != nil {
		return true
	}

	return false
}

// SetAllowFido2Devices gets a reference to the given bool and assigns it to the AllowFido2Devices field.
func (o *LicenseMfa) SetAllowFido2Devices(v bool) {
	o.AllowFido2Devices = &v
}

func (o LicenseMfa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowPushNotification != nil {
		toSerialize["allowPushNotification"] = o.AllowPushNotification
	}
	if o.AllowMfaNotificationsOutsideWhitelist != nil {
		toSerialize["allowMfaNotificationsOutsideWhitelist"] = o.AllowMfaNotificationsOutsideWhitelist
	}
	if o.AllowFido2Devices != nil {
		toSerialize["allowFido2Devices"] = o.AllowFido2Devices
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseMfa struct {
	value *LicenseMfa
	isSet bool
}

func (v NullableLicenseMfa) Get() *LicenseMfa {
	return v.value
}

func (v *NullableLicenseMfa) Set(val *LicenseMfa) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseMfa) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseMfa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseMfa(val *LicenseMfa) *NullableLicenseMfa {
	return &NullableLicenseMfa{value: val, isSet: true}
}

func (v NullableLicenseMfa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseMfa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


