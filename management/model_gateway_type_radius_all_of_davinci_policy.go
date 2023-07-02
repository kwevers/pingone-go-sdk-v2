/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the GatewayTypeRADIUSAllOfDavinciPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTypeRADIUSAllOfDavinciPolicy{}

// GatewayTypeRADIUSAllOfDavinciPolicy struct for GatewayTypeRADIUSAllOfDavinciPolicy
type GatewayTypeRADIUSAllOfDavinciPolicy struct {
	// The ID of the Davinci flow policy to use.
	Id string `json:"id"`
}

// NewGatewayTypeRADIUSAllOfDavinciPolicy instantiates a new GatewayTypeRADIUSAllOfDavinciPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTypeRADIUSAllOfDavinciPolicy(id string) *GatewayTypeRADIUSAllOfDavinciPolicy {
	this := GatewayTypeRADIUSAllOfDavinciPolicy{}
	this.Id = id
	return &this
}

// NewGatewayTypeRADIUSAllOfDavinciPolicyWithDefaults instantiates a new GatewayTypeRADIUSAllOfDavinciPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTypeRADIUSAllOfDavinciPolicyWithDefaults() *GatewayTypeRADIUSAllOfDavinciPolicy {
	this := GatewayTypeRADIUSAllOfDavinciPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *GatewayTypeRADIUSAllOfDavinciPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeRADIUSAllOfDavinciPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GatewayTypeRADIUSAllOfDavinciPolicy) SetId(v string) {
	o.Id = v
}

func (o GatewayTypeRADIUSAllOfDavinciPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTypeRADIUSAllOfDavinciPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGatewayTypeRADIUSAllOfDavinciPolicy struct {
	value *GatewayTypeRADIUSAllOfDavinciPolicy
	isSet bool
}

func (v NullableGatewayTypeRADIUSAllOfDavinciPolicy) Get() *GatewayTypeRADIUSAllOfDavinciPolicy {
	return v.value
}

func (v *NullableGatewayTypeRADIUSAllOfDavinciPolicy) Set(val *GatewayTypeRADIUSAllOfDavinciPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTypeRADIUSAllOfDavinciPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTypeRADIUSAllOfDavinciPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTypeRADIUSAllOfDavinciPolicy(val *GatewayTypeRADIUSAllOfDavinciPolicy) *NullableGatewayTypeRADIUSAllOfDavinciPolicy {
	return &NullableGatewayTypeRADIUSAllOfDavinciPolicy{value: val, isSet: true}
}

func (v NullableGatewayTypeRADIUSAllOfDavinciPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTypeRADIUSAllOfDavinciPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


