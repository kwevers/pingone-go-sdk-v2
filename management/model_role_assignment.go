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

// RoleAssignment struct for RoleAssignment
type RoleAssignment struct {
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Gateway *GatewayInstanceGateway `json:"gateway,omitempty"`
	// A string that specifies the user role assignment ID.
	Id *string `json:"id,omitempty"`
	// A boolean that specifies whether this role assignment can be deleted by the current actor.
	ReadOnly *bool `json:"readOnly,omitempty"`
	Role RoleAssignmentRole `json:"role"`
	Scope RoleAssignmentScope `json:"scope"`
}

// NewRoleAssignment instantiates a new RoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignment(role RoleAssignmentRole, scope RoleAssignmentScope) *RoleAssignment {
	this := RoleAssignment{}
	this.Role = role
	this.Scope = scope
	return &this
}

// NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentWithDefaults() *RoleAssignment {
	this := RoleAssignment{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *RoleAssignment) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *RoleAssignment) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *RoleAssignment) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *RoleAssignment) GetGateway() GatewayInstanceGateway {
	if o == nil || o.Gateway == nil {
		var ret GatewayInstanceGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetGatewayOk() (*GatewayInstanceGateway, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *RoleAssignment) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given GatewayInstanceGateway and assigns it to the Gateway field.
func (o *RoleAssignment) SetGateway(v GatewayInstanceGateway) {
	o.Gateway = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleAssignment) SetId(v string) {
	o.Id = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *RoleAssignment) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *RoleAssignment) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *RoleAssignment) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRole returns the Role field value
func (o *RoleAssignment) GetRole() RoleAssignmentRole {
	if o == nil {
		var ret RoleAssignmentRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetRoleOk() (*RoleAssignmentRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleAssignment) SetRole(v RoleAssignmentRole) {
	o.Role = v
}

// GetScope returns the Scope field value
func (o *RoleAssignment) GetScope() RoleAssignmentScope {
	if o == nil {
		var ret RoleAssignmentScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetScopeOk() (*RoleAssignmentScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *RoleAssignment) SetScope(v RoleAssignmentScope) {
	o.Scope = v
}

func (o RoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignment struct {
	value *RoleAssignment
	isSet bool
}

func (v NullableRoleAssignment) Get() *RoleAssignment {
	return v.value
}

func (v *NullableRoleAssignment) Set(val *RoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignment(val *RoleAssignment) *NullableRoleAssignment {
	return &NullableRoleAssignment{value: val, isSet: true}
}

func (v NullableRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


