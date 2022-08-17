/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionIDFirst struct for SignOnPolicyActionIDFirst
type SignOnPolicyActionIDFirst struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Condition *SignOnPolicyActionCommonConditionOrOrInner `json:"condition,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority int32 `json:"priority"`
	SignOnPolicy *SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy,omitempty"`
	Type EnumSignOnPolicyType `json:"type"`
	// The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language
	DiscoveryRules []SignOnPolicyActionIDFirstAllOfDiscoveryRules `json:"discoveryRules,omitempty"`
	// A boolean that if set to true and if the user's account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented.
	EnforceLockoutForIdentityProviders *bool `json:"enforceLockoutForIdentityProviders,omitempty"`
	Recovery *SignOnPolicyActionLoginAllOfRecovery `json:"recovery,omitempty"`
	Registration *SignOnPolicyActionLoginAllOfRegistration `json:"registration,omitempty"`
	// An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow.
	SocialProviders []SignOnPolicyActionLoginAllOfSocialProviders `json:"socialProviders,omitempty"`
}

// NewSignOnPolicyActionIDFirst instantiates a new SignOnPolicyActionIDFirst object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDFirst(priority int32, type_ EnumSignOnPolicyType) *SignOnPolicyActionIDFirst {
	this := SignOnPolicyActionIDFirst{}
	this.Priority = priority
	this.Type = type_
	return &this
}

// NewSignOnPolicyActionIDFirstWithDefaults instantiates a new SignOnPolicyActionIDFirst object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDFirstWithDefaults() *SignOnPolicyActionIDFirst {
	this := SignOnPolicyActionIDFirst{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *SignOnPolicyActionIDFirst) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetCondition() SignOnPolicyActionCommonConditionOrOrInner {
	if o == nil || o.Condition == nil {
		var ret SignOnPolicyActionCommonConditionOrOrInner
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SignOnPolicyActionCommonConditionOrOrInner and assigns it to the Condition field.
func (o *SignOnPolicyActionIDFirst) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner) {
	o.Condition = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionIDFirst) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionIDFirst) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionIDFirst) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionIDFirst) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || o.SignOnPolicy == nil {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || o.SignOnPolicy == nil {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasSignOnPolicy() bool {
	if o != nil && o.SignOnPolicy != nil {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionIDFirst) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionIDFirst) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionIDFirst) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetDiscoveryRules returns the DiscoveryRules field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetDiscoveryRules() []SignOnPolicyActionIDFirstAllOfDiscoveryRules {
	if o == nil || o.DiscoveryRules == nil {
		var ret []SignOnPolicyActionIDFirstAllOfDiscoveryRules
		return ret
	}
	return o.DiscoveryRules
}

// GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetDiscoveryRulesOk() ([]SignOnPolicyActionIDFirstAllOfDiscoveryRules, bool) {
	if o == nil || o.DiscoveryRules == nil {
		return nil, false
	}
	return o.DiscoveryRules, true
}

// HasDiscoveryRules returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasDiscoveryRules() bool {
	if o != nil && o.DiscoveryRules != nil {
		return true
	}

	return false
}

// SetDiscoveryRules gets a reference to the given []SignOnPolicyActionIDFirstAllOfDiscoveryRules and assigns it to the DiscoveryRules field.
func (o *SignOnPolicyActionIDFirst) SetDiscoveryRules(v []SignOnPolicyActionIDFirstAllOfDiscoveryRules) {
	o.DiscoveryRules = v
}

// GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetEnforceLockoutForIdentityProviders() bool {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		var ret bool
		return ret
	}
	return *o.EnforceLockoutForIdentityProviders
}

// GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool) {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		return nil, false
	}
	return o.EnforceLockoutForIdentityProviders, true
}

// HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasEnforceLockoutForIdentityProviders() bool {
	if o != nil && o.EnforceLockoutForIdentityProviders != nil {
		return true
	}

	return false
}

// SetEnforceLockoutForIdentityProviders gets a reference to the given bool and assigns it to the EnforceLockoutForIdentityProviders field.
func (o *SignOnPolicyActionIDFirst) SetEnforceLockoutForIdentityProviders(v bool) {
	o.EnforceLockoutForIdentityProviders = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetRecovery() SignOnPolicyActionLoginAllOfRecovery {
	if o == nil || o.Recovery == nil {
		var ret SignOnPolicyActionLoginAllOfRecovery
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool) {
	if o == nil || o.Recovery == nil {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasRecovery() bool {
	if o != nil && o.Recovery != nil {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given SignOnPolicyActionLoginAllOfRecovery and assigns it to the Recovery field.
func (o *SignOnPolicyActionIDFirst) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetRegistration() SignOnPolicyActionLoginAllOfRegistration {
	if o == nil || o.Registration == nil {
		var ret SignOnPolicyActionLoginAllOfRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool) {
	if o == nil || o.Registration == nil {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasRegistration() bool {
	if o != nil && o.Registration != nil {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionLoginAllOfRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionIDFirst) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration) {
	o.Registration = &v
}

// GetSocialProviders returns the SocialProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirst) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders {
	if o == nil || o.SocialProviders == nil {
		var ret []SignOnPolicyActionLoginAllOfSocialProviders
		return ret
	}
	return o.SocialProviders
}

// GetSocialProvidersOk returns a tuple with the SocialProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirst) GetSocialProvidersOk() ([]SignOnPolicyActionLoginAllOfSocialProviders, bool) {
	if o == nil || o.SocialProviders == nil {
		return nil, false
	}
	return o.SocialProviders, true
}

// HasSocialProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirst) HasSocialProviders() bool {
	if o != nil && o.SocialProviders != nil {
		return true
	}

	return false
}

// SetSocialProviders gets a reference to the given []SignOnPolicyActionLoginAllOfSocialProviders and assigns it to the SocialProviders field.
func (o *SignOnPolicyActionIDFirst) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders) {
	o.SocialProviders = v
}

func (o SignOnPolicyActionIDFirst) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if o.SignOnPolicy != nil {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.DiscoveryRules != nil {
		toSerialize["discoveryRules"] = o.DiscoveryRules
	}
	if o.EnforceLockoutForIdentityProviders != nil {
		toSerialize["enforceLockoutForIdentityProviders"] = o.EnforceLockoutForIdentityProviders
	}
	if o.Recovery != nil {
		toSerialize["recovery"] = o.Recovery
	}
	if o.Registration != nil {
		toSerialize["registration"] = o.Registration
	}
	if o.SocialProviders != nil {
		toSerialize["socialProviders"] = o.SocialProviders
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDFirst struct {
	value *SignOnPolicyActionIDFirst
	isSet bool
}

func (v NullableSignOnPolicyActionIDFirst) Get() *SignOnPolicyActionIDFirst {
	return v.value
}

func (v *NullableSignOnPolicyActionIDFirst) Set(val *SignOnPolicyActionIDFirst) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDFirst) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDFirst) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDFirst(val *SignOnPolicyActionIDFirst) *NullableSignOnPolicyActionIDFirst {
	return &NullableSignOnPolicyActionIDFirst{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDFirst) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDFirst) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


