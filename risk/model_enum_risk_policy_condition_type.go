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

// EnumRiskPolicyConditionType Indicates the type of policy you are defining. Can be one of the following values: `AGGREGATED_SCORES`, `AGGREGATED_WEIGHTS`, `VALUE_COMPARISON`, and `IP_RANGE` (for override policies or custom predictors). 
type EnumRiskPolicyConditionType string

// List of EnumRiskPolicyConditionType
const (
	ENUMRISKPOLICYCONDITIONTYPE_AGGREGATED_WEIGHTS EnumRiskPolicyConditionType = "AGGREGATED_WEIGHTS"
	ENUMRISKPOLICYCONDITIONTYPE_AGGREGATED_SCORES EnumRiskPolicyConditionType = "AGGREGATED_SCORES"
	ENUMRISKPOLICYCONDITIONTYPE_VALUE_COMPARISON EnumRiskPolicyConditionType = "VALUE_COMPARISON"
	ENUMRISKPOLICYCONDITIONTYPE_IP_RANGE EnumRiskPolicyConditionType = "IP_RANGE"
)

// All allowed values of EnumRiskPolicyConditionType enum
var AllowedEnumRiskPolicyConditionTypeEnumValues = []EnumRiskPolicyConditionType{
	"AGGREGATED_WEIGHTS",
	"AGGREGATED_SCORES",
	"VALUE_COMPARISON",
	"IP_RANGE",
}

func (v *EnumRiskPolicyConditionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRiskPolicyConditionType(value)
	for _, existing := range AllowedEnumRiskPolicyConditionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumRiskPolicyConditionType", value)
}

// NewEnumRiskPolicyConditionTypeFromValue returns a pointer to a valid EnumRiskPolicyConditionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRiskPolicyConditionTypeFromValue(v string) (*EnumRiskPolicyConditionType, error) {
	ev := EnumRiskPolicyConditionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRiskPolicyConditionType: valid values are %v", v, AllowedEnumRiskPolicyConditionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRiskPolicyConditionType) IsValid() bool {
	for _, existing := range AllowedEnumRiskPolicyConditionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRiskPolicyConditionType value
func (v EnumRiskPolicyConditionType) Ptr() *EnumRiskPolicyConditionType {
	return &v
}

type NullableEnumRiskPolicyConditionType struct {
	value *EnumRiskPolicyConditionType
	isSet bool
}

func (v NullableEnumRiskPolicyConditionType) Get() *EnumRiskPolicyConditionType {
	return v.value
}

func (v *NullableEnumRiskPolicyConditionType) Set(val *EnumRiskPolicyConditionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRiskPolicyConditionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRiskPolicyConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRiskPolicyConditionType(val *EnumRiskPolicyConditionType) *NullableEnumRiskPolicyConditionType {
	return &NullableEnumRiskPolicyConditionType{value: val, isSet: true}
}

func (v NullableEnumRiskPolicyConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRiskPolicyConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
