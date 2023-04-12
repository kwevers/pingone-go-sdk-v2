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

// EnumNotificationsPolicyCountryLimitDeliveryMethod The delivery method that the defined limitation should be applied to.  Options are `SMS`, `Voice`.
type EnumNotificationsPolicyCountryLimitDeliveryMethod string

// List of EnumNotificationsPolicyCountryLimitDeliveryMethod
const (
	ENUMNOTIFICATIONSPOLICYCOUNTRYLIMITDELIVERYMETHOD_SMS EnumNotificationsPolicyCountryLimitDeliveryMethod = "SMS"
	ENUMNOTIFICATIONSPOLICYCOUNTRYLIMITDELIVERYMETHOD_VOICE EnumNotificationsPolicyCountryLimitDeliveryMethod = "Voice"
)

// All allowed values of EnumNotificationsPolicyCountryLimitDeliveryMethod enum
var AllowedEnumNotificationsPolicyCountryLimitDeliveryMethodEnumValues = []EnumNotificationsPolicyCountryLimitDeliveryMethod{
	"SMS",
	"Voice",
}

func (v *EnumNotificationsPolicyCountryLimitDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumNotificationsPolicyCountryLimitDeliveryMethod(value)
	for _, existing := range AllowedEnumNotificationsPolicyCountryLimitDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumNotificationsPolicyCountryLimitDeliveryMethod", value)
}

// NewEnumNotificationsPolicyCountryLimitDeliveryMethodFromValue returns a pointer to a valid EnumNotificationsPolicyCountryLimitDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumNotificationsPolicyCountryLimitDeliveryMethodFromValue(v string) (*EnumNotificationsPolicyCountryLimitDeliveryMethod, error) {
	ev := EnumNotificationsPolicyCountryLimitDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumNotificationsPolicyCountryLimitDeliveryMethod: valid values are %v", v, AllowedEnumNotificationsPolicyCountryLimitDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumNotificationsPolicyCountryLimitDeliveryMethod) IsValid() bool {
	for _, existing := range AllowedEnumNotificationsPolicyCountryLimitDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumNotificationsPolicyCountryLimitDeliveryMethod value
func (v EnumNotificationsPolicyCountryLimitDeliveryMethod) Ptr() *EnumNotificationsPolicyCountryLimitDeliveryMethod {
	return &v
}

type NullableEnumNotificationsPolicyCountryLimitDeliveryMethod struct {
	value *EnumNotificationsPolicyCountryLimitDeliveryMethod
	isSet bool
}

func (v NullableEnumNotificationsPolicyCountryLimitDeliveryMethod) Get() *EnumNotificationsPolicyCountryLimitDeliveryMethod {
	return v.value
}

func (v *NullableEnumNotificationsPolicyCountryLimitDeliveryMethod) Set(val *EnumNotificationsPolicyCountryLimitDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumNotificationsPolicyCountryLimitDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumNotificationsPolicyCountryLimitDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumNotificationsPolicyCountryLimitDeliveryMethod(val *EnumNotificationsPolicyCountryLimitDeliveryMethod) *NullableEnumNotificationsPolicyCountryLimitDeliveryMethod {
	return &NullableEnumNotificationsPolicyCountryLimitDeliveryMethod{value: val, isSet: true}
}

func (v NullableEnumNotificationsPolicyCountryLimitDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumNotificationsPolicyCountryLimitDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

