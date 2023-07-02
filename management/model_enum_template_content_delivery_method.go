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

// EnumTemplateContentDeliveryMethod The content's delivery method. Possible values are `Email`, `SMS`, `Voice` or `Push`. Cannot be changed after it is initially set in `POST /environments/{{envID}}/templates/{{templateName}}/contents`.
type EnumTemplateContentDeliveryMethod string

// List of EnumTemplateContentDeliveryMethod
const (
	ENUMTEMPLATECONTENTDELIVERYMETHOD_EMAIL EnumTemplateContentDeliveryMethod = "Email"
	ENUMTEMPLATECONTENTDELIVERYMETHOD_SMS EnumTemplateContentDeliveryMethod = "SMS"
	ENUMTEMPLATECONTENTDELIVERYMETHOD_VOICE EnumTemplateContentDeliveryMethod = "Voice"
	ENUMTEMPLATECONTENTDELIVERYMETHOD_PUSH EnumTemplateContentDeliveryMethod = "Push"
)

// All allowed values of EnumTemplateContentDeliveryMethod enum
var AllowedEnumTemplateContentDeliveryMethodEnumValues = []EnumTemplateContentDeliveryMethod{
	"Email",
	"SMS",
	"Voice",
	"Push",
}

func (v *EnumTemplateContentDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumTemplateContentDeliveryMethod(value)
	for _, existing := range AllowedEnumTemplateContentDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumTemplateContentDeliveryMethod", value)
}

// NewEnumTemplateContentDeliveryMethodFromValue returns a pointer to a valid EnumTemplateContentDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumTemplateContentDeliveryMethodFromValue(v string) (*EnumTemplateContentDeliveryMethod, error) {
	ev := EnumTemplateContentDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumTemplateContentDeliveryMethod: valid values are %v", v, AllowedEnumTemplateContentDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumTemplateContentDeliveryMethod) IsValid() bool {
	for _, existing := range AllowedEnumTemplateContentDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumTemplateContentDeliveryMethod value
func (v EnumTemplateContentDeliveryMethod) Ptr() *EnumTemplateContentDeliveryMethod {
	return &v
}

type NullableEnumTemplateContentDeliveryMethod struct {
	value *EnumTemplateContentDeliveryMethod
	isSet bool
}

func (v NullableEnumTemplateContentDeliveryMethod) Get() *EnumTemplateContentDeliveryMethod {
	return v.value
}

func (v *NullableEnumTemplateContentDeliveryMethod) Set(val *EnumTemplateContentDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTemplateContentDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTemplateContentDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTemplateContentDeliveryMethod(val *EnumTemplateContentDeliveryMethod) *NullableEnumTemplateContentDeliveryMethod {
	return &NullableEnumTemplateContentDeliveryMethod{value: val, isSet: true}
}

func (v NullableEnumTemplateContentDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTemplateContentDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

