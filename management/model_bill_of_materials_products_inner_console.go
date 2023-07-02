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

// checks if the BillOfMaterialsProductsInnerConsole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillOfMaterialsProductsInnerConsole{}

// BillOfMaterialsProductsInnerConsole struct for BillOfMaterialsProductsInnerConsole
type BillOfMaterialsProductsInnerConsole struct {
	// Primary console link for certain products
	Href string `json:"href"`
}

// NewBillOfMaterialsProductsInnerConsole instantiates a new BillOfMaterialsProductsInnerConsole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillOfMaterialsProductsInnerConsole(href string) *BillOfMaterialsProductsInnerConsole {
	this := BillOfMaterialsProductsInnerConsole{}
	this.Href = href
	return &this
}

// NewBillOfMaterialsProductsInnerConsoleWithDefaults instantiates a new BillOfMaterialsProductsInnerConsole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillOfMaterialsProductsInnerConsoleWithDefaults() *BillOfMaterialsProductsInnerConsole {
	this := BillOfMaterialsProductsInnerConsole{}
	return &this
}

// GetHref returns the Href field value
func (o *BillOfMaterialsProductsInnerConsole) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInnerConsole) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *BillOfMaterialsProductsInnerConsole) SetHref(v string) {
	o.Href = v
}

func (o BillOfMaterialsProductsInnerConsole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillOfMaterialsProductsInnerConsole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	return toSerialize, nil
}

type NullableBillOfMaterialsProductsInnerConsole struct {
	value *BillOfMaterialsProductsInnerConsole
	isSet bool
}

func (v NullableBillOfMaterialsProductsInnerConsole) Get() *BillOfMaterialsProductsInnerConsole {
	return v.value
}

func (v *NullableBillOfMaterialsProductsInnerConsole) Set(val *BillOfMaterialsProductsInnerConsole) {
	v.value = val
	v.isSet = true
}

func (v NullableBillOfMaterialsProductsInnerConsole) IsSet() bool {
	return v.isSet
}

func (v *NullableBillOfMaterialsProductsInnerConsole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillOfMaterialsProductsInnerConsole(val *BillOfMaterialsProductsInnerConsole) *NullableBillOfMaterialsProductsInnerConsole {
	return &NullableBillOfMaterialsProductsInnerConsole{value: val, isSet: true}
}

func (v NullableBillOfMaterialsProductsInnerConsole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillOfMaterialsProductsInnerConsole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


