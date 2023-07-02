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

// EnumFormCategory A string that specifies the type of form. The `CUSTOM` form type allows the form to be built with fields that do not map specifically to the PingOne directory attributes.
type EnumFormCategory string

// List of EnumFormCategory
const (
	ENUMFORMCATEGORY_CUSTOM EnumFormCategory = "CUSTOM"
)

// All allowed values of EnumFormCategory enum
var AllowedEnumFormCategoryEnumValues = []EnumFormCategory{
	"CUSTOM",
}

func (v *EnumFormCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFormCategory(value)
	for _, existing := range AllowedEnumFormCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFormCategory", value)
}

// NewEnumFormCategoryFromValue returns a pointer to a valid EnumFormCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFormCategoryFromValue(v string) (*EnumFormCategory, error) {
	ev := EnumFormCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFormCategory: valid values are %v", v, AllowedEnumFormCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFormCategory) IsValid() bool {
	for _, existing := range AllowedEnumFormCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFormCategory value
func (v EnumFormCategory) Ptr() *EnumFormCategory {
	return &v
}

type NullableEnumFormCategory struct {
	value *EnumFormCategory
	isSet bool
}

func (v NullableEnumFormCategory) Get() *EnumFormCategory {
	return v.value
}

func (v *NullableEnumFormCategory) Set(val *EnumFormCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFormCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFormCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFormCategory(val *EnumFormCategory) *NullableEnumFormCategory {
	return &NullableEnumFormCategory{value: val, isSet: true}
}

func (v NullableEnumFormCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFormCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

