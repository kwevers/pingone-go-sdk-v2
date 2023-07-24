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

// checks if the RiskPredictorUserRiskBehavior type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorUserRiskBehavior{}

// RiskPredictorUserRiskBehavior struct for RiskPredictorUserRiskBehavior
type RiskPredictorUserRiskBehavior struct {
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
	PredictionModel RiskPredictorUserRiskBehaviorAllOfPredictionModel `json:"predictionModel"`
}

// NewRiskPredictorUserRiskBehavior instantiates a new RiskPredictorUserRiskBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorUserRiskBehavior(name string, compactName string, type_ EnumPredictorType, predictionModel RiskPredictorUserRiskBehaviorAllOfPredictionModel) *RiskPredictorUserRiskBehavior {
	this := RiskPredictorUserRiskBehavior{}
	this.Name = name
	this.CompactName = compactName
	this.Type = type_
	this.PredictionModel = predictionModel
	return &this
}

// NewRiskPredictorUserRiskBehaviorWithDefaults instantiates a new RiskPredictorUserRiskBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorUserRiskBehaviorWithDefaults() *RiskPredictorUserRiskBehavior {
	this := RiskPredictorUserRiskBehavior{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *RiskPredictorUserRiskBehavior) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPredictorUserRiskBehavior) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RiskPredictorUserRiskBehavior) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPredictorUserRiskBehavior) SetName(v string) {
	o.Name = v
}

// GetCompactName returns the CompactName field value
func (o *RiskPredictorUserRiskBehavior) GetCompactName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompactName
}

// GetCompactNameOk returns a tuple with the CompactName field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetCompactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompactName, true
}

// SetCompactName sets field value
func (o *RiskPredictorUserRiskBehavior) SetCompactName(v string) {
	o.CompactName = v
}

// GetType returns the Type field value
func (o *RiskPredictorUserRiskBehavior) GetType() EnumPredictorType {
	if o == nil {
		var ret EnumPredictorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetTypeOk() (*EnumPredictorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskPredictorUserRiskBehavior) SetType(v EnumPredictorType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPredictorUserRiskBehavior) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RiskPredictorUserRiskBehavior) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RiskPredictorUserRiskBehavior) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *RiskPredictorUserRiskBehavior) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetDeletable() bool {
	if o == nil || IsNil(o.Deletable) {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetDeletableOk() (*bool, bool) {
	if o == nil || IsNil(o.Deletable) {
		return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasDeletable() bool {
	if o != nil && !IsNil(o.Deletable) {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *RiskPredictorUserRiskBehavior) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetDefault() RiskPredictorCommonDefault {
	if o == nil || IsNil(o.Default) {
		var ret RiskPredictorCommonDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetDefaultOk() (*RiskPredictorCommonDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given RiskPredictorCommonDefault and assigns it to the Default field.
func (o *RiskPredictorUserRiskBehavior) SetDefault(v RiskPredictorCommonDefault) {
	o.Default = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RiskPredictorUserRiskBehavior) GetCondition() RiskPredictorCommonCondition {
	if o == nil || IsNil(o.Condition) {
		var ret RiskPredictorCommonCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetConditionOk() (*RiskPredictorCommonCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RiskPredictorUserRiskBehavior) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given RiskPredictorCommonCondition and assigns it to the Condition field.
func (o *RiskPredictorUserRiskBehavior) SetCondition(v RiskPredictorCommonCondition) {
	o.Condition = &v
}

// GetPredictionModel returns the PredictionModel field value
func (o *RiskPredictorUserRiskBehavior) GetPredictionModel() RiskPredictorUserRiskBehaviorAllOfPredictionModel {
	if o == nil {
		var ret RiskPredictorUserRiskBehaviorAllOfPredictionModel
		return ret
	}

	return o.PredictionModel
}

// GetPredictionModelOk returns a tuple with the PredictionModel field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehavior) GetPredictionModelOk() (*RiskPredictorUserRiskBehaviorAllOfPredictionModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PredictionModel, true
}

// SetPredictionModel sets field value
func (o *RiskPredictorUserRiskBehavior) SetPredictionModel(v RiskPredictorUserRiskBehaviorAllOfPredictionModel) {
	o.PredictionModel = v
}

func (o RiskPredictorUserRiskBehavior) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorUserRiskBehavior) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["compactName"] = o.CompactName
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	// skip: createdAt is readOnly
	// skip: updatedAt is readOnly
	// skip: licensed is readOnly
	// skip: deletable is readOnly
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	toSerialize["predictionModel"] = o.PredictionModel
	return toSerialize, nil
}

type NullableRiskPredictorUserRiskBehavior struct {
	value *RiskPredictorUserRiskBehavior
	isSet bool
}

func (v NullableRiskPredictorUserRiskBehavior) Get() *RiskPredictorUserRiskBehavior {
	return v.value
}

func (v *NullableRiskPredictorUserRiskBehavior) Set(val *RiskPredictorUserRiskBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorUserRiskBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorUserRiskBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorUserRiskBehavior(val *RiskPredictorUserRiskBehavior) *NullableRiskPredictorUserRiskBehavior {
	return &NullableRiskPredictorUserRiskBehavior{value: val, isSet: true}
}

func (v NullableRiskPredictorUserRiskBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorUserRiskBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


