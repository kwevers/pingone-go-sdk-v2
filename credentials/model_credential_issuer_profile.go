/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"time"
)

// checks if the CredentialIssuerProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuerProfile{}

// CredentialIssuerProfile struct for CredentialIssuerProfile
type CredentialIssuerProfile struct {
	ApplicationInstance *CredentialIssuerProfileApplicationInstance `json:"applicationInstance,omitempty"`
	// A string that specifies the date and time the issuer profile was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the identifier (UUID) of the credential issuer.
	Id *string `json:"id,omitempty"`
	// The name of the credential issuer. This will be included in credentials issued.
	Name string `json:"name"`
	// A string that specifies the date and time the credential issuer profile was last updated; can be null.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A string that specifies the base URL associated with the credential issuer.
	SiteUrl *string `json:"siteUrl,omitempty"`
	// A string that specifies the default notification template used in credential issuance notifications. Deprecated.
	CustomEmailTemplate *string `json:"customEmailTemplate,omitempty"`
}

// NewCredentialIssuerProfile instantiates a new CredentialIssuerProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuerProfile(name string) *CredentialIssuerProfile {
	this := CredentialIssuerProfile{}
	this.Name = name
	return &this
}

// NewCredentialIssuerProfileWithDefaults instantiates a new CredentialIssuerProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuerProfileWithDefaults() *CredentialIssuerProfile {
	this := CredentialIssuerProfile{}
	return &this
}

// GetApplicationInstance returns the ApplicationInstance field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetApplicationInstance() CredentialIssuerProfileApplicationInstance {
	if o == nil || IsNil(o.ApplicationInstance) {
		var ret CredentialIssuerProfileApplicationInstance
		return ret
	}
	return *o.ApplicationInstance
}

// GetApplicationInstanceOk returns a tuple with the ApplicationInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetApplicationInstanceOk() (*CredentialIssuerProfileApplicationInstance, bool) {
	if o == nil || IsNil(o.ApplicationInstance) {
		return nil, false
	}
	return o.ApplicationInstance, true
}

// HasApplicationInstance returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasApplicationInstance() bool {
	if o != nil && !IsNil(o.ApplicationInstance) {
		return true
	}

	return false
}

// SetApplicationInstance gets a reference to the given CredentialIssuerProfileApplicationInstance and assigns it to the ApplicationInstance field.
func (o *CredentialIssuerProfile) SetApplicationInstance(v CredentialIssuerProfileApplicationInstance) {
	o.ApplicationInstance = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CredentialIssuerProfile) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *CredentialIssuerProfile) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialIssuerProfile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CredentialIssuerProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CredentialIssuerProfile) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CredentialIssuerProfile) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSiteUrl returns the SiteUrl field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetSiteUrl() string {
	if o == nil || IsNil(o.SiteUrl) {
		var ret string
		return ret
	}
	return *o.SiteUrl
}

// GetSiteUrlOk returns a tuple with the SiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetSiteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SiteUrl) {
		return nil, false
	}
	return o.SiteUrl, true
}

// HasSiteUrl returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasSiteUrl() bool {
	if o != nil && !IsNil(o.SiteUrl) {
		return true
	}

	return false
}

// SetSiteUrl gets a reference to the given string and assigns it to the SiteUrl field.
func (o *CredentialIssuerProfile) SetSiteUrl(v string) {
	o.SiteUrl = &v
}

// GetCustomEmailTemplate returns the CustomEmailTemplate field value if set, zero value otherwise.
func (o *CredentialIssuerProfile) GetCustomEmailTemplate() string {
	if o == nil || IsNil(o.CustomEmailTemplate) {
		var ret string
		return ret
	}
	return *o.CustomEmailTemplate
}

// GetCustomEmailTemplateOk returns a tuple with the CustomEmailTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerProfile) GetCustomEmailTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEmailTemplate) {
		return nil, false
	}
	return o.CustomEmailTemplate, true
}

// HasCustomEmailTemplate returns a boolean if a field has been set.
func (o *CredentialIssuerProfile) HasCustomEmailTemplate() bool {
	if o != nil && !IsNil(o.CustomEmailTemplate) {
		return true
	}

	return false
}

// SetCustomEmailTemplate gets a reference to the given string and assigns it to the CustomEmailTemplate field.
func (o *CredentialIssuerProfile) SetCustomEmailTemplate(v string) {
	o.CustomEmailTemplate = &v
}

func (o CredentialIssuerProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuerProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationInstance) {
		toSerialize["applicationInstance"] = o.ApplicationInstance
	}
	// skip: createdAt is readOnly
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: updatedAt is readOnly
	if !IsNil(o.SiteUrl) {
		toSerialize["siteUrl"] = o.SiteUrl
	}
	if !IsNil(o.CustomEmailTemplate) {
		toSerialize["customEmailTemplate"] = o.CustomEmailTemplate
	}
	return toSerialize, nil
}

type NullableCredentialIssuerProfile struct {
	value *CredentialIssuerProfile
	isSet bool
}

func (v NullableCredentialIssuerProfile) Get() *CredentialIssuerProfile {
	return v.value
}

func (v *NullableCredentialIssuerProfile) Set(val *CredentialIssuerProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuerProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuerProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuerProfile(val *CredentialIssuerProfile) *NullableCredentialIssuerProfile {
	return &NullableCredentialIssuerProfile{value: val, isSet: true}
}

func (v NullableCredentialIssuerProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuerProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


