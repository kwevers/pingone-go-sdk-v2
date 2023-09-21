/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// checks if the ResourceScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceScope{}

// ResourceScope struct for ResourceScope
type ResourceScope struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the resource scope name.
	Name string `json:"name"`
	// A string that specifies the description of the scope.
	Description *string `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// An array that specifies the user schema attributes that can be read or updated for the specified PingOne access control scope. The value is an array of schema attribute paths (such as `username`, `name.given`, `shirtSize`) that the scope controls. This property is supported only for the `p1:read:user`, `p1:update:user` and `p1:read:user:{suffix}` and `p1:update:user:{suffix}` scopes. No other PingOne platform scopes allow this behavior. Any attributes not listed in the attribute array are excluded from the read or update action. The wildcard path (*) in the array includes all attributes and cannot be used in conjunction with any other user schema attribute paths
	SchemaAttributes []string `json:"schemaAttributes,omitempty"`
	// A list of custom resource attribute IDs. This property applies only for the resource with its type property set to `OPENID_CONNECT`. Moreover, this property does not display predefined OpenID Connect (OIDC) mappings, such as the `email` claim in the OIDC `email` scope or the `name` claim in the `profile` scope. You can create custom attributes, and these custom attributes can be added to `mappedClaims` and will display in the response.
	MappedClaims []string `json:"mappedClaims,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewResourceScope instantiates a new ResourceScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceScope(name string) *ResourceScope {
	this := ResourceScope{}
	this.Name = name
	return &this
}

// NewResourceScopeWithDefaults instantiates a new ResourceScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceScopeWithDefaults() *ResourceScope {
	this := ResourceScope{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceScope) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceScope) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *ResourceScope) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceScope) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceScope) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceScope) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ResourceScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceScope) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceScope) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceScope) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceScope) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ResourceScope) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ResourceScope) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ResourceScope) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetSchemaAttributes returns the SchemaAttributes field value if set, zero value otherwise.
func (o *ResourceScope) GetSchemaAttributes() []string {
	if o == nil || IsNil(o.SchemaAttributes) {
		var ret []string
		return ret
	}
	return o.SchemaAttributes
}

// GetSchemaAttributesOk returns a tuple with the SchemaAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetSchemaAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.SchemaAttributes) {
		return nil, false
	}
	return o.SchemaAttributes, true
}

// HasSchemaAttributes returns a boolean if a field has been set.
func (o *ResourceScope) HasSchemaAttributes() bool {
	if o != nil && !IsNil(o.SchemaAttributes) {
		return true
	}

	return false
}

// SetSchemaAttributes gets a reference to the given []string and assigns it to the SchemaAttributes field.
func (o *ResourceScope) SetSchemaAttributes(v []string) {
	o.SchemaAttributes = v
}

// GetMappedClaims returns the MappedClaims field value if set, zero value otherwise.
func (o *ResourceScope) GetMappedClaims() []string {
	if o == nil || IsNil(o.MappedClaims) {
		var ret []string
		return ret
	}
	return o.MappedClaims
}

// GetMappedClaimsOk returns a tuple with the MappedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetMappedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.MappedClaims) {
		return nil, false
	}
	return o.MappedClaims, true
}

// HasMappedClaims returns a boolean if a field has been set.
func (o *ResourceScope) HasMappedClaims() bool {
	if o != nil && !IsNil(o.MappedClaims) {
		return true
	}

	return false
}

// SetMappedClaims gets a reference to the given []string and assigns it to the MappedClaims field.
func (o *ResourceScope) SetMappedClaims(v []string) {
	o.MappedClaims = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ResourceScope) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ResourceScope) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ResourceScope) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ResourceScope) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ResourceScope) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ResourceScope) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ResourceScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.SchemaAttributes) {
		toSerialize["schemaAttributes"] = o.SchemaAttributes
	}
	if !IsNil(o.MappedClaims) {
		toSerialize["mappedClaims"] = o.MappedClaims
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableResourceScope struct {
	value *ResourceScope
	isSet bool
}

func (v NullableResourceScope) Get() *ResourceScope {
	return v.value
}

func (v *NullableResourceScope) Set(val *ResourceScope) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceScope) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceScope(val *ResourceScope) *NullableResourceScope {
	return &NullableResourceScope{value: val, isSet: true}
}

func (v NullableResourceScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


