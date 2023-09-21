/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"time"
)

// checks if the RiskPredictorIPReputation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorIPReputation{}

// RiskPredictorIPReputation struct for RiskPredictorIPReputation
type RiskPredictorIPReputation struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights.
	Name string `json:"name"`
	// A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details).
	CompactName string `json:"compactName"`
	Type EnumPredictorType `json:"type"`
	// A string type. This specifies the description of the risk predictor. Maximum length is 1024 characters.
	Description *string `json:"description,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Indicates whether PingOne Risk is licensed for the environment.
	Licensed *bool `json:"licensed,omitempty"`
	// A boolean to indicate whether the predictor is deletable in the environment.
	Deletable *bool `json:"deletable,omitempty"`
	Default *RiskPredictorCommonDefault `json:"default,omitempty"`
	Condition *RiskPredictorCommonCondition `json:"condition,omitempty"`
	// A list of IP addresses (CDIRs) that are ignored for the predictor results.
	WhiteList []string `json:"whiteList,omitempty"`
}

// NewRiskPredictorIPReputation instantiates a new RiskPredictorIPReputation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorIPReputation(name string, compactName string, type_ EnumPredictorType) *RiskPredictorIPReputation {
	this := RiskPredictorIPReputation{}
	this.Name = name
	this.CompactName = compactName
	this.Type = type_
	return &this
}

// NewRiskPredictorIPReputationWithDefaults instantiates a new RiskPredictorIPReputation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorIPReputationWithDefaults() *RiskPredictorIPReputation {
	this := RiskPredictorIPReputation{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *RiskPredictorIPReputation) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPredictorIPReputation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RiskPredictorIPReputation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPredictorIPReputation) SetName(v string) {
	o.Name = v
}

// GetCompactName returns the CompactName field value
func (o *RiskPredictorIPReputation) GetCompactName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompactName
}

// GetCompactNameOk returns a tuple with the CompactName field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetCompactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompactName, true
}

// SetCompactName sets field value
func (o *RiskPredictorIPReputation) SetCompactName(v string) {
	o.CompactName = v
}

// GetType returns the Type field value
func (o *RiskPredictorIPReputation) GetType() EnumPredictorType {
	if o == nil {
		var ret EnumPredictorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetTypeOk() (*EnumPredictorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskPredictorIPReputation) SetType(v EnumPredictorType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPredictorIPReputation) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RiskPredictorIPReputation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RiskPredictorIPReputation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *RiskPredictorIPReputation) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetDeletable() bool {
	if o == nil || IsNil(o.Deletable) {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetDeletableOk() (*bool, bool) {
	if o == nil || IsNil(o.Deletable) {
		return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasDeletable() bool {
	if o != nil && !IsNil(o.Deletable) {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *RiskPredictorIPReputation) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetDefault() RiskPredictorCommonDefault {
	if o == nil || IsNil(o.Default) {
		var ret RiskPredictorCommonDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetDefaultOk() (*RiskPredictorCommonDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given RiskPredictorCommonDefault and assigns it to the Default field.
func (o *RiskPredictorIPReputation) SetDefault(v RiskPredictorCommonDefault) {
	o.Default = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetCondition() RiskPredictorCommonCondition {
	if o == nil || IsNil(o.Condition) {
		var ret RiskPredictorCommonCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetConditionOk() (*RiskPredictorCommonCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given RiskPredictorCommonCondition and assigns it to the Condition field.
func (o *RiskPredictorIPReputation) SetCondition(v RiskPredictorCommonCondition) {
	o.Condition = &v
}

// GetWhiteList returns the WhiteList field value if set, zero value otherwise.
func (o *RiskPredictorIPReputation) GetWhiteList() []string {
	if o == nil || IsNil(o.WhiteList) {
		var ret []string
		return ret
	}
	return o.WhiteList
}

// GetWhiteListOk returns a tuple with the WhiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorIPReputation) GetWhiteListOk() ([]string, bool) {
	if o == nil || IsNil(o.WhiteList) {
		return nil, false
	}
	return o.WhiteList, true
}

// HasWhiteList returns a boolean if a field has been set.
func (o *RiskPredictorIPReputation) HasWhiteList() bool {
	if o != nil && !IsNil(o.WhiteList) {
		return true
	}

	return false
}

// SetWhiteList gets a reference to the given []string and assigns it to the WhiteList field.
func (o *RiskPredictorIPReputation) SetWhiteList(v []string) {
	o.WhiteList = v
}

func (o RiskPredictorIPReputation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorIPReputation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["compactName"] = o.CompactName
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Licensed) {
		toSerialize["licensed"] = o.Licensed
	}
	if !IsNil(o.Deletable) {
		toSerialize["deletable"] = o.Deletable
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.WhiteList) {
		toSerialize["whiteList"] = o.WhiteList
	}
	return toSerialize, nil
}

type NullableRiskPredictorIPReputation struct {
	value *RiskPredictorIPReputation
	isSet bool
}

func (v NullableRiskPredictorIPReputation) Get() *RiskPredictorIPReputation {
	return v.value
}

func (v *NullableRiskPredictorIPReputation) Set(val *RiskPredictorIPReputation) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorIPReputation) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorIPReputation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorIPReputation(val *RiskPredictorIPReputation) *NullableRiskPredictorIPReputation {
	return &NullableRiskPredictorIPReputation{value: val, isSet: true}
}

func (v NullableRiskPredictorIPReputation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorIPReputation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


