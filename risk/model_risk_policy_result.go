/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// RiskPolicyResult struct for RiskPolicyResult
type RiskPolicyResult struct {
	Level EnumRiskLevel `json:"level"`
	Type *EnumResultType `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewRiskPolicyResult instantiates a new RiskPolicyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicyResult(level EnumRiskLevel) *RiskPolicyResult {
	this := RiskPolicyResult{}
	this.Level = level
	return &this
}

// NewRiskPolicyResultWithDefaults instantiates a new RiskPolicyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicyResultWithDefaults() *RiskPolicyResult {
	this := RiskPolicyResult{}
	return &this
}

// GetLevel returns the Level field value
func (o *RiskPolicyResult) GetLevel() EnumRiskLevel {
	if o == nil {
		var ret EnumRiskLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *RiskPolicyResult) GetLevelOk() (*EnumRiskLevel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *RiskPolicyResult) SetLevel(v EnumRiskLevel) {
	o.Level = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPolicyResult) GetType() EnumResultType {
	if o == nil || isNil(o.Type) {
		var ret EnumResultType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyResult) GetTypeOk() (*EnumResultType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPolicyResult) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumResultType and assigns it to the Type field.
func (o *RiskPolicyResult) SetType(v EnumResultType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskPolicyResult) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyResult) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskPolicyResult) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskPolicyResult) SetValue(v string) {
	o.Value = &v
}

func (o RiskPolicyResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPolicyResult struct {
	value *RiskPolicyResult
	isSet bool
}

func (v NullableRiskPolicyResult) Get() *RiskPolicyResult {
	return v.value
}

func (v *NullableRiskPolicyResult) Set(val *RiskPolicyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicyResult(val *RiskPolicyResult) *NullableRiskPolicyResult {
	return &NullableRiskPolicyResult{value: val, isSet: true}
}

func (v NullableRiskPolicyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


