/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"time"
)

// checks if the MFAPushCredentialResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFAPushCredentialResponse{}

// MFAPushCredentialResponse struct for MFAPushCredentialResponse
type MFAPushCredentialResponse struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// A string that specifies the push credential ID.
	Id *string `json:"id,omitempty"`
	Type *EnumMFAPushCredentialAttrType `json:"type,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewMFAPushCredentialResponse instantiates a new MFAPushCredentialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredentialResponse() *MFAPushCredentialResponse {
	this := MFAPushCredentialResponse{}
	return &this
}

// NewMFAPushCredentialResponseWithDefaults instantiates a new MFAPushCredentialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialResponseWithDefaults() *MFAPushCredentialResponse {
	this := MFAPushCredentialResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *MFAPushCredentialResponse) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MFAPushCredentialResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetType() EnumMFAPushCredentialAttrType {
	if o == nil || IsNil(o.Type) {
		var ret EnumMFAPushCredentialAttrType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumMFAPushCredentialAttrType and assigns it to the Type field.
func (o *MFAPushCredentialResponse) SetType(v EnumMFAPushCredentialAttrType) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MFAPushCredentialResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MFAPushCredentialResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o MFAPushCredentialResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFAPushCredentialResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableMFAPushCredentialResponse struct {
	value *MFAPushCredentialResponse
	isSet bool
}

func (v NullableMFAPushCredentialResponse) Get() *MFAPushCredentialResponse {
	return v.value
}

func (v *NullableMFAPushCredentialResponse) Set(val *MFAPushCredentialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialResponse(val *MFAPushCredentialResponse) *NullableMFAPushCredentialResponse {
	return &NullableMFAPushCredentialResponse{value: val, isSet: true}
}

func (v NullableMFAPushCredentialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


