/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// EnumFIDO2PolicyDiscoverableCredentials Use this field to specify when registered users can authenticate without providing credentials. The possible values are: - `DISCOURAGED` - discoverable credentials are not used, even when supported by the FIDO device. In cases where use of discoverable credentials is required by the FIDO device itself, this setting does not override the device setting. - `REQUIRED` - require the use of discoverable credentials. This option is required for usernameless authentication. - `PREFERRED` - use discoverable credentials where possible. 
type EnumFIDO2PolicyDiscoverableCredentials string

// List of EnumFIDO2PolicyDiscoverableCredentials
const (
	ENUMFIDO2POLICYDISCOVERABLECREDENTIALS_DISCOURAGED EnumFIDO2PolicyDiscoverableCredentials = "DISCOURAGED"
	ENUMFIDO2POLICYDISCOVERABLECREDENTIALS_REQUIRED EnumFIDO2PolicyDiscoverableCredentials = "REQUIRED"
	ENUMFIDO2POLICYDISCOVERABLECREDENTIALS_PREFERRED EnumFIDO2PolicyDiscoverableCredentials = "PREFERRED"
)

// All allowed values of EnumFIDO2PolicyDiscoverableCredentials enum
var AllowedEnumFIDO2PolicyDiscoverableCredentialsEnumValues = []EnumFIDO2PolicyDiscoverableCredentials{
	"DISCOURAGED",
	"REQUIRED",
	"PREFERRED",
}

func (v *EnumFIDO2PolicyDiscoverableCredentials) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFIDO2PolicyDiscoverableCredentials(value)
	for _, existing := range AllowedEnumFIDO2PolicyDiscoverableCredentialsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFIDO2PolicyDiscoverableCredentials", value)
}

// NewEnumFIDO2PolicyDiscoverableCredentialsFromValue returns a pointer to a valid EnumFIDO2PolicyDiscoverableCredentials
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFIDO2PolicyDiscoverableCredentialsFromValue(v string) (*EnumFIDO2PolicyDiscoverableCredentials, error) {
	ev := EnumFIDO2PolicyDiscoverableCredentials(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFIDO2PolicyDiscoverableCredentials: valid values are %v", v, AllowedEnumFIDO2PolicyDiscoverableCredentialsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFIDO2PolicyDiscoverableCredentials) IsValid() bool {
	for _, existing := range AllowedEnumFIDO2PolicyDiscoverableCredentialsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFIDO2PolicyDiscoverableCredentials value
func (v EnumFIDO2PolicyDiscoverableCredentials) Ptr() *EnumFIDO2PolicyDiscoverableCredentials {
	return &v
}

type NullableEnumFIDO2PolicyDiscoverableCredentials struct {
	value *EnumFIDO2PolicyDiscoverableCredentials
	isSet bool
}

func (v NullableEnumFIDO2PolicyDiscoverableCredentials) Get() *EnumFIDO2PolicyDiscoverableCredentials {
	return v.value
}

func (v *NullableEnumFIDO2PolicyDiscoverableCredentials) Set(val *EnumFIDO2PolicyDiscoverableCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFIDO2PolicyDiscoverableCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFIDO2PolicyDiscoverableCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFIDO2PolicyDiscoverableCredentials(val *EnumFIDO2PolicyDiscoverableCredentials) *NullableEnumFIDO2PolicyDiscoverableCredentials {
	return &NullableEnumFIDO2PolicyDiscoverableCredentials{value: val, isSet: true}
}

func (v NullableEnumFIDO2PolicyDiscoverableCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFIDO2PolicyDiscoverableCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
