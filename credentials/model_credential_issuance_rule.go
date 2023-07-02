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

// checks if the CredentialIssuanceRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRule{}

// CredentialIssuanceRule struct for CredentialIssuanceRule
type CredentialIssuanceRule struct {
	Automation CredentialIssuanceRuleAutomation `json:"automation"`
	// A string that specifies the date and time the credential issuance rule was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CredentialType *CredentialIssuanceRuleCredentialType `json:"credentialType,omitempty"`
	DigitalWalletApplication *CredentialIssuanceRuleDigitalWalletApplication `json:"digitalWalletApplication,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Filter *CredentialIssuanceRuleFilter `json:"filter,omitempty"`
	// A string that specifies the identifier (UUID) of the credential issuance rule.
	Id *string `json:"id,omitempty"`
	Notification *CredentialIssuanceRuleNotification `json:"notification,omitempty"`
	Status EnumCredentialIssuanceRuleStatus `json:"status"`
	// A string that specifies the date and time the credential issuance rule was last updated; can be null.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewCredentialIssuanceRule instantiates a new CredentialIssuanceRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRule(automation CredentialIssuanceRuleAutomation, status EnumCredentialIssuanceRuleStatus) *CredentialIssuanceRule {
	this := CredentialIssuanceRule{}
	this.Automation = automation
	this.Status = status
	return &this
}

// NewCredentialIssuanceRuleWithDefaults instantiates a new CredentialIssuanceRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleWithDefaults() *CredentialIssuanceRule {
	this := CredentialIssuanceRule{}
	return &this
}

// GetAutomation returns the Automation field value
func (o *CredentialIssuanceRule) GetAutomation() CredentialIssuanceRuleAutomation {
	if o == nil {
		var ret CredentialIssuanceRuleAutomation
		return ret
	}

	return o.Automation
}

// GetAutomationOk returns a tuple with the Automation field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetAutomationOk() (*CredentialIssuanceRuleAutomation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Automation, true
}

// SetAutomation sets field value
func (o *CredentialIssuanceRule) SetAutomation(v CredentialIssuanceRuleAutomation) {
	o.Automation = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CredentialIssuanceRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetCredentialType() CredentialIssuanceRuleCredentialType {
	if o == nil || IsNil(o.CredentialType) {
		var ret CredentialIssuanceRuleCredentialType
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetCredentialTypeOk() (*CredentialIssuanceRuleCredentialType, bool) {
	if o == nil || IsNil(o.CredentialType) {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasCredentialType() bool {
	if o != nil && !IsNil(o.CredentialType) {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given CredentialIssuanceRuleCredentialType and assigns it to the CredentialType field.
func (o *CredentialIssuanceRule) SetCredentialType(v CredentialIssuanceRuleCredentialType) {
	o.CredentialType = &v
}

// GetDigitalWalletApplication returns the DigitalWalletApplication field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetDigitalWalletApplication() CredentialIssuanceRuleDigitalWalletApplication {
	if o == nil || IsNil(o.DigitalWalletApplication) {
		var ret CredentialIssuanceRuleDigitalWalletApplication
		return ret
	}
	return *o.DigitalWalletApplication
}

// GetDigitalWalletApplicationOk returns a tuple with the DigitalWalletApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetDigitalWalletApplicationOk() (*CredentialIssuanceRuleDigitalWalletApplication, bool) {
	if o == nil || IsNil(o.DigitalWalletApplication) {
		return nil, false
	}
	return o.DigitalWalletApplication, true
}

// HasDigitalWalletApplication returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasDigitalWalletApplication() bool {
	if o != nil && !IsNil(o.DigitalWalletApplication) {
		return true
	}

	return false
}

// SetDigitalWalletApplication gets a reference to the given CredentialIssuanceRuleDigitalWalletApplication and assigns it to the DigitalWalletApplication field.
func (o *CredentialIssuanceRule) SetDigitalWalletApplication(v CredentialIssuanceRuleDigitalWalletApplication) {
	o.DigitalWalletApplication = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *CredentialIssuanceRule) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetFilter() CredentialIssuanceRuleFilter {
	if o == nil || IsNil(o.Filter) {
		var ret CredentialIssuanceRuleFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetFilterOk() (*CredentialIssuanceRuleFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given CredentialIssuanceRuleFilter and assigns it to the Filter field.
func (o *CredentialIssuanceRule) SetFilter(v CredentialIssuanceRuleFilter) {
	o.Filter = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialIssuanceRule) SetId(v string) {
	o.Id = &v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetNotification() CredentialIssuanceRuleNotification {
	if o == nil || IsNil(o.Notification) {
		var ret CredentialIssuanceRuleNotification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetNotificationOk() (*CredentialIssuanceRuleNotification, bool) {
	if o == nil || IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasNotification() bool {
	if o != nil && !IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given CredentialIssuanceRuleNotification and assigns it to the Notification field.
func (o *CredentialIssuanceRule) SetNotification(v CredentialIssuanceRuleNotification) {
	o.Notification = &v
}

// GetStatus returns the Status field value
func (o *CredentialIssuanceRule) GetStatus() EnumCredentialIssuanceRuleStatus {
	if o == nil {
		var ret EnumCredentialIssuanceRuleStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetStatusOk() (*EnumCredentialIssuanceRuleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CredentialIssuanceRule) SetStatus(v EnumCredentialIssuanceRuleStatus) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CredentialIssuanceRule) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CredentialIssuanceRule) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CredentialIssuanceRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CredentialIssuanceRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["automation"] = o.Automation
	// skip: createdAt is readOnly
	if !IsNil(o.CredentialType) {
		toSerialize["credentialType"] = o.CredentialType
	}
	if !IsNil(o.DigitalWalletApplication) {
		toSerialize["digitalWalletApplication"] = o.DigitalWalletApplication
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	// skip: id is readOnly
	if !IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	toSerialize["status"] = o.Status
	// skip: updatedAt is readOnly
	return toSerialize, nil
}

type NullableCredentialIssuanceRule struct {
	value *CredentialIssuanceRule
	isSet bool
}

func (v NullableCredentialIssuanceRule) Get() *CredentialIssuanceRule {
	return v.value
}

func (v *NullableCredentialIssuanceRule) Set(val *CredentialIssuanceRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRule(val *CredentialIssuanceRule) *NullableCredentialIssuanceRule {
	return &NullableCredentialIssuanceRule{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


