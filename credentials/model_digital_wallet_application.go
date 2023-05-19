/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"time"
)

// checks if the DigitalWalletApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletApplication{}

// DigitalWalletApplication struct for DigitalWalletApplication
type DigitalWalletApplication struct {
	Application ObjectApplication `json:"application"`
	// A string that specifies the URL sent in notifications to the user to communicate with the service.
	AppOpenUrl string `json:"appOpenUrl"`
	// A string that specifies the date and time the credential digital wallet app was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the identifier (UUID) associated with the credential digital wallet app.
	Id *string `json:"id,omitempty"`
	// A string that specifies the name associated with the digital wallet app.
	Name string `json:"name"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A boolean that specifies whether the user's wallet app uses the PingOne Wallet SDK.
	UsesPingOneWalletSDK *bool `json:"usesPingOneWalletSDK,omitempty"`
}

// NewDigitalWalletApplication instantiates a new DigitalWalletApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletApplication(application ObjectApplication, appOpenUrl string, name string) *DigitalWalletApplication {
	this := DigitalWalletApplication{}
	this.Application = application
	this.AppOpenUrl = appOpenUrl
	this.Name = name
	return &this
}

// NewDigitalWalletApplicationWithDefaults instantiates a new DigitalWalletApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletApplicationWithDefaults() *DigitalWalletApplication {
	this := DigitalWalletApplication{}
	return &this
}

// GetApplication returns the Application field value
func (o *DigitalWalletApplication) GetApplication() ObjectApplication {
	if o == nil {
		var ret ObjectApplication
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetApplicationOk() (*ObjectApplication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *DigitalWalletApplication) SetApplication(v ObjectApplication) {
	o.Application = v
}

// GetAppOpenUrl returns the AppOpenUrl field value
func (o *DigitalWalletApplication) GetAppOpenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppOpenUrl
}

// GetAppOpenUrlOk returns a tuple with the AppOpenUrl field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetAppOpenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppOpenUrl, true
}

// SetAppOpenUrl sets field value
func (o *DigitalWalletApplication) SetAppOpenUrl(v string) {
	o.AppOpenUrl = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DigitalWalletApplication) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DigitalWalletApplication) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DigitalWalletApplication) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *DigitalWalletApplication) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *DigitalWalletApplication) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *DigitalWalletApplication) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DigitalWalletApplication) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DigitalWalletApplication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DigitalWalletApplication) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DigitalWalletApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DigitalWalletApplication) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DigitalWalletApplication) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DigitalWalletApplication) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DigitalWalletApplication) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsesPingOneWalletSDK returns the UsesPingOneWalletSDK field value if set, zero value otherwise.
func (o *DigitalWalletApplication) GetUsesPingOneWalletSDK() bool {
	if o == nil || IsNil(o.UsesPingOneWalletSDK) {
		var ret bool
		return ret
	}
	return *o.UsesPingOneWalletSDK
}

// GetUsesPingOneWalletSDKOk returns a tuple with the UsesPingOneWalletSDK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletApplication) GetUsesPingOneWalletSDKOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesPingOneWalletSDK) {
		return nil, false
	}
	return o.UsesPingOneWalletSDK, true
}

// HasUsesPingOneWalletSDK returns a boolean if a field has been set.
func (o *DigitalWalletApplication) HasUsesPingOneWalletSDK() bool {
	if o != nil && !IsNil(o.UsesPingOneWalletSDK) {
		return true
	}

	return false
}

// SetUsesPingOneWalletSDK gets a reference to the given bool and assigns it to the UsesPingOneWalletSDK field.
func (o *DigitalWalletApplication) SetUsesPingOneWalletSDK(v bool) {
	o.UsesPingOneWalletSDK = &v
}

func (o DigitalWalletApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application"] = o.Application
	toSerialize["appOpenUrl"] = o.AppOpenUrl
	// skip: createdAt is readOnly
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: updatedAt is readOnly
	if !IsNil(o.UsesPingOneWalletSDK) {
		toSerialize["usesPingOneWalletSDK"] = o.UsesPingOneWalletSDK
	}
	return toSerialize, nil
}

type NullableDigitalWalletApplication struct {
	value *DigitalWalletApplication
	isSet bool
}

func (v NullableDigitalWalletApplication) Get() *DigitalWalletApplication {
	return v.value
}

func (v *NullableDigitalWalletApplication) Set(val *DigitalWalletApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletApplication(val *DigitalWalletApplication) *NullableDigitalWalletApplication {
	return &NullableDigitalWalletApplication{value: val, isSet: true}
}

func (v NullableDigitalWalletApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

