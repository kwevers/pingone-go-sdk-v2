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

// EnumNotificationsSettingsDeliveryMode the model 'EnumNotificationsSettingsDeliveryMode'
type EnumNotificationsSettingsDeliveryMode string

// List of EnumNotificationsSettingsDeliveryMode
const (
	ENUMNOTIFICATIONSSETTINGSDELIVERYMODE_ALL EnumNotificationsSettingsDeliveryMode = "ALL"
)

// All allowed values of EnumNotificationsSettingsDeliveryMode enum
var AllowedEnumNotificationsSettingsDeliveryModeEnumValues = []EnumNotificationsSettingsDeliveryMode{
	"ALL",
}

func (v *EnumNotificationsSettingsDeliveryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumNotificationsSettingsDeliveryMode(value)
	for _, existing := range AllowedEnumNotificationsSettingsDeliveryModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumNotificationsSettingsDeliveryMode", value)
}

// NewEnumNotificationsSettingsDeliveryModeFromValue returns a pointer to a valid EnumNotificationsSettingsDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumNotificationsSettingsDeliveryModeFromValue(v string) (*EnumNotificationsSettingsDeliveryMode, error) {
	ev := EnumNotificationsSettingsDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumNotificationsSettingsDeliveryMode: valid values are %v", v, AllowedEnumNotificationsSettingsDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumNotificationsSettingsDeliveryMode) IsValid() bool {
	for _, existing := range AllowedEnumNotificationsSettingsDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumNotificationsSettingsDeliveryMode value
func (v EnumNotificationsSettingsDeliveryMode) Ptr() *EnumNotificationsSettingsDeliveryMode {
	return &v
}

type NullableEnumNotificationsSettingsDeliveryMode struct {
	value *EnumNotificationsSettingsDeliveryMode
	isSet bool
}

func (v NullableEnumNotificationsSettingsDeliveryMode) Get() *EnumNotificationsSettingsDeliveryMode {
	return v.value
}

func (v *NullableEnumNotificationsSettingsDeliveryMode) Set(val *EnumNotificationsSettingsDeliveryMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumNotificationsSettingsDeliveryMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumNotificationsSettingsDeliveryMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumNotificationsSettingsDeliveryMode(val *EnumNotificationsSettingsDeliveryMode) *NullableEnumNotificationsSettingsDeliveryMode {
	return &NullableEnumNotificationsSettingsDeliveryMode{value: val, isSet: true}
}

func (v NullableEnumNotificationsSettingsDeliveryMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumNotificationsSettingsDeliveryMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

