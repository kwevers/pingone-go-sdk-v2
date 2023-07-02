/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the UpdateEnvironmentTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEnvironmentTypeRequest{}

// UpdateEnvironmentTypeRequest struct for UpdateEnvironmentTypeRequest
type UpdateEnvironmentTypeRequest struct {
	Type *EnumEnvironmentType `json:"type,omitempty"`
}

// NewUpdateEnvironmentTypeRequest instantiates a new UpdateEnvironmentTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnvironmentTypeRequest() *UpdateEnvironmentTypeRequest {
	this := UpdateEnvironmentTypeRequest{}
	return &this
}

// NewUpdateEnvironmentTypeRequestWithDefaults instantiates a new UpdateEnvironmentTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnvironmentTypeRequestWithDefaults() *UpdateEnvironmentTypeRequest {
	this := UpdateEnvironmentTypeRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateEnvironmentTypeRequest) GetType() EnumEnvironmentType {
	if o == nil || IsNil(o.Type) {
		var ret EnumEnvironmentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentTypeRequest) GetTypeOk() (*EnumEnvironmentType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateEnvironmentTypeRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumEnvironmentType and assigns it to the Type field.
func (o *UpdateEnvironmentTypeRequest) SetType(v EnumEnvironmentType) {
	o.Type = &v
}

func (o UpdateEnvironmentTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEnvironmentTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUpdateEnvironmentTypeRequest struct {
	value *UpdateEnvironmentTypeRequest
	isSet bool
}

func (v NullableUpdateEnvironmentTypeRequest) Get() *UpdateEnvironmentTypeRequest {
	return v.value
}

func (v *NullableUpdateEnvironmentTypeRequest) Set(val *UpdateEnvironmentTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnvironmentTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnvironmentTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnvironmentTypeRequest(val *UpdateEnvironmentTypeRequest) *NullableUpdateEnvironmentTypeRequest {
	return &NullableUpdateEnvironmentTypeRequest{value: val, isSet: true}
}

func (v NullableUpdateEnvironmentTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnvironmentTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


