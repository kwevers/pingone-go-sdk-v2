/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the P1ErrorDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &P1ErrorDetailsInner{}

// P1ErrorDetailsInner struct for P1ErrorDetailsInner
type P1ErrorDetailsInner struct {
	// A general fault code which the client must handle to provide all exception handling routines and to localize messages for users. This code is common across all PingOne services and is human readable (such as a defined constant rather than a number).
	Code *string `json:"code,omitempty"`
	// The item that caused the error (such as a form field ID or an attribute inside a JSON object).
	Target *string `json:"target,omitempty"`
	// A short description of the error. This message is intended to assist with debugging and is returned in English only.
	Message *string `json:"message,omitempty"`
	InnerError *P1ErrorDetailsInnerInnerError `json:"innerError,omitempty"`
}

// NewP1ErrorDetailsInner instantiates a new P1ErrorDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP1ErrorDetailsInner() *P1ErrorDetailsInner {
	this := P1ErrorDetailsInner{}
	return &this
}

// NewP1ErrorDetailsInnerWithDefaults instantiates a new P1ErrorDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP1ErrorDetailsInnerWithDefaults() *P1ErrorDetailsInner {
	this := P1ErrorDetailsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *P1ErrorDetailsInner) SetCode(v string) {
	o.Code = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *P1ErrorDetailsInner) SetTarget(v string) {
	o.Target = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *P1ErrorDetailsInner) SetMessage(v string) {
	o.Message = &v
}

// GetInnerError returns the InnerError field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetInnerError() P1ErrorDetailsInnerInnerError {
	if o == nil || IsNil(o.InnerError) {
		var ret P1ErrorDetailsInnerInnerError
		return ret
	}
	return *o.InnerError
}

// GetInnerErrorOk returns a tuple with the InnerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetInnerErrorOk() (*P1ErrorDetailsInnerInnerError, bool) {
	if o == nil || IsNil(o.InnerError) {
		return nil, false
	}
	return o.InnerError, true
}

// HasInnerError returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasInnerError() bool {
	if o != nil && !IsNil(o.InnerError) {
		return true
	}

	return false
}

// SetInnerError gets a reference to the given P1ErrorDetailsInnerInnerError and assigns it to the InnerError field.
func (o *P1ErrorDetailsInner) SetInnerError(v P1ErrorDetailsInnerInnerError) {
	o.InnerError = &v
}

func (o P1ErrorDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o P1ErrorDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.InnerError) {
		toSerialize["innerError"] = o.InnerError
	}
	return toSerialize, nil
}

type NullableP1ErrorDetailsInner struct {
	value *P1ErrorDetailsInner
	isSet bool
}

func (v NullableP1ErrorDetailsInner) Get() *P1ErrorDetailsInner {
	return v.value
}

func (v *NullableP1ErrorDetailsInner) Set(val *P1ErrorDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableP1ErrorDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableP1ErrorDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP1ErrorDetailsInner(val *P1ErrorDetailsInner) *NullableP1ErrorDetailsInner {
	return &NullableP1ErrorDetailsInner{value: val, isSet: true}
}

func (v NullableP1ErrorDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP1ErrorDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


