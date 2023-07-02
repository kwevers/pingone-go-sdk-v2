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

// checks if the BillOfMaterials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillOfMaterials{}

// BillOfMaterials struct for BillOfMaterials
type BillOfMaterials struct {
	SolutionType *EnumSolutionType `json:"solutionType,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// An array that specifies the products associated with this bill of materials
	Products []BillOfMaterialsProductsInner `json:"products"`
}

// NewBillOfMaterials instantiates a new BillOfMaterials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillOfMaterials(products []BillOfMaterialsProductsInner) *BillOfMaterials {
	this := BillOfMaterials{}
	this.Products = products
	return &this
}

// NewBillOfMaterialsWithDefaults instantiates a new BillOfMaterials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillOfMaterialsWithDefaults() *BillOfMaterials {
	this := BillOfMaterials{}
	return &this
}

// GetSolutionType returns the SolutionType field value if set, zero value otherwise.
func (o *BillOfMaterials) GetSolutionType() EnumSolutionType {
	if o == nil || IsNil(o.SolutionType) {
		var ret EnumSolutionType
		return ret
	}
	return *o.SolutionType
}

// GetSolutionTypeOk returns a tuple with the SolutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterials) GetSolutionTypeOk() (*EnumSolutionType, bool) {
	if o == nil || IsNil(o.SolutionType) {
		return nil, false
	}
	return o.SolutionType, true
}

// HasSolutionType returns a boolean if a field has been set.
func (o *BillOfMaterials) HasSolutionType() bool {
	if o != nil && !IsNil(o.SolutionType) {
		return true
	}

	return false
}

// SetSolutionType gets a reference to the given EnumSolutionType and assigns it to the SolutionType field.
func (o *BillOfMaterials) SetSolutionType(v EnumSolutionType) {
	o.SolutionType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BillOfMaterials) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterials) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BillOfMaterials) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *BillOfMaterials) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BillOfMaterials) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterials) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BillOfMaterials) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *BillOfMaterials) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetProducts returns the Products field value
func (o *BillOfMaterials) GetProducts() []BillOfMaterialsProductsInner {
	if o == nil {
		var ret []BillOfMaterialsProductsInner
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *BillOfMaterials) GetProductsOk() ([]BillOfMaterialsProductsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *BillOfMaterials) SetProducts(v []BillOfMaterialsProductsInner) {
	o.Products = v
}

func (o BillOfMaterials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillOfMaterials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SolutionType) {
		toSerialize["solutionType"] = o.SolutionType
	}
	// skip: createdAt is readOnly
	// skip: updatedAt is readOnly
	toSerialize["products"] = o.Products
	return toSerialize, nil
}

type NullableBillOfMaterials struct {
	value *BillOfMaterials
	isSet bool
}

func (v NullableBillOfMaterials) Get() *BillOfMaterials {
	return v.value
}

func (v *NullableBillOfMaterials) Set(val *BillOfMaterials) {
	v.value = val
	v.isSet = true
}

func (v NullableBillOfMaterials) IsSet() bool {
	return v.isSet
}

func (v *NullableBillOfMaterials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillOfMaterials(val *BillOfMaterials) *NullableBillOfMaterials {
	return &NullableBillOfMaterials{value: val, isSet: true}
}

func (v NullableBillOfMaterials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillOfMaterials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


