/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumPropagationStoreTypeSCIMUniqueUserIdentifier Specifies the unique user identifier to use. Options are either `userName` or `workEmail`.
type EnumPropagationStoreTypeSCIMUniqueUserIdentifier string

// List of EnumPropagationStoreTypeSCIMUniqueUserIdentifier
const (
	ENUMPROPAGATIONSTORETYPESCIMUNIQUEUSERIDENTIFIER_USER_NAME EnumPropagationStoreTypeSCIMUniqueUserIdentifier = "userName"
	ENUMPROPAGATIONSTORETYPESCIMUNIQUEUSERIDENTIFIER_WORK_EMAIL EnumPropagationStoreTypeSCIMUniqueUserIdentifier = "workEmail"
)

// All allowed values of EnumPropagationStoreTypeSCIMUniqueUserIdentifier enum
var AllowedEnumPropagationStoreTypeSCIMUniqueUserIdentifierEnumValues = []EnumPropagationStoreTypeSCIMUniqueUserIdentifier{
	"userName",
	"workEmail",
}

func (v *EnumPropagationStoreTypeSCIMUniqueUserIdentifier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPropagationStoreTypeSCIMUniqueUserIdentifier(value)
	for _, existing := range AllowedEnumPropagationStoreTypeSCIMUniqueUserIdentifierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumPropagationStoreTypeSCIMUniqueUserIdentifier(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumPropagationStoreTypeSCIMUniqueUserIdentifierFromValue returns a pointer to a valid EnumPropagationStoreTypeSCIMUniqueUserIdentifier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPropagationStoreTypeSCIMUniqueUserIdentifierFromValue(v string) (*EnumPropagationStoreTypeSCIMUniqueUserIdentifier, error) {
	ev := EnumPropagationStoreTypeSCIMUniqueUserIdentifier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPropagationStoreTypeSCIMUniqueUserIdentifier: valid values are %v", v, AllowedEnumPropagationStoreTypeSCIMUniqueUserIdentifierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPropagationStoreTypeSCIMUniqueUserIdentifier) IsValid() bool {
	for _, existing := range AllowedEnumPropagationStoreTypeSCIMUniqueUserIdentifierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPropagationStoreTypeSCIMUniqueUserIdentifier value
func (v EnumPropagationStoreTypeSCIMUniqueUserIdentifier) Ptr() *EnumPropagationStoreTypeSCIMUniqueUserIdentifier {
	return &v
}

type NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier struct {
	value *EnumPropagationStoreTypeSCIMUniqueUserIdentifier
	isSet bool
}

func (v NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier) Get() *EnumPropagationStoreTypeSCIMUniqueUserIdentifier {
	return v.value
}

func (v *NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier) Set(val *EnumPropagationStoreTypeSCIMUniqueUserIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier(val *EnumPropagationStoreTypeSCIMUniqueUserIdentifier) *NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier {
	return &NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier{value: val, isSet: true}
}

func (v NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPropagationStoreTypeSCIMUniqueUserIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
