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

// EnumDistanceUnit A string that specifies the unit of distance. Options are `kilometers`, `miles`.
type EnumDistanceUnit string

// List of EnumDistanceUnit
const (
	ENUMDISTANCEUNIT_KILOMETERS EnumDistanceUnit = "kilometers"
	ENUMDISTANCEUNIT_MILES EnumDistanceUnit = "miles"
)

// All allowed values of EnumDistanceUnit enum
var AllowedEnumDistanceUnitEnumValues = []EnumDistanceUnit{
	"kilometers",
	"miles",
}

func (v *EnumDistanceUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumDistanceUnit(value)
	for _, existing := range AllowedEnumDistanceUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumDistanceUnit", value)
}

// NewEnumDistanceUnitFromValue returns a pointer to a valid EnumDistanceUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumDistanceUnitFromValue(v string) (*EnumDistanceUnit, error) {
	ev := EnumDistanceUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumDistanceUnit: valid values are %v", v, AllowedEnumDistanceUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumDistanceUnit) IsValid() bool {
	for _, existing := range AllowedEnumDistanceUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumDistanceUnit value
func (v EnumDistanceUnit) Ptr() *EnumDistanceUnit {
	return &v
}

type NullableEnumDistanceUnit struct {
	value *EnumDistanceUnit
	isSet bool
}

func (v NullableEnumDistanceUnit) Get() *EnumDistanceUnit {
	return v.value
}

func (v *NullableEnumDistanceUnit) Set(val *EnumDistanceUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumDistanceUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumDistanceUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumDistanceUnit(val *EnumDistanceUnit) *NullableEnumDistanceUnit {
	return &NullableEnumDistanceUnit{value: val, isSet: true}
}

func (v NullableEnumDistanceUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumDistanceUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
