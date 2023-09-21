/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
	"time"
)

// checks if the VerifyPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyPolicy{}

// VerifyPolicy struct for VerifyPolicy
type VerifyPolicy struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// Name displayed in PingOne Admin UI.
	Name string `json:"name"`
	// Description displayed in PingOne Admin UI, 1-1024 characters.
	Description *string `json:"description,omitempty"`
	// Required as true to set this verify policy as the default policy for the environment; otherwise optional and defaults to false.
	Default *bool `json:"default,omitempty"`
	GovernmentId *GovernmentIdConfiguration `json:"governmentId,omitempty"`
	FacialComparison *FacialComparisonConfiguration `json:"facialComparison,omitempty"`
	Liveness *LivenessConfiguration `json:"liveness,omitempty"`
	Email *OTPDeviceConfiguration `json:"email,omitempty"`
	Phone *OTPDeviceConfiguration `json:"phone,omitempty"`
	Transaction *TransactionConfiguration `json:"transaction,omitempty"`
	Voice *VoiceConfiguration `json:"voice,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewVerifyPolicy instantiates a new VerifyPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPolicy(name string) *VerifyPolicy {
	this := VerifyPolicy{}
	this.Name = name
	return &this
}

// NewVerifyPolicyWithDefaults instantiates a new VerifyPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPolicyWithDefaults() *VerifyPolicy {
	this := VerifyPolicy{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VerifyPolicy) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VerifyPolicy) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *VerifyPolicy) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VerifyPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VerifyPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VerifyPolicy) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *VerifyPolicy) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *VerifyPolicy) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *VerifyPolicy) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetName returns the Name field value
func (o *VerifyPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VerifyPolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VerifyPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VerifyPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VerifyPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *VerifyPolicy) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *VerifyPolicy) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *VerifyPolicy) SetDefault(v bool) {
	o.Default = &v
}

// GetGovernmentId returns the GovernmentId field value if set, zero value otherwise.
func (o *VerifyPolicy) GetGovernmentId() GovernmentIdConfiguration {
	if o == nil || IsNil(o.GovernmentId) {
		var ret GovernmentIdConfiguration
		return ret
	}
	return *o.GovernmentId
}

// GetGovernmentIdOk returns a tuple with the GovernmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetGovernmentIdOk() (*GovernmentIdConfiguration, bool) {
	if o == nil || IsNil(o.GovernmentId) {
		return nil, false
	}
	return o.GovernmentId, true
}

// HasGovernmentId returns a boolean if a field has been set.
func (o *VerifyPolicy) HasGovernmentId() bool {
	if o != nil && !IsNil(o.GovernmentId) {
		return true
	}

	return false
}

// SetGovernmentId gets a reference to the given GovernmentIdConfiguration and assigns it to the GovernmentId field.
func (o *VerifyPolicy) SetGovernmentId(v GovernmentIdConfiguration) {
	o.GovernmentId = &v
}

// GetFacialComparison returns the FacialComparison field value if set, zero value otherwise.
func (o *VerifyPolicy) GetFacialComparison() FacialComparisonConfiguration {
	if o == nil || IsNil(o.FacialComparison) {
		var ret FacialComparisonConfiguration
		return ret
	}
	return *o.FacialComparison
}

// GetFacialComparisonOk returns a tuple with the FacialComparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetFacialComparisonOk() (*FacialComparisonConfiguration, bool) {
	if o == nil || IsNil(o.FacialComparison) {
		return nil, false
	}
	return o.FacialComparison, true
}

// HasFacialComparison returns a boolean if a field has been set.
func (o *VerifyPolicy) HasFacialComparison() bool {
	if o != nil && !IsNil(o.FacialComparison) {
		return true
	}

	return false
}

// SetFacialComparison gets a reference to the given FacialComparisonConfiguration and assigns it to the FacialComparison field.
func (o *VerifyPolicy) SetFacialComparison(v FacialComparisonConfiguration) {
	o.FacialComparison = &v
}

// GetLiveness returns the Liveness field value if set, zero value otherwise.
func (o *VerifyPolicy) GetLiveness() LivenessConfiguration {
	if o == nil || IsNil(o.Liveness) {
		var ret LivenessConfiguration
		return ret
	}
	return *o.Liveness
}

// GetLivenessOk returns a tuple with the Liveness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetLivenessOk() (*LivenessConfiguration, bool) {
	if o == nil || IsNil(o.Liveness) {
		return nil, false
	}
	return o.Liveness, true
}

// HasLiveness returns a boolean if a field has been set.
func (o *VerifyPolicy) HasLiveness() bool {
	if o != nil && !IsNil(o.Liveness) {
		return true
	}

	return false
}

// SetLiveness gets a reference to the given LivenessConfiguration and assigns it to the Liveness field.
func (o *VerifyPolicy) SetLiveness(v LivenessConfiguration) {
	o.Liveness = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *VerifyPolicy) GetEmail() OTPDeviceConfiguration {
	if o == nil || IsNil(o.Email) {
		var ret OTPDeviceConfiguration
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetEmailOk() (*OTPDeviceConfiguration, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *VerifyPolicy) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given OTPDeviceConfiguration and assigns it to the Email field.
func (o *VerifyPolicy) SetEmail(v OTPDeviceConfiguration) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *VerifyPolicy) GetPhone() OTPDeviceConfiguration {
	if o == nil || IsNil(o.Phone) {
		var ret OTPDeviceConfiguration
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetPhoneOk() (*OTPDeviceConfiguration, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *VerifyPolicy) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given OTPDeviceConfiguration and assigns it to the Phone field.
func (o *VerifyPolicy) SetPhone(v OTPDeviceConfiguration) {
	o.Phone = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *VerifyPolicy) GetTransaction() TransactionConfiguration {
	if o == nil || IsNil(o.Transaction) {
		var ret TransactionConfiguration
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetTransactionOk() (*TransactionConfiguration, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *VerifyPolicy) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given TransactionConfiguration and assigns it to the Transaction field.
func (o *VerifyPolicy) SetTransaction(v TransactionConfiguration) {
	o.Transaction = &v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *VerifyPolicy) GetVoice() VoiceConfiguration {
	if o == nil || IsNil(o.Voice) {
		var ret VoiceConfiguration
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetVoiceOk() (*VoiceConfiguration, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *VerifyPolicy) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given VoiceConfiguration and assigns it to the Voice field.
func (o *VerifyPolicy) SetVoice(v VoiceConfiguration) {
	o.Voice = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VerifyPolicy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VerifyPolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VerifyPolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VerifyPolicy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VerifyPolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VerifyPolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o VerifyPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyPolicy) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.GovernmentId) {
		toSerialize["governmentId"] = o.GovernmentId
	}
	if !IsNil(o.FacialComparison) {
		toSerialize["facialComparison"] = o.FacialComparison
	}
	if !IsNil(o.Liveness) {
		toSerialize["liveness"] = o.Liveness
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableVerifyPolicy struct {
	value *VerifyPolicy
	isSet bool
}

func (v NullableVerifyPolicy) Get() *VerifyPolicy {
	return v.value
}

func (v *NullableVerifyPolicy) Set(val *VerifyPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPolicy(val *VerifyPolicy) *NullableVerifyPolicy {
	return &NullableVerifyPolicy{value: val, isSet: true}
}

func (v NullableVerifyPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


