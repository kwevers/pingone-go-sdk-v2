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

// checks if the FormFieldRecaptchaV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldRecaptchaV2{}

// FormFieldRecaptchaV2 struct for FormFieldRecaptchaV2
type FormFieldRecaptchaV2 struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
	// A string that specifies an identifier for the field component.
	Key string `json:"key"`
	Size EnumFormRecaptchaV2Size `json:"size"`
	Theme EnumFormRecaptchaV2Theme `json:"theme"`
	Alignment EnumFormItemAlignment `json:"alignment"`
}

// NewFormFieldRecaptchaV2 instantiates a new FormFieldRecaptchaV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldRecaptchaV2(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, size EnumFormRecaptchaV2Size, theme EnumFormRecaptchaV2Theme, alignment EnumFormItemAlignment) *FormFieldRecaptchaV2 {
	this := FormFieldRecaptchaV2{}
	this.Type = type_
	this.Position = position
	this.Key = key
	this.Size = size
	this.Theme = theme
	this.Alignment = alignment
	return &this
}

// NewFormFieldRecaptchaV2WithDefaults instantiates a new FormFieldRecaptchaV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldRecaptchaV2WithDefaults() *FormFieldRecaptchaV2 {
	this := FormFieldRecaptchaV2{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldRecaptchaV2) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldRecaptchaV2) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldRecaptchaV2) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldRecaptchaV2) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldRecaptchaV2) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldRecaptchaV2) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

// GetKey returns the Key field value
func (o *FormFieldRecaptchaV2) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FormFieldRecaptchaV2) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *FormFieldRecaptchaV2) SetKey(v string) {
	o.Key = v
}

// GetSize returns the Size field value
func (o *FormFieldRecaptchaV2) GetSize() EnumFormRecaptchaV2Size {
	if o == nil {
		var ret EnumFormRecaptchaV2Size
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FormFieldRecaptchaV2) GetSizeOk() (*EnumFormRecaptchaV2Size, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FormFieldRecaptchaV2) SetSize(v EnumFormRecaptchaV2Size) {
	o.Size = v
}

// GetTheme returns the Theme field value
func (o *FormFieldRecaptchaV2) GetTheme() EnumFormRecaptchaV2Theme {
	if o == nil {
		var ret EnumFormRecaptchaV2Theme
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *FormFieldRecaptchaV2) GetThemeOk() (*EnumFormRecaptchaV2Theme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *FormFieldRecaptchaV2) SetTheme(v EnumFormRecaptchaV2Theme) {
	o.Theme = v
}

// GetAlignment returns the Alignment field value
func (o *FormFieldRecaptchaV2) GetAlignment() EnumFormItemAlignment {
	if o == nil {
		var ret EnumFormItemAlignment
		return ret
	}

	return o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value
// and a boolean to check if the value has been set.
func (o *FormFieldRecaptchaV2) GetAlignmentOk() (*EnumFormItemAlignment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alignment, true
}

// SetAlignment sets field value
func (o *FormFieldRecaptchaV2) SetAlignment(v EnumFormItemAlignment) {
	o.Alignment = v
}

func (o FormFieldRecaptchaV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldRecaptchaV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	toSerialize["key"] = o.Key
	toSerialize["size"] = o.Size
	toSerialize["theme"] = o.Theme
	toSerialize["alignment"] = o.Alignment
	return toSerialize, nil
}

type NullableFormFieldRecaptchaV2 struct {
	value *FormFieldRecaptchaV2
	isSet bool
}

func (v NullableFormFieldRecaptchaV2) Get() *FormFieldRecaptchaV2 {
	return v.value
}

func (v *NullableFormFieldRecaptchaV2) Set(val *FormFieldRecaptchaV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldRecaptchaV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldRecaptchaV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldRecaptchaV2(val *FormFieldRecaptchaV2) *NullableFormFieldRecaptchaV2 {
	return &NullableFormFieldRecaptchaV2{value: val, isSet: true}
}

func (v NullableFormFieldRecaptchaV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldRecaptchaV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


