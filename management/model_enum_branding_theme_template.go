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

// EnumBrandingThemeTemplate The template name of the branding theme associated with the environment. Options are default, focus, mural, slate, and split.
type EnumBrandingThemeTemplate string

// List of EnumBrandingThemeTemplate
const (
	ENUMBRANDINGTHEMETEMPLATE_DEFAULT EnumBrandingThemeTemplate = "default"
	ENUMBRANDINGTHEMETEMPLATE_FOCUS EnumBrandingThemeTemplate = "focus"
	ENUMBRANDINGTHEMETEMPLATE_MURAL EnumBrandingThemeTemplate = "mural"
	ENUMBRANDINGTHEMETEMPLATE_SLATE EnumBrandingThemeTemplate = "slate"
	ENUMBRANDINGTHEMETEMPLATE_SPLIT EnumBrandingThemeTemplate = "split"
)

// All allowed values of EnumBrandingThemeTemplate enum
var AllowedEnumBrandingThemeTemplateEnumValues = []EnumBrandingThemeTemplate{
	"default",
	"focus",
	"mural",
	"slate",
	"split",
}

func (v *EnumBrandingThemeTemplate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumBrandingThemeTemplate(value)
	for _, existing := range AllowedEnumBrandingThemeTemplateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumBrandingThemeTemplate", value)
}

// NewEnumBrandingThemeTemplateFromValue returns a pointer to a valid EnumBrandingThemeTemplate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumBrandingThemeTemplateFromValue(v string) (*EnumBrandingThemeTemplate, error) {
	ev := EnumBrandingThemeTemplate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumBrandingThemeTemplate: valid values are %v", v, AllowedEnumBrandingThemeTemplateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumBrandingThemeTemplate) IsValid() bool {
	for _, existing := range AllowedEnumBrandingThemeTemplateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumBrandingThemeTemplate value
func (v EnumBrandingThemeTemplate) Ptr() *EnumBrandingThemeTemplate {
	return &v
}

type NullableEnumBrandingThemeTemplate struct {
	value *EnumBrandingThemeTemplate
	isSet bool
}

func (v NullableEnumBrandingThemeTemplate) Get() *EnumBrandingThemeTemplate {
	return v.value
}

func (v *NullableEnumBrandingThemeTemplate) Set(val *EnumBrandingThemeTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumBrandingThemeTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumBrandingThemeTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumBrandingThemeTemplate(val *EnumBrandingThemeTemplate) *NullableEnumBrandingThemeTemplate {
	return &NullableEnumBrandingThemeTemplate{value: val, isSet: true}
}

func (v NullableEnumBrandingThemeTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumBrandingThemeTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
