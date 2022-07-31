/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumIdentityProviderOIDCTokenAuthMethod A string that specifies the OIDC identity provider's token endpoint authentication method. Options are CLIENT_SECRET_BASIC (default), CLIENT_SECRET_POST, and NONE. This is a required property.
type EnumIdentityProviderOIDCTokenAuthMethod string

// List of EnumIdentityProviderOIDCTokenAuthMethod
const (
	ENUMIDENTITYPROVIDEROIDCTOKENAUTHMETHOD_CLIENT_SECRET_BASIC EnumIdentityProviderOIDCTokenAuthMethod = "CLIENT_SECRET_BASIC"
	ENUMIDENTITYPROVIDEROIDCTOKENAUTHMETHOD_CLIENT_SECRET_POST EnumIdentityProviderOIDCTokenAuthMethod = "CLIENT_SECRET_POST"
	ENUMIDENTITYPROVIDEROIDCTOKENAUTHMETHOD_NONE EnumIdentityProviderOIDCTokenAuthMethod = "NONE"
)

// All allowed values of EnumIdentityProviderOIDCTokenAuthMethod enum
var AllowedEnumIdentityProviderOIDCTokenAuthMethodEnumValues = []EnumIdentityProviderOIDCTokenAuthMethod{
	"CLIENT_SECRET_BASIC",
	"CLIENT_SECRET_POST",
	"NONE",
}

func (v *EnumIdentityProviderOIDCTokenAuthMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIdentityProviderOIDCTokenAuthMethod(value)
	for _, existing := range AllowedEnumIdentityProviderOIDCTokenAuthMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumIdentityProviderOIDCTokenAuthMethod", value)
}

// NewEnumIdentityProviderOIDCTokenAuthMethodFromValue returns a pointer to a valid EnumIdentityProviderOIDCTokenAuthMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIdentityProviderOIDCTokenAuthMethodFromValue(v string) (*EnumIdentityProviderOIDCTokenAuthMethod, error) {
	ev := EnumIdentityProviderOIDCTokenAuthMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIdentityProviderOIDCTokenAuthMethod: valid values are %v", v, AllowedEnumIdentityProviderOIDCTokenAuthMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIdentityProviderOIDCTokenAuthMethod) IsValid() bool {
	for _, existing := range AllowedEnumIdentityProviderOIDCTokenAuthMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIdentityProviderOIDCTokenAuthMethod value
func (v EnumIdentityProviderOIDCTokenAuthMethod) Ptr() *EnumIdentityProviderOIDCTokenAuthMethod {
	return &v
}

type NullableEnumIdentityProviderOIDCTokenAuthMethod struct {
	value *EnumIdentityProviderOIDCTokenAuthMethod
	isSet bool
}

func (v NullableEnumIdentityProviderOIDCTokenAuthMethod) Get() *EnumIdentityProviderOIDCTokenAuthMethod {
	return v.value
}

func (v *NullableEnumIdentityProviderOIDCTokenAuthMethod) Set(val *EnumIdentityProviderOIDCTokenAuthMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIdentityProviderOIDCTokenAuthMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIdentityProviderOIDCTokenAuthMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIdentityProviderOIDCTokenAuthMethod(val *EnumIdentityProviderOIDCTokenAuthMethod) *NullableEnumIdentityProviderOIDCTokenAuthMethod {
	return &NullableEnumIdentityProviderOIDCTokenAuthMethod{value: val, isSet: true}
}

func (v NullableEnumIdentityProviderOIDCTokenAuthMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIdentityProviderOIDCTokenAuthMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

