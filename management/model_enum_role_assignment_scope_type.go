/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumRoleAssignmentScopeType A string that specifies the type of resource defining the scope of the Role assignment.
type EnumRoleAssignmentScopeType string

// List of EnumRoleAssignmentScopeType
const (
	ENUMROLEASSIGNMENTSCOPETYPE_ORGANIZATION EnumRoleAssignmentScopeType = "ORGANIZATION"
	ENUMROLEASSIGNMENTSCOPETYPE_ENVIRONMENT EnumRoleAssignmentScopeType = "ENVIRONMENT"
	ENUMROLEASSIGNMENTSCOPETYPE_POPULATION EnumRoleAssignmentScopeType = "POPULATION"
)

// All allowed values of EnumRoleAssignmentScopeType enum
var AllowedEnumRoleAssignmentScopeTypeEnumValues = []EnumRoleAssignmentScopeType{
	"ORGANIZATION",
	"ENVIRONMENT",
	"POPULATION",
}

func (v *EnumRoleAssignmentScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRoleAssignmentScopeType(value)
	for _, existing := range AllowedEnumRoleAssignmentScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumRoleAssignmentScopeType", value)
}

// NewEnumRoleAssignmentScopeTypeFromValue returns a pointer to a valid EnumRoleAssignmentScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRoleAssignmentScopeTypeFromValue(v string) (*EnumRoleAssignmentScopeType, error) {
	ev := EnumRoleAssignmentScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRoleAssignmentScopeType: valid values are %v", v, AllowedEnumRoleAssignmentScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRoleAssignmentScopeType) IsValid() bool {
	for _, existing := range AllowedEnumRoleAssignmentScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRoleAssignmentScopeType value
func (v EnumRoleAssignmentScopeType) Ptr() *EnumRoleAssignmentScopeType {
	return &v
}

type NullableEnumRoleAssignmentScopeType struct {
	value *EnumRoleAssignmentScopeType
	isSet bool
}

func (v NullableEnumRoleAssignmentScopeType) Get() *EnumRoleAssignmentScopeType {
	return v.value
}

func (v *NullableEnumRoleAssignmentScopeType) Set(val *EnumRoleAssignmentScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRoleAssignmentScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRoleAssignmentScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRoleAssignmentScopeType(val *EnumRoleAssignmentScopeType) *NullableEnumRoleAssignmentScopeType {
	return &NullableEnumRoleAssignmentScopeType{value: val, isSet: true}
}

func (v NullableEnumRoleAssignmentScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRoleAssignmentScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

