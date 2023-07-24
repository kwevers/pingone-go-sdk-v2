/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the LinksHATEOASNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksHATEOASNext{}

// LinksHATEOASNext An object that describes the next page of results. This property is present only if there is a next page of results and the `limit` parameter is used.
type LinksHATEOASNext struct {
	// The URI of the resource.
	Href *string `json:"href,omitempty"`
}

// NewLinksHATEOASNext instantiates a new LinksHATEOASNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksHATEOASNext() *LinksHATEOASNext {
	this := LinksHATEOASNext{}
	return &this
}

// NewLinksHATEOASNextWithDefaults instantiates a new LinksHATEOASNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksHATEOASNextWithDefaults() *LinksHATEOASNext {
	this := LinksHATEOASNext{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LinksHATEOASNext) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksHATEOASNext) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinksHATEOASNext) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinksHATEOASNext) SetHref(v string) {
	o.Href = &v
}

func (o LinksHATEOASNext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksHATEOASNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableLinksHATEOASNext struct {
	value *LinksHATEOASNext
	isSet bool
}

func (v NullableLinksHATEOASNext) Get() *LinksHATEOASNext {
	return v.value
}

func (v *NullableLinksHATEOASNext) Set(val *LinksHATEOASNext) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksHATEOASNext) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksHATEOASNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksHATEOASNext(val *LinksHATEOASNext) *NullableLinksHATEOASNext {
	return &NullableLinksHATEOASNext{value: val, isSet: true}
}

func (v NullableLinksHATEOASNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksHATEOASNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


