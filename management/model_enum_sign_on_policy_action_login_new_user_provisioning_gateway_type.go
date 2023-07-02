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

// EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType A string identifying the type of gateway. Currently, only `LDAP` is supported.
type EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType string

// List of EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType
const (
	ENUMSIGNONPOLICYACTIONLOGINNEWUSERPROVISIONINGGATEWAYTYPE_LDAP EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType = "LDAP"
)

// All allowed values of EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType enum
var AllowedEnumSignOnPolicyActionLoginNewUserProvisioningGatewayTypeEnumValues = []EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType{
	"LDAP",
}

func (v *EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType(value)
	for _, existing := range AllowedEnumSignOnPolicyActionLoginNewUserProvisioningGatewayTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType", value)
}

// NewEnumSignOnPolicyActionLoginNewUserProvisioningGatewayTypeFromValue returns a pointer to a valid EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumSignOnPolicyActionLoginNewUserProvisioningGatewayTypeFromValue(v string) (*EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType, error) {
	ev := EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType: valid values are %v", v, AllowedEnumSignOnPolicyActionLoginNewUserProvisioningGatewayTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) IsValid() bool {
	for _, existing := range AllowedEnumSignOnPolicyActionLoginNewUserProvisioningGatewayTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType value
func (v EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) Ptr() *EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType {
	return &v
}

type NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType struct {
	value *EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType
	isSet bool
}

func (v NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) Get() *EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType {
	return v.value
}

func (v *NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) Set(val *EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType(val *EnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) *NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType {
	return &NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType{value: val, isSet: true}
}

func (v NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumSignOnPolicyActionLoginNewUserProvisioningGatewayType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

