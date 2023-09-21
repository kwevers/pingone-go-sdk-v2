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

// checks if the ProvisionedCredentialUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionedCredentialUser{}

// ProvisionedCredentialUser struct for ProvisionedCredentialUser
type ProvisionedCredentialUser struct {
	// A string that specifies the identifier (UUID) of the user associated with the provisioned credential.
	Id *string `json:"id,omitempty"`
}

// NewProvisionedCredentialUser instantiates a new ProvisionedCredentialUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionedCredentialUser() *ProvisionedCredentialUser {
	this := ProvisionedCredentialUser{}
	return &this
}

// NewProvisionedCredentialUserWithDefaults instantiates a new ProvisionedCredentialUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionedCredentialUserWithDefaults() *ProvisionedCredentialUser {
	this := ProvisionedCredentialUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProvisionedCredentialUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredentialUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProvisionedCredentialUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProvisionedCredentialUser) SetId(v string) {
	o.Id = &v
}

func (o ProvisionedCredentialUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionedCredentialUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableProvisionedCredentialUser struct {
	value *ProvisionedCredentialUser
	isSet bool
}

func (v NullableProvisionedCredentialUser) Get() *ProvisionedCredentialUser {
	return v.value
}

func (v *NullableProvisionedCredentialUser) Set(val *ProvisionedCredentialUser) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionedCredentialUser) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionedCredentialUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionedCredentialUser(val *ProvisionedCredentialUser) *NullableProvisionedCredentialUser {
	return &NullableProvisionedCredentialUser{value: val, isSet: true}
}

func (v NullableProvisionedCredentialUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionedCredentialUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


