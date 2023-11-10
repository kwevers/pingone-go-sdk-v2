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

// EnumPropagationStoreTypeLDAPGatewayLDAPType Type of LDAP gateway; can be `PingDirectory` or `Microsoft Active Directory`.
type EnumPropagationStoreTypeLDAPGatewayLDAPType string

// List of EnumPropagationStoreTypeLDAPGatewayLDAPType
const (
	ENUMPROPAGATIONSTORETYPELDAPGATEWAYLDAPTYPE_PING_DIRECTORY EnumPropagationStoreTypeLDAPGatewayLDAPType = "PingDirectory"
	ENUMPROPAGATIONSTORETYPELDAPGATEWAYLDAPTYPE_MICROSOFT_ACTIVE_DIRECTORY EnumPropagationStoreTypeLDAPGatewayLDAPType = "Microsoft Active Directory"
)

// All allowed values of EnumPropagationStoreTypeLDAPGatewayLDAPType enum
var AllowedEnumPropagationStoreTypeLDAPGatewayLDAPTypeEnumValues = []EnumPropagationStoreTypeLDAPGatewayLDAPType{
	"PingDirectory",
	"Microsoft Active Directory",
}

func (v *EnumPropagationStoreTypeLDAPGatewayLDAPType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPropagationStoreTypeLDAPGatewayLDAPType(value)
	for _, existing := range AllowedEnumPropagationStoreTypeLDAPGatewayLDAPTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumPropagationStoreTypeLDAPGatewayLDAPType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumPropagationStoreTypeLDAPGatewayLDAPTypeFromValue returns a pointer to a valid EnumPropagationStoreTypeLDAPGatewayLDAPType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPropagationStoreTypeLDAPGatewayLDAPTypeFromValue(v string) (*EnumPropagationStoreTypeLDAPGatewayLDAPType, error) {
	ev := EnumPropagationStoreTypeLDAPGatewayLDAPType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPropagationStoreTypeLDAPGatewayLDAPType: valid values are %v", v, AllowedEnumPropagationStoreTypeLDAPGatewayLDAPTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPropagationStoreTypeLDAPGatewayLDAPType) IsValid() bool {
	for _, existing := range AllowedEnumPropagationStoreTypeLDAPGatewayLDAPTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPropagationStoreTypeLDAPGatewayLDAPType value
func (v EnumPropagationStoreTypeLDAPGatewayLDAPType) Ptr() *EnumPropagationStoreTypeLDAPGatewayLDAPType {
	return &v
}

type NullableEnumPropagationStoreTypeLDAPGatewayLDAPType struct {
	value *EnumPropagationStoreTypeLDAPGatewayLDAPType
	isSet bool
}

func (v NullableEnumPropagationStoreTypeLDAPGatewayLDAPType) Get() *EnumPropagationStoreTypeLDAPGatewayLDAPType {
	return v.value
}

func (v *NullableEnumPropagationStoreTypeLDAPGatewayLDAPType) Set(val *EnumPropagationStoreTypeLDAPGatewayLDAPType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPropagationStoreTypeLDAPGatewayLDAPType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPropagationStoreTypeLDAPGatewayLDAPType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPropagationStoreTypeLDAPGatewayLDAPType(val *EnumPropagationStoreTypeLDAPGatewayLDAPType) *NullableEnumPropagationStoreTypeLDAPGatewayLDAPType {
	return &NullableEnumPropagationStoreTypeLDAPGatewayLDAPType{value: val, isSet: true}
}

func (v NullableEnumPropagationStoreTypeLDAPGatewayLDAPType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPropagationStoreTypeLDAPGatewayLDAPType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
