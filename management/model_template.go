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

// checks if the Template type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Template{}

// Template struct for Template
type Template struct {
	// The template id
	Id *string `json:"id,omitempty"`
	// The template’s display name.
	DisplayName string `json:"displayName"`
	// The delivery methods supported for this template. Valid values are `SMS`, `Voice`, `Email` and `Push`.
	DeliveryMethods []string `json:"deliveryMethods"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The description of the template.
	Description *string `json:"description,omitempty"`
	// Lists the variables you can use in each template content. The `required` property indicates whether the variable is required in template content. If `required` is `true`, the `requiredForDeliveryMethods` property lists the `deliveryMethods` types that require the variable. Note that if `required` is `true`, but `requiredForDeliveryMethods` is not returned, all `deliveryMethods` types are required. For more information, see [Variables](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-variables).
	Variables map[string]interface{} `json:"variables"`
	// Specifies whether dynamic variables can be used in the template's contents. For more information, see [Dynamic variables](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-dynamic-variables).
	AllowDynamicVariables *bool `json:"allowDynamicVariables,omitempty"`
}

// NewTemplate instantiates a new Template object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplate(displayName string, deliveryMethods []string, variables map[string]interface{}) *Template {
	this := Template{}
	this.DisplayName = displayName
	this.DeliveryMethods = deliveryMethods
	this.Variables = variables
	return &this
}

// NewTemplateWithDefaults instantiates a new Template object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateWithDefaults() *Template {
	this := Template{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Template) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Template) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Template) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value
func (o *Template) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Template) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Template) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDeliveryMethods returns the DeliveryMethods field value
func (o *Template) GetDeliveryMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeliveryMethods
}

// GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field value
// and a boolean to check if the value has been set.
func (o *Template) GetDeliveryMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryMethods, true
}

// SetDeliveryMethods sets field value
func (o *Template) SetDeliveryMethods(v []string) {
	o.DeliveryMethods = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Template) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Template) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Template) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Template) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Template) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Template) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Template) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Template) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Template) SetDescription(v string) {
	o.Description = &v
}

// GetVariables returns the Variables field value
func (o *Template) GetVariables() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *Template) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// SetVariables sets field value
func (o *Template) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

// GetAllowDynamicVariables returns the AllowDynamicVariables field value if set, zero value otherwise.
func (o *Template) GetAllowDynamicVariables() bool {
	if o == nil || IsNil(o.AllowDynamicVariables) {
		var ret bool
		return ret
	}
	return *o.AllowDynamicVariables
}

// GetAllowDynamicVariablesOk returns a tuple with the AllowDynamicVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetAllowDynamicVariablesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDynamicVariables) {
		return nil, false
	}
	return o.AllowDynamicVariables, true
}

// HasAllowDynamicVariables returns a boolean if a field has been set.
func (o *Template) HasAllowDynamicVariables() bool {
	if o != nil && !IsNil(o.AllowDynamicVariables) {
		return true
	}

	return false
}

// SetAllowDynamicVariables gets a reference to the given bool and assigns it to the AllowDynamicVariables field.
func (o *Template) SetAllowDynamicVariables(v bool) {
	o.AllowDynamicVariables = &v
}

func (o Template) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Template) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["displayName"] = o.DisplayName
	toSerialize["deliveryMethods"] = o.DeliveryMethods
	// skip: createdAt is readOnly
	// skip: updatedAt is readOnly
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["variables"] = o.Variables
	if !IsNil(o.AllowDynamicVariables) {
		toSerialize["allowDynamicVariables"] = o.AllowDynamicVariables
	}
	return toSerialize, nil
}

type NullableTemplate struct {
	value *Template
	isSet bool
}

func (v NullableTemplate) Get() *Template {
	return v.value
}

func (v *NullableTemplate) Set(val *Template) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplate(val *Template) *NullableTemplate {
	return &NullableTemplate{value: val, isSet: true}
}

func (v NullableTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


