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

// checks if the Group type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Group{}

// Group struct for Group
type Group struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// The unique identifier for the group. Search all groups for a specific group ID with a SCIM filter on GET /environments/{environmentID}/groups. Retrieve all the group IDs associated with a user with GET /environments/{environmentID}/users/{userID}?include=memberOfGroupIDs.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// For external groups, set during user creation/update. For groups created on PingOne, this parameter is identical to `name`.
	DisplayName *string `json:"displayName,omitempty"`
	Population *GroupPopulation `json:"population,omitempty"`
	// The group name. A group name can be reused across populations, but the same user cannot belong to multiple groups with the same group name. Population groups cannot share a group name with an environment group. Search all groups for a specific group name with a SCIM filter on GET /environments/{environmentID}/groups. Retrieve all the group names associated with a user with GET /environments/{environmentID}/users/{userID}?include=memberOfGroupNames. Use this operation to determine group membership in attribute mappings for claims and assertions.
	Name string `json:"name"`
	// A SCIM filter that determines which users are dynamically added to the group. For more information, see Adding users to a group and Removing users from a group.
	UserFilter *string `json:"userFilter,omitempty"`
	// The group description.
	Description *string `json:"description,omitempty"`
	// A user-defined identifier for the group. Use this propertry to syncronize a group in PingOne with the same group in an external system. PingOne does not directly use this property. Search all groups for a specific external ID with a SCIM filter on GET /environments/{environmentID}/groups
	ExternalId *string `json:"externalId,omitempty"`
	// Optional User-defined custom data.
	CustomData map[string]interface{} `json:"customData,omitempty"`
	// External groups only. Set during user creation/update.
	SourceId *string `json:"sourceId,omitempty"`
	SourceType *EnumGroupSourceType `json:"sourceType,omitempty"`
	DirectMemberCounts *GroupDirectMemberCounts `json:"directMemberCounts,omitempty"`
	TotalMemberCounts *GroupTotalMemberCounts `json:"totalMemberCounts,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(name string) *Group {
	this := Group{}
	this.Name = name
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Group) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Group) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *Group) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Group) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Group) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Group) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Group) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Group) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Group) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPopulation returns the Population field value if set, zero value otherwise.
func (o *Group) GetPopulation() GroupPopulation {
	if o == nil || IsNil(o.Population) {
		var ret GroupPopulation
		return ret
	}
	return *o.Population
}

// GetPopulationOk returns a tuple with the Population field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetPopulationOk() (*GroupPopulation, bool) {
	if o == nil || IsNil(o.Population) {
		return nil, false
	}
	return o.Population, true
}

// HasPopulation returns a boolean if a field has been set.
func (o *Group) HasPopulation() bool {
	if o != nil && !IsNil(o.Population) {
		return true
	}

	return false
}

// SetPopulation gets a reference to the given GroupPopulation and assigns it to the Population field.
func (o *Group) SetPopulation(v GroupPopulation) {
	o.Population = &v
}

// GetName returns the Name field value
func (o *Group) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Group) SetName(v string) {
	o.Name = v
}

// GetUserFilter returns the UserFilter field value if set, zero value otherwise.
func (o *Group) GetUserFilter() string {
	if o == nil || IsNil(o.UserFilter) {
		var ret string
		return ret
	}
	return *o.UserFilter
}

// GetUserFilterOk returns a tuple with the UserFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUserFilterOk() (*string, bool) {
	if o == nil || IsNil(o.UserFilter) {
		return nil, false
	}
	return o.UserFilter, true
}

// HasUserFilter returns a boolean if a field has been set.
func (o *Group) HasUserFilter() bool {
	if o != nil && !IsNil(o.UserFilter) {
		return true
	}

	return false
}

// SetUserFilter gets a reference to the given string and assigns it to the UserFilter field.
func (o *Group) SetUserFilter(v string) {
	o.UserFilter = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Group) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Group) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Group) SetDescription(v string) {
	o.Description = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Group) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Group) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Group) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *Group) GetCustomData() map[string]interface{} {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomData) {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *Group) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]interface{} and assigns it to the CustomData field.
func (o *Group) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *Group) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *Group) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *Group) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *Group) GetSourceType() EnumGroupSourceType {
	if o == nil || IsNil(o.SourceType) {
		var ret EnumGroupSourceType
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSourceTypeOk() (*EnumGroupSourceType, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *Group) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given EnumGroupSourceType and assigns it to the SourceType field.
func (o *Group) SetSourceType(v EnumGroupSourceType) {
	o.SourceType = &v
}

// GetDirectMemberCounts returns the DirectMemberCounts field value if set, zero value otherwise.
func (o *Group) GetDirectMemberCounts() GroupDirectMemberCounts {
	if o == nil || IsNil(o.DirectMemberCounts) {
		var ret GroupDirectMemberCounts
		return ret
	}
	return *o.DirectMemberCounts
}

// GetDirectMemberCountsOk returns a tuple with the DirectMemberCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDirectMemberCountsOk() (*GroupDirectMemberCounts, bool) {
	if o == nil || IsNil(o.DirectMemberCounts) {
		return nil, false
	}
	return o.DirectMemberCounts, true
}

// HasDirectMemberCounts returns a boolean if a field has been set.
func (o *Group) HasDirectMemberCounts() bool {
	if o != nil && !IsNil(o.DirectMemberCounts) {
		return true
	}

	return false
}

// SetDirectMemberCounts gets a reference to the given GroupDirectMemberCounts and assigns it to the DirectMemberCounts field.
func (o *Group) SetDirectMemberCounts(v GroupDirectMemberCounts) {
	o.DirectMemberCounts = &v
}

// GetTotalMemberCounts returns the TotalMemberCounts field value if set, zero value otherwise.
func (o *Group) GetTotalMemberCounts() GroupTotalMemberCounts {
	if o == nil || IsNil(o.TotalMemberCounts) {
		var ret GroupTotalMemberCounts
		return ret
	}
	return *o.TotalMemberCounts
}

// GetTotalMemberCountsOk returns a tuple with the TotalMemberCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetTotalMemberCountsOk() (*GroupTotalMemberCounts, bool) {
	if o == nil || IsNil(o.TotalMemberCounts) {
		return nil, false
	}
	return o.TotalMemberCounts, true
}

// HasTotalMemberCounts returns a boolean if a field has been set.
func (o *Group) HasTotalMemberCounts() bool {
	if o != nil && !IsNil(o.TotalMemberCounts) {
		return true
	}

	return false
}

// SetTotalMemberCounts gets a reference to the given GroupTotalMemberCounts and assigns it to the TotalMemberCounts field.
func (o *Group) SetTotalMemberCounts(v GroupTotalMemberCounts) {
	o.TotalMemberCounts = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Population) {
		toSerialize["population"] = o.Population
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.UserFilter) {
		toSerialize["userFilter"] = o.UserFilter
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.DirectMemberCounts) {
		toSerialize["directMemberCounts"] = o.DirectMemberCounts
	}
	if !IsNil(o.TotalMemberCounts) {
		toSerialize["totalMemberCounts"] = o.TotalMemberCounts
	}
	return toSerialize, nil
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


