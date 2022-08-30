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

// SignOnPolicyActionIDPAllOfRegistration Specifies the account registration options.
type SignOnPolicyActionIDPAllOfRegistration struct {
	// A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user's profile during account creation. This is an optional property. If omitted, the default value is set to false.
	ConfirmIdentityProviderAttributes *bool `json:"confirmIdentityProviderAttributes,omitempty"`
	// A boolean that specifies the enabled/disabled state of the policy action. The property is disabled by default when creating a new policy. When enabled, it allows the use of the new user registration flow. This attribute should be set to true when implementing the social login sign-on policy option.
	Enabled bool `json:"enabled"`
	Population *SignOnPolicyActionLoginAllOfRegistrationPopulation `json:"population,omitempty"`
}

// NewSignOnPolicyActionIDPAllOfRegistration instantiates a new SignOnPolicyActionIDPAllOfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDPAllOfRegistration(enabled bool) *SignOnPolicyActionIDPAllOfRegistration {
	this := SignOnPolicyActionIDPAllOfRegistration{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	this.Enabled = enabled
	return &this
}

// NewSignOnPolicyActionIDPAllOfRegistrationWithDefaults instantiates a new SignOnPolicyActionIDPAllOfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDPAllOfRegistrationWithDefaults() *SignOnPolicyActionIDPAllOfRegistration {
	this := SignOnPolicyActionIDPAllOfRegistration{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDPAllOfRegistration) GetConfirmIdentityProviderAttributes() bool {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmIdentityProviderAttributes
}

// GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOfRegistration) GetConfirmIdentityProviderAttributesOk() (*bool, bool) {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		return nil, false
	}
	return o.ConfirmIdentityProviderAttributes, true
}

// HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDPAllOfRegistration) HasConfirmIdentityProviderAttributes() bool {
	if o != nil && o.ConfirmIdentityProviderAttributes != nil {
		return true
	}

	return false
}

// SetConfirmIdentityProviderAttributes gets a reference to the given bool and assigns it to the ConfirmIdentityProviderAttributes field.
func (o *SignOnPolicyActionIDPAllOfRegistration) SetConfirmIdentityProviderAttributes(v bool) {
	o.ConfirmIdentityProviderAttributes = &v
}

// GetEnabled returns the Enabled field value
func (o *SignOnPolicyActionIDPAllOfRegistration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOfRegistration) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SignOnPolicyActionIDPAllOfRegistration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPopulation returns the Population field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDPAllOfRegistration) GetPopulation() SignOnPolicyActionLoginAllOfRegistrationPopulation {
	if o == nil || o.Population == nil {
		var ret SignOnPolicyActionLoginAllOfRegistrationPopulation
		return ret
	}
	return *o.Population
}

// GetPopulationOk returns a tuple with the Population field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOfRegistration) GetPopulationOk() (*SignOnPolicyActionLoginAllOfRegistrationPopulation, bool) {
	if o == nil || o.Population == nil {
		return nil, false
	}
	return o.Population, true
}

// HasPopulation returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDPAllOfRegistration) HasPopulation() bool {
	if o != nil && o.Population != nil {
		return true
	}

	return false
}

// SetPopulation gets a reference to the given SignOnPolicyActionLoginAllOfRegistrationPopulation and assigns it to the Population field.
func (o *SignOnPolicyActionIDPAllOfRegistration) SetPopulation(v SignOnPolicyActionLoginAllOfRegistrationPopulation) {
	o.Population = &v
}

func (o SignOnPolicyActionIDPAllOfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfirmIdentityProviderAttributes != nil {
		toSerialize["confirmIdentityProviderAttributes"] = o.ConfirmIdentityProviderAttributes
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Population != nil {
		toSerialize["population"] = o.Population
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDPAllOfRegistration struct {
	value *SignOnPolicyActionIDPAllOfRegistration
	isSet bool
}

func (v NullableSignOnPolicyActionIDPAllOfRegistration) Get() *SignOnPolicyActionIDPAllOfRegistration {
	return v.value
}

func (v *NullableSignOnPolicyActionIDPAllOfRegistration) Set(val *SignOnPolicyActionIDPAllOfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDPAllOfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDPAllOfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDPAllOfRegistration(val *SignOnPolicyActionIDPAllOfRegistration) *NullableSignOnPolicyActionIDPAllOfRegistration {
	return &NullableSignOnPolicyActionIDPAllOfRegistration{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDPAllOfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDPAllOfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


