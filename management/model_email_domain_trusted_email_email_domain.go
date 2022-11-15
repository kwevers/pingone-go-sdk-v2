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

// EmailDomainTrustedEmailEmailDomain struct for EmailDomainTrustedEmailEmailDomain
type EmailDomainTrustedEmailEmailDomain struct {
	// A string that specifies the trusted email domain resource’s unique identifier associated with the resource.
	Id *string `json:"id,omitempty"`
}

// NewEmailDomainTrustedEmailEmailDomain instantiates a new EmailDomainTrustedEmailEmailDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainTrustedEmailEmailDomain() *EmailDomainTrustedEmailEmailDomain {
	this := EmailDomainTrustedEmailEmailDomain{}
	return &this
}

// NewEmailDomainTrustedEmailEmailDomainWithDefaults instantiates a new EmailDomainTrustedEmailEmailDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainTrustedEmailEmailDomainWithDefaults() *EmailDomainTrustedEmailEmailDomain {
	this := EmailDomainTrustedEmailEmailDomain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailDomainTrustedEmailEmailDomain) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmailEmailDomain) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailDomainTrustedEmailEmailDomain) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailDomainTrustedEmailEmailDomain) SetId(v string) {
	o.Id = &v
}

func (o EmailDomainTrustedEmailEmailDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDomainTrustedEmailEmailDomain struct {
	value *EmailDomainTrustedEmailEmailDomain
	isSet bool
}

func (v NullableEmailDomainTrustedEmailEmailDomain) Get() *EmailDomainTrustedEmailEmailDomain {
	return v.value
}

func (v *NullableEmailDomainTrustedEmailEmailDomain) Set(val *EmailDomainTrustedEmailEmailDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainTrustedEmailEmailDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainTrustedEmailEmailDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainTrustedEmailEmailDomain(val *EmailDomainTrustedEmailEmailDomain) *NullableEmailDomainTrustedEmailEmailDomain {
	return &NullableEmailDomainTrustedEmailEmailDomain{value: val, isSet: true}
}

func (v NullableEmailDomainTrustedEmailEmailDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainTrustedEmailEmailDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


