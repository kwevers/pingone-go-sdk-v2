/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// EnumFIDO2PolicyAuthenticatorAttachment Used to control the type of authenticators that are allowed. Can be one of the following values: - `PLATFORM` - only allow the use of FIDO device authenticators that contain an internal authenticator (such as a face or fingerprint scanner) - `CROSS_PLATFORM` - allow use of cross-platform authenticators, which are external to the accessing device (such as a security key) - `BOTH` - allow both categories of authenticators 
type EnumFIDO2PolicyAuthenticatorAttachment string

// List of EnumFIDO2PolicyAuthenticatorAttachment
const (
	ENUMFIDO2POLICYAUTHENTICATORATTACHMENT_PLATFORM EnumFIDO2PolicyAuthenticatorAttachment = "PLATFORM"
	ENUMFIDO2POLICYAUTHENTICATORATTACHMENT_CROSS_PLATFORM EnumFIDO2PolicyAuthenticatorAttachment = "CROSS_PLATFORM"
	ENUMFIDO2POLICYAUTHENTICATORATTACHMENT_BOTH EnumFIDO2PolicyAuthenticatorAttachment = "BOTH"
)

// All allowed values of EnumFIDO2PolicyAuthenticatorAttachment enum
var AllowedEnumFIDO2PolicyAuthenticatorAttachmentEnumValues = []EnumFIDO2PolicyAuthenticatorAttachment{
	"PLATFORM",
	"CROSS_PLATFORM",
	"BOTH",
}

func (v *EnumFIDO2PolicyAuthenticatorAttachment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFIDO2PolicyAuthenticatorAttachment(value)
	for _, existing := range AllowedEnumFIDO2PolicyAuthenticatorAttachmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFIDO2PolicyAuthenticatorAttachment", value)
}

// NewEnumFIDO2PolicyAuthenticatorAttachmentFromValue returns a pointer to a valid EnumFIDO2PolicyAuthenticatorAttachment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFIDO2PolicyAuthenticatorAttachmentFromValue(v string) (*EnumFIDO2PolicyAuthenticatorAttachment, error) {
	ev := EnumFIDO2PolicyAuthenticatorAttachment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFIDO2PolicyAuthenticatorAttachment: valid values are %v", v, AllowedEnumFIDO2PolicyAuthenticatorAttachmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFIDO2PolicyAuthenticatorAttachment) IsValid() bool {
	for _, existing := range AllowedEnumFIDO2PolicyAuthenticatorAttachmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFIDO2PolicyAuthenticatorAttachment value
func (v EnumFIDO2PolicyAuthenticatorAttachment) Ptr() *EnumFIDO2PolicyAuthenticatorAttachment {
	return &v
}

type NullableEnumFIDO2PolicyAuthenticatorAttachment struct {
	value *EnumFIDO2PolicyAuthenticatorAttachment
	isSet bool
}

func (v NullableEnumFIDO2PolicyAuthenticatorAttachment) Get() *EnumFIDO2PolicyAuthenticatorAttachment {
	return v.value
}

func (v *NullableEnumFIDO2PolicyAuthenticatorAttachment) Set(val *EnumFIDO2PolicyAuthenticatorAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFIDO2PolicyAuthenticatorAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFIDO2PolicyAuthenticatorAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFIDO2PolicyAuthenticatorAttachment(val *EnumFIDO2PolicyAuthenticatorAttachment) *NullableEnumFIDO2PolicyAuthenticatorAttachment {
	return &NullableEnumFIDO2PolicyAuthenticatorAttachment{value: val, isSet: true}
}

func (v NullableEnumFIDO2PolicyAuthenticatorAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFIDO2PolicyAuthenticatorAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
