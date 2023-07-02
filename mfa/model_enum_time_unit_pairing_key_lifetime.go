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

// EnumTimeUnitPairingKeyLifetime The time unit for the `pairingKeyLifetime.duration` parameter.
type EnumTimeUnitPairingKeyLifetime string

// List of EnumTimeUnitPairingKeyLifetime
const (
	ENUMTIMEUNITPAIRINGKEYLIFETIME_MINUTES EnumTimeUnitPairingKeyLifetime = "MINUTES"
	ENUMTIMEUNITPAIRINGKEYLIFETIME_HOURS EnumTimeUnitPairingKeyLifetime = "HOURS"
)

// All allowed values of EnumTimeUnitPairingKeyLifetime enum
var AllowedEnumTimeUnitPairingKeyLifetimeEnumValues = []EnumTimeUnitPairingKeyLifetime{
	"MINUTES",
	"HOURS",
}

func (v *EnumTimeUnitPairingKeyLifetime) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumTimeUnitPairingKeyLifetime(value)
	for _, existing := range AllowedEnumTimeUnitPairingKeyLifetimeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumTimeUnitPairingKeyLifetime", value)
}

// NewEnumTimeUnitPairingKeyLifetimeFromValue returns a pointer to a valid EnumTimeUnitPairingKeyLifetime
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumTimeUnitPairingKeyLifetimeFromValue(v string) (*EnumTimeUnitPairingKeyLifetime, error) {
	ev := EnumTimeUnitPairingKeyLifetime(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumTimeUnitPairingKeyLifetime: valid values are %v", v, AllowedEnumTimeUnitPairingKeyLifetimeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumTimeUnitPairingKeyLifetime) IsValid() bool {
	for _, existing := range AllowedEnumTimeUnitPairingKeyLifetimeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumTimeUnitPairingKeyLifetime value
func (v EnumTimeUnitPairingKeyLifetime) Ptr() *EnumTimeUnitPairingKeyLifetime {
	return &v
}

type NullableEnumTimeUnitPairingKeyLifetime struct {
	value *EnumTimeUnitPairingKeyLifetime
	isSet bool
}

func (v NullableEnumTimeUnitPairingKeyLifetime) Get() *EnumTimeUnitPairingKeyLifetime {
	return v.value
}

func (v *NullableEnumTimeUnitPairingKeyLifetime) Set(val *EnumTimeUnitPairingKeyLifetime) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTimeUnitPairingKeyLifetime) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTimeUnitPairingKeyLifetime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTimeUnitPairingKeyLifetime(val *EnumTimeUnitPairingKeyLifetime) *NullableEnumTimeUnitPairingKeyLifetime {
	return &NullableEnumTimeUnitPairingKeyLifetime{value: val, isSet: true}
}

func (v NullableEnumTimeUnitPairingKeyLifetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTimeUnitPairingKeyLifetime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

