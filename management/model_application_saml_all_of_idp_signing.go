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

// checks if the ApplicationSAMLAllOfIdpSigning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSAMLAllOfIdpSigning{}

// ApplicationSAMLAllOfIdpSigning struct for ApplicationSAMLAllOfIdpSigning
type ApplicationSAMLAllOfIdpSigning struct {
	Key ApplicationSAMLAllOfIdpSigningKey `json:"key"`
	Algorithm *EnumCertificateKeySignagureAlgorithm `json:"algorithm,omitempty"`
}

// NewApplicationSAMLAllOfIdpSigning instantiates a new ApplicationSAMLAllOfIdpSigning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSAMLAllOfIdpSigning(key ApplicationSAMLAllOfIdpSigningKey) *ApplicationSAMLAllOfIdpSigning {
	this := ApplicationSAMLAllOfIdpSigning{}
	this.Key = key
	return &this
}

// NewApplicationSAMLAllOfIdpSigningWithDefaults instantiates a new ApplicationSAMLAllOfIdpSigning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSAMLAllOfIdpSigningWithDefaults() *ApplicationSAMLAllOfIdpSigning {
	this := ApplicationSAMLAllOfIdpSigning{}
	return &this
}

// GetKey returns the Key field value
func (o *ApplicationSAMLAllOfIdpSigning) GetKey() ApplicationSAMLAllOfIdpSigningKey {
	if o == nil {
		var ret ApplicationSAMLAllOfIdpSigningKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOfIdpSigning) GetKeyOk() (*ApplicationSAMLAllOfIdpSigningKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ApplicationSAMLAllOfIdpSigning) SetKey(v ApplicationSAMLAllOfIdpSigningKey) {
	o.Key = v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOfIdpSigning) GetAlgorithm() EnumCertificateKeySignagureAlgorithm {
	if o == nil || IsNil(o.Algorithm) {
		var ret EnumCertificateKeySignagureAlgorithm
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOfIdpSigning) GetAlgorithmOk() (*EnumCertificateKeySignagureAlgorithm, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOfIdpSigning) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given EnumCertificateKeySignagureAlgorithm and assigns it to the Algorithm field.
func (o *ApplicationSAMLAllOfIdpSigning) SetAlgorithm(v EnumCertificateKeySignagureAlgorithm) {
	o.Algorithm = &v
}

func (o ApplicationSAMLAllOfIdpSigning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSAMLAllOfIdpSigning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	return toSerialize, nil
}

type NullableApplicationSAMLAllOfIdpSigning struct {
	value *ApplicationSAMLAllOfIdpSigning
	isSet bool
}

func (v NullableApplicationSAMLAllOfIdpSigning) Get() *ApplicationSAMLAllOfIdpSigning {
	return v.value
}

func (v *NullableApplicationSAMLAllOfIdpSigning) Set(val *ApplicationSAMLAllOfIdpSigning) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSAMLAllOfIdpSigning) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSAMLAllOfIdpSigning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSAMLAllOfIdpSigning(val *ApplicationSAMLAllOfIdpSigning) *NullableApplicationSAMLAllOfIdpSigning {
	return &NullableApplicationSAMLAllOfIdpSigning{value: val, isSet: true}
}

func (v NullableApplicationSAMLAllOfIdpSigning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSAMLAllOfIdpSigning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


