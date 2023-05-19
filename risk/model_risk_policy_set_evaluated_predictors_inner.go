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

// checks if the RiskPolicySetEvaluatedPredictorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicySetEvaluatedPredictorsInner{}

// RiskPolicySetEvaluatedPredictorsInner struct for RiskPolicySetEvaluatedPredictorsInner
type RiskPolicySetEvaluatedPredictorsInner struct {
	// A string that specifies the resource’s unique identifier.
	Id string `json:"id"`
}

// NewRiskPolicySetEvaluatedPredictorsInner instantiates a new RiskPolicySetEvaluatedPredictorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicySetEvaluatedPredictorsInner(id string) *RiskPolicySetEvaluatedPredictorsInner {
	this := RiskPolicySetEvaluatedPredictorsInner{}
	this.Id = id
	return &this
}

// NewRiskPolicySetEvaluatedPredictorsInnerWithDefaults instantiates a new RiskPolicySetEvaluatedPredictorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicySetEvaluatedPredictorsInnerWithDefaults() *RiskPolicySetEvaluatedPredictorsInner {
	this := RiskPolicySetEvaluatedPredictorsInner{}
	return &this
}

// GetId returns the Id field value
func (o *RiskPolicySetEvaluatedPredictorsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetEvaluatedPredictorsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RiskPolicySetEvaluatedPredictorsInner) SetId(v string) {
	o.Id = v
}

func (o RiskPolicySetEvaluatedPredictorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicySetEvaluatedPredictorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableRiskPolicySetEvaluatedPredictorsInner struct {
	value *RiskPolicySetEvaluatedPredictorsInner
	isSet bool
}

func (v NullableRiskPolicySetEvaluatedPredictorsInner) Get() *RiskPolicySetEvaluatedPredictorsInner {
	return v.value
}

func (v *NullableRiskPolicySetEvaluatedPredictorsInner) Set(val *RiskPolicySetEvaluatedPredictorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySetEvaluatedPredictorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySetEvaluatedPredictorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySetEvaluatedPredictorsInner(val *RiskPolicySetEvaluatedPredictorsInner) *NullableRiskPolicySetEvaluatedPredictorsInner {
	return &NullableRiskPolicySetEvaluatedPredictorsInner{value: val, isSet: true}
}

func (v NullableRiskPolicySetEvaluatedPredictorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySetEvaluatedPredictorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


