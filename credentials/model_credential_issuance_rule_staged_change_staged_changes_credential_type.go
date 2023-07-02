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

// checks if the CredentialIssuanceRuleStagedChangeStagedChangesCredentialType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleStagedChangeStagedChangesCredentialType{}

// CredentialIssuanceRuleStagedChangeStagedChangesCredentialType struct for CredentialIssuanceRuleStagedChangeStagedChangesCredentialType
type CredentialIssuanceRuleStagedChangeStagedChangesCredentialType struct {
	// A string that specifies the identifier (UUID) of the credential type with which this credential issuance rule is associated.
	Id *string `json:"id,omitempty"`
}

// NewCredentialIssuanceRuleStagedChangeStagedChangesCredentialType instantiates a new CredentialIssuanceRuleStagedChangeStagedChangesCredentialType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleStagedChangeStagedChangesCredentialType() *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType {
	this := CredentialIssuanceRuleStagedChangeStagedChangesCredentialType{}
	return &this
}

// NewCredentialIssuanceRuleStagedChangeStagedChangesCredentialTypeWithDefaults instantiates a new CredentialIssuanceRuleStagedChangeStagedChangesCredentialType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleStagedChangeStagedChangesCredentialTypeWithDefaults() *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType {
	this := CredentialIssuanceRuleStagedChangeStagedChangesCredentialType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) SetId(v string) {
	o.Id = &v
}

func (o CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType struct {
	value *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType
	isSet bool
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType) Get() *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType {
	return v.value
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType) Set(val *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType(val *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) *NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType {
	return &NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChangesCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


