/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ApplicationSAMLAllOfIdpSigningtype struct for ApplicationSAMLAllOfIdpSigningtype
type ApplicationSAMLAllOfIdpSigningtype struct {
	Key ApplicationSAMLAllOfIdpSigningtypeKey `json:"key"`
}

// NewApplicationSAMLAllOfIdpSigningtype instantiates a new ApplicationSAMLAllOfIdpSigningtype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSAMLAllOfIdpSigningtype(key ApplicationSAMLAllOfIdpSigningtypeKey) *ApplicationSAMLAllOfIdpSigningtype {
	this := ApplicationSAMLAllOfIdpSigningtype{}
	this.Key = key
	return &this
}

// NewApplicationSAMLAllOfIdpSigningtypeWithDefaults instantiates a new ApplicationSAMLAllOfIdpSigningtype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSAMLAllOfIdpSigningtypeWithDefaults() *ApplicationSAMLAllOfIdpSigningtype {
	this := ApplicationSAMLAllOfIdpSigningtype{}
	return &this
}

// GetKey returns the Key field value
func (o *ApplicationSAMLAllOfIdpSigningtype) GetKey() ApplicationSAMLAllOfIdpSigningtypeKey {
	if o == nil {
		var ret ApplicationSAMLAllOfIdpSigningtypeKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOfIdpSigningtype) GetKeyOk() (*ApplicationSAMLAllOfIdpSigningtypeKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ApplicationSAMLAllOfIdpSigningtype) SetKey(v ApplicationSAMLAllOfIdpSigningtypeKey) {
	o.Key = v
}

func (o ApplicationSAMLAllOfIdpSigningtype) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationSAMLAllOfIdpSigningtype struct {
	value *ApplicationSAMLAllOfIdpSigningtype
	isSet bool
}

func (v NullableApplicationSAMLAllOfIdpSigningtype) Get() *ApplicationSAMLAllOfIdpSigningtype {
	return v.value
}

func (v *NullableApplicationSAMLAllOfIdpSigningtype) Set(val *ApplicationSAMLAllOfIdpSigningtype) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSAMLAllOfIdpSigningtype) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSAMLAllOfIdpSigningtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSAMLAllOfIdpSigningtype(val *ApplicationSAMLAllOfIdpSigningtype) *NullableApplicationSAMLAllOfIdpSigningtype {
	return &NullableApplicationSAMLAllOfIdpSigningtype{value: val, isSet: true}
}

func (v NullableApplicationSAMLAllOfIdpSigningtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSAMLAllOfIdpSigningtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


