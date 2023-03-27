/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// EnumRiskLevel Enum string that specifies a given risk level.  Options are `LOW`, `MEDIUM`, `HIGH`.
type EnumRiskLevel string

// List of EnumRiskLevel
const (
	ENUMRISKLEVEL_LOW EnumRiskLevel = "LOW"
	ENUMRISKLEVEL_MEDIUM EnumRiskLevel = "MEDIUM"
	ENUMRISKLEVEL_HIGH EnumRiskLevel = "HIGH"
)

// All allowed values of EnumRiskLevel enum
var AllowedEnumRiskLevelEnumValues = []EnumRiskLevel{
	"LOW",
	"MEDIUM",
	"HIGH",
}

func (v *EnumRiskLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRiskLevel(value)
	for _, existing := range AllowedEnumRiskLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumRiskLevel", value)
}

// NewEnumRiskLevelFromValue returns a pointer to a valid EnumRiskLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRiskLevelFromValue(v string) (*EnumRiskLevel, error) {
	ev := EnumRiskLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRiskLevel: valid values are %v", v, AllowedEnumRiskLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRiskLevel) IsValid() bool {
	for _, existing := range AllowedEnumRiskLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRiskLevel value
func (v EnumRiskLevel) Ptr() *EnumRiskLevel {
	return &v
}

type NullableEnumRiskLevel struct {
	value *EnumRiskLevel
	isSet bool
}

func (v NullableEnumRiskLevel) Get() *EnumRiskLevel {
	return v.value
}

func (v *NullableEnumRiskLevel) Set(val *EnumRiskLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRiskLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRiskLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRiskLevel(val *EnumRiskLevel) *NullableEnumRiskLevel {
	return &NullableEnumRiskLevel{value: val, isSet: true}
}

func (v NullableEnumRiskLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRiskLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

