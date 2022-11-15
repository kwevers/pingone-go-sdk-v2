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

// LanguageLocalizationStatus struct for LanguageLocalizationStatus
type LanguageLocalizationStatus struct {
	// The time the language localization status resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	Locale *LanguageLocalizationStatusLocale `json:"locale,omitempty"`
	// The name of the service associated with this resource.
	Service string `json:"service"`
	LocalizationComplete *bool `json:"localizationComplete,omitempty"`
	StatusDetails *string `json:"statusDetails,omitempty"`
	// The time the language localization status resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewLanguageLocalizationStatus instantiates a new LanguageLocalizationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageLocalizationStatus(service string) *LanguageLocalizationStatus {
	this := LanguageLocalizationStatus{}
	this.Service = service
	return &this
}

// NewLanguageLocalizationStatusWithDefaults instantiates a new LanguageLocalizationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageLocalizationStatusWithDefaults() *LanguageLocalizationStatus {
	this := LanguageLocalizationStatus{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LanguageLocalizationStatus) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *LanguageLocalizationStatus) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LanguageLocalizationStatus) SetId(v string) {
	o.Id = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetLocale() LanguageLocalizationStatusLocale {
	if o == nil || isNil(o.Locale) {
		var ret LanguageLocalizationStatusLocale
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetLocaleOk() (*LanguageLocalizationStatusLocale, bool) {
	if o == nil || isNil(o.Locale) {
    return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasLocale() bool {
	if o != nil && !isNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given LanguageLocalizationStatusLocale and assigns it to the Locale field.
func (o *LanguageLocalizationStatus) SetLocale(v LanguageLocalizationStatusLocale) {
	o.Locale = &v
}

// GetService returns the Service field value
func (o *LanguageLocalizationStatus) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetServiceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *LanguageLocalizationStatus) SetService(v string) {
	o.Service = v
}

// GetLocalizationComplete returns the LocalizationComplete field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetLocalizationComplete() bool {
	if o == nil || isNil(o.LocalizationComplete) {
		var ret bool
		return ret
	}
	return *o.LocalizationComplete
}

// GetLocalizationCompleteOk returns a tuple with the LocalizationComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetLocalizationCompleteOk() (*bool, bool) {
	if o == nil || isNil(o.LocalizationComplete) {
    return nil, false
	}
	return o.LocalizationComplete, true
}

// HasLocalizationComplete returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasLocalizationComplete() bool {
	if o != nil && !isNil(o.LocalizationComplete) {
		return true
	}

	return false
}

// SetLocalizationComplete gets a reference to the given bool and assigns it to the LocalizationComplete field.
func (o *LanguageLocalizationStatus) SetLocalizationComplete(v bool) {
	o.LocalizationComplete = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetStatusDetails() string {
	if o == nil || isNil(o.StatusDetails) {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetStatusDetailsOk() (*string, bool) {
	if o == nil || isNil(o.StatusDetails) {
    return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasStatusDetails() bool {
	if o != nil && !isNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *LanguageLocalizationStatus) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LanguageLocalizationStatus) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLocalizationStatus) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LanguageLocalizationStatus) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LanguageLocalizationStatus) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o LanguageLocalizationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.LocalizationComplete) {
		toSerialize["localizationComplete"] = o.LocalizationComplete
	}
	if !isNil(o.StatusDetails) {
		toSerialize["statusDetails"] = o.StatusDetails
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLanguageLocalizationStatus struct {
	value *LanguageLocalizationStatus
	isSet bool
}

func (v NullableLanguageLocalizationStatus) Get() *LanguageLocalizationStatus {
	return v.value
}

func (v *NullableLanguageLocalizationStatus) Set(val *LanguageLocalizationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageLocalizationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageLocalizationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageLocalizationStatus(val *LanguageLocalizationStatus) *NullableLanguageLocalizationStatus {
	return &NullableLanguageLocalizationStatus{value: val, isSet: true}
}

func (v NullableLanguageLocalizationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageLocalizationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


