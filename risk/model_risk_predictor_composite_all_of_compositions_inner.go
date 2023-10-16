/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// checks if the RiskPredictorCompositeAllOfCompositionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCompositeAllOfCompositionsInner{}

// RiskPredictorCompositeAllOfCompositionsInner struct for RiskPredictorCompositeAllOfCompositionsInner
type RiskPredictorCompositeAllOfCompositionsInner struct {
	Condition RiskPredictorCompositeConditionBase `json:"condition"`
	Level EnumRiskLevel `json:"level"`
}

// NewRiskPredictorCompositeAllOfCompositionsInner instantiates a new RiskPredictorCompositeAllOfCompositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCompositeAllOfCompositionsInner(condition RiskPredictorCompositeConditionBase, level EnumRiskLevel) *RiskPredictorCompositeAllOfCompositionsInner {
	this := RiskPredictorCompositeAllOfCompositionsInner{}
	this.Condition = condition
	this.Level = level
	return &this
}

// NewRiskPredictorCompositeAllOfCompositionsInnerWithDefaults instantiates a new RiskPredictorCompositeAllOfCompositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCompositeAllOfCompositionsInnerWithDefaults() *RiskPredictorCompositeAllOfCompositionsInner {
	this := RiskPredictorCompositeAllOfCompositionsInner{}
	return &this
}

// GetCondition returns the Condition field value
func (o *RiskPredictorCompositeAllOfCompositionsInner) GetCondition() RiskPredictorCompositeConditionBase {
	if o == nil {
		var ret RiskPredictorCompositeConditionBase
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeAllOfCompositionsInner) GetConditionOk() (*RiskPredictorCompositeConditionBase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *RiskPredictorCompositeAllOfCompositionsInner) SetCondition(v RiskPredictorCompositeConditionBase) {
	o.Condition = v
}

// GetLevel returns the Level field value
func (o *RiskPredictorCompositeAllOfCompositionsInner) GetLevel() EnumRiskLevel {
	if o == nil {
		var ret EnumRiskLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeAllOfCompositionsInner) GetLevelOk() (*EnumRiskLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *RiskPredictorCompositeAllOfCompositionsInner) SetLevel(v EnumRiskLevel) {
	o.Level = v
}

func (o RiskPredictorCompositeAllOfCompositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCompositeAllOfCompositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition"] = o.Condition
	toSerialize["level"] = o.Level
	return toSerialize, nil
}

type NullableRiskPredictorCompositeAllOfCompositionsInner struct {
	value *RiskPredictorCompositeAllOfCompositionsInner
	isSet bool
}

func (v NullableRiskPredictorCompositeAllOfCompositionsInner) Get() *RiskPredictorCompositeAllOfCompositionsInner {
	return v.value
}

func (v *NullableRiskPredictorCompositeAllOfCompositionsInner) Set(val *RiskPredictorCompositeAllOfCompositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCompositeAllOfCompositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCompositeAllOfCompositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCompositeAllOfCompositionsInner(val *RiskPredictorCompositeAllOfCompositionsInner) *NullableRiskPredictorCompositeAllOfCompositionsInner {
	return &NullableRiskPredictorCompositeAllOfCompositionsInner{value: val, isSet: true}
}

func (v NullableRiskPredictorCompositeAllOfCompositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCompositeAllOfCompositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

