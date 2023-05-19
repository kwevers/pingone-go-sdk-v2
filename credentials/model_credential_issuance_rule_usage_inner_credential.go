/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the CredentialIssuanceRuleUsageInnerCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleUsageInnerCredential{}

// CredentialIssuanceRuleUsageInnerCredential struct for CredentialIssuanceRuleUsageInnerCredential
type CredentialIssuanceRuleUsageInnerCredential struct {
	// A string representing the identifier (UUID) of the credential subject to the issue action identified by the credential issuance rule.
	Id string `json:"id"`
}

// NewCredentialIssuanceRuleUsageInnerCredential instantiates a new CredentialIssuanceRuleUsageInnerCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleUsageInnerCredential(id string) *CredentialIssuanceRuleUsageInnerCredential {
	this := CredentialIssuanceRuleUsageInnerCredential{}
	this.Id = id
	return &this
}

// NewCredentialIssuanceRuleUsageInnerCredentialWithDefaults instantiates a new CredentialIssuanceRuleUsageInnerCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleUsageInnerCredentialWithDefaults() *CredentialIssuanceRuleUsageInnerCredential {
	this := CredentialIssuanceRuleUsageInnerCredential{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialIssuanceRuleUsageInnerCredential) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleUsageInnerCredential) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CredentialIssuanceRuleUsageInnerCredential) SetId(v string) {
	o.Id = v
}

func (o CredentialIssuanceRuleUsageInnerCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleUsageInnerCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleUsageInnerCredential struct {
	value *CredentialIssuanceRuleUsageInnerCredential
	isSet bool
}

func (v NullableCredentialIssuanceRuleUsageInnerCredential) Get() *CredentialIssuanceRuleUsageInnerCredential {
	return v.value
}

func (v *NullableCredentialIssuanceRuleUsageInnerCredential) Set(val *CredentialIssuanceRuleUsageInnerCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleUsageInnerCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleUsageInnerCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleUsageInnerCredential(val *CredentialIssuanceRuleUsageInnerCredential) *NullableCredentialIssuanceRuleUsageInnerCredential {
	return &NullableCredentialIssuanceRuleUsageInnerCredential{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleUsageInnerCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleUsageInnerCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

