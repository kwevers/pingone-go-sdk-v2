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

// checks if the Schema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schema{}

// Schema struct for Schema
type Schema struct {
	// Indicates whether or not the `contains` operator can be used. You can use the `contains` operator in a maximum of 5 custom attributes.
	AllowsContainsOperator *bool `json:"allowsContainsOperator,omitempty"`
	Attributes []SchemaAttribute `json:"attributes,omitempty"`
	// A string that specifies the description of the schema.
	Description *string `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the resource name.
	Name *string `json:"name,omitempty"`
}

// NewSchema instantiates a new Schema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchema() *Schema {
	this := Schema{}
	return &this
}

// NewSchemaWithDefaults instantiates a new Schema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaWithDefaults() *Schema {
	this := Schema{}
	return &this
}

// GetAllowsContainsOperator returns the AllowsContainsOperator field value if set, zero value otherwise.
func (o *Schema) GetAllowsContainsOperator() bool {
	if o == nil || IsNil(o.AllowsContainsOperator) {
		var ret bool
		return ret
	}
	return *o.AllowsContainsOperator
}

// GetAllowsContainsOperatorOk returns a tuple with the AllowsContainsOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetAllowsContainsOperatorOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowsContainsOperator) {
		return nil, false
	}
	return o.AllowsContainsOperator, true
}

// HasAllowsContainsOperator returns a boolean if a field has been set.
func (o *Schema) HasAllowsContainsOperator() bool {
	if o != nil && !IsNil(o.AllowsContainsOperator) {
		return true
	}

	return false
}

// SetAllowsContainsOperator gets a reference to the given bool and assigns it to the AllowsContainsOperator field.
func (o *Schema) SetAllowsContainsOperator(v bool) {
	o.AllowsContainsOperator = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Schema) GetAttributes() []SchemaAttribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []SchemaAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetAttributesOk() ([]SchemaAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Schema) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []SchemaAttribute and assigns it to the Attributes field.
func (o *Schema) SetAttributes(v []SchemaAttribute) {
	o.Attributes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Schema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Schema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Schema) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Schema) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Schema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Schema) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Schema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Schema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Schema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Schema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Schema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Schema) SetName(v string) {
	o.Name = &v
}

func (o Schema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: allowsContainsOperator is readOnly
	// skip: attributes is readOnly
	// skip: description is readOnly
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	// skip: name is readOnly
	return toSerialize, nil
}

type NullableSchema struct {
	value *Schema
	isSet bool
}

func (v NullableSchema) Get() *Schema {
	return v.value
}

func (v *NullableSchema) Set(val *Schema) {
	v.value = val
	v.isSet = true
}

func (v NullableSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchema(val *Schema) *NullableSchema {
	return &NullableSchema{value: val, isSet: true}
}

func (v NullableSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


