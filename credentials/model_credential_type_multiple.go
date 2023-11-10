/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the CredentialTypeMultiple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialTypeMultiple{}

// CredentialTypeMultiple An object containing directives that enable the issuance of multiple credentials to a user based on a directory attribute.
type CredentialTypeMultiple struct {
	// A PingOne Expression Language (PEL) expression evaluated by the P1 Credentials service on issuance.If an array, calculates the array length for the count. Populates the limit to a variable, __ITERATOR__, available to PEL expressions in metadata.fields.attribute.
	Expression string `json:"expression"`
}

// NewCredentialTypeMultiple instantiates a new CredentialTypeMultiple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialTypeMultiple(expression string) *CredentialTypeMultiple {
	this := CredentialTypeMultiple{}
	this.Expression = expression
	return &this
}

// NewCredentialTypeMultipleWithDefaults instantiates a new CredentialTypeMultiple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialTypeMultipleWithDefaults() *CredentialTypeMultiple {
	this := CredentialTypeMultiple{}
	return &this
}

// GetExpression returns the Expression field value
func (o *CredentialTypeMultiple) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CredentialTypeMultiple) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *CredentialTypeMultiple) SetExpression(v string) {
	o.Expression = v
}

func (o CredentialTypeMultiple) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialTypeMultiple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableCredentialTypeMultiple struct {
	value *CredentialTypeMultiple
	isSet bool
}

func (v NullableCredentialTypeMultiple) Get() *CredentialTypeMultiple {
	return v.value
}

func (v *NullableCredentialTypeMultiple) Set(val *CredentialTypeMultiple) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialTypeMultiple) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialTypeMultiple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialTypeMultiple(val *CredentialTypeMultiple) *NullableCredentialTypeMultiple {
	return &NullableCredentialTypeMultiple{value: val, isSet: true}
}

func (v NullableCredentialTypeMultiple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialTypeMultiple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

