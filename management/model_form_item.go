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

// checks if the FormItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormItem{}

// FormItem struct for FormItem
type FormItem struct {
	// A string that specifies the field content.
	Content *string `json:"content,omitempty"`
}

// NewFormItem instantiates a new FormItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormItem() *FormItem {
	this := FormItem{}
	return &this
}

// NewFormItemWithDefaults instantiates a new FormItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormItemWithDefaults() *FormItem {
	this := FormItem{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *FormItem) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormItem) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *FormItem) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *FormItem) SetContent(v string) {
	o.Content = &v
}

func (o FormItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableFormItem struct {
	value *FormItem
	isSet bool
}

func (v NullableFormItem) Get() *FormItem {
	return v.value
}

func (v *NullableFormItem) Set(val *FormItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFormItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFormItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormItem(val *FormItem) *NullableFormItem {
	return &NullableFormItem{value: val, isSet: true}
}

func (v NullableFormItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


