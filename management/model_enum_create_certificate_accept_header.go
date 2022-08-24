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

// EnumCreateCertificateAcceptHeader the model 'EnumCreateCertificateAcceptHeader'
type EnumCreateCertificateAcceptHeader string

// List of EnumCreateCertificateAcceptHeader
const (
	ENUMCREATECERTIFICATEACCEPTHEADER_APPLICATION_X_PKCS7_CERTIFICATES EnumCreateCertificateAcceptHeader = "application/x-pkcs7-certificates"
)

// All allowed values of EnumCreateCertificateAcceptHeader enum
var AllowedEnumCreateCertificateAcceptHeaderEnumValues = []EnumCreateCertificateAcceptHeader{
	"application/x-pkcs7-certificates",
}

func (v *EnumCreateCertificateAcceptHeader) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCreateCertificateAcceptHeader(value)
	for _, existing := range AllowedEnumCreateCertificateAcceptHeaderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCreateCertificateAcceptHeader", value)
}

// NewEnumCreateCertificateAcceptHeaderFromValue returns a pointer to a valid EnumCreateCertificateAcceptHeader
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCreateCertificateAcceptHeaderFromValue(v string) (*EnumCreateCertificateAcceptHeader, error) {
	ev := EnumCreateCertificateAcceptHeader(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCreateCertificateAcceptHeader: valid values are %v", v, AllowedEnumCreateCertificateAcceptHeaderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCreateCertificateAcceptHeader) IsValid() bool {
	for _, existing := range AllowedEnumCreateCertificateAcceptHeaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCreateCertificateAcceptHeader value
func (v EnumCreateCertificateAcceptHeader) Ptr() *EnumCreateCertificateAcceptHeader {
	return &v
}

type NullableEnumCreateCertificateAcceptHeader struct {
	value *EnumCreateCertificateAcceptHeader
	isSet bool
}

func (v NullableEnumCreateCertificateAcceptHeader) Get() *EnumCreateCertificateAcceptHeader {
	return v.value
}

func (v *NullableEnumCreateCertificateAcceptHeader) Set(val *EnumCreateCertificateAcceptHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCreateCertificateAcceptHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCreateCertificateAcceptHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCreateCertificateAcceptHeader(val *EnumCreateCertificateAcceptHeader) *NullableEnumCreateCertificateAcceptHeader {
	return &NullableEnumCreateCertificateAcceptHeader{value: val, isSet: true}
}

func (v NullableEnumCreateCertificateAcceptHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCreateCertificateAcceptHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

