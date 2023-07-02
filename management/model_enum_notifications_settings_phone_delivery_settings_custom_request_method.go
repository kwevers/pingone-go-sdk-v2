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

// EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod The type of HTTP request method. Possible values are `GET` and `POST`.
type EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod string

// List of EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod
const (
	ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMREQUESTMETHOD_GET EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod = "GET"
	ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMREQUESTMETHOD_POST EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod = "POST"
)

// All allowed values of EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod enum
var AllowedEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethodEnumValues = []EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod{
	"GET",
	"POST",
}

func (v *EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod(value)
	for _, existing := range AllowedEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod", value)
}

// NewEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethodFromValue returns a pointer to a valid EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethodFromValue(v string) (*EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod, error) {
	ev := EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod: valid values are %v", v, AllowedEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) IsValid() bool {
	for _, existing := range AllowedEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod value
func (v EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) Ptr() *EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod {
	return &v
}

type NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod struct {
	value *EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod
	isSet bool
}

func (v NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) Get() *EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod {
	return v.value
}

func (v *NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) Set(val *EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod(val *EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) *NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod {
	return &NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod{value: val, isSet: true}
}

func (v NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

